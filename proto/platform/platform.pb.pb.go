// Code generated by protoc-gen-go. DO NOT EDIT.
// source: platform.pb

package platformsvr

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// Request 请求数据格式
type Request struct {
	MainId               int32    `protobuf:"varint,1,opt,name=mainId,proto3" json:"mainId,omitempty"`
	SubId                int32    `protobuf:"varint,2,opt,name=subId,proto3" json:"subId,omitempty"`
	RequestId            int32    `protobuf:"varint,3,opt,name=requestId,proto3" json:"requestId,omitempty"`
	Input                string   `protobuf:"bytes,4,opt,name=input,proto3" json:"input,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a88c425e0dc4cae, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetMainId() int32 {
	if m != nil {
		return m.MainId
	}
	return 0
}

func (m *Request) GetSubId() int32 {
	if m != nil {
		return m.SubId
	}
	return 0
}

func (m *Request) GetRequestId() int32 {
	if m != nil {
		return m.RequestId
	}
	return 0
}

func (m *Request) GetInput() string {
	if m != nil {
		return m.Input
	}
	return ""
}

// Response 响应数据格式
type Response struct {
	MainId               int32    `protobuf:"varint,1,opt,name=mainId,proto3" json:"mainId,omitempty"`
	SubId                int32    `protobuf:"varint,2,opt,name=subId,proto3" json:"subId,omitempty"`
	RequestId            int32    `protobuf:"varint,3,opt,name=requestId,proto3" json:"requestId,omitempty"`
	Output               string   `protobuf:"bytes,4,opt,name=output,proto3" json:"output,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a88c425e0dc4cae, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMainId() int32 {
	if m != nil {
		return m.MainId
	}
	return 0
}

func (m *Response) GetSubId() int32 {
	if m != nil {
		return m.SubId
	}
	return 0
}

func (m *Response) GetRequestId() int32 {
	if m != nil {
		return m.RequestId
	}
	return 0
}

func (m *Response) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

