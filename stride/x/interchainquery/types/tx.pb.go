// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: interchainquery/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

// MsgSubmitQueryResponse represents a message type to fulfil a query request.
type MsgSubmitQueryResponse struct {
	ChainId     string `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty" yaml:"chain_id"`
	QueryId     string `protobuf:"bytes,2,opt,name=query_id,json=queryId,proto3" json:"query_id,omitempty" yaml:"query_id"`
	Result      []byte `protobuf:"bytes,3,opt,name=result,proto3" json:"result,omitempty" yaml:"result"`
	Height      int64  `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty" yaml:"height"`
	FromAddress string `protobuf:"bytes,5,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`
}

func (m *MsgSubmitQueryResponse) Reset()         { *m = MsgSubmitQueryResponse{} }
func (m *MsgSubmitQueryResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitQueryResponse) ProtoMessage()    {}
func (*MsgSubmitQueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_81f8142dad1607ce, []int{0}
}
func (m *MsgSubmitQueryResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitQueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitQueryResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitQueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitQueryResponse.Merge(m, src)
}
func (m *MsgSubmitQueryResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitQueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitQueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitQueryResponse proto.InternalMessageInfo

// MsgSubmitQueryResponseResponse defines the MsgSubmitQueryResponse response
// type.
type MsgSubmitQueryResponseResponse struct {
}

func (m *MsgSubmitQueryResponseResponse) Reset()         { *m = MsgSubmitQueryResponseResponse{} }
func (m *MsgSubmitQueryResponseResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitQueryResponseResponse) ProtoMessage()    {}
func (*MsgSubmitQueryResponseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_81f8142dad1607ce, []int{1}
}
func (m *MsgSubmitQueryResponseResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitQueryResponseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitQueryResponseResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitQueryResponseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitQueryResponseResponse.Merge(m, src)
}
func (m *MsgSubmitQueryResponseResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitQueryResponseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitQueryResponseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitQueryResponseResponse proto.InternalMessageInfo

// MsgQueryBalance defines the payload for Msg/QueryBalance
type MsgQueryBalance struct {
	ChainId      string `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty" yaml:"chain_id"`
	Address      string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Denom        string `protobuf:"bytes,3,opt,name=denom,proto3" json:"denom,omitempty"`
	ConnectionId string `protobuf:"bytes,4,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty" yaml:"connection_id"`
	Caller       string `protobuf:"bytes,5,opt,name=caller,proto3" json:"caller,omitempty"`
	Height       int64  `protobuf:"varint,6,opt,name=height,proto3" json:"height,omitempty" yaml:"height"`
}

func (m *MsgQueryBalance) Reset()         { *m = MsgQueryBalance{} }
func (m *MsgQueryBalance) String() string { return proto.CompactTextString(m) }
func (*MsgQueryBalance) ProtoMessage()    {}
func (*MsgQueryBalance) Descriptor() ([]byte, []int) {
	return fileDescriptor_81f8142dad1607ce, []int{2}
}
func (m *MsgQueryBalance) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgQueryBalance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgQueryBalance.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgQueryBalance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgQueryBalance.Merge(m, src)
}
func (m *MsgQueryBalance) XXX_Size() int {
	return m.Size()
}
func (m *MsgQueryBalance) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgQueryBalance.DiscardUnknown(m)
}

var xxx_messageInfo_MsgQueryBalance proto.InternalMessageInfo

type MsgQueryBalanceResponse struct {
}

func (m *MsgQueryBalanceResponse) Reset()         { *m = MsgQueryBalanceResponse{} }
func (m *MsgQueryBalanceResponse) String() string { return proto.CompactTextString(m) }
func (*MsgQueryBalanceResponse) ProtoMessage()    {}
func (*MsgQueryBalanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_81f8142dad1607ce, []int{3}
}
func (m *MsgQueryBalanceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgQueryBalanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgQueryBalanceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgQueryBalanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgQueryBalanceResponse.Merge(m, src)
}
func (m *MsgQueryBalanceResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgQueryBalanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgQueryBalanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgQueryBalanceResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSubmitQueryResponse)(nil), "stride.interchainquery.MsgSubmitQueryResponse")
	proto.RegisterType((*MsgSubmitQueryResponseResponse)(nil), "stride.interchainquery.MsgSubmitQueryResponseResponse")
	proto.RegisterType((*MsgQueryBalance)(nil), "stride.interchainquery.MsgQueryBalance")
	proto.RegisterType((*MsgQueryBalanceResponse)(nil), "stride.interchainquery.MsgQueryBalanceResponse")
}

