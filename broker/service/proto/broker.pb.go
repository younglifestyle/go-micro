// Code generated by protoc-gen-go. DO NOT EDIT.
// source: broker/service/proto/broker.proto

package go_micro_broker

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_df4d8f04292cf3fe, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type PublishRequest struct {
	Topic                string   `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Message              *Message `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishRequest) Reset()         { *m = PublishRequest{} }
func (m *PublishRequest) String() string { return proto.CompactTextString(m) }
func (*PublishRequest) ProtoMessage()    {}
func (*PublishRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_df4d8f04292cf3fe, []int{1}
}

func (m *PublishRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishRequest.Unmarshal(m, b)
}
func (m *PublishRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishRequest.Marshal(b, m, deterministic)
}
func (m *PublishRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishRequest.Merge(m, src)
}
func (m *PublishRequest) XXX_Size() int {
	return xxx_messageInfo_PublishRequest.Size(m)
}
func (m *PublishRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PublishRequest proto.InternalMessageInfo

func (m *PublishRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *PublishRequest) GetMessage() *Message {
	if m != nil {
		return m.Message
	}
	return nil
}

type SubscribeRequest struct {
	Topic                string   `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Queue                string   `protobuf:"bytes,2,opt,name=queue,proto3" json:"queue,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeRequest) Reset()         { *m = SubscribeRequest{} }
func (m *SubscribeRequest) String() string { return proto.CompactTextString(m) }
func (*SubscribeRequest) ProtoMessage()    {}
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_df4d8f04292cf3fe, []int{2}
}

func (m *SubscribeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeRequest.Unmarshal(m, b)
}
func (m *SubscribeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeRequest.Marshal(b, m, deterministic)
}
func (m *SubscribeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeRequest.Merge(m, src)
}
func (m *SubscribeRequest) XXX_Size() int {
	return xxx_messageInfo_SubscribeRequest.Size(m)
}
func (m *SubscribeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeRequest proto.InternalMessageInfo

func (m *SubscribeRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *SubscribeRequest) GetQueue() string {
	if m != nil {
		return m.Queue
	}
	return ""
}

type Message struct {
	Header               map[string]string `protobuf:"bytes,1,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body                 []byte            `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_df4d8f04292cf3fe, []int{3}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetHeader() map[string]string {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Message) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "go.micro.broker.Empty")
	proto.RegisterType((*PublishRequest)(nil), "go.micro.broker.PublishRequest")
	proto.RegisterType((*SubscribeRequest)(nil), "go.micro.broker.SubscribeRequest")
	proto.RegisterType((*Message)(nil), "go.micro.broker.Message")
	proto.RegisterMapType((map[string]string)(nil), "go.micro.broker.Message.HeaderEntry")
}

func init() { proto.RegisterFile("broker/service/proto/broker.proto", fileDescriptor_df4d8f04292cf3fe) }

var fileDescriptor_df4d8f04292cf3fe = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0x4d, 0x4b, 0xc3, 0x40,
	0x14, 0xec, 0xb6, 0xb6, 0xa1, 0xaf, 0xa2, 0x65, 0x29, 0x12, 0x7a, 0x31, 0x0d, 0x1e, 0x72, 0xda,
	0x48, 0xbc, 0xa8, 0x88, 0x07, 0xb1, 0xe0, 0x41, 0x41, 0xd6, 0x9b, 0xb7, 0x6c, 0xfa, 0x68, 0x43,
	0x1b, 0x37, 0xdd, 0x4d, 0x0a, 0xf9, 0x23, 0x9e, 0xfc, 0xb1, 0xd2, 0xdd, 0xf8, 0xd5, 0x50, 0x6f,
	0x6f, 0xde, 0xce, 0xce, 0x1b, 0x66, 0x60, 0x22, 0x94, 0x5c, 0xa2, 0x0a, 0x35, 0xaa, 0x4d, 0x9a,
	0x60, 0x98, 0x2b, 0x59, 0xc8, 0xd0, 0x2e, 0x99, 0x01, 0xf4, 0x78, 0x2e, 0x59, 0x96, 0x26, 0x4a,
	0x32, 0xbb, 0xf6, 0x1d, 0xe8, 0x4e, 0xb3, 0xbc, 0xa8, 0xfc, 0x57, 0x38, 0x7a, 0x2e, 0xc5, 0x2a,
	0xd5, 0x0b, 0x8e, 0xeb, 0x12, 0x75, 0x41, 0x47, 0xd0, 0x2d, 0x64, 0x9e, 0x26, 0x2e, 0xf1, 0x48,
	0xd0, 0xe7, 0x16, 0xd0, 0x08, 0x9c, 0x0c, 0xb5, 0x8e, 0xe7, 0xe8, 0xb6, 0x3d, 0x12, 0x0c, 0x22,
	0x97, 0xed, 0x68, 0xb2, 0x27, 0xfb, 0xce, 0xbf, 0x88, 0xfe, 0x2d, 0x0c, 0x5f, 0x4a, 0xa1, 0x13,
	0x95, 0x0a, 0xfc, 0x5f, 0x7d, 0x04, 0xdd, 0x75, 0x89, 0xa5, 0xd5, 0xee, 0x73, 0x0b, 0xfc, 0x77,
	0x02, 0x4e, 0x2d, 0x4a, 0x6f, 0xa0, 0xb7, 0xc0, 0x78, 0x86, 0xca, 0x25, 0x5e, 0x27, 0x18, 0x44,
	0x67, 0xfb, 0xce, 0xb3, 0x07, 0x43, 0x9b, 0xbe, 0x15, 0xaa, 0xe2, 0xf5, 0x1f, 0x4a, 0xe1, 0x40,
	0xc8, 0x59, 0x65, 0xe4, 0x0f, 0xb9, 0x99, 0xc7, 0x57, 0x30, 0xf8, 0x45, 0xa5, 0x43, 0xe8, 0x2c,
	0xb1, 0xaa, 0x6d, 0x6d, 0xc7, 0xad, 0xa9, 0x4d, 0xbc, 0xfa, 0x31, 0x65, 0xc0, 0x75, 0xfb, 0x92,
	0x44, 0x1f, 0x04, 0x7a, 0x77, 0xe6, 0x2a, 0xbd, 0x07, 0xa7, 0xce, 0x8f, 0x9e, 0x36, 0x2c, 0xfd,
	0x4d, 0x76, 0x7c, 0xd2, 0x20, 0xd8, 0x0e, 0x5a, 0xf4, 0x11, 0xfa, 0xdf, 0x49, 0xd1, 0x49, 0x83,
	0xb6, 0x9b, 0xe2, 0x78, 0x6f, 0xf8, 0x7e, 0xeb, 0x9c, 0x88, 0x9e, 0x29, 0xfd, 0xe2, 0x33, 0x00,
	0x00, 0xff, 0xff, 0x19, 0x9f, 0x10, 0x75, 0x19, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BrokerClient is the client API for Broker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BrokerClient interface {
	Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*Empty, error)
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (Broker_SubscribeClient, error)
}

