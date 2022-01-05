// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: internalapi/central/local_scanner.proto

package central

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/stackrox/rox/generated/storage"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type LocalScannerCertificates struct {
	Ca                   []byte   `protobuf:"bytes,1,opt,name=ca,proto3" json:"ca,omitempty"`
	Cert                 []byte   `protobuf:"bytes,2,opt,name=cert,proto3" json:"cert,omitempty"`
	Key                  []byte   `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LocalScannerCertificates) Reset()         { *m = LocalScannerCertificates{} }
func (m *LocalScannerCertificates) String() string { return proto.CompactTextString(m) }
func (*LocalScannerCertificates) ProtoMessage()    {}
func (*LocalScannerCertificates) Descriptor() ([]byte, []int) {
	return fileDescriptor_856923c76f63cf0a, []int{0}
}
func (m *LocalScannerCertificates) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LocalScannerCertificates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LocalScannerCertificates.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LocalScannerCertificates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalScannerCertificates.Merge(m, src)
}
func (m *LocalScannerCertificates) XXX_Size() int {
	return m.Size()
}
func (m *LocalScannerCertificates) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalScannerCertificates.DiscardUnknown(m)
}

var xxx_messageInfo_LocalScannerCertificates proto.InternalMessageInfo

func (m *LocalScannerCertificates) GetCa() []byte {
	if m != nil {
		return m.Ca
	}
	return nil
}

func (m *LocalScannerCertificates) GetCert() []byte {
	if m != nil {
		return m.Cert
	}
	return nil
}

func (m *LocalScannerCertificates) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *LocalScannerCertificates) MessageClone() proto.Message {
	return m.Clone()
}
func (m *LocalScannerCertificates) Clone() *LocalScannerCertificates {
	if m == nil {
		return nil
	}
	cloned := new(LocalScannerCertificates)
	*cloned = *m

	if m.Ca != nil {
		cloned.Ca = make([]byte, len(m.Ca))
		copy(cloned.Ca, m.Ca)
	}
	if m.Cert != nil {
		cloned.Cert = make([]byte, len(m.Cert))
		copy(cloned.Cert, m.Cert)
	}
	if m.Key != nil {
		cloned.Key = make([]byte, len(m.Key))
		copy(cloned.Key, m.Key)
	}
	return cloned
}

type IssueLocalScannerCertsRequest struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IssueLocalScannerCertsRequest) Reset()         { *m = IssueLocalScannerCertsRequest{} }
func (m *IssueLocalScannerCertsRequest) String() string { return proto.CompactTextString(m) }
func (*IssueLocalScannerCertsRequest) ProtoMessage()    {}
func (*IssueLocalScannerCertsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_856923c76f63cf0a, []int{1}
}
func (m *IssueLocalScannerCertsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IssueLocalScannerCertsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IssueLocalScannerCertsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IssueLocalScannerCertsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IssueLocalScannerCertsRequest.Merge(m, src)
}
func (m *IssueLocalScannerCertsRequest) XXX_Size() int {
	return m.Size()
}
func (m *IssueLocalScannerCertsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IssueLocalScannerCertsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IssueLocalScannerCertsRequest proto.InternalMessageInfo

func (m *IssueLocalScannerCertsRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *IssueLocalScannerCertsRequest) MessageClone() proto.Message {
	return m.Clone()
}
func (m *IssueLocalScannerCertsRequest) Clone() *IssueLocalScannerCertsRequest {
	if m == nil {
		return nil
	}
	cloned := new(IssueLocalScannerCertsRequest)
	*cloned = *m

	return cloned
}

type IssueLocalScannerCertsResponse struct {
	ScannerCerts         *LocalScannerCertificates `protobuf:"bytes,1,opt,name=scanner_certs,json=scannerCerts,proto3" json:"scanner_certs,omitempty"`
	ScannerDbCerts       *LocalScannerCertificates `protobuf:"bytes,2,opt,name=scanner_db_certs,json=scannerDbCerts,proto3" json:"scanner_db_certs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *IssueLocalScannerCertsResponse) Reset()         { *m = IssueLocalScannerCertsResponse{} }
