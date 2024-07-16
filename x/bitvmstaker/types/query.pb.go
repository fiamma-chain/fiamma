// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fiamma/bitvmstaker/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_684fd52f1c274172, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_684fd52f1c274172, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QueryAllStakerInfoRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllStakerInfoRequest) Reset()         { *m = QueryAllStakerInfoRequest{} }
func (m *QueryAllStakerInfoRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllStakerInfoRequest) ProtoMessage()    {}
func (*QueryAllStakerInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_684fd52f1c274172, []int{2}
}
func (m *QueryAllStakerInfoRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllStakerInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllStakerInfoRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllStakerInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllStakerInfoRequest.Merge(m, src)
}
func (m *QueryAllStakerInfoRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllStakerInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllStakerInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllStakerInfoRequest proto.InternalMessageInfo

func (m *QueryAllStakerInfoRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllStakerInfoResponse struct {
	AllStakerInfo []StakerInfo        `protobuf:"bytes,1,rep,name=all_staker_info,json=allStakerInfo,proto3" json:"all_staker_info"`
	Pagination    *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllStakerInfoResponse) Reset()         { *m = QueryAllStakerInfoResponse{} }
func (m *QueryAllStakerInfoResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllStakerInfoResponse) ProtoMessage()    {}
func (*QueryAllStakerInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_684fd52f1c274172, []int{3}
}
func (m *QueryAllStakerInfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllStakerInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllStakerInfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllStakerInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllStakerInfoResponse.Merge(m, src)
}
func (m *QueryAllStakerInfoResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllStakerInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllStakerInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllStakerInfoResponse proto.InternalMessageInfo

func (m *QueryAllStakerInfoResponse) GetAllStakerInfo() []StakerInfo {
	if m != nil {
		return m.AllStakerInfo
	}
	return nil
}

func (m *QueryAllStakerInfoResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryCommitteeAddressRequest struct {
}

func (m *QueryCommitteeAddressRequest) Reset()         { *m = QueryCommitteeAddressRequest{} }
func (m *QueryCommitteeAddressRequest) String() string { return proto.CompactTextString(m) }
func (*QueryCommitteeAddressRequest) ProtoMessage()    {}
func (*QueryCommitteeAddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_684fd52f1c274172, []int{4}
}
func (m *QueryCommitteeAddressRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCommitteeAddressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCommitteeAddressRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCommitteeAddressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCommitteeAddressRequest.Merge(m, src)
}
func (m *QueryCommitteeAddressRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryCommitteeAddressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCommitteeAddressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCommitteeAddressRequest proto.InternalMessageInfo

type QueryCommitteeAddressResponse struct {
	CommitteeAddress string `protobuf:"bytes,1,opt,name=committeeAddress,proto3" json:"committeeAddress,omitempty"`
}

func (m *QueryCommitteeAddressResponse) Reset()         { *m = QueryCommitteeAddressResponse{} }
func (m *QueryCommitteeAddressResponse) String() string { return proto.CompactTextString(m) }
func (*QueryCommitteeAddressResponse) ProtoMessage()    {}
func (*QueryCommitteeAddressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_684fd52f1c274172, []int{5}
}
func (m *QueryCommitteeAddressResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCommitteeAddressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCommitteeAddressResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCommitteeAddressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCommitteeAddressResponse.Merge(m, src)
}
func (m *QueryCommitteeAddressResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryCommitteeAddressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCommitteeAddressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCommitteeAddressResponse proto.InternalMessageInfo

func (m *QueryCommitteeAddressResponse) GetCommitteeAddress() string {
	if m != nil {
		return m.CommitteeAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "fiamma.bitvmstaker.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "fiamma.bitvmstaker.QueryParamsResponse")
	proto.RegisterType((*QueryAllStakerInfoRequest)(nil), "fiamma.bitvmstaker.QueryAllStakerInfoRequest")
	proto.RegisterType((*QueryAllStakerInfoResponse)(nil), "fiamma.bitvmstaker.QueryAllStakerInfoResponse")
	proto.RegisterType((*QueryCommitteeAddressRequest)(nil), "fiamma.bitvmstaker.QueryCommitteeAddressRequest")
	proto.RegisterType((*QueryCommitteeAddressResponse)(nil), "fiamma.bitvmstaker.QueryCommitteeAddressResponse")
}

func init() { proto.RegisterFile("fiamma/bitvmstaker/query.proto", fileDescriptor_684fd52f1c274172) }

var fileDescriptor_684fd52f1c274172 = []byte{
	// 527 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x41, 0x6b, 0x14, 0x31,
	0x14, 0xc7, 0x37, 0xad, 0x2e, 0x34, 0xa5, 0x58, 0x63, 0x0f, 0x35, 0xac, 0x69, 0x19, 0xb4, 0x5b,
	0x17, 0x9a, 0xb8, 0xab, 0x57, 0x0f, 0x5d, 0x41, 0x11, 0x3d, 0xd4, 0xd1, 0x93, 0x97, 0x25, 0xbb,
	0xcd, 0x0e, 0x83, 0x33, 0x93, 0xe9, 0x24, 0x2d, 0xf6, 0xe0, 0xc5, 0x4f, 0x20, 0xf8, 0x01, 0xbc,
	0x89, 0x47, 0xc1, 0x2f, 0xd1, 0x63, 0xc5, 0x8b, 0x27, 0x91, 0x5d, 0xc1, 0xaf, 0x21, 0x9b, 0xc4,
	0x76, 0x67, 0x37, 0x43, 0xf5, 0x32, 0x84, 0xbc, 0xf7, 0xff, 0xbf, 0xdf, 0x7b, 0x2f, 0x0c, 0x24,
	0xc3, 0x98, 0xa7, 0x29, 0x67, 0xfd, 0x58, 0x1f, 0xa5, 0x4a, 0xf3, 0x57, 0xa2, 0x60, 0x07, 0x87,
	0xa2, 0x38, 0xa6, 0x79, 0x21, 0xb5, 0x44, 0xc8, 0xc6, 0xe9, 0x54, 0x1c, 0x5f, 0xe5, 0x69, 0x9c,
	0x49, 0x66, 0xbe, 0x36, 0x0d, 0xaf, 0x45, 0x32, 0x92, 0xe6, 0xc8, 0x26, 0x27, 0x77, 0xdb, 0x88,
	0xa4, 0x8c, 0x12, 0xc1, 0x78, 0x1e, 0x33, 0x9e, 0x65, 0x52, 0x73, 0x1d, 0xcb, 0x4c, 0xb9, 0x68,
	0x6b, 0x20, 0x55, 0x2a, 0x15, 0xeb, 0x73, 0x25, 0x6c, 0x4d, 0x76, 0xd4, 0xee, 0x0b, 0xcd, 0xdb,
	0x2c, 0xe7, 0x51, 0x9c, 0x99, 0x64, 0x97, 0xbb, 0xe1, 0xc1, 0xcc, 0x79, 0xc1, 0xd3, 0xbf, 0x66,
	0x37, 0x3d, 0x09, 0x53, 0x67, 0x9b, 0x15, 0xac, 0x41, 0xf4, 0x6c, 0x52, 0x68, 0xcf, 0x48, 0x43,
	0x71, 0x70, 0x28, 0x94, 0x0e, 0x5e, 0xc0, 0x6b, 0xa5, 0x5b, 0x95, 0xcb, 0x4c, 0x09, 0x74, 0x1f,
	0xd6, 0x6d, 0x89, 0x75, 0xb0, 0x09, 0xb6, 0x97, 0x3b, 0x98, 0xce, 0xcf, 0x82, 0x5a, 0x4d, 0x77,
	0xe9, 0xe4, 0xc7, 0x46, 0xed, 0xd3, 0xef, 0xcf, 0x2d, 0x10, 0x3a, 0x51, 0x30, 0x80, 0xd7, 0x8d,
	0xeb, 0x6e, 0x92, 0x3c, 0x37, 0xb9, 0x8f, 0xb3, 0xa1, 0x74, 0x25, 0xd1, 0x43, 0x08, 0xcf, 0x7b,
	0x74, 0xfe, 0x5b, 0xd4, 0x0e, 0x84, 0x4e, 0x06, 0x42, 0xed, 0x12, 0xdc, 0x40, 0xe8, 0x1e, 0x8f,
	0x84, 0xd3, 0x86, 0x53, 0xca, 0xe0, 0x0b, 0x80, 0xd8, 0x57, 0xc5, 0xb5, 0xf0, 0x14, 0x5e, 0xe1,
	0x49, 0xd2, 0xb3, 0xac, 0xbd, 0x38, 0x1b, 0xca, 0x75, 0xb0, 0xb9, 0xb8, 0xbd, 0xdc, 0x21, 0xbe,
	0x5e, 0xce, 0x0d, 0xba, 0x97, 0x26, 0xfd, 0x84, 0x2b, 0x7c, 0xda, 0x15, 0x3d, 0x2a, 0x41, 0x2f,
	0x18, 0xe8, 0xe6, 0x85, 0xd0, 0x16, 0xa5, 0x44, 0x4d, 0x60, 0xc3, 0x40, 0x3f, 0x90, 0x69, 0x1a,
	0x6b, 0x2d, 0xc4, 0xee, 0xfe, 0x7e, 0x21, 0xd4, 0xd9, 0x42, 0x9e, 0xc0, 0x1b, 0x15, 0x71, 0xd7,
	0x57, 0x0b, 0xae, 0x0e, 0x66, 0x62, 0x66, 0x88, 0x4b, 0xe1, 0xdc, 0x7d, 0xe7, 0xeb, 0x22, 0xbc,
	0x6c, 0xdc, 0xd0, 0x1b, 0x58, 0xb7, 0xeb, 0x42, 0x5b, 0xbe, 0xf6, 0xe7, 0x5f, 0x06, 0x6e, 0x5e,
	0x98, 0x67, 0x81, 0x82, 0xe0, 0xed, 0xb7, 0x5f, 0xef, 0x17, 0x1a, 0x08, 0xb3, 0xca, 0x87, 0x8a,
	0x3e, 0x00, 0xb8, 0x52, 0x5a, 0x13, 0xda, 0xa9, 0xb4, 0xf7, 0x3d, 0x1a, 0x4c, 0xff, 0x35, 0xdd,
	0x41, 0x31, 0x03, 0x75, 0x1b, 0x35, 0x7d, 0x50, 0x91, 0xd0, 0xbd, 0x99, 0xb7, 0x81, 0x3e, 0x02,
	0xb8, 0x3a, 0x3b, 0x73, 0x74, 0xa7, 0xb2, 0x6a, 0xc5, 0xfa, 0x70, 0xfb, 0x3f, 0x14, 0x0e, 0x75,
	0xc7, 0xa0, 0x36, 0xd1, 0x2d, 0x1f, 0xea, 0xd9, 0x4a, 0x7b, 0xdc, 0xca, 0xba, 0xf7, 0x4e, 0x46,
	0x04, 0x9c, 0x8e, 0x08, 0xf8, 0x39, 0x22, 0xe0, 0xdd, 0x98, 0xd4, 0x4e, 0xc7, 0xa4, 0xf6, 0x7d,
	0x4c, 0x6a, 0x2f, 0xb1, 0xd3, 0xbf, 0x2e, 0x39, 0xe8, 0xe3, 0x5c, 0xa8, 0x7e, 0xdd, 0xfc, 0x04,
	0xee, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0xcc, 0xa7, 0x66, 0xe8, 0xf4, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// QueryAllStakerInfoRequest is the request type for the Query/AllStakerInfo RPC method.
	AllStakerInfo(ctx context.Context, in *QueryAllStakerInfoRequest, opts ...grpc.CallOption) (*QueryAllStakerInfoResponse, error)
	// QueryCommitteeAddressRequest is the request type for the Query/CommitteeAddress RPC method.
	CommitteeAddress(ctx context.Context, in *QueryCommitteeAddressRequest, opts ...grpc.CallOption) (*QueryCommitteeAddressResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/fiamma.bitvmstaker.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AllStakerInfo(ctx context.Context, in *QueryAllStakerInfoRequest, opts ...grpc.CallOption) (*QueryAllStakerInfoResponse, error) {
	out := new(QueryAllStakerInfoResponse)
	err := c.cc.Invoke(ctx, "/fiamma.bitvmstaker.Query/AllStakerInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) CommitteeAddress(ctx context.Context, in *QueryCommitteeAddressRequest, opts ...grpc.CallOption) (*QueryCommitteeAddressResponse, error) {
	out := new(QueryCommitteeAddressResponse)
	err := c.cc.Invoke(ctx, "/fiamma.bitvmstaker.Query/CommitteeAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// QueryAllStakerInfoRequest is the request type for the Query/AllStakerInfo RPC method.
	AllStakerInfo(context.Context, *QueryAllStakerInfoRequest) (*QueryAllStakerInfoResponse, error)
	// QueryCommitteeAddressRequest is the request type for the Query/CommitteeAddress RPC method.
	CommitteeAddress(context.Context, *QueryCommitteeAddressRequest) (*QueryCommitteeAddressResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) AllStakerInfo(ctx context.Context, req *QueryAllStakerInfoRequest) (*QueryAllStakerInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllStakerInfo not implemented")
}
func (*UnimplementedQueryServer) CommitteeAddress(ctx context.Context, req *QueryCommitteeAddressRequest) (*QueryCommitteeAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommitteeAddress not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fiamma.bitvmstaker.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AllStakerInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllStakerInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AllStakerInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fiamma.bitvmstaker.Query/AllStakerInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AllStakerInfo(ctx, req.(*QueryAllStakerInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_CommitteeAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCommitteeAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CommitteeAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fiamma.bitvmstaker.Query/CommitteeAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CommitteeAddress(ctx, req.(*QueryCommitteeAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fiamma.bitvmstaker.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "AllStakerInfo",
			Handler:    _Query_AllStakerInfo_Handler,
		},
		{
			MethodName: "CommitteeAddress",
			Handler:    _Query_CommitteeAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fiamma/bitvmstaker/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryAllStakerInfoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllStakerInfoRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllStakerInfoRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllStakerInfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllStakerInfoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllStakerInfoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.AllStakerInfo) > 0 {
		for iNdEx := len(m.AllStakerInfo) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AllStakerInfo[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryCommitteeAddressRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCommitteeAddressRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCommitteeAddressRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryCommitteeAddressResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCommitteeAddressResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCommitteeAddressResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CommitteeAddress) > 0 {
		i -= len(m.CommitteeAddress)
		copy(dAtA[i:], m.CommitteeAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.CommitteeAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAllStakerInfoRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllStakerInfoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.AllStakerInfo) > 0 {
		for _, e := range m.AllStakerInfo {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryCommitteeAddressRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryCommitteeAddressResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CommitteeAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
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
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllStakerInfoRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllStakerInfoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllStakerInfoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllStakerInfoResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllStakerInfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllStakerInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllStakerInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AllStakerInfo = append(m.AllStakerInfo, StakerInfo{})
			if err := m.AllStakerInfo[len(m.AllStakerInfo)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryCommitteeAddressRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryCommitteeAddressRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCommitteeAddressRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryCommitteeAddressResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryCommitteeAddressResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCommitteeAddressResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommitteeAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CommitteeAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