type brokerClient struct {
	cc *grpc.ClientConn
}

func NewBrokerClient(cc *grpc.ClientConn) BrokerClient {
	return &brokerClient{cc}
}

func (c *brokerClient) Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/go.micro.broker.Broker/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brokerClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (Broker_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Broker_serviceDesc.Streams[0], "/go.micro.broker.Broker/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &brokerSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Broker_SubscribeClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type brokerSubscribeClient struct {
	grpc.ClientStream
}

func (x *brokerSubscribeClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BrokerServer is the server API for Broker service.
type BrokerServer interface {
	Publish(context.Context, *PublishRequest) (*Empty, error)
	Subscribe(*SubscribeRequest, Broker_SubscribeServer) error
}

// UnimplementedBrokerServer can be embedded to have forward compatible implementations.
type UnimplementedBrokerServer struct {
}

func (*UnimplementedBrokerServer) Publish(ctx context.Context, req *PublishRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (*UnimplementedBrokerServer) Subscribe(req *SubscribeRequest, srv Broker_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}

func RegisterBrokerServer(s *grpc.Server, srv BrokerServer) {
	s.RegisterService(&_Broker_serviceDesc, srv)
}

func _Broker_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.broker.Broker/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerServer).Publish(ctx, req.(*PublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Broker_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BrokerServer).Subscribe(m, &brokerSubscribeServer{stream})
}

type Broker_SubscribeServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type brokerSubscribeServer struct {
	grpc.ServerStream
}

func (x *brokerSubscribeServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

var _Broker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "go.micro.broker.Broker",
	HandlerType: (*BrokerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _Broker_Publish_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _Broker_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "broker/service/proto/broker.proto",
}