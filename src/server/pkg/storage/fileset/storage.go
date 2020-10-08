package fileset

import (
	"context"
	"fmt"
	"math"
	"path"
	"time"

	units "github.com/docker/go-units"
	"github.com/pachyderm/pachyderm/src/client/pkg/errors"
	"github.com/pachyderm/pachyderm/src/server/pkg/storage/chunk"
	"github.com/pachyderm/pachyderm/src/server/pkg/storage/fileset/index"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
)

const (
	// TODO Not sure if these are the tags we should use, but the header and padding tag should show up before and after respectively in the
	// lexicographical ordering of file content tags.
	// headerTag is the tag used for the tar header bytes.
	headerTag = ""
	// paddingTag is the tag used for the padding bytes at the end of a tar entry.
	paddingTag = "~"
	// DefaultMemoryThreshold is the default for the memory threshold that must
	// be met before a file set part is serialized (excluding close).
	DefaultMemoryThreshold = 1024 * units.MB
	// DefaultShardThreshold is the default for the size threshold that must
	// be met before a shard is created by the shard function.
	DefaultShardThreshold = 1024 * units.MB
	// DefaultLevelZeroSize is the default size for level zero in the compacted
	// representation of a file set.
	DefaultLevelZeroSize = 1 * units.MB
	// DefaultLevelSizeBase is the default base of the exponential growth function
	// for level sizes in the compacted representation of a file set.
	DefaultLevelSizeBase = 10
	// Diff is the suffix of a path that points to the diff of the prefix.
	Diff = "diff"
	// Compacted is the suffix of a path that points to the compaction of the prefix.
	Compacted = "compacted"
)

var (
	// ErrNoFileSetFound is returned by the methods on Storage when a fileset does not exist
	ErrNoFileSetFound = errors.Errorf("no fileset found")
)

// Storage is the abstraction that manages fileset storage.
type Storage struct {
	paths                        PathStore
	chunks                       *chunk.Storage
	memThreshold, shardThreshold int64
	levelZeroSize                int64
	levelSizeBase                int
	filesetSem                   *semaphore.Weighted
}

