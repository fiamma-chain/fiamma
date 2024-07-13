// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fiamma/bitvmstaker/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgUpdateParams is the Msg/UpdateParams request type.
type MsgUpdateParams struct {
	// authority is the address that controls the module (defaults to x/gov unless overwritten).
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// NOTE: All parameters must be supplied.
	Params Params `protobuf:"bytes,2,opt,name=params,proto3" json:"params"`
}

func (m *MsgUpdateParams) Reset()         { *m = MsgUpdateParams{} }
func (m *MsgUpdateParams) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateParams) ProtoMessage()    {}
func (*MsgUpdateParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca62a47b6be03f96, []int{0}
}
func (m *MsgUpdateParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateParams.Merge(m, src)
}
func (m *MsgUpdateParams) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateParams) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateParams.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateParams proto.InternalMessageInfo

func (m *MsgUpdateParams) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgUpdateParams) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// MsgUpdateParamsResponse defines the response structure for executing a
// MsgUpdateParams message.
type MsgUpdateParamsResponse struct {
}

func (m *MsgUpdateParamsResponse) Reset()         { *m = MsgUpdateParamsResponse{} }
func (m *MsgUpdateParamsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateParamsResponse) ProtoMessage()    {}
func (*MsgUpdateParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca62a47b6be03f96, []int{1}
}
func (m *MsgUpdateParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateParamsResponse.Merge(m, src)
}
func (m *MsgUpdateParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateParamsResponse proto.InternalMessageInfo

type MsgCreateStaker struct {
	Creator       string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	StakerAddress string `protobuf:"bytes,2,opt,name=stakerAddress,proto3" json:"stakerAddress,omitempty"`
}

func (m *MsgCreateStaker) Reset()         { *m = MsgCreateStaker{} }
func (m *MsgCreateStaker) String() string { return proto.CompactTextString(m) }
func (*MsgCreateStaker) ProtoMessage()    {}
func (*MsgCreateStaker) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca62a47b6be03f96, []int{2}
}
func (m *MsgCreateStaker) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateStaker) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateStaker.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateStaker) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateStaker.Merge(m, src)
}
func (m *MsgCreateStaker) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateStaker) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateStaker.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateStaker proto.InternalMessageInfo

func (m *MsgCreateStaker) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreateStaker) GetStakerAddress() string {
	if m != nil {
		return m.StakerAddress
	}
	return ""
}

type MsgCreateStakerResponse struct {
}

func (m *MsgCreateStakerResponse) Reset()         { *m = MsgCreateStakerResponse{} }
func (m *MsgCreateStakerResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateStakerResponse) ProtoMessage()    {}
func (*MsgCreateStakerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca62a47b6be03f96, []int{3}
}
func (m *MsgCreateStakerResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateStakerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateStakerResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateStakerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateStakerResponse.Merge(m, src)
}
func (m *MsgCreateStakerResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateStakerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateStakerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateStakerResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgUpdateParams)(nil), "fiamma.bitvmstaker.MsgUpdateParams")
	proto.RegisterType((*MsgUpdateParamsResponse)(nil), "fiamma.bitvmstaker.MsgUpdateParamsResponse")
	proto.RegisterType((*MsgCreateStaker)(nil), "fiamma.bitvmstaker.MsgCreateStaker")
	proto.RegisterType((*MsgCreateStakerResponse)(nil), "fiamma.bitvmstaker.MsgCreateStakerResponse")
}

func init() { proto.RegisterFile("fiamma/bitvmstaker/tx.proto", fileDescriptor_ca62a47b6be03f96) }

