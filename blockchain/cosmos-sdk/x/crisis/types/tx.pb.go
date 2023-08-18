// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/crisis/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// MsgVerifyInvariant represents a message to verify a particular invariance.
type MsgVerifyInvariant struct {
	// sender is the account address of private key to send coins to fee collector account.
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	// name of the invariant module.
	InvariantModuleName string `protobuf:"bytes,2,opt,name=invariant_module_name,json=invariantModuleName,proto3" json:"invariant_module_name,omitempty"`
	// invariant_route is the msg's invariant route.
	InvariantRoute string `protobuf:"bytes,3,opt,name=invariant_route,json=invariantRoute,proto3" json:"invariant_route,omitempty"`
}

func (m *MsgVerifyInvariant) Reset()         { *m = MsgVerifyInvariant{} }
func (m *MsgVerifyInvariant) String() string { return proto.CompactTextString(m) }
func (*MsgVerifyInvariant) ProtoMessage()    {}
func (*MsgVerifyInvariant) Descriptor() ([]byte, []int) {
	return fileDescriptor_61276163172fe867, []int{0}
}
func (m *MsgVerifyInvariant) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgVerifyInvariant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgVerifyInvariant.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgVerifyInvariant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgVerifyInvariant.Merge(m, src)
}
func (m *MsgVerifyInvariant) XXX_Size() int {
	return m.Size()
}
func (m *MsgVerifyInvariant) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgVerifyInvariant.DiscardUnknown(m)
}

var xxx_messageInfo_MsgVerifyInvariant proto.InternalMessageInfo

// MsgVerifyInvariantResponse defines the Msg/VerifyInvariant response type.
type MsgVerifyInvariantResponse struct {
}