// NewStorage creates a new Storage.
func NewStorage(pathStore PathStore, chunks *chunk.Storage, opts ...StorageOption) *Storage {
	s := &Storage{
		paths:          pathStore,
		chunks:         chunks,
		memThreshold:   DefaultMemoryThreshold,
		shardThreshold: DefaultShardThreshold,
		levelZeroSize:  DefaultLevelZeroSize,
		levelSizeBase:  DefaultLevelSizeBase,
		filesetSem:     semaphore.NewWeighted(math.MaxInt64),
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

// ChunkStorage returns the underlying chunk storage instance for this storage instance.
func (s *Storage) ChunkStorage() *chunk.Storage {
	return s.chunks
}

// New creates a new in-memory fileset.
func (s *Storage) New(ctx context.Context, fileSet, defaultTag string, opts ...UnorderedWriterOption) (*UnorderedWriter, error) {
	return newUnorderedWriter(ctx, s, fileSet, s.memThreshold, defaultTag, opts...)
}

// NewWriter makes a Writer backed by the path `fileSet` in object storage.
func (s *Storage) NewWriter(ctx context.Context, fileSet string, opts ...WriterOption) *Writer {
	return s.newWriter(ctx, fileSet, opts...)
}

func (s *Storage) newWriter(ctx context.Context, fileSet string, opts ...WriterOption) *Writer {
	return newWriter(ctx, s.paths, s.chunks, fileSet, opts...)
}

// NewReader makes a Reader backed by the path `fileSet` in object storage.
// TODO Expose some notion of read ahead (read a certain number of chunks in parallel).
// this will be necessary to speed up reading large files.
func (s *Storage) NewReader(ctx context.Context, fileSet string, opts ...index.Option) (*Reader, error) {
	return newReader(ctx, s.paths, s.chunks, fileSet, opts...)
}

// NewMergeReader returns a merge reader for a set for filesets.
func (s *Storage) NewMergeReader(ctx context.Context, fileSets []string, opts ...index.Option) (*MergeReader, error) {
	return s.newMergeReader(ctx, fileSets, opts...)
}

func (s *Storage) newMergeReader(ctx context.Context, fileSets []string, opts ...index.Option) (*MergeReader, error) {
	var rs []*Reader
	for _, fileSet := range fileSets {
		if err := s.paths.Walk(ctx, fileSet, func(name string) error {
			r, err := s.NewReader(ctx, name, opts...)
			if err != nil {
				return err
			}
			rs = append(rs, r)
			return nil
		}); err != nil {
			return nil, err
		}
	}
	return newMergeReader(rs), nil
}

// OpenFileSet makes a source which will iterate over the prefix fileSet
func (s *Storage) OpenFileSet(ctx context.Context, fileSet string, opts ...index.Option) FileSet {
	return &mergeSource{
		s: s,
		getReader: func() (*MergeReader, error) {
			return s.NewMergeReader(ctx, []string{fileSet}, opts...)
		},
	}
}

// Shard shards the merge of the file sets with the passed in prefix into file ranges.
// TODO This should be extended to be more configurable (different criteria
// for creating shards).
func (s *Storage) Shard(ctx context.Context, fileSets []string, shardFunc ShardFunc) error {
	mr, err := s.NewMergeReader(ctx, fileSets)
	if err != nil {
		return err
	}
	return shard(mr, s.shardThreshold, shardFunc)
}

// Copy copies the fileset at srcPrefix to dstPrefix. It does *not* perform compaction
// ttl sets the time to live on the keys under dstPrefix if ttl == 0, it is ignored
func (s *Storage) Copy(ctx context.Context, srcPrefix, dstPrefix string, ttl time.Duration) error {
	// TODO: perform this atomically with postgres
	return s.paths.Walk(ctx, srcPrefix, func(srcPath string) error {
		dstPath := dstPrefix + srcPath[len(srcPrefix):]
		idx, err := s.paths.GetIndex(ctx, srcPath)
		if err != nil {
			return err
		}
		return s.paths.PutIndex(ctx, dstPath, idx, ttl)
	})
}

// CompactStats contains information about what was compacted.
type CompactStats struct {
	OutputSize int64
}

// Compact compacts a set of filesets into an output fileset.
func (s *Storage) Compact(ctx context.Context, outputFileSet string, inputFileSets []string, ttl time.Duration, opts ...index.Option) (*CompactStats, error) {
	var size int64
	w := s.newWriter(ctx, outputFileSet, WithIndexCallback(func(idx *index.Index) error {
		size += idx.SizeBytes
		return nil
	}))
	mr, err := s.newMergeReader(ctx, inputFileSets, opts...)
	if err != nil {
		return nil, err
	}
	if err := mr.WriteTo(w); err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}
	if ttl > 0 {
		if _, err := s.SetTTL(ctx, outputFileSet, ttl); err != nil {
			return nil, err
		}
	}
	return &CompactStats{OutputSize: size}, nil
}

// CompactSpec specifies the input and output for a compaction operation.
type CompactSpec struct {
	Output string
	Input  []string
}

// CompactSpec returns a compaction specification that determines the input filesets (the diff file set and potentially
// compacted filesets) and output fileset.
func (s *Storage) CompactSpec(ctx context.Context, fileSet string, compactedFileSet ...string) (*CompactSpec, error) {
	if len(compactedFileSet) > 1 {
		return nil, errors.Errorf("multiple compacted FileSets")
	}
	spec, err := s.compactSpec(ctx, fileSet, compactedFileSet...)
	if err != nil {
		return nil, err
	}
	return spec, nil
}

func (s *Storage) compactSpec(ctx context.Context, fileSet string, compactedFileSet ...string) (ret *CompactSpec, retErr error) {
	idx, err := s.paths.GetIndex(ctx, path.Join(fileSet, Diff))
	if err != nil {
		return nil, err
	}
	size := idx.SizeBytes
	spec := &CompactSpec{
		Input: []string{path.Join(fileSet, Diff)},
	}
	var level int
	// Handle first commit being compacted.
	if len(compactedFileSet) == 0 {
		for size > s.levelSize(level) {
			level++
		}
		spec.Output = path.Join(fileSet, Compacted, levelName(level))
		return spec, nil
	}
	// While we can't fit it all in the current level
	for {
		levelPath := path.Join(compactedFileSet[0], Compacted, levelName(level))
		idx, err := s.paths.GetIndex(ctx, levelPath)
		if err != nil {
			if err == ErrPathNotExists {
				return nil, err
			}
		} else {
			spec.Input = append(spec.Input, levelPath)
			size += idx.SizeBytes
		}
		if size <= s.levelSize(level) {
			break
		}
		level++
	}
	// Now we know the output level
	spec.Output = path.Join(fileSet, Compacted, levelName(level))
	// Copy the other levels that may exist
	if err := s.paths.Walk(ctx, path.Join(compactedFileSet[0], Compacted), func(src string) error {
		lName := path.Base(src)
		l, err := parseLevel(lName)
		if err != nil {
			return err
		}
		if l > level {
			dst := path.Join(fileSet, Compacted, levelName(l))
			if err := copyPath(ctx, s.paths, s.paths, src, dst); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}
	// Inputs should be ordered with priority from least to greatest.
	for i := 0; i < len(spec.Input)/2; i++ {
		spec.Input[i], spec.Input[len(spec.Input)-1-i] = spec.Input[len(spec.Input)-1-i], spec.Input[i]
	}
	return spec, nil
}

// Delete deletes a fileset.
func (s *Storage) Delete(ctx context.Context, fileSet string) error {
	return s.paths.Walk(ctx, fileSet, func(p string) error {
		return s.paths.Delete(ctx, fileSet)
	})
}

// WalkFileSet calls f with the path of every primitive fileSet under prefix.
func (s *Storage) WalkFileSet(ctx context.Context, prefix string, f func(string) error) error {
	return s.paths.Walk(ctx, prefix, f)
}

// SetTTL sets the time-to-live for the path p.
// if no fileset is found SetTTL returns ErrNoFileSetFound
func (s *Storage) SetTTL(ctx context.Context, p string, ttl time.Duration) (time.Time, error) {
	expiresAt, err := s.paths.SetTTL(ctx, p, ttl)
	switch {
	case err == ErrPathNotExists:
		err = ErrNoFileSetFound
	}
	return expiresAt, err
}

// WithRenewer calls cb with a Renewer, and a context which will be canceled if the renewer is unable to renew a path.
func (s *Storage) WithRenewer(ctx context.Context, ttl time.Duration, cb func(context.Context, *Renewer) error) error {
	r := newRenewer(s, ttl)
	cancelCtx, cf := context.WithCancel(ctx)
	eg, errCtx := errgroup.WithContext(cancelCtx)
	eg.Go(func() error {
		return r.run(errCtx)
	})
	eg.Go(func() error {
		defer cf()
		return cb(errCtx, r)
	})
	return eg.Wait()
}

func (s *Storage) levelSize(i int) int64 {
	return s.levelZeroSize * int64(math.Pow(float64(s.levelSizeBase), float64(i)))
}

const subFileSetFmt = "%020d"
const levelFmt = "level_" + subFileSetFmt

// SubFileSetStr returns the string representation of a subfileset.
func SubFileSetStr(subFileSet int64) string {
	return fmt.Sprintf(subFileSetFmt, subFileSet)
}

func levelName(i int) string {
	return fmt.Sprintf(levelFmt, i)
}

func parseLevel(x string) (int, error) {
	var y int
	_, err := fmt.Sscanf(x, levelFmt, &y)
	return y, err
}
