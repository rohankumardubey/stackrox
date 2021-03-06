// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storage/comments.proto

package storage

import (
	fmt "fmt"
	types "github.com/gogo/protobuf/types"
	proto "github.com/golang/protobuf/proto"
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

type ResourceType int32

const (
	ResourceType_UNSET_RESOURCE_TYPE ResourceType = 0
	ResourceType_ALERT               ResourceType = 1
	ResourceType_PROCESS             ResourceType = 2
)

var ResourceType_name = map[int32]string{
	0: "UNSET_RESOURCE_TYPE",
	1: "ALERT",
	2: "PROCESS",
}

var ResourceType_value = map[string]int32{
	"UNSET_RESOURCE_TYPE": 0,
	"ALERT":               1,
	"PROCESS":             2,
}

func (x ResourceType) String() string {
	return proto.EnumName(ResourceType_name, int32(x))
}

func (ResourceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3273dfba75e9f263, []int{0}
}

type Comment struct {
	ResourceType         ResourceType     `protobuf:"varint,1,opt,name=resource_type,json=resourceType,proto3,enum=storage.ResourceType" json:"resource_type,omitempty"`
	ResourceId           string           `protobuf:"bytes,2,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	CommentId            string           `protobuf:"bytes,3,opt,name=comment_id,json=commentId,proto3" json:"comment_id,omitempty"`
	CommentMessage       string           `protobuf:"bytes,4,opt,name=comment_message,json=commentMessage,proto3" json:"comment_message,omitempty"`
	User                 *Comment_User    `protobuf:"bytes,5,opt,name=user,proto3" json:"user,omitempty"`
	CreatedAt            *types.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	LastModified         *types.Timestamp `protobuf:"bytes,7,opt,name=last_modified,json=lastModified,proto3" json:"last_modified,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Comment) Reset()         { *m = Comment{} }
func (m *Comment) String() string { return proto.CompactTextString(m) }
func (*Comment) ProtoMessage()    {}
func (*Comment) Descriptor() ([]byte, []int) {
	return fileDescriptor_3273dfba75e9f263, []int{0}
}
func (m *Comment) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Comment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Comment.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Comment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Comment.Merge(m, src)
}
func (m *Comment) XXX_Size() int {
	return m.Size()
}
func (m *Comment) XXX_DiscardUnknown() {
	xxx_messageInfo_Comment.DiscardUnknown(m)
}

var xxx_messageInfo_Comment proto.InternalMessageInfo

func (m *Comment) GetResourceType() ResourceType {
	if m != nil {
		return m.ResourceType
	}
	return ResourceType_UNSET_RESOURCE_TYPE
}

func (m *Comment) GetResourceId() string {
	if m != nil {
		return m.ResourceId
	}
	return ""
}

func (m *Comment) GetCommentId() string {
	if m != nil {
		return m.CommentId
	}
	return ""
}

func (m *Comment) GetCommentMessage() string {
	if m != nil {
		return m.CommentMessage
	}
	return ""
}

func (m *Comment) GetUser() *Comment_User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Comment) GetCreatedAt() *types.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Comment) GetLastModified() *types.Timestamp {
	if m != nil {
		return m.LastModified
	}
	return nil
}

func (m *Comment) MessageClone() proto.Message {
	return m.Clone()
}
func (m *Comment) Clone() *Comment {
	if m == nil {
		return nil
	}
	cloned := new(Comment)
	*cloned = *m

	cloned.User = m.User.Clone()
	cloned.CreatedAt = m.CreatedAt.Clone()
	cloned.LastModified = m.LastModified.Clone()
	return cloned
}

type Comment_User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Comment_User) Reset()         { *m = Comment_User{} }
func (m *Comment_User) String() string { return proto.CompactTextString(m) }
func (*Comment_User) ProtoMessage()    {}
func (*Comment_User) Descriptor() ([]byte, []int) {
	return fileDescriptor_3273dfba75e9f263, []int{0, 0}
}
func (m *Comment_User) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Comment_User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Comment_User.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Comment_User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Comment_User.Merge(m, src)
}
func (m *Comment_User) XXX_Size() int {
	return m.Size()
}
func (m *Comment_User) XXX_DiscardUnknown() {
	xxx_messageInfo_Comment_User.DiscardUnknown(m)
}

var xxx_messageInfo_Comment_User proto.InternalMessageInfo

func (m *Comment_User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Comment_User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Comment_User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Comment_User) MessageClone() proto.Message {
	return m.Clone()
}
func (m *Comment_User) Clone() *Comment_User {
	if m == nil {
		return nil
	}
	cloned := new(Comment_User)
	*cloned = *m

	return cloned
}

func init() {
	proto.RegisterEnum("storage.ResourceType", ResourceType_name, ResourceType_value)
	proto.RegisterType((*Comment)(nil), "storage.Comment")
	proto.RegisterType((*Comment_User)(nil), "storage.Comment.User")
}

