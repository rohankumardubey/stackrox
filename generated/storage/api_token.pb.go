// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storage/api_token.proto

package storage

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

type TokenMetadata struct {
	Id                   string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" sql:"pk"`
	Name                 string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Roles                []string         `protobuf:"bytes,7,rep,name=roles,proto3" json:"roles,omitempty"`
	IssuedAt             *types.Timestamp `protobuf:"bytes,4,opt,name=issued_at,json=issuedAt,proto3" json:"issued_at,omitempty"`
	Expiration           *types.Timestamp `protobuf:"bytes,5,opt,name=expiration,proto3" json:"expiration,omitempty"`
	Revoked              bool             `protobuf:"varint,6,opt,name=revoked,proto3" json:"revoked,omitempty"`
	Role                 string           `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty"` // Deprecated: Do not use.
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TokenMetadata) Reset()         { *m = TokenMetadata{} }
func (m *TokenMetadata) String() string { return proto.CompactTextString(m) }
func (*TokenMetadata) ProtoMessage()    {}
func (*TokenMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_c11d10095315801c, []int{0}
}
func (m *TokenMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TokenMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TokenMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TokenMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenMetadata.Merge(m, src)
}
func (m *TokenMetadata) XXX_Size() int {
	return m.Size()
}
func (m *TokenMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_TokenMetadata proto.InternalMessageInfo

func (m *TokenMetadata) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TokenMetadata) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TokenMetadata) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *TokenMetadata) GetIssuedAt() *types.Timestamp {
	if m != nil {
		return m.IssuedAt
	}
	return nil
}

func (m *TokenMetadata) GetExpiration() *types.Timestamp {
	if m != nil {
		return m.Expiration
	}
	return nil
}

func (m *TokenMetadata) GetRevoked() bool {
	if m != nil {
		return m.Revoked
	}
	return false
}

// Deprecated: Do not use.
func (m *TokenMetadata) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *TokenMetadata) MessageClone() proto.Message {
	return m.Clone()
}
func (m *TokenMetadata) Clone() *TokenMetadata {
	if m == nil {
		return nil
	}
	cloned := new(TokenMetadata)
	*cloned = *m

	if m.Roles != nil {
		cloned.Roles = make([]string, len(m.Roles))
		copy(cloned.Roles, m.Roles)
	}
	cloned.IssuedAt = m.IssuedAt.Clone()
	cloned.Expiration = m.Expiration.Clone()
	return cloned
}

func init() {
	proto.RegisterType((*TokenMetadata)(nil), "storage.TokenMetadata")
}

func init() { proto.RegisterFile("storage/api_token.proto", fileDescriptor_c11d10095315801c) }

var fileDescriptor_c11d10095315801c = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0x71, 0xfa, 0x6f, 0x60, 0xb1, 0x2a, 0x30, 0x15, 0x4a, 0xa3, 0x4e, 0x99, 0x1c, 0x09,
	0x90, 0x90, 0xba, 0xd1, 0x9d, 0xc5, 0xea, 0xc4, 0x52, 0xb9, 0xc4, 0x44, 0x56, 0x7e, 0x6e, 0xb0,
	0x5d, 0xd4, 0x47, 0xe1, 0x91, 0x18, 0x79, 0x02, 0x84, 0xc2, 0x1b, 0xf0, 0x00, 0x08, 0xc5, 0x6e,
	0x24, 0x36, 0xb6, 0x73, 0xce, 0x3d, 0xb6, 0x3f, 0x5f, 0x7c, 0x6e, 0x2c, 0x68, 0x91, 0xc9, 0x44,
	0xd4, 0x6a, 0x63, 0x21, 0x97, 0x15, 0xab, 0x35, 0x58, 0x20, 0xa3, 0xc3, 0x60, 0x36, 0xcf, 0x00,
	0xb2, 0x42, 0x26, 0x2e, 0xde, 0xee, 0x9e, 0x12, 0xab, 0x4a, 0x69, 0xac, 0x28, 0x6b, 0xdf, 0x9c,
	0x4d, 0x33, 0xc8, 0xc0, 0xc9, 0xa4, 0x55, 0x3e, 0x5d, 0xfc, 0x20, 0x7c, 0xba, 0x6e, 0xef, 0xbb,
	0x97, 0x56, 0xa4, 0xc2, 0x0a, 0x72, 0x89, 0x03, 0x95, 0x52, 0x14, 0xa1, 0x78, 0xb2, 0x3a, 0xf9,
	0xfe, 0x98, 0x8f, 0xcd, 0x73, 0xb1, 0x5c, 0xd4, 0xf9, 0x82, 0x07, 0x2a, 0x25, 0x04, 0xf7, 0x2b,
	0x51, 0x4a, 0x1a, 0xb4, 0x73, 0xee, 0x34, 0x99, 0xe2, 0x81, 0x86, 0x42, 0x1a, 0x3a, 0x8a, 0x7a,
	0xf1, 0x84, 0x7b, 0x43, 0x6e, 0xf1, 0x44, 0x19, 0xb3, 0x93, 0xe9, 0x46, 0x58, 0xda, 0x8f, 0x50,
	0x7c, 0x7c, 0x35, 0x63, 0x1e, 0x92, 0x75, 0x90, 0x6c, 0xdd, 0x41, 0xf2, 0xb1, 0x2f, 0xdf, 0x59,
	0xb2, 0xc4, 0x58, 0xee, 0x6b, 0xa5, 0x85, 0x55, 0x50, 0xd1, 0xc1, 0xbf, 0x27, 0xff, 0xb4, 0x09,
	0xc5, 0x23, 0x2d, 0x5f, 0x20, 0x97, 0x29, 0x1d, 0x46, 0x28, 0x1e, 0xf3, 0xce, 0x92, 0x33, 0xdc,
	0x6f, 0xb9, 0x68, 0xcf, 0x7d, 0x2c, 0xa0, 0x88, 0x3b, 0xbf, 0xba, 0x79, 0x6b, 0x42, 0xf4, 0xde,
	0x84, 0xe8, 0xb3, 0x09, 0xd1, 0xeb, 0x57, 0x78, 0x84, 0x2f, 0x14, 0x30, 0x63, 0xc5, 0x63, 0xae,
	0x61, 0xef, 0xdf, 0x63, 0x87, 0x25, 0x3f, 0x74, 0xdb, 0xde, 0x0e, 0x5d, 0x7e, 0xfd, 0x1b, 0x00,
	0x00, 0xff, 0xff, 0xdd, 0x5a, 0xc9, 0xd0, 0x98, 0x01, 0x00, 0x00,
}