func (m *MsgVerifyInvariantResponse) Reset()         { *m = MsgVerifyInvariantResponse{} }
func (m *MsgVerifyInvariantResponse) String() string { return proto.CompactTextString(m) }
func (*MsgVerifyInvariantResponse) ProtoMessage()    {}
func (*MsgVerifyInvariantResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_61276163172fe867, []int{1}
}
func (m *MsgVerifyInvariantResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgVerifyInvariantResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgVerifyInvariantResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgVerifyInvariantResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgVerifyInvariantResponse.Merge(m, src)
}
func (m *MsgVerifyInvariantResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgVerifyInvariantResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgVerifyInvariantResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgVerifyInvariantResponse proto.InternalMessageInfo

// MsgUpdateParams is the Msg/UpdateParams request type.
//
// Since: cosmos-sdk 0.47
type MsgUpdateParams struct {
	// authority is the address that controls the module (defaults to x/gov unless overwritten).
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// constant_fee defines the x/crisis parameter.
	ConstantFee types.Coin `protobuf:"bytes,2,opt,name=constant_fee,json=constantFee,proto3" json:"constant_fee"`
}

func (m *MsgUpdateParams) Reset()         { *m = MsgUpdateParams{} }
func (m *MsgUpdateParams) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateParams) ProtoMessage()    {}
func (*MsgUpdateParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_61276163172fe867, []int{2}
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

func (m *MsgUpdateParams) GetConstantFee() types.Coin {
	if m != nil {
		return m.ConstantFee
	}
	return types.Coin{}
}

// MsgUpdateParamsResponse defines the response structure for executing a
// MsgUpdateParams message.
//
// Since: cosmos-sdk 0.47
type MsgUpdateParamsResponse struct {
}

func (m *MsgUpdateParamsResponse) Reset()         { *m = MsgUpdateParamsResponse{} }
func (m *MsgUpdateParamsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateParamsResponse) ProtoMessage()    {}
func (*MsgUpdateParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_61276163172fe867, []int{3}
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

func init() {
	proto.RegisterType((*MsgVerifyInvariant)(nil), "cosmos.crisis.v1beta1.MsgVerifyInvariant")
	proto.RegisterType((*MsgVerifyInvariantResponse)(nil), "cosmos.crisis.v1beta1.MsgVerifyInvariantResponse")
	proto.RegisterType((*MsgUpdateParams)(nil), "cosmos.crisis.v1beta1.MsgUpdateParams")
	proto.RegisterType((*MsgUpdateParamsResponse)(nil), "cosmos.crisis.v1beta1.MsgUpdateParamsResponse")
}

func init() { proto.RegisterFile("cosmos/crisis/v1beta1/tx.proto", fileDescriptor_61276163172fe867) }

var fileDescriptor_61276163172fe867 = []byte{
	// 506 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcf, 0x6b, 0x13, 0x41,
	0x14, 0xde, 0xb5, 0x58, 0xc8, 0xb4, 0x18, 0x5c, 0x5b, 0x9a, 0x2c, 0xba, 0x91, 0x08, 0xfe, 0x88,
	0x74, 0xd7, 0x44, 0xec, 0xa1, 0x37, 0x23, 0x2a, 0x1e, 0x22, 0xb2, 0xa2, 0x07, 0x2f, 0x61, 0xb2,
	0x3b, 0xd9, 0x0e, 0xba, 0x33, 0x61, 0xde, 0x24, 0x34, 0x37, 0xf1, 0x24, 0x9e, 0xfc, 0x13, 0x7a,
	0xf4, 0x98, 0x83, 0x7f, 0x44, 0xf1, 0x54, 0x3c, 0x89, 0x07, 0x91, 0x04, 0x89, 0x7f, 0x86, 0xec,
	0xce, 0x4c, 0x52, 0x93, 0x8a, 0xb9, 0x24, 0xcb, 0xfb, 0xbe, 0xf7, 0xde, 0xf7, 0x7d, 0x33, 0x83,
	0xbc, 0x88, 0x43, 0xca, 0x21, 0x88, 0x04, 0x05, 0x0a, 0xc1, 0xa0, 0xde, 0x21, 0x12, 0xd7, 0x03,
	0x79, 0xe8, 0xf7, 0x04, 0x97, 0xdc, 0xd9, 0x56, 0xb8, 0xaf, 0x70, 0x5f, 0xe3, 0xee, 0x56, 0xc2,
	0x13, 0x9e, 0x33, 0x82, 0xec, 0x4b, 0x91, 0xdd, 0xb2, 0x22, 0xb7, 0x15, 0xa0, 0x3b, 0x15, 0xb4,
	0xa3, 0xf7, 0xa4, 0x90, 0x04, 0x83, 0x7a, 0xf6, 0xa7, 0x81, 0x8b, 0x38, 0xa5, 0x8c, 0x07, 0xf9,
	0xaf, 0x2e, 0x19, 0x4d, 0x1d, 0x0c, 0x64, 0xa6, 0x28, 0xe2, 0x94, 0x29, 0xbc, 0xfa, 0xdd, 0x46,
	0x4e, 0x0b, 0x92, 0x97, 0x44, 0xd0, 0xee, 0xf0, 0x09, 0x1b, 0x60, 0x41, 0x31, 0x93, 0xce, 0x1d,
	0xb4, 0x0e, 0x84, 0xc5, 0x44, 0x94, 0xec, 0xab, 0xf6, 0xcd, 0x42, 0xb3, 0xf4, 0xf5, 0xf3, 0xee,
	0x96, 0x16, 0x71, 0x3f, 0x8e, 0x05, 0x01, 0x78, 0x2e, 0x05, 0x65, 0x49, 0xa8, 0x79, 0x4e, 0x03,
	0x6d, 0x53, 0xd3, 0xde, 0x4e, 0x79, 0xdc, 0x7f, 0x43, 0xda, 0x0c, 0xa7, 0xa4, 0x74, 0x2e, 0x1b,
	0x10, 0x5e, 0x9a, 0x81, 0xad, 0x1c, 0x7b, 0x8a, 0x53, 0xe2, 0xdc, 0x40, 0xc5, 0x79, 0x8f, 0xe0,
	0x7d, 0x49, 0x4a, 0x6b, 0x39, 0xfb, 0xc2, 0xac, 0x1c, 0x66, 0xd5, 0xfd, 0x7b, 0xef, 0x8f, 0x2a,
	0xd6, 0xef, 0xa3, 0x8a, 0xf5, 0x6e, 0x3a, 0xaa, 0xe9, 0x8d, 0x1f, 0xa6, 0xa3, 0xda, 0x15, 0x25,
	0x69, 0x17, 0xe2, 0xd7, 0xc1, 0xb2, 0x8b, 0xea, 0x65, 0xe4, 0x2e, 0x57, 0x43, 0x02, 0x3d, 0xce,
	0x80, 0x54, 0xbf, 0xd8, 0xa8, 0xd8, 0x82, 0xe4, 0x45, 0x2f, 0xc6, 0x92, 0x3c, 0xc3, 0x02, 0xa7,
	0xe0, 0xec, 0xa1, 0x02, 0xee, 0xcb, 0x03, 0x2e, 0xa8, 0x1c, 0xfe, 0xd7, 0xfa, 0x9c, 0xea, 0x3c,
	0x46, 0x9b, 0x11, 0x67, 0x20, 0x33, 0x23, 0x5d, 0xa2, 0x4c, 0x6f, 0x34, 0xca, 0xbe, 0xee, 0xcb,
	0xd2, 0x37, 0xe7, 0xed, 0x3f, 0xe0, 0x94, 0x35, 0x0b, 0xc7, 0x3f, 0x2a, 0xd6, 0xa7, 0xe9, 0xa8,
	0x66, 0x87, 0x1b, 0xa6, 0xf3, 0x11, 0x21, 0xfb, 0x7b, 0x99, 0xc3, 0xf9, 0xe0, 0xcc, 0xe4, 0xb5,
	0x53, 0x26, 0x0f, 0xcd, 0xe5, 0x5a, 0x10, 0x5e, 0x2d, 0xa3, 0x9d, 0x85, 0x92, 0xf1, 0xd9, 0xf8,
	0x65, 0xa3, 0xb5, 0x16, 0x24, 0x0e, 0x47, 0xc5, 0xc5, 0x63, 0xbe, 0xe5, 0x9f, 0x79, 0x25, 0xfd,
	0xe5, 0xd4, 0xdc, 0xfa, 0xca, 0x54, 0xb3, 0xd8, 0xe9, 0xa2, 0xcd, 0xbf, 0xc2, 0xbd, 0xfe, 0xef,
	0x11, 0xa7, 0x79, 0xae, 0xbf, 0x1a, 0xcf, 0xec, 0x71, 0xcf, 0xbf, 0xcd, 0x72, 0x6c, 0x3e, 0x3c,
	0x1e, 0x7b, 0xf6, 0xc9, 0xd8, 0xb3, 0x7f, 0x8e, 0x3d, 0xfb, 0xe3, 0xc4, 0xb3, 0x4e, 0x26, 0x9e,
	0xf5, 0x6d, 0xe2, 0x59, 0xaf, 0x6e, 0x27, 0x54, 0x1e, 0xf4, 0x3b, 0x7e, 0xc4, 0xd3, 0xc0, 0xbc,
	0xd1, 0x33, 0x32, 0x95, 0xc3, 0x1e, 0x81, 0xce, 0x7a, 0xfe, 0x30, 0xee, 0xfe, 0x09, 0x00, 0x00,
	0xff, 0xff, 0x2a, 0xe4, 0x45, 0xe4, 0xce, 0x03, 0x00, 0x00,
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
	// VerifyInvariant defines a method to verify a particular invariant.
	VerifyInvariant(ctx context.Context, in *MsgVerifyInvariant, opts ...grpc.CallOption) (*MsgVerifyInvariantResponse, error)
	// UpdateParams defines a governance operation for updating the x/crisis module
	// parameters. The authority is defined in the keeper.
	//
	// Since: cosmos-sdk 0.47
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) VerifyInvariant(ctx context.Context, in *MsgVerifyInvariant, opts ...grpc.CallOption) (*MsgVerifyInvariantResponse, error) {
	out := new(MsgVerifyInvariantResponse)
	err := c.cc.Invoke(ctx, "/cosmos.crisis.v1beta1.Msg/VerifyInvariant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.crisis.v1beta1.Msg/UpdateParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// VerifyInvariant defines a method to verify a particular invariant.
	VerifyInvariant(context.Context, *MsgVerifyInvariant) (*MsgVerifyInvariantResponse, error)
	// UpdateParams defines a governance operation for updating the x/crisis module
	// parameters. The authority is defined in the keeper.
	//
	// Since: cosmos-sdk 0.47
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) VerifyInvariant(ctx context.Context, req *MsgVerifyInvariant) (*MsgVerifyInvariantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyInvariant not implemented")
}
func (*UnimplementedMsgServer) UpdateParams(ctx context.Context, req *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_VerifyInvariant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgVerifyInvariant)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).VerifyInvariant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.crisis.v1beta1.Msg/VerifyInvariant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).VerifyInvariant(ctx, req.(*MsgVerifyInvariant))
	}
	return interceptor(ctx, in, info, handler)
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
		FullMethod: "/cosmos.crisis.v1beta1.Msg/UpdateParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.crisis.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "VerifyInvariant",
			Handler:    _Msg_VerifyInvariant_Handler,
		},
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/crisis/v1beta1/tx.proto",
}