func init() { proto.RegisterFile("storage/comments.proto", fileDescriptor_3273dfba75e9f263) }

var fileDescriptor_3273dfba75e9f263 = []byte{
	// 400 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0xc7, 0xeb, 0x34, 0x69, 0xb4, 0x93, 0x34, 0x44, 0xe6, 0x6b, 0x89, 0x44, 0x1a, 0x71, 0x21,
	0x70, 0x70, 0x24, 0xe0, 0x02, 0x97, 0x52, 0xa2, 0x3d, 0x54, 0xa2, 0xb4, 0x72, 0x36, 0x07, 0xb8,
	0xac, 0xdc, 0x78, 0xba, 0xb2, 0x88, 0xeb, 0x95, 0xed, 0x48, 0xf4, 0x49, 0xe0, 0x91, 0x38, 0xf2,
	0x08, 0x28, 0xbc, 0x08, 0x8a, 0xd7, 0x8b, 0x72, 0xe3, 0xb6, 0xf3, 0xff, 0x58, 0xcd, 0xfc, 0x64,
	0x78, 0xe4, 0xbc, 0xb1, 0xa2, 0xc4, 0xd9, 0xca, 0x68, 0x8d, 0xb7, 0xde, 0xb1, 0xca, 0x1a, 0x6f,
	0x68, 0x37, 0xea, 0xa3, 0x93, 0xd2, 0x98, 0x72, 0x8d, 0xb3, 0x20, 0x5f, 0x6f, 0x6e, 0x66, 0x5e,
	0x69, 0x74, 0x5e, 0xe8, 0xaa, 0x4e, 0x3e, 0xfb, 0x7e, 0x08, 0xdd, 0x79, 0x5d, 0xa6, 0xef, 0xe0,
	0xd8, 0xa2, 0x33, 0x1b, 0xbb, 0xc2, 0xc2, 0xdf, 0x55, 0x98, 0x92, 0x09, 0x99, 0x0e, 0x5e, 0x3d,
	0x64, 0xf1, 0x6f, 0x8c, 0x47, 0x37, 0xbf, 0xab, 0x90, 0xf7, 0xed, 0xde, 0x44, 0x4f, 0xa0, 0xf7,
	0xaf, 0xab, 0x64, 0xda, 0x9a, 0x90, 0x69, 0xc2, 0xa1, 0x91, 0xce, 0x25, 0x7d, 0x0a, 0x10, 0x97,
	0xdc, 0xf9, 0x87, 0xc1, 0x4f, 0xa2, 0x72, 0x2e, 0xe9, 0x73, 0xb8, 0xd7, 0xd8, 0x1a, 0x9d, 0x13,
	0x25, 0xa6, 0xed, 0x90, 0x19, 0x44, 0xf9, 0xa2, 0x56, 0xe9, 0x0b, 0x68, 0x6f, 0x1c, 0xda, 0xb4,
	0x33, 0x21, 0xd3, 0xde, 0xde, 0x6e, 0xf1, 0x08, 0xb6, 0x74, 0x68, 0x79, 0x88, 0xd0, 0xb7, 0x00,
	0x2b, 0x8b, 0xc2, 0xa3, 0x2c, 0x84, 0x4f, 0x8f, 0x42, 0x61, 0xc4, 0x6a, 0x22, 0xac, 0x21, 0xc2,
	0xf2, 0x86, 0x08, 0x4f, 0x62, 0xfa, 0xcc, 0xd3, 0x53, 0x38, 0x5e, 0x0b, 0xe7, 0x0b, 0x6d, 0xa4,
	0xba, 0x51, 0x28, 0xd3, 0xee, 0x7f, 0xdb, 0xfd, 0x5d, 0xe1, 0x22, 0xe6, 0x47, 0xef, 0xa1, 0xbd,
	0xdb, 0x84, 0x0e, 0xa0, 0xa5, 0x64, 0x00, 0x99, 0xf0, 0x96, 0x92, 0x94, 0x42, 0xfb, 0x56, 0x68,
	0x8c, 0x80, 0xc2, 0x37, 0x7d, 0x00, 0x1d, 0xd4, 0x42, 0xad, 0x23, 0x95, 0x7a, 0x78, 0x79, 0x0a,
	0xfd, 0x7d, 0xde, 0xf4, 0x31, 0xdc, 0x5f, 0x7e, 0x5a, 0x64, 0x79, 0xc1, 0xb3, 0xc5, 0xe5, 0x92,
	0xcf, 0xb3, 0x22, 0xff, 0x7c, 0x95, 0x0d, 0x0f, 0x68, 0x02, 0x9d, 0xb3, 0x8f, 0x19, 0xcf, 0x87,
	0x84, 0xf6, 0xa0, 0x7b, 0xc5, 0x2f, 0xe7, 0xd9, 0x62, 0x31, 0x6c, 0x7d, 0x78, 0xf3, 0x73, 0x3b,
	0x26, 0xbf, 0xb6, 0x63, 0xf2, 0x7b, 0x3b, 0x26, 0x3f, 0xfe, 0x8c, 0x0f, 0xe0, 0x89, 0x32, 0xcc,
	0x79, 0xb1, 0xfa, 0x6a, 0xcd, 0xb7, 0xfa, 0x84, 0x06, 0xdf, 0x97, 0xe6, 0xc5, 0x5c, 0x1f, 0x05,
	0xfd, 0xf5, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x47, 0x59, 0x9b, 0x9f, 0x5b, 0x02, 0x00, 0x00,
}