func init() { proto.RegisterFile("interchainquery/tx.proto", fileDescriptor_81f8142dad1607ce) }

var fileDescriptor_81f8142dad1607ce = []byte{
	// 563 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xb1, 0x6f, 0xd3, 0x40,
	0x14, 0xc6, 0x63, 0xa7, 0x4d, 0xda, 0xc3, 0x55, 0x85, 0x1b, 0x15, 0x27, 0x42, 0x76, 0x64, 0x21,
	0x11, 0x2a, 0xd5, 0x46, 0x45, 0x30, 0x04, 0x31, 0x90, 0x2d, 0x12, 0x11, 0xc2, 0xd9, 0x58, 0xa2,
	0x8b, 0x7d, 0x75, 0x4e, 0xb2, 0xef, 0x82, 0xef, 0x82, 0x9a, 0x95, 0x89, 0x11, 0x89, 0x05, 0xc4,
	0x92, 0x95, 0x11, 0x89, 0x3f, 0x82, 0xb1, 0x82, 0x85, 0x29, 0x42, 0x09, 0x03, 0x73, 0xfe, 0x02,
	0xe4, 0x3b, 0xbb, 0x84, 0x34, 0x22, 0x74, 0xf2, 0xbd, 0xfb, 0x7e, 0xf6, 0x7b, 0xfe, 0x3e, 0x9f,
	0x81, 0x81, 0x09, 0x47, 0x89, 0x3f, 0x80, 0x98, 0xbc, 0x18, 0xa1, 0x64, 0xec, 0xf2, 0x33, 0x67,
	0x98, 0x50, 0x4e, 0xf5, 0x43, 0xc6, 0x13, 0x1c, 0x20, 0x67, 0x05, 0xa8, 0x55, 0x42, 0x1a, 0x52,
	0x81, 0xb8, 0xe9, 0x4a, 0xd2, 0xb5, 0xaa, 0x4f, 0x59, 0x4c, 0x59, 0x4f, 0x0a, 0xb2, 0xc8, 0xa4,
	0x9b, 0x21, 0xa5, 0x61, 0x84, 0x5c, 0x38, 0xc4, 0x2e, 0x24, 0x84, 0x72, 0xc8, 0x31, 0x25, 0xb9,
	0x5a, 0xcd, 0x54, 0x51, 0xf5, 0x47, 0xa7, 0x2e, 0x24, 0x63, 0x29, 0xd9, 0xef, 0x55, 0x70, 0xd8,
	0x61, 0x61, 0x77, 0xd4, 0x8f, 0x31, 0x7f, 0x96, 0x36, 0xf7, 0x10, 0x1b, 0x52, 0xc2, 0x90, 0xee,
	0x80, 0x1d, 0x31, 0x52, 0x0f, 0x07, 0x86, 0x52, 0x57, 0x1a, 0xbb, 0xad, 0x83, 0xc5, 0xd4, 0xda,
	0x1f, 0xc3, 0x38, 0x6a, 0xda, 0xb9, 0x62, 0x7b, 0x65, 0xb1, 0x6c, 0x07, 0x29, 0x2f, 0xa6, 0x4f,
	0x79, 0x75, 0x95, 0xcf, 0x15, 0xdb, 0x2b, 0x8b, 0x65, 0x3b, 0xd0, 0xef, 0x80, 0x52, 0x82, 0xd8,
	0x28, 0xe2, 0x46, 0xb1, 0xae, 0x34, 0xb4, 0xd6, 0xf5, 0xc5, 0xd4, 0xda, 0x93, 0xb4, 0xdc, 0xb7,
	0xbd, 0x0c, 0x48, 0xd1, 0x01, 0xc2, 0xe1, 0x80, 0x1b, 0x5b, 0x75, 0xa5, 0x51, 0x5c, 0x46, 0xe5,
	0xbe, 0xed, 0x65, 0x80, 0xfe, 0x10, 0x68, 0xa7, 0x09, 0x8d, 0x7b, 0x30, 0x08, 0x12, 0xc4, 0x98,
	0xb1, 0x2d, 0x26, 0x31, 0xbe, 0x7e, 0x3e, 0xae, 0x64, 0x8e, 0x3d, 0x96, 0x4a, 0x97, 0x27, 0x98,
	0x84, 0xde, 0xb5, 0x94, 0xce, 0xb6, 0x9a, 0xda, 0xeb, 0x89, 0x55, 0x78, 0x37, 0xb1, 0x94, 0x5f,
	0x13, 0xab, 0x60, 0xd7, 0x81, 0xb9, 0xde, 0x9a, 0xfc, 0x6a, 0x7f, 0x52, 0xc1, 0x7e, 0x87, 0x85,
	0x42, 0x6c, 0xc1, 0x08, 0x12, 0xff, 0xea, 0xb6, 0x9d, 0x80, 0x72, 0x3e, 0xab, 0xba, 0x61, 0xd6,
	0x1c, 0xd4, 0x2b, 0x60, 0x3b, 0x40, 0x84, 0xc6, 0xc2, 0xb9, 0x5d, 0x4f, 0x16, 0xfa, 0x23, 0xb0,
	0xe7, 0x53, 0x42, 0x90, 0x9f, 0x66, 0x9f, 0xb6, 0xdf, 0x92, 0xcf, 0x5b, 0x4c, 0xad, 0x4a, 0xd6,
	0x7e, 0x59, 0xb6, 0x3d, 0xed, 0x4f, 0xdd, 0x0e, 0xf4, 0xbb, 0xa0, 0xe4, 0xc3, 0x28, 0x42, 0xc9,
	0x46, 0xcf, 0x32, 0x6e, 0x29, 0x96, 0xd2, 0x86, 0x58, 0x9a, 0x3b, 0xa9, 0xb3, 0xc2, 0xd5, 0x2a,
	0xb8, 0xb1, 0x62, 0x59, 0x6e, 0xe7, 0xc9, 0x07, 0x15, 0x14, 0x3b, 0x2c, 0xd4, 0x3f, 0x2a, 0xe0,
	0x60, 0xed, 0x17, 0xe9, 0xac, 0x3f, 0x2f, 0xce, 0xfa, 0x98, 0x6a, 0x0f, 0xae, 0xc6, 0x5f, 0xc4,
	0x7a, 0xf4, 0xea, 0xdb, 0xcf, 0xb7, 0xea, 0x2d, 0xdb, 0x72, 0x2f, 0x9f, 0x5c, 0x97, 0x89, 0x1b,
	0x45, 0xd9, 0x54, 0x8e, 0xf4, 0x01, 0xd0, 0xfe, 0x8a, 0xff, 0xf6, 0x3f, 0x7a, 0x2e, 0x83, 0x35,
	0xf7, 0x3f, 0xc1, 0x0b, 0x77, 0x34, 0x00, 0xd2, 0xb9, 0x51, 0xf2, 0x12, 0xfb, 0xa8, 0xf5, 0xf4,
	0xcb, 0xcc, 0x54, 0xce, 0x67, 0xa6, 0xf2, 0x63, 0x66, 0x2a, 0x6f, 0xe6, 0x66, 0xe1, 0x7c, 0x6e,
	0x16, 0xbe, 0xcf, 0xcd, 0xc2, 0xf3, 0xfb, 0x21, 0xe6, 0x83, 0x51, 0xdf, 0xf1, 0x69, 0xec, 0x76,
	0x45, 0x8b, 0xe3, 0x27, 0xb0, 0xcf, 0x5c, 0xd9, 0xce, 0x3d, 0xbb, 0xfc, 0x52, 0xe3, 0x21, 0x62,
	0xfd, 0x92, 0xf8, 0x21, 0xdc, 0xfb, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x05, 0xf3, 0x83, 0xae,
	0x04, 0x00, 0x00,
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
	// SubmitQueryResponse defines a method for submit query responses.
	SubmitQueryResponse(ctx context.Context, in *MsgSubmitQueryResponse, opts ...grpc.CallOption) (*MsgSubmitQueryResponseResponse, error)
	QueryBalance(ctx context.Context, in *MsgQueryBalance, opts ...grpc.CallOption) (*MsgQueryBalanceResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SubmitQueryResponse(ctx context.Context, in *MsgSubmitQueryResponse, opts ...grpc.CallOption) (*MsgSubmitQueryResponseResponse, error) {
	out := new(MsgSubmitQueryResponseResponse)
	err := c.cc.Invoke(ctx, "/stride.interchainquery.Msg/SubmitQueryResponse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) QueryBalance(ctx context.Context, in *MsgQueryBalance, opts ...grpc.CallOption) (*MsgQueryBalanceResponse, error) {
	out := new(MsgQueryBalanceResponse)
	err := c.cc.Invoke(ctx, "/stride.interchainquery.Msg/QueryBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// SubmitQueryResponse defines a method for submit query responses.
	SubmitQueryResponse(context.Context, *MsgSubmitQueryResponse) (*MsgSubmitQueryResponseResponse, error)
	QueryBalance(context.Context, *MsgQueryBalance) (*MsgQueryBalanceResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SubmitQueryResponse(ctx context.Context, req *MsgSubmitQueryResponse) (*MsgSubmitQueryResponseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitQueryResponse not implemented")
}
func (*UnimplementedMsgServer) QueryBalance(ctx context.Context, req *MsgQueryBalance) (*MsgQueryBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryBalance not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SubmitQueryResponse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSubmitQueryResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SubmitQueryResponse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stride.interchainquery.Msg/SubmitQueryResponse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitQueryResponse(ctx, req.(*MsgSubmitQueryResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_QueryBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgQueryBalance)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).QueryBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stride.interchainquery.Msg/QueryBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).QueryBalance(ctx, req.(*MsgQueryBalance))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stride.interchainquery.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitQueryResponse",
			Handler:    _Msg_SubmitQueryResponse_Handler,
		},
		{
			MethodName: "QueryBalance",
			Handler:    _Msg_QueryBalance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "interchainquery/tx.proto",
}

// MsgServiceClient is the client API for MsgService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgServiceClient interface {
}

type msgServiceClient struct {
	cc grpc1.ClientConn
}

func NewMsgServiceClient(cc grpc1.ClientConn) MsgServiceClient {
	return &msgServiceClient{cc}
}

// MsgServiceServer is the server API for MsgService service.
type MsgServiceServer interface {
}

// UnimplementedMsgServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServiceServer struct {
}

func RegisterMsgServiceServer(s grpc1.Server, srv MsgServiceServer) {
	s.RegisterService(&_MsgService_serviceDesc, srv)
}

var _MsgService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stride.interchainquery.MsgService",
	HandlerType: (*MsgServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "interchainquery/tx.proto",
}

func (m *MsgSubmitQueryResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitQueryResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitQueryResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FromAddress) > 0 {
		i -= len(m.FromAddress)
		copy(dAtA[i:], m.FromAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.FromAddress)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Height != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Result) > 0 {
		i -= len(m.Result)
		copy(dAtA[i:], m.Result)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Result)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.QueryId) > 0 {
		i -= len(m.QueryId)
		copy(dAtA[i:], m.QueryId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.QueryId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSubmitQueryResponseResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitQueryResponseResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitQueryResponseResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgQueryBalance) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgQueryBalance) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgQueryBalance) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Height != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Caller) > 0 {
		i -= len(m.Caller)
		copy(dAtA[i:], m.Caller)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Caller)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ConnectionId) > 0 {
		i -= len(m.ConnectionId)
		copy(dAtA[i:], m.ConnectionId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ConnectionId)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgQueryBalanceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgQueryBalanceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgQueryBalanceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgSubmitQueryResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.QueryId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Result)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Height != 0 {
		n += 1 + sovTx(uint64(m.Height))
	}
	l = len(m.FromAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSubmitQueryResponseResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgQueryBalance) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ConnectionId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Caller)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Height != 0 {
		n += 1 + sovTx(uint64(m.Height))
	}
	return n
}

func (m *MsgQueryBalanceResponse) Size() (n int) {
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
func (m *MsgSubmitQueryResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSubmitQueryResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitQueryResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
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
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueryId", wireType)
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
			m.QueryId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Result = append(m.Result[:0], dAtA[iNdEx:postIndex]...)
			if m.Result == nil {
				m.Result = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromAddress", wireType)
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
			m.FromAddress = string(dAtA[iNdEx:postIndex])
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
func (m *MsgSubmitQueryResponseResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSubmitQueryResponseResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitQueryResponseResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgQueryBalance) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgQueryBalance: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgQueryBalance: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
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
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
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
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionId", wireType)
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
			m.ConnectionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Caller", wireType)
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
			m.Caller = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *MsgQueryBalanceResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgQueryBalanceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgQueryBalanceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
