// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: client/admin/1_7/hashtree/hashtree.proto

package hashtree

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	pfs "github.com/pachyderm/pachyderm/src/client/admin/1_7/pfs"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// FileNodeProto is a node corresponding to a file (which is also a leaf node).
type FileNodeProto struct {
	// Object references an object in the object store which contains the content
	// of the data.
	Objects              []*pfs.Object `protobuf:"bytes,4,rep,name=objects,proto3" json:"objects,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *FileNodeProto) Reset()         { *m = FileNodeProto{} }
func (m *FileNodeProto) String() string { return proto.CompactTextString(m) }
func (*FileNodeProto) ProtoMessage()    {}
func (*FileNodeProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a4a4341d52933f6, []int{0}
}
func (m *FileNodeProto) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FileNodeProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FileNodeProto.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FileNodeProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileNodeProto.Merge(m, src)
}
func (m *FileNodeProto) XXX_Size() int {
	return m.Size()
}
func (m *FileNodeProto) XXX_DiscardUnknown() {
	xxx_messageInfo_FileNodeProto.DiscardUnknown(m)
}

var xxx_messageInfo_FileNodeProto proto.InternalMessageInfo

func (m *FileNodeProto) GetObjects() []*pfs.Object {
	if m != nil {
		return m.Objects
	}
	return nil
}

// DirectoryNodeProto is a node corresponding to a directory.
type DirectoryNodeProto struct {
	// Children of this directory. Note that paths are relative, so if "/foo/bar"
	// has a child "baz", that means that there is a file at "/foo/bar/baz".
	//
	// 'Children' is ordered alphabetically, to quickly check if a new file is
	// overwriting an existing one.
	Children             []string    `protobuf:"bytes,3,rep,name=children,proto3" json:"children,omitempty"`
	Header               *pfs.Object `protobuf:"bytes,4,opt,name=header,proto3" json:"header,omitempty"`
	Footer               *pfs.Object `protobuf:"bytes,5,opt,name=footer,proto3" json:"footer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *DirectoryNodeProto) Reset()         { *m = DirectoryNodeProto{} }
func (m *DirectoryNodeProto) String() string { return proto.CompactTextString(m) }
func (*DirectoryNodeProto) ProtoMessage()    {}
func (*DirectoryNodeProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a4a4341d52933f6, []int{1}
}
func (m *DirectoryNodeProto) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DirectoryNodeProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DirectoryNodeProto.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DirectoryNodeProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DirectoryNodeProto.Merge(m, src)
}
func (m *DirectoryNodeProto) XXX_Size() int {
	return m.Size()
}
func (m *DirectoryNodeProto) XXX_DiscardUnknown() {
	xxx_messageInfo_DirectoryNodeProto.DiscardUnknown(m)
}

var xxx_messageInfo_DirectoryNodeProto proto.InternalMessageInfo

func (m *DirectoryNodeProto) GetChildren() []string {
	if m != nil {
		return m.Children
	}
	return nil
}

func (m *DirectoryNodeProto) GetHeader() *pfs.Object {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DirectoryNodeProto) GetFooter() *pfs.Object {
	if m != nil {
		return m.Footer
	}
	return nil
}

// NodeProto is a node in the file tree (either a file or a directory)
type NodeProto struct {
	// Name is the name (not path) of the file/directory (e.g. /lib).
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Hash is a hash of the node's name and contents (which includes the
	// BlockRefs of a file and the Children of a directory). This can be used to
	// detect if the name or contents have changed between versions.
	Hash []byte `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	// subtree_size is the of the subtree under node; i.e. if this is a directory,
	// subtree_size includes all children.
	SubtreeSize int64 `protobuf:"varint,3,opt,name=subtree_size,json=subtreeSize,proto3" json:"subtree_size,omitempty"`
	// Exactly one of the following fields must be set. The type of this node will
	// be determined by which field is set.
	FileNode             *FileNodeProto      `protobuf:"bytes,4,opt,name=file_node,json=fileNode,proto3" json:"file_node,omitempty"`
	DirNode              *DirectoryNodeProto `protobuf:"bytes,5,opt,name=dir_node,json=dirNode,proto3" json:"dir_node,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *NodeProto) Reset()         { *m = NodeProto{} }
func (m *NodeProto) String() string { return proto.CompactTextString(m) }
func (*NodeProto) ProtoMessage()    {}
func (*NodeProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a4a4341d52933f6, []int{2}
}
func (m *NodeProto) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NodeProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NodeProto.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NodeProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeProto.Merge(m, src)
}
func (m *NodeProto) XXX_Size() int {
	return m.Size()
}
func (m *NodeProto) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeProto.DiscardUnknown(m)
}