func (m *IssueLocalScannerCertsResponse) String() string { return proto.CompactTextString(m) }
func (*IssueLocalScannerCertsResponse) ProtoMessage()    {}
func (*IssueLocalScannerCertsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_856923c76f63cf0a, []int{2}
}
func (m *IssueLocalScannerCertsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IssueLocalScannerCertsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IssueLocalScannerCertsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IssueLocalScannerCertsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IssueLocalScannerCertsResponse.Merge(m, src)
}
func (m *IssueLocalScannerCertsResponse) XXX_Size() int {
	return m.Size()
}
func (m *IssueLocalScannerCertsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IssueLocalScannerCertsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IssueLocalScannerCertsResponse proto.InternalMessageInfo

func (m *IssueLocalScannerCertsResponse) GetScannerCerts() *LocalScannerCertificates {
	if m != nil {
		return m.ScannerCerts
	}
	return nil
}

func (m *IssueLocalScannerCertsResponse) GetScannerDbCerts() *LocalScannerCertificates {
	if m != nil {
		return m.ScannerDbCerts
	}
	return nil
}

func (m *IssueLocalScannerCertsResponse) MessageClone() proto.Message {
	return m.Clone()
}
func (m *IssueLocalScannerCertsResponse) Clone() *IssueLocalScannerCertsResponse {
	if m == nil {
		return nil
	}
	cloned := new(IssueLocalScannerCertsResponse)
	*cloned = *m

	cloned.ScannerCerts = m.ScannerCerts.Clone()
	cloned.ScannerDbCerts = m.ScannerDbCerts.Clone()
	return cloned
}

func init() {
	proto.RegisterType((*LocalScannerCertificates)(nil), "central.LocalScannerCertificates")
	proto.RegisterType((*IssueLocalScannerCertsRequest)(nil), "central.IssueLocalScannerCertsRequest")
	proto.RegisterType((*IssueLocalScannerCertsResponse)(nil), "central.IssueLocalScannerCertsResponse")
}

func init() {
	proto.RegisterFile("internalapi/central/local_scanner.proto", fileDescriptor_856923c76f63cf0a)
}

var fileDescriptor_856923c76f63cf0a = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x4a, 0xf3, 0x40,
	0x14, 0x85, 0xff, 0x49, 0x7f, 0x94, 0x5e, 0x6b, 0x29, 0x23, 0x48, 0x2c, 0x3a, 0x68, 0x17, 0xd6,
	0x55, 0x0b, 0x75, 0xed, 0x46, 0x45, 0x10, 0x5d, 0x48, 0xba, 0x73, 0x53, 0x26, 0xd3, 0xab, 0x0c,
	0xc6, 0x49, 0x9c, 0x3b, 0x15, 0xba, 0xf3, 0x31, 0x7c, 0x11, 0xdf, 0xc1, 0xa5, 0x8f, 0x20, 0xf1,
	0x45, 0x24, 0x93, 0x54, 0x8b, 0x58, 0x75, 0x77, 0xb9, 0x39, 0xe7, 0xcb, 0x9c, 0x33, 0x03, 0x5d,
	0x6d, 0x1c, 0x5a, 0x23, 0x13, 0x99, 0xe9, 0xbe, 0x42, 0xe3, 0xac, 0x4c, 0xfa, 0x49, 0xaa, 0x64,
	0x32, 0x22, 0x25, 0x8d, 0x41, 0xdb, 0xcb, 0x6c, 0xea, 0x52, 0xbe, 0x5c, 0x7d, 0x6c, 0x0b, 0x72,
	0xa9, 0x95, 0xd7, 0xd8, 0x27, 0xb4, 0xf7, 0x5a, 0xe1, 0x48, 0x8f, 0xd1, 0x38, 0xed, 0xa6, 0xa5,
	0xb0, 0x73, 0x01, 0xe1, 0x79, 0xe1, 0x1f, 0x96, 0xf6, 0x23, 0xb4, 0x4e, 0x5f, 0x69, 0x25, 0x1d,
	0x12, 0x6f, 0x42, 0xa0, 0x64, 0xc8, 0xb6, 0xd9, 0x5e, 0x23, 0x0a, 0x94, 0xe4, 0x1c, 0xfe, 0x2b,
	0xb4, 0x2e, 0x0c, 0xfc, 0xc6, 0xcf, 0xbc, 0x05, 0xb5, 0x1b, 0x9c, 0x86, 0x35, 0xbf, 0x2a, 0xc6,
	0xce, 0x01, 0x6c, 0x9d, 0x12, 0x4d, 0xf0, 0x2b, 0x96, 0x22, 0xbc, 0x9b, 0x20, 0x39, 0xbe, 0x09,
	0x75, 0x23, 0x6f, 0x91, 0x32, 0xa9, 0xd0, 0xd3, 0xeb, 0xd1, 0xe7, 0xa2, 0xf3, 0xc4, 0x40, 0x2c,
	0xf2, 0x53, 0x96, 0x1a, 0x42, 0x7e, 0x02, 0xab, 0x55, 0xda, 0x51, 0x71, 0x06, 0xf2, 0x90, 0x95,
	0xc1, 0x4e, 0xaf, 0x0a, 0xdd, 0x5b, 0x94, 0x28, 0x6a, 0xd0, 0x1c, 0x8f, 0x9f, 0x41, 0x6b, 0xc6,
	0x19, 0xc7, 0x15, 0x2a, 0xf8, 0x2b, 0xaa, 0x59, 0x59, 0x8f, 0x63, 0x0f, 0x1b, 0x3c, 0x30, 0x58,
	0x9b, 0x17, 0x0f, 0xcb, 0xbe, 0xb9, 0x86, 0xf5, 0xef, 0xe3, 0xf0, 0xdd, 0x8f, 0x9f, 0xfc, 0xd8,
	0x57, 0xbb, 0xfb, 0xab, 0xae, 0xec, 0xe5, 0x70, 0xe3, 0x39, 0x17, 0xec, 0x25, 0x17, 0xec, 0x35,
	0x17, 0xec, 0xf1, 0x4d, 0xfc, 0xbb, 0x9c, 0x3d, 0x83, 0x78, 0xc9, 0xdf, 0xf6, 0xfe, 0x7b, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x16, 0xc6, 0x41, 0xdf, 0x41, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LocalScannerServiceClient is the client API for LocalScannerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConnInterface.NewStream.
type LocalScannerServiceClient interface {
	IssueLocalScannerCerts(ctx context.Context, in *IssueLocalScannerCertsRequest, opts ...grpc.CallOption) (*IssueLocalScannerCertsResponse, error)
}

type localScannerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLocalScannerServiceClient(cc grpc.ClientConnInterface) LocalScannerServiceClient {
	return &localScannerServiceClient{cc}
}

func (c *localScannerServiceClient) IssueLocalScannerCerts(ctx context.Context, in *IssueLocalScannerCertsRequest, opts ...grpc.CallOption) (*IssueLocalScannerCertsResponse, error) {
	out := new(IssueLocalScannerCertsResponse)
	err := c.cc.Invoke(ctx, "/central.LocalScannerService/IssueLocalScannerCerts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LocalScannerServiceServer is the server API for LocalScannerService service.
type LocalScannerServiceServer interface {
	IssueLocalScannerCerts(context.Context, *IssueLocalScannerCertsRequest) (*IssueLocalScannerCertsResponse, error)
}

// UnimplementedLocalScannerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLocalScannerServiceServer struct {
}

func (*UnimplementedLocalScannerServiceServer) IssueLocalScannerCerts(ctx context.Context, req *IssueLocalScannerCertsRequest) (*IssueLocalScannerCertsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IssueLocalScannerCerts not implemented")
}

func RegisterLocalScannerServiceServer(s *grpc.Server, srv LocalScannerServiceServer) {
	s.RegisterService(&_LocalScannerService_serviceDesc, srv)
}

func _LocalScannerService_IssueLocalScannerCerts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IssueLocalScannerCertsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocalScannerServiceServer).IssueLocalScannerCerts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/central.LocalScannerService/IssueLocalScannerCerts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocalScannerServiceServer).IssueLocalScannerCerts(ctx, req.(*IssueLocalScannerCertsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LocalScannerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "central.LocalScannerService",
	HandlerType: (*LocalScannerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IssueLocalScannerCerts",
			Handler:    _LocalScannerService_IssueLocalScannerCerts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internalapi/central/local_scanner.proto",
}

func (m *LocalScannerCertificates) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LocalScannerCertificates) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LocalScannerCertificates) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintLocalScanner(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Cert) > 0 {
		i -= len(m.Cert)
		copy(dAtA[i:], m.Cert)
		i = encodeVarintLocalScanner(dAtA, i, uint64(len(m.Cert)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Ca) > 0 {
		i -= len(m.Ca)
		copy(dAtA[i:], m.Ca)
		i = encodeVarintLocalScanner(dAtA, i, uint64(len(m.Ca)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *IssueLocalScannerCertsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IssueLocalScannerCertsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IssueLocalScannerCertsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Namespace) > 0 {
		i -= len(m.Namespace)
		copy(dAtA[i:], m.Namespace)
		i = encodeVarintLocalScanner(dAtA, i, uint64(len(m.Namespace)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *IssueLocalScannerCertsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IssueLocalScannerCertsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IssueLocalScannerCertsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.ScannerDbCerts != nil {
		{
			size, err := m.ScannerDbCerts.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintLocalScanner(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.ScannerCerts != nil {
		{
			size, err := m.ScannerCerts.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintLocalScanner(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintLocalScanner(dAtA []byte, offset int, v uint64) int {
	offset -= sovLocalScanner(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LocalScannerCertificates) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Ca)
	if l > 0 {
		n += 1 + l + sovLocalScanner(uint64(l))
	}
	l = len(m.Cert)
	if l > 0 {
		n += 1 + l + sovLocalScanner(uint64(l))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovLocalScanner(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *IssueLocalScannerCertsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovLocalScanner(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *IssueLocalScannerCertsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ScannerCerts != nil {
		l = m.ScannerCerts.Size()
		n += 1 + l + sovLocalScanner(uint64(l))
	}
	if m.ScannerDbCerts != nil {
		l = m.ScannerDbCerts.Size()
		n += 1 + l + sovLocalScanner(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovLocalScanner(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLocalScanner(x uint64) (n int) {
	return sovLocalScanner(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LocalScannerCertificates) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLocalScanner
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
			return fmt.Errorf("proto: LocalScannerCertificates: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LocalScannerCertificates: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ca", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLocalScanner
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
				return ErrInvalidLengthLocalScanner
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthLocalScanner
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ca = append(m.Ca[:0], dAtA[iNdEx:postIndex]...)
			if m.Ca == nil {
				m.Ca = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cert", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLocalScanner
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
				return ErrInvalidLengthLocalScanner
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthLocalScanner
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cert = append(m.Cert[:0], dAtA[iNdEx:postIndex]...)
			if m.Cert == nil {
				m.Cert = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLocalScanner
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
				return ErrInvalidLengthLocalScanner
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthLocalScanner
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLocalScanner(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLocalScanner
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
func (m *IssueLocalScannerCertsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLocalScanner
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
			return fmt.Errorf("proto: IssueLocalScannerCertsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IssueLocalScannerCertsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLocalScanner
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
				return ErrInvalidLengthLocalScanner
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLocalScanner
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLocalScanner(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLocalScanner
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
func (m *IssueLocalScannerCertsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLocalScanner
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
			return fmt.Errorf("proto: IssueLocalScannerCertsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IssueLocalScannerCertsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScannerCerts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLocalScanner
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
				return ErrInvalidLengthLocalScanner
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLocalScanner
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ScannerCerts == nil {
				m.ScannerCerts = &LocalScannerCertificates{}
			}
			if err := m.ScannerCerts.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScannerDbCerts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLocalScanner
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
				return ErrInvalidLengthLocalScanner
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLocalScanner
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ScannerDbCerts == nil {
				m.ScannerDbCerts = &LocalScannerCertificates{}
			}
			if err := m.ScannerDbCerts.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLocalScanner(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLocalScanner
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
func skipLocalScanner(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLocalScanner
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
					return 0, ErrIntOverflowLocalScanner
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
					return 0, ErrIntOverflowLocalScanner
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
				return 0, ErrInvalidLengthLocalScanner
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLocalScanner
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLocalScanner
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLocalScanner        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLocalScanner          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLocalScanner = fmt.Errorf("proto: unexpected end of group")
)