func (m *MsgVerifyInvariant) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgVerifyInvariant) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgVerifyInvariant) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.InvariantRoute) > 0 {
		i -= len(m.InvariantRoute)
		copy(dAtA[i:], m.InvariantRoute)
		i = encodeVarintTx(dAtA, i, uint64(len(m.InvariantRoute)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.InvariantModuleName) > 0 {
		i -= len(m.InvariantModuleName)
		copy(dAtA[i:], m.InvariantModuleName)
		i = encodeVarintTx(dAtA, i, uint64(len(m.InvariantModuleName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgVerifyInvariantResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgVerifyInvariantResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgVerifyInvariantResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
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
		size, err := m.ConstantFee.MarshalToSizedBuffer(dAtA[:i])
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
func (m *MsgVerifyInvariant) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.InvariantModuleName)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.InvariantRoute)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgVerifyInvariantResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
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
	l = m.ConstantFee.Size()
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

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgVerifyInvariant) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgVerifyInvariant: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgVerifyInvariant: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
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
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InvariantModuleName", wireType)
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
			m.InvariantModuleName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InvariantRoute", wireType)
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
			m.InvariantRoute = string(dAtA[iNdEx:postIndex])
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
func (m *MsgVerifyInvariantResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgVerifyInvariantResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgVerifyInvariantResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field ConstantFee", wireType)
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
			if err := m.ConstantFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