var fileDescriptor_ca62a47b6be03f96 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4e, 0xcb, 0x4c, 0xcc,
	0xcd, 0x4d, 0xd4, 0x4f, 0xca, 0x2c, 0x29, 0xcb, 0x2d, 0x2e, 0x49, 0xcc, 0x4e, 0x2d, 0xd2, 0x2f,
	0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x82, 0x48, 0xea, 0x21, 0x49, 0x4a, 0x09,
	0x26, 0xe6, 0x66, 0xe6, 0xe5, 0xeb, 0x83, 0x49, 0x88, 0x32, 0x29, 0xf1, 0xe4, 0xfc, 0xe2, 0xdc,
	0xfc, 0x62, 0xfd, 0xdc, 0xe2, 0x74, 0xfd, 0x32, 0x43, 0x10, 0x05, 0x95, 0x90, 0x84, 0x48, 0xc4,
	0x83, 0x79, 0xfa, 0x10, 0x0e, 0x54, 0x4a, 0x24, 0x3d, 0x3f, 0x3d, 0x1f, 0x22, 0x0e, 0x62, 0x41,
	0x45, 0xe5, 0xb1, 0xb8, 0xa6, 0x20, 0xb1, 0x28, 0x31, 0x17, 0xaa, 0x4d, 0xe9, 0x20, 0x23, 0x17,
	0xbf, 0x6f, 0x71, 0x7a, 0x68, 0x41, 0x4a, 0x62, 0x49, 0x6a, 0x00, 0x58, 0x46, 0xc8, 0x8c, 0x8b,
	0x33, 0xb1, 0xb4, 0x24, 0x23, 0xbf, 0x28, 0xb3, 0xa4, 0x52, 0x82, 0x51, 0x81, 0x51, 0x83, 0xd3,
	0x49, 0xe2, 0xd2, 0x16, 0x5d, 0x11, 0xa8, 0x7d, 0x8e, 0x29, 0x29, 0x45, 0xa9, 0xc5, 0xc5, 0xc1,
	0x25, 0x45, 0x99, 0x79, 0xe9, 0x41, 0x08, 0xa5, 0x42, 0xb6, 0x5c, 0x6c, 0x10, 0xb3, 0x25, 0x98,
	0x14, 0x18, 0x35, 0xb8, 0x8d, 0xa4, 0xf4, 0x30, 0xbd, 0xab, 0x07, 0xb1, 0xc3, 0x89, 0xf3, 0xc4,
	0x3d, 0x79, 0x86, 0x15, 0xcf, 0x37, 0x68, 0x31, 0x06, 0x41, 0x35, 0x59, 0x99, 0x37, 0x3d, 0xdf,
	0xa0, 0x85, 0x30, 0xae, 0xeb, 0xf9, 0x06, 0x2d, 0x15, 0xa8, 0xf3, 0x2b, 0x50, 0x3c, 0x80, 0xe6,
	0x5e, 0x25, 0x49, 0x2e, 0x71, 0x34, 0xa1, 0xa0, 0xd4, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0xa5,
	0x78, 0xb0, 0xef, 0x9c, 0x8b, 0x52, 0x13, 0x4b, 0x52, 0x83, 0xc1, 0xba, 0x85, 0x24, 0xb8, 0xd8,
	0x93, 0x41, 0xfc, 0xfc, 0x22, 0x88, 0xdf, 0x82, 0x60, 0x5c, 0x21, 0x15, 0x2e, 0x5e, 0x88, 0x0d,
	0x50, 0x1f, 0x82, 0xbd, 0xc1, 0x19, 0x84, 0x2a, 0x68, 0xc5, 0x03, 0x72, 0x26, 0x4c, 0x0f, 0xd4,
	0x6e, 0x64, 0x0b, 0x60, 0x76, 0x1b, 0x5d, 0x60, 0xe4, 0x62, 0xf6, 0x2d, 0x4e, 0x17, 0x4a, 0xe0,
	0xe2, 0x41, 0x09, 0x5e, 0x65, 0x6c, 0xc1, 0x82, 0xe6, 0x01, 0x29, 0x6d, 0x22, 0x14, 0xc1, 0x6c,
	0x02, 0xd9, 0x80, 0xe2, 0x45, 0x5c, 0x36, 0x20, 0x2b, 0xc2, 0x69, 0x03, 0x36, 0xbf, 0x48, 0xb1,
	0x36, 0x80, 0xa2, 0xca, 0xc9, 0xe4, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c,
	0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2,
	0xa4, 0xb0, 0xc6, 0x54, 0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0x38, 0xa9, 0x19, 0x03, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xa9, 0x62, 0xcc, 0xfe, 0x1b, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	CreateStaker(ctx context.Context, in *MsgCreateStaker, opts ...grpc.CallOption) (*MsgCreateStakerResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, "/fiamma.bitvmstaker.Msg/UpdateParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateStaker(ctx context.Context, in *MsgCreateStaker, opts ...grpc.CallOption) (*MsgCreateStakerResponse, error) {
	out := new(MsgCreateStakerResponse)
	err := c.cc.Invoke(ctx, "/fiamma.bitvmstaker.Msg/CreateStaker", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	CreateStaker(context.Context, *MsgCreateStaker) (*MsgCreateStakerResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) UpdateParams(ctx context.Context, req *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (*UnimplementedMsgServer) CreateStaker(ctx context.Context, req *MsgCreateStaker) (*MsgCreateStakerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStaker not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fiamma.bitvmstaker.Msg/UpdateParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateStaker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateStaker)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateStaker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fiamma.bitvmstaker.Msg/CreateStaker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateStaker(ctx, req.(*MsgCreateStaker))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fiamma.bitvmstaker.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "CreateStaker",
			Handler:    _Msg_CreateStaker_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fiamma/bitvmstaker/tx.proto",
}

func (m *MsgUpdateParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgCreateStaker) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateStaker) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateStaker) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.StakerAddress) > 0 {
		i -= len(m.StakerAddress)
		copy(dAtA[i:], m.StakerAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.StakerAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateStakerResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateStakerResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateStakerResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgUpdateParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Params.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgUpdateParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgCreateStaker) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.StakerAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateStakerResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgUpdateParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUpdateParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUpdateParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUpdateParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgCreateStaker) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateStaker: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateStaker: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StakerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgCreateStakerResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateStakerResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateStakerResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
