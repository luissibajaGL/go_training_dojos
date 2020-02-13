// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pi.proto

package pb_pi_v1

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type PiRequest struct {
	A                    float64  `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PiRequest) Reset()         { *m = PiRequest{} }
func (m *PiRequest) String() string { return proto.CompactTextString(m) }
func (*PiRequest) ProtoMessage()    {}
func (*PiRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_42e306efa434d900, []int{0}
}

func (m *PiRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PiRequest.Unmarshal(m, b)
}
func (m *PiRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PiRequest.Marshal(b, m, deterministic)
}
func (m *PiRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PiRequest.Merge(m, src)
}
func (m *PiRequest) XXX_Size() int {
	return xxx_messageInfo_PiRequest.Size(m)
}
func (m *PiRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PiRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PiRequest proto.InternalMessageInfo

func (m *PiRequest) GetA() float64 {
	if m != nil {
		return m.A
	}
	return 0
}

type PiResponse struct {
	Result               float64  `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PiResponse) Reset()         { *m = PiResponse{} }
func (m *PiResponse) String() string { return proto.CompactTextString(m) }
func (*PiResponse) ProtoMessage()    {}
func (*PiResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_42e306efa434d900, []int{1}
}

func (m *PiResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PiResponse.Unmarshal(m, b)
}
func (m *PiResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PiResponse.Marshal(b, m, deterministic)
}
func (m *PiResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PiResponse.Merge(m, src)
}
func (m *PiResponse) XXX_Size() int {
	return xxx_messageInfo_PiResponse.Size(m)
}
func (m *PiResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PiResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PiResponse proto.InternalMessageInfo

func (m *PiResponse) GetResult() float64 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*PiRequest)(nil), "pb.pi.v1.PiRequest")
	proto.RegisterType((*PiResponse)(nil), "pb.pi.v1.PiResponse")
}

func init() { proto.RegisterFile("pi.proto", fileDescriptor_42e306efa434d900) }

var fileDescriptor_42e306efa434d900 = []byte{
	// 140 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x28, 0xc8, 0xd4, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x28, 0x48, 0xd2, 0x2b, 0xc8, 0xd4, 0x2b, 0x33, 0x54, 0x92,
	0xe4, 0xe2, 0x0c, 0xc8, 0x0c, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0xe2, 0xe1, 0x62, 0x4c,
	0x94, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x09, 0x62, 0x4c, 0x54, 0x52, 0xe1, 0xe2, 0x02, 0x49, 0x15,
	0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0x89, 0x71, 0xb1, 0x15, 0xa5, 0x16, 0x97, 0xe6, 0x94, 0x40,
	0x15, 0x40, 0x79, 0x46, 0xce, 0x20, 0x03, 0x82, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x85, 0xcc,
	0xb8, 0xd8, 0x9d, 0xf3, 0x73, 0x0b, 0x4a, 0x4b, 0x52, 0x85, 0x84, 0xf5, 0x60, 0x76, 0xe8, 0xc1,
	0x2d, 0x90, 0x12, 0x41, 0x15, 0x84, 0x18, 0xad, 0xc4, 0x90, 0xc4, 0x06, 0x76, 0x96, 0x31, 0x20,
	0x00, 0x00, 0xff, 0xff, 0xad, 0x9c, 0xdd, 0x58, 0xa2, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PiServiceClient is the client API for PiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PiServiceClient interface {
	Compute(ctx context.Context, in *PiRequest, opts ...grpc.CallOption) (*PiResponse, error)
}

type piServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPiServiceClient(cc grpc.ClientConnInterface) PiServiceClient {
	return &piServiceClient{cc}
}

func (c *piServiceClient) Compute(ctx context.Context, in *PiRequest, opts ...grpc.CallOption) (*PiResponse, error) {
	out := new(PiResponse)
	err := c.cc.Invoke(ctx, "/pb.pi.v1.PiService/Compute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PiServiceServer is the server API for PiService service.
type PiServiceServer interface {
	Compute(context.Context, *PiRequest) (*PiResponse, error)
}

// UnimplementedPiServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPiServiceServer struct {
}

func (*UnimplementedPiServiceServer) Compute(ctx context.Context, req *PiRequest) (*PiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Compute not implemented")
}

func RegisterPiServiceServer(s *grpc.Server, srv PiServiceServer) {
	s.RegisterService(&_PiService_serviceDesc, srv)
}

func _PiService_Compute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PiServiceServer).Compute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.pi.v1.PiService/Compute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PiServiceServer).Compute(ctx, req.(*PiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PiService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.pi.v1.PiService",
	HandlerType: (*PiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Compute",
			Handler:    _PiService_Compute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pi.proto",
}