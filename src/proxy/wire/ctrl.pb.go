// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ctrl.proto

package wire

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	math "math"
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

type IDResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IDResponse) Reset()         { *m = IDResponse{} }
func (m *IDResponse) String() string { return proto.CompactTextString(m) }
func (*IDResponse) ProtoMessage()    {}
func (*IDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0572e205f89e843, []int{0}
}

func (m *IDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IDResponse.Unmarshal(m, b)
}
func (m *IDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IDResponse.Marshal(b, m, deterministic)
}
func (m *IDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IDResponse.Merge(m, src)
}
func (m *IDResponse) XXX_Size() int {
	return xxx_messageInfo_IDResponse.Size(m)
}
func (m *IDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IDResponse proto.InternalMessageInfo

func (m *IDResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type StakeResponse struct {
	Value                float64  `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StakeResponse) Reset()         { *m = StakeResponse{} }
func (m *StakeResponse) String() string { return proto.CompactTextString(m) }
func (*StakeResponse) ProtoMessage()    {}
func (*StakeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0572e205f89e843, []int{1}
}

func (m *StakeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StakeResponse.Unmarshal(m, b)
}
func (m *StakeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StakeResponse.Marshal(b, m, deterministic)
}
func (m *StakeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StakeResponse.Merge(m, src)
}
func (m *StakeResponse) XXX_Size() int {
	return xxx_messageInfo_StakeResponse.Size(m)
}
func (m *StakeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StakeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StakeResponse proto.InternalMessageInfo

func (m *StakeResponse) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type InternalTxnRequest struct {
	Amount               uint64   `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Receiver             string   `protobuf:"bytes,2,opt,name=receiver,proto3" json:"receiver,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InternalTxnRequest) Reset()         { *m = InternalTxnRequest{} }
func (m *InternalTxnRequest) String() string { return proto.CompactTextString(m) }
func (*InternalTxnRequest) ProtoMessage()    {}
func (*InternalTxnRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0572e205f89e843, []int{2}
}

func (m *InternalTxnRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InternalTxnRequest.Unmarshal(m, b)
}
func (m *InternalTxnRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InternalTxnRequest.Marshal(b, m, deterministic)
}
func (m *InternalTxnRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InternalTxnRequest.Merge(m, src)
}
func (m *InternalTxnRequest) XXX_Size() int {
	return xxx_messageInfo_InternalTxnRequest.Size(m)
}
func (m *InternalTxnRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InternalTxnRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InternalTxnRequest proto.InternalMessageInfo

func (m *InternalTxnRequest) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *InternalTxnRequest) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func init() {
	proto.RegisterType((*IDResponse)(nil), "wire.IDResponse")
	proto.RegisterType((*StakeResponse)(nil), "wire.StakeResponse")
	proto.RegisterType((*InternalTxnRequest)(nil), "wire.InternalTxnRequest")
}

func init() { proto.RegisterFile("ctrl.proto", fileDescriptor_a0572e205f89e843) }

var fileDescriptor_a0572e205f89e843 = []byte{
	// 246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x93, 0x90, 0x16, 0x1d, 0x51, 0x64, 0x94, 0x12, 0xa2, 0x07, 0x59, 0x10, 0x3c, 0x6d,
	0x45, 0xc1, 0xbb, 0x58, 0xc1, 0x5c, 0x57, 0x5f, 0x20, 0x6d, 0xc7, 0x12, 0xdc, 0xee, 0xc6, 0xcd,
	0x6c, 0xd5, 0xe7, 0xf2, 0x05, 0xa5, 0xbb, 0xb5, 0x2a, 0x92, 0xe3, 0x3f, 0x7c, 0x33, 0xf3, 0xfd,
	0x00, 0x33, 0x76, 0x5a, 0xb6, 0xce, 0xb2, 0xc5, 0xfc, 0xad, 0x71, 0x54, 0x9e, 0x2c, 0xac, 0x5d,
	0x68, 0x1a, 0x87, 0xd9, 0xd4, 0x3f, 0x8f, 0x69, 0xd9, 0xf2, 0x47, 0x44, 0xc4, 0x29, 0x40, 0x35,
	0x51, 0xd4, 0xb5, 0xd6, 0x74, 0x84, 0x07, 0x90, 0x35, 0xf3, 0x22, 0x3d, 0x4b, 0x2f, 0x76, 0x55,
	0xd6, 0xcc, 0xc5, 0x39, 0xec, 0x3f, 0x72, 0xfd, 0x42, 0x5b, 0xe0, 0x18, 0x06, 0xab, 0x5a, 0x7b,
	0x0a, 0x4c, 0xaa, 0x62, 0x10, 0x0f, 0x80, 0x95, 0x61, 0x72, 0xa6, 0xd6, 0x4f, 0xef, 0x46, 0xd1,
	0xab, 0xa7, 0x8e, 0x71, 0x04, 0xc3, 0x7a, 0x69, 0xbd, 0xe1, 0x00, 0xe7, 0x6a, 0x93, 0xb0, 0x84,
	0x1d, 0x47, 0x33, 0x6a, 0x56, 0xe4, 0x8a, 0x2c, 0xbc, 0xda, 0xe6, 0xab, 0xcf, 0x14, 0xf2, 0x3b,
	0x76, 0x1a, 0x2f, 0x21, 0xab, 0x26, 0x38, 0x92, 0xd1, 0x5d, 0x7e, 0xbb, 0xcb, 0xfb, 0xb5, 0x7b,
	0x79, 0x28, 0xd7, 0xcd, 0xe4, 0x8f, 0xb9, 0x48, 0xf0, 0x06, 0x06, 0xc1, 0xb5, 0x77, 0xe9, 0x28,
	0x2e, 0xfd, 0x29, 0x24, 0x12, 0xbc, 0x85, 0xbd, 0x5f, 0xf2, 0x58, 0x6c, 0x4e, 0xff, 0xeb, 0x53,
	0xf6, 0xdc, 0x15, 0xc9, 0x74, 0x18, 0x26, 0xd7, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfd, 0xc7,
	0x6b, 0x56, 0x7c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CtrlClient is the client API for Ctrl service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CtrlClient interface {
	ID(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*IDResponse, error)
	Stake(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*StakeResponse, error)
	InternalTxn(ctx context.Context, in *InternalTxnRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type ctrlClient struct {
	cc *grpc.ClientConn
}

func NewCtrlClient(cc *grpc.ClientConn) CtrlClient {
	return &ctrlClient{cc}
}

func (c *ctrlClient) ID(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*IDResponse, error) {
	out := new(IDResponse)
	err := c.cc.Invoke(ctx, "/wire.Ctrl/ID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ctrlClient) Stake(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*StakeResponse, error) {
	out := new(StakeResponse)
	err := c.cc.Invoke(ctx, "/wire.Ctrl/Stake", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ctrlClient) InternalTxn(ctx context.Context, in *InternalTxnRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/wire.Ctrl/InternalTxn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CtrlServer is the server API for Ctrl service.
type CtrlServer interface {
	ID(context.Context, *empty.Empty) (*IDResponse, error)
	Stake(context.Context, *empty.Empty) (*StakeResponse, error)
	InternalTxn(context.Context, *InternalTxnRequest) (*empty.Empty, error)
}

func RegisterCtrlServer(s *grpc.Server, srv CtrlServer) {
	s.RegisterService(&_Ctrl_serviceDesc, srv)
}

func _Ctrl_ID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CtrlServer).ID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wire.Ctrl/ID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CtrlServer).ID(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ctrl_Stake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CtrlServer).Stake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wire.Ctrl/Stake",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CtrlServer).Stake(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ctrl_InternalTxn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InternalTxnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CtrlServer).InternalTxn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wire.Ctrl/InternalTxn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CtrlServer).InternalTxn(ctx, req.(*InternalTxnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Ctrl_serviceDesc = grpc.ServiceDesc{
	ServiceName: "wire.Ctrl",
	HandlerType: (*CtrlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ID",
			Handler:    _Ctrl_ID_Handler,
		},
		{
			MethodName: "Stake",
			Handler:    _Ctrl_Stake_Handler,
		},
		{
			MethodName: "InternalTxn",
			Handler:    _Ctrl_InternalTxn_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ctrl.proto",
}