type RequestRegisterService struct {
	ServiceId            int32    `protobuf:"varint,1,opt,name=serviceId,proto3" json:"serviceId,omitempty"`
	ServiceName          string   `protobuf:"bytes,2,opt,name=serviceName,proto3" json:"serviceName,omitempty"`
	ServiceVersion       string   `protobuf:"bytes,3,opt,name=serviceVersion,proto3" json:"serviceVersion,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestRegisterService) Reset()         { *m = RequestRegisterService{} }
func (m *RequestRegisterService) String() string { return proto.CompactTextString(m) }
func (*RequestRegisterService) ProtoMessage()    {}
func (*RequestRegisterService) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a88c425e0dc4cae, []int{2}
}

func (m *RequestRegisterService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestRegisterService.Unmarshal(m, b)
}
func (m *RequestRegisterService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestRegisterService.Marshal(b, m, deterministic)
}
func (m *RequestRegisterService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestRegisterService.Merge(m, src)
}
func (m *RequestRegisterService) XXX_Size() int {
	return xxx_messageInfo_RequestRegisterService.Size(m)
}
func (m *RequestRegisterService) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestRegisterService.DiscardUnknown(m)
}

var xxx_messageInfo_RequestRegisterService proto.InternalMessageInfo

func (m *RequestRegisterService) GetServiceId() int32 {
	if m != nil {
		return m.ServiceId
	}
	return 0
}

func (m *RequestRegisterService) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *RequestRegisterService) GetServiceVersion() string {
	if m != nil {
		return m.ServiceVersion
	}
	return ""
}

type ResponseRegisterService struct {
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseRegisterService) Reset()         { *m = ResponseRegisterService{} }
func (m *ResponseRegisterService) String() string { return proto.CompactTextString(m) }
func (*ResponseRegisterService) ProtoMessage()    {}
func (*ResponseRegisterService) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a88c425e0dc4cae, []int{3}
}

func (m *ResponseRegisterService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseRegisterService.Unmarshal(m, b)
}
func (m *ResponseRegisterService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseRegisterService.Marshal(b, m, deterministic)
}
func (m *ResponseRegisterService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseRegisterService.Merge(m, src)
}
func (m *ResponseRegisterService) XXX_Size() int {
	return xxx_messageInfo_ResponseRegisterService.Size(m)
}
func (m *ResponseRegisterService) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseRegisterService.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseRegisterService proto.InternalMessageInfo

func (m *ResponseRegisterService) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ResponseRegisterService) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "platformsvr.Request")
	proto.RegisterType((*Response)(nil), "platformsvr.Response")
	proto.RegisterType((*RequestRegisterService)(nil), "platformsvr.RequestRegisterService")
	proto.RegisterType((*ResponseRegisterService)(nil), "platformsvr.ResponseRegisterService")
}

func init() { proto.RegisterFile("platform.pb", fileDescriptor_1a88c425e0dc4cae) }

var fileDescriptor_1a88c425e0dc4cae = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0x31, 0x4f, 0xfb, 0x30,
	0x10, 0xc5, 0xeb, 0x7f, 0xff, 0x0d, 0xcd, 0x55, 0x02, 0x64, 0x95, 0x10, 0x55, 0x0c, 0x95, 0x41,
	0x28, 0x53, 0x84, 0x60, 0x67, 0x80, 0x29, 0x0b, 0x42, 0xae, 0xc4, 0x88, 0x94, 0x10, 0x53, 0x59,
	0x90, 0x38, 0xf8, 0x9c, 0xce, 0x7c, 0x1e, 0x3e, 0x25, 0x4a, 0xec, 0x36, 0x95, 0xd5, 0x91, 0xcd,
	0xbf, 0x97, 0xbb, 0x7b, 0xef, 0x72, 0x30, 0x6b, 0x3e, 0x73, 0xf3, 0xae, 0x74, 0x95, 0x36, 0x05,
	0xdd, 0x01, 0x6e, 0x34, 0xfb, 0x80, 0x23, 0x2e, 0xbe, 0x5a, 0x81, 0x86, 0x46, 0x10, 0x54, 0xb9,
	0xac, 0xb3, 0x32, 0x26, 0x4b, 0x92, 0x4c, 0xb8, 0x23, 0x3a, 0x87, 0x09, 0xb6, 0x45, 0x56, 0xc6,
	0xff, 0x7a, 0xd9, 0x02, 0xbd, 0x80, 0x50, 0xdb, 0xc6, 0xac, 0x8c, 0xc7, 0xfd, 0x97, 0x41, 0xe8,
	0x7a, 0x64, 0xdd, 0xb4, 0x26, 0xfe, 0xbf, 0x24, 0x49, 0xc8, 0x2d, 0xb0, 0x1a, 0xa6, 0x5c, 0x60,
	0xa3, 0x6a, 0x14, 0x7f, 0xea, 0x16, 0x41, 0xa0, 0x5a, 0x33, 0xd8, 0x39, 0x62, 0xdf, 0x04, 0x22,
	0xb7, 0x1d, 0x17, 0x6b, 0x89, 0x46, 0xe8, 0x95, 0xd0, 0x1b, 0xf9, 0x26, 0xba, 0x81, 0x68, 0x9f,
	0xbb, 0x04, 0x83, 0x40, 0x97, 0x30, 0x73, 0xf0, 0x94, 0x57, 0xa2, 0x8f, 0x12, 0xf2, 0x7d, 0x89,
	0x5e, 0xc3, 0xb1, 0xc3, 0x17, 0xa1, 0x51, 0xaa, 0xba, 0x4f, 0x15, 0x72, 0x4f, 0x65, 0x8f, 0x70,
	0xbe, 0x5d, 0xd9, 0x8f, 0x10, 0x41, 0x80, 0x26, 0x37, 0x2d, 0x6e, 0xff, 0x80, 0x25, 0x7a, 0x0a,
	0xe3, 0x0a, 0xd7, 0xce, 0xb4, 0x7b, 0xde, 0xfe, 0x10, 0x98, 0x3e, 0xbb, 0xa3, 0xd1, 0x7b, 0x08,
	0x1f, 0x64, 0xb9, 0x32, 0x5a, 0xe4, 0x15, 0x9d, 0xa7, 0x7b, 0xc7, 0x4c, 0xdd, 0xae, 0x8b, 0x33,
	0x4f, 0xb5, 0xfe, 0x6c, 0x94, 0x90, 0x1b, 0x42, 0x5f, 0xe1, 0xc4, 0x4f, 0x72, 0x79, 0x68, 0x8a,
	0x57, 0xb4, 0xb8, 0x3a, 0x38, 0xd4, 0xab, 0x62, 0xa3, 0x22, 0x68, 0xb4, 0x32, 0xea, 0xee, 0x37,
	0x00, 0x00, 0xff, 0xff, 0x27, 0x6d, 0xc8, 0xa9, 0x74, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PlatformClient is the client API for Platform service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PlatformClient interface {
	BidStream(ctx context.Context, opts ...grpc.CallOption) (Platform_BidStreamClient, error)
	RegisterService(ctx context.Context, in *RequestRegisterService, opts ...grpc.CallOption) (*ResponseRegisterService, error)
}

type platformClient struct {
	cc *grpc.ClientConn
}

func NewPlatformClient(cc *grpc.ClientConn) PlatformClient {
	return &platformClient{cc}
}

func (c *platformClient) BidStream(ctx context.Context, opts ...grpc.CallOption) (Platform_BidStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Platform_serviceDesc.Streams[0], "/platformsvr.Platform/BidStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &platformBidStreamClient{stream}
	return x, nil
}

type Platform_BidStreamClient interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type platformBidStreamClient struct {
	grpc.ClientStream
}

func (x *platformBidStreamClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *platformBidStreamClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *platformClient) RegisterService(ctx context.Context, in *RequestRegisterService, opts ...grpc.CallOption) (*ResponseRegisterService, error) {
	out := new(ResponseRegisterService)
	err := c.cc.Invoke(ctx, "/platformsvr.Platform/RegisterService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlatformServer is the server API for Platform service.
type PlatformServer interface {
	BidStream(Platform_BidStreamServer) error
	RegisterService(context.Context, *RequestRegisterService) (*ResponseRegisterService, error)
}

func RegisterPlatformServer(s *grpc.Server, srv PlatformServer) {
	s.RegisterService(&_Platform_serviceDesc, srv)
}

func _Platform_BidStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PlatformServer).BidStream(&platformBidStreamServer{stream})
}

type Platform_BidStreamServer interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type platformBidStreamServer struct {
	grpc.ServerStream
}

func (x *platformBidStreamServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *platformBidStreamServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Platform_RegisterService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestRegisterService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServer).RegisterService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/platformsvr.Platform/RegisterService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServer).RegisterService(ctx, req.(*RequestRegisterService))
	}
	return interceptor(ctx, in, info, handler)
}

var _Platform_serviceDesc = grpc.ServiceDesc{
	ServiceName: "platformsvr.Platform",
	HandlerType: (*PlatformServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterService",
			Handler:    _Platform_RegisterService_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BidStream",
			Handler:       _Platform_BidStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "platform.pb",
}