func (m *Comment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Comment) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Comment) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.LastModified != nil {
		{
			size, err := m.LastModified.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintComments(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.CreatedAt != nil {
		{
			size, err := m.CreatedAt.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintComments(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if m.User != nil {
		{
			size, err := m.User.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintComments(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.CommentMessage) > 0 {
		i -= len(m.CommentMessage)
		copy(dAtA[i:], m.CommentMessage)
		i = encodeVarintComments(dAtA, i, uint64(len(m.CommentMessage)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.CommentId) > 0 {
		i -= len(m.CommentId)
		copy(dAtA[i:], m.CommentId)
		i = encodeVarintComments(dAtA, i, uint64(len(m.CommentId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ResourceId) > 0 {
		i -= len(m.ResourceId)
		copy(dAtA[i:], m.ResourceId)
		i = encodeVarintComments(dAtA, i, uint64(len(m.ResourceId)))
		i--
		dAtA[i] = 0x12
	}
	if m.ResourceType != 0 {
		i = encodeVarintComments(dAtA, i, uint64(m.ResourceType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Comment_User) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Comment_User) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Comment_User) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Email) > 0 {
		i -= len(m.Email)
		copy(dAtA[i:], m.Email)
		i = encodeVarintComments(dAtA, i, uint64(len(m.Email)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintComments(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintComments(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintComments(dAtA []byte, offset int, v uint64) int {
	offset -= sovComments(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Comment) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ResourceType != 0 {
		n += 1 + sovComments(uint64(m.ResourceType))
	}
	l = len(m.ResourceId)
	if l > 0 {
		n += 1 + l + sovComments(uint64(l))
	}
	l = len(m.CommentId)
	if l > 0 {
		n += 1 + l + sovComments(uint64(l))
	}
	l = len(m.CommentMessage)
	if l > 0 {
		n += 1 + l + sovComments(uint64(l))
	}
	if m.User != nil {
		l = m.User.Size()
		n += 1 + l + sovComments(uint64(l))
	}
	if m.CreatedAt != nil {
		l = m.CreatedAt.Size()
		n += 1 + l + sovComments(uint64(l))
	}
	if m.LastModified != nil {
		l = m.LastModified.Size()
		n += 1 + l + sovComments(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Comment_User) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovComments(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovComments(uint64(l))
	}
	l = len(m.Email)
	if l > 0 {
		n += 1 + l + sovComments(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovComments(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozComments(x uint64) (n int) {
	return sovComments(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Comment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowComments
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
			return fmt.Errorf("proto: Comment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Comment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResourceType", wireType)
			}
			m.ResourceType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComments
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ResourceType |= ResourceType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResourceId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComments
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
				return ErrInvalidLengthComments
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComments
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResourceId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommentId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComments
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
				return ErrInvalidLengthComments
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComments
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CommentId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommentMessage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComments
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
				return ErrInvalidLengthComments
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComments
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CommentMessage = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComments
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
				return ErrInvalidLengthComments
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthComments
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.User == nil {
				m.User = &Comment_User{}
			}
			if err := m.User.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComments
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
				return ErrInvalidLengthComments
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthComments
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CreatedAt == nil {
				m.CreatedAt = &types.Timestamp{}
			}
			if err := m.CreatedAt.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastModified", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComments
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
				return ErrInvalidLengthComments
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthComments
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LastModified == nil {
				m.LastModified = &types.Timestamp{}
			}
			if err := m.LastModified.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipComments(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthComments
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
func (m *Comment_User) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowComments
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
			return fmt.Errorf("proto: User: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: User: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComments
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
				return ErrInvalidLengthComments
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComments
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComments
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
				return ErrInvalidLengthComments
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComments
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Email", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComments
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
				return ErrInvalidLengthComments
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComments
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Email = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipComments(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthComments
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
func skipComments(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowComments
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
					return 0, ErrIntOverflowComments
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowComments
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
				return 0, ErrInvalidLengthComments
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupComments
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthComments
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthComments        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowComments          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupComments = fmt.Errorf("proto: unexpected end of group")
)