func (m *TokenMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TokenMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Roles) > 0 {
		for iNdEx := len(m.Roles) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Roles[iNdEx])
			copy(dAtA[i:], m.Roles[iNdEx])
			i = encodeVarintApiToken(dAtA, i, uint64(len(m.Roles[iNdEx])))
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.Revoked {
		i--
		if m.Revoked {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if m.Expiration != nil {
		{
			size, err := m.Expiration.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintApiToken(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.IssuedAt != nil {
		{
			size, err := m.IssuedAt.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintApiToken(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Role) > 0 {
		i -= len(m.Role)
		copy(dAtA[i:], m.Role)
		i = encodeVarintApiToken(dAtA, i, uint64(len(m.Role)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintApiToken(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintApiToken(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintApiToken(dAtA []byte, offset int, v uint64) int {
	offset -= sovApiToken(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TokenMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovApiToken(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovApiToken(uint64(l))
	}
	l = len(m.Role)
	if l > 0 {
		n += 1 + l + sovApiToken(uint64(l))
	}
	if m.IssuedAt != nil {
		l = m.IssuedAt.Size()
		n += 1 + l + sovApiToken(uint64(l))
	}
	if m.Expiration != nil {
		l = m.Expiration.Size()
		n += 1 + l + sovApiToken(uint64(l))
	}
	if m.Revoked {
		n += 2
	}
	if len(m.Roles) > 0 {
		for _, s := range m.Roles {
			l = len(s)
			n += 1 + l + sovApiToken(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovApiToken(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozApiToken(x uint64) (n int) {
	return sovApiToken(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TokenMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApiToken
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
			return fmt.Errorf("proto: TokenMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiToken
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
				return ErrInvalidLengthApiToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApiToken
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
					return ErrIntOverflowApiToken
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
				return ErrInvalidLengthApiToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApiToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiToken
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
				return ErrInvalidLengthApiToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApiToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Role = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IssuedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiToken
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
				return ErrInvalidLengthApiToken
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthApiToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.IssuedAt == nil {
				m.IssuedAt = &types.Timestamp{}
			}
			if err := m.IssuedAt.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiToken
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
				return ErrInvalidLengthApiToken
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthApiToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Expiration == nil {
				m.Expiration = &types.Timestamp{}
			}
			if err := m.Expiration.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Revoked", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiToken
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Revoked = bool(v != 0)
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Roles", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiToken
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
				return ErrInvalidLengthApiToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApiToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Roles = append(m.Roles, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApiToken(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthApiToken
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
func skipApiToken(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApiToken
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
					return 0, ErrIntOverflowApiToken
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
					return 0, ErrIntOverflowApiToken
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
				return 0, ErrInvalidLengthApiToken
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupApiToken
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthApiToken
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthApiToken        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApiToken          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupApiToken = fmt.Errorf("proto: unexpected end of group")
)