var xxx_messageInfo_NodeProto proto.InternalMessageInfo

func (m *NodeProto) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NodeProto) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *NodeProto) GetSubtreeSize() int64 {
	if m != nil {
		return m.SubtreeSize
	}
	return 0
}

func (m *NodeProto) GetFileNode() *FileNodeProto {
	if m != nil {
		return m.FileNode
	}
	return nil
}

func (m *NodeProto) GetDirNode() *DirectoryNodeProto {
	if m != nil {
		return m.DirNode
	}
	return nil
}

// HashTreeProto is a tree corresponding to the complete file contents of a
// pachyderm repo at a given commit (based on a Merkle Tree). We store one
// HashTree for every PFS commit.
type HashTreeProto struct {
	// Version is an arbitrary version number, set by the corresponding library
	// in hashtree.go.  This ensures that if the hash function used to create
	// these trees is changed, we won't run into errors when deserializing old
	// trees. The current version is 1.
	Version int32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// Fs maps each node's path to the NodeProto with that node's details.
	// See "Potential Optimizations" at the end for a compression scheme that
	// could be useful if this map gets too large.
	//
	// Note that the key must end in "/" if an only if the value has .dir_node set
	// (i.e. iff the path points to a directory).
	Fs                   map[string]*NodeProto `protobuf:"bytes,2,rep,name=fs,proto3" json:"fs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *HashTreeProto) Reset()         { *m = HashTreeProto{} }
func (m *HashTreeProto) String() string { return proto.CompactTextString(m) }
func (*HashTreeProto) ProtoMessage()    {}
func (*HashTreeProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a4a4341d52933f6, []int{3}
}
func (m *HashTreeProto) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HashTreeProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HashTreeProto.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HashTreeProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HashTreeProto.Merge(m, src)
}
func (m *HashTreeProto) XXX_Size() int {
	return m.Size()
}
func (m *HashTreeProto) XXX_DiscardUnknown() {
	xxx_messageInfo_HashTreeProto.DiscardUnknown(m)
}

var xxx_messageInfo_HashTreeProto proto.InternalMessageInfo

func (m *HashTreeProto) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *HashTreeProto) GetFs() map[string]*NodeProto {
	if m != nil {
		return m.Fs
	}
	return nil
}

func init() {
	proto.RegisterType((*FileNodeProto)(nil), "hashtree_1_7.FileNodeProto")
	proto.RegisterType((*DirectoryNodeProto)(nil), "hashtree_1_7.DirectoryNodeProto")
	proto.RegisterType((*NodeProto)(nil), "hashtree_1_7.NodeProto")
	proto.RegisterType((*HashTreeProto)(nil), "hashtree_1_7.HashTreeProto")
	proto.RegisterMapType((map[string]*NodeProto)(nil), "hashtree_1_7.HashTreeProto.FsEntry")
}

func init() {
	proto.RegisterFile("client/admin/1_7/hashtree/hashtree.proto", fileDescriptor_1a4a4341d52933f6)
}

var fileDescriptor_1a4a4341d52933f6 = []byte{
	// 440 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x65, 0xed, 0xa4, 0x49, 0x26, 0xa9, 0x40, 0x7b, 0x61, 0x15, 0x24, 0xcb, 0x84, 0x03, 0xe6,
	0x80, 0xad, 0xb6, 0x87, 0x56, 0x85, 0x13, 0x82, 0x8a, 0x53, 0x40, 0x0b, 0x27, 0x2e, 0x91, 0x63,
	0x8f, 0xf1, 0x82, 0xe3, 0x8d, 0x76, 0x37, 0x95, 0xdc, 0x23, 0x5f, 0xc1, 0x57, 0xf0, 0x1d, 0x5c,
	0x90, 0xf8, 0x04, 0x14, 0x7e, 0x04, 0x79, 0xed, 0x34, 0xb5, 0x4a, 0x0f, 0x96, 0xde, 0x9b, 0x99,
	0xf7, 0x76, 0xf6, 0x79, 0x21, 0x48, 0x0a, 0x81, 0xa5, 0x89, 0xe2, 0x74, 0x25, 0xca, 0xe8, 0x68,
	0x71, 0x1a, 0xe5, 0xb1, 0xce, 0x8d, 0x42, 0xbc, 0x06, 0xe1, 0x5a, 0x49, 0x23, 0xe9, 0x64, 0xc7,
	0x17, 0x47, 0x8b, 0xd3, 0xa9, 0x77, 0x4b, 0xb7, 0xce, 0x74, 0xfd, 0x35, 0xd3, 0xb3, 0x73, 0x38,
	0xbc, 0x10, 0x05, 0xce, 0x65, 0x8a, 0xef, 0xad, 0xfc, 0x19, 0x0c, 0xe4, 0xf2, 0x0b, 0x26, 0x46,
	0xb3, 0x9e, 0xef, 0x06, 0xe3, 0xe3, 0xfb, 0xe1, 0x3a, 0xd3, 0xb5, 0x57, 0xf8, 0xce, 0xd6, 0xf9,
	0xae, 0x3f, 0xfb, 0x46, 0x80, 0xbe, 0x16, 0x0a, 0x13, 0x23, 0x55, 0xb5, 0x77, 0x98, 0xc2, 0x30,
	0xc9, 0x45, 0x91, 0x2a, 0x2c, 0x99, 0xeb, 0xbb, 0xc1, 0x88, 0x5f, 0x73, 0xfa, 0x14, 0x0e, 0x72,
	0x8c, 0x53, 0x54, 0xac, 0xe7, 0x93, 0xff, 0x99, 0xb7, 0xed, 0x7a, 0x30, 0x93, 0xd2, 0xa0, 0x62,
	0xfd, 0x3b, 0x06, 0x9b, 0xf6, 0xec, 0x17, 0x81, 0xd1, 0xfe, 0x6c, 0x0a, 0xbd, 0x32, 0x5e, 0x21,
	0x23, 0x3e, 0x09, 0x46, 0xdc, 0xe2, 0xba, 0x56, 0x47, 0xc2, 0x1c, 0x9f, 0x04, 0x13, 0x6e, 0x31,
	0x7d, 0x0c, 0x13, 0xbd, 0x59, 0xda, 0x94, 0xb4, 0xb8, 0x42, 0xe6, 0xfa, 0x24, 0x70, 0xf9, 0xb8,
	0xad, 0x7d, 0x10, 0x57, 0x48, 0xcf, 0x60, 0x94, 0x89, 0x02, 0x17, 0xa5, 0x4c, 0xb1, 0xdd, 0xf6,
	0x51, 0x78, 0x33, 0xdb, 0xb0, 0x13, 0x1c, 0x1f, 0x66, 0x2d, 0xa5, 0x2f, 0x60, 0x98, 0x0a, 0xd5,
	0x08, 0x9b, 0xed, 0xfd, 0xae, 0xf0, 0x76, 0x68, 0x7c, 0x90, 0x0a, 0x55, 0xb3, 0xd9, 0x0f, 0x02,
	0x87, 0x6f, 0x63, 0x9d, 0x7f, 0x54, 0xd8, 0xde, 0x89, 0xc1, 0xe0, 0x12, 0x95, 0x16, 0xb2, 0xb4,
	0xd7, 0xea, 0xf3, 0x1d, 0xa5, 0x27, 0xe0, 0x64, 0x9a, 0x39, 0xf6, 0x37, 0x3d, 0xe9, 0x1e, 0xd1,
	0xb1, 0x08, 0x2f, 0xf4, 0x9b, 0xd2, 0xa8, 0x8a, 0x3b, 0x99, 0x9e, 0xce, 0x61, 0xd0, 0x52, 0xfa,
	0x00, 0xdc, 0xaf, 0x58, 0xb5, 0x61, 0xd5, 0x90, 0x3e, 0x87, 0xfe, 0x65, 0x5c, 0x6c, 0xd0, 0x86,
	0x35, 0x3e, 0x7e, 0xd8, 0x35, 0xdd, 0xaf, 0xdb, 0x4c, 0x9d, 0x3b, 0x67, 0xe4, 0xd5, 0xfc, 0xe7,
	0xd6, 0x23, 0xbf, 0xb7, 0x1e, 0xf9, 0xb3, 0xf5, 0xc8, 0xf7, 0xbf, 0xde, 0xbd, 0x4f, 0x2f, 0x3f,
	0x0b, 0x93, 0x6f, 0x96, 0x61, 0x22, 0x57, 0xd1, 0x3a, 0x4e, 0xf2, 0x2a, 0x45, 0x75, 0x13, 0x69,
	0x95, 0x44, 0x77, 0x3e, 0xe7, 0xe5, 0x81, 0x7d, 0x98, 0x27, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x0a, 0x0e, 0x46, 0x18, 0xf2, 0x02, 0x00, 0x00,
}

func (m *FileNodeProto) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FileNodeProto) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FileNodeProto) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Objects) > 0 {
		for iNdEx := len(m.Objects) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Objects[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintHashtree(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	return len(dAtA) - i, nil
}

func (m *DirectoryNodeProto) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DirectoryNodeProto) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DirectoryNodeProto) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Footer != nil {
		{
			size, err := m.Footer.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHashtree(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHashtree(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Children) > 0 {
		for iNdEx := len(m.Children) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Children[iNdEx])
			copy(dAtA[i:], m.Children[iNdEx])
			i = encodeVarintHashtree(dAtA, i, uint64(len(m.Children[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	return len(dAtA) - i, nil
}

func (m *NodeProto) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NodeProto) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NodeProto) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.DirNode != nil {
		{
			size, err := m.DirNode.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHashtree(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.FileNode != nil {
		{
			size, err := m.FileNode.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHashtree(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.SubtreeSize != 0 {
		i = encodeVarintHashtree(dAtA, i, uint64(m.SubtreeSize))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintHashtree(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintHashtree(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *HashTreeProto) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HashTreeProto) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HashTreeProto) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Fs) > 0 {
		for k := range m.Fs {
			v := m.Fs[k]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintHashtree(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintHashtree(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintHashtree(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Version != 0 {
		i = encodeVarintHashtree(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintHashtree(dAtA []byte, offset int, v uint64) int {
	offset -= sovHashtree(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FileNodeProto) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Objects) > 0 {
		for _, e := range m.Objects {
			l = e.Size()
			n += 1 + l + sovHashtree(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DirectoryNodeProto) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Children) > 0 {
		for _, s := range m.Children {
			l = len(s)
			n += 1 + l + sovHashtree(uint64(l))
		}
	}
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovHashtree(uint64(l))
	}
	if m.Footer != nil {
		l = m.Footer.Size()
		n += 1 + l + sovHashtree(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *NodeProto) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovHashtree(uint64(l))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovHashtree(uint64(l))
	}
	if m.SubtreeSize != 0 {
		n += 1 + sovHashtree(uint64(m.SubtreeSize))
	}
	if m.FileNode != nil {
		l = m.FileNode.Size()
		n += 1 + l + sovHashtree(uint64(l))
	}
	if m.DirNode != nil {
		l = m.DirNode.Size()
		n += 1 + l + sovHashtree(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *HashTreeProto) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Version != 0 {
		n += 1 + sovHashtree(uint64(m.Version))
	}
	if len(m.Fs) > 0 {
		for k, v := range m.Fs {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovHashtree(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovHashtree(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovHashtree(uint64(mapEntrySize))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovHashtree(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHashtree(x uint64) (n int) {
	return sovHashtree(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FileNodeProto) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHashtree
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FileNodeProto: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FileNodeProto: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Objects", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHashtree
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHashtree
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHashtree
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Objects = append(m.Objects, &pfs.Object{})
			if err := m.Objects[len(m.Objects)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHashtree(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHashtree
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHashtree
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DirectoryNodeProto) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHashtree
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DirectoryNodeProto: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DirectoryNodeProto: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Children", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHashtree
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHashtree
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHashtree
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Children = append(m.Children, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHashtree
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHashtree
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHashtree
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &pfs.Object{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Footer", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHashtree
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHashtree
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHashtree
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Footer == nil {
				m.Footer = &pfs.Object{}
			}
			if err := m.Footer.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHashtree(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHashtree
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHashtree
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *NodeProto) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHashtree
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NodeProto: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NodeProto: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHashtree
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHashtree
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHashtree
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHashtree
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthHashtree
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthHashtree
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubtreeSize", wireType)
			}
			m.SubtreeSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHashtree
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SubtreeSize |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FileNode", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHashtree
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHashtree
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHashtree
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FileNode == nil {
				m.FileNode = &FileNodeProto{}
			}
			if err := m.FileNode.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DirNode", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHashtree
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHashtree
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHashtree
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DirNode == nil {
				m.DirNode = &DirectoryNodeProto{}
			}
			if err := m.DirNode.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHashtree(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHashtree
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHashtree
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HashTreeProto) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHashtree
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HashTreeProto: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HashTreeProto: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHashtree
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHashtree
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHashtree
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHashtree
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Fs == nil {
				m.Fs = make(map[string]*NodeProto)
			}
			var mapkey string
			var mapvalue *NodeProto
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowHashtree
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowHashtree
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthHashtree
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthHashtree
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowHashtree
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthHashtree
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthHashtree
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &NodeProto{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipHashtree(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthHashtree
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Fs[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHashtree(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHashtree
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHashtree
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipHashtree(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHashtree
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowHashtree
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowHashtree
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthHashtree
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthHashtree
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHashtree
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipHashtree(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthHashtree
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthHashtree = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHashtree   = fmt.Errorf("proto: integer overflow")
)