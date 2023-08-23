// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peer/peer.proto

package peer

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

func init() { proto.RegisterFile("peer/peer.proto", fileDescriptor_c302117fbb08ad42) }

var fileDescriptor_c302117fbb08ad42 = []byte{
	// 175 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0xb1, 0x0a, 0xc2, 0x30,
	0x10, 0x86, 0x37, 0x91, 0x2c, 0x85, 0x0a, 0x22, 0xc5, 0xc9, 0xcd, 0x25, 0x01, 0x7d, 0x03, 0xc5,
	0xb9, 0x45, 0x37, 0x17, 0x69, 0x9b, 0x33, 0x0d, 0xd4, 0x5c, 0xb8, 0xab, 0x83, 0x6f, 0x2f, 0xed,
	0x35, 0xa0, 0x4b, 0x02, 0xdf, 0xff, 0xdd, 0x71, 0xbf, 0xca, 0x22, 0x00, 0x99, 0xf1, 0xd1, 0x91,
	0x70, 0xc0, 0x7c, 0x31, 0x7d, 0x5c, 0xac, 0x24, 0x20, 0x8c, 0xc8, 0x75, 0x2f, 0x61, 0xb1, 0xfd,
	0x83, 0x0f, 0x02, 0x8e, 0x18, 0x18, 0x24, 0x3d, 0x94, 0x6a, 0x79, 0x09, 0x16, 0x89, 0x81, 0xf2,
	0xb3, 0xca, 0x2a, 0xc2, 0x16, 0x98, 0xab, 0xd9, 0xce, 0xd7, 0xa2, 0xb1, 0xbe, 0x79, 0x17, 0xc0,
	0x26, 0x5e, 0x6c, 0x12, 0x4f, 0xe4, 0x3a, 0xaf, 0x3d, 0x95, 0x6a, 0x87, 0xe4, 0x74, 0xf7, 0x89,
	0x40, 0x3d, 0x58, 0x07, 0xa4, 0x9f, 0x75, 0x43, 0xbe, 0x4d, 0x13, 0xe3, 0x39, 0xf7, 0xbd, 0xf3,
	0x43, 0xf7, 0x6e, 0x74, 0x8b, 0x2f, 0xf3, 0xa3, 0x1a, 0x51, 0x8d, 0xa8, 0x53, 0xc5, 0x46, 0xca,
	0x1d, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x59, 0x1b, 0x23, 0x6a, 0xf6, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EndorserClient is the client API for Endorser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EndorserClient interface {
	ProcessProposal(ctx context.Context, in *SignedProposal, opts ...grpc.CallOption) (*ProposalResponse, error)
}

type endorserClient struct {
	cc *grpc.ClientConn
}

func NewEndorserClient(cc *grpc.ClientConn) EndorserClient {
	return &endorserClient{cc}
}

func (c *endorserClient) ProcessProposal(ctx context.Context, in *SignedProposal, opts ...grpc.CallOption) (*ProposalResponse, error) {
	out := new(ProposalResponse)
	err := c.cc.Invoke(ctx, "/protos.Endorser/ProcessProposal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EndorserServer is the server API for Endorser service.
type EndorserServer interface {
	ProcessProposal(context.Context, *SignedProposal) (*ProposalResponse, error)
}

func RegisterEndorserServer(s *grpc.Server, srv EndorserServer) {
	s.RegisterService(&_Endorser_serviceDesc, srv)
}

func _Endorser_ProcessProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignedProposal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndorserServer).ProcessProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Endorser/ProcessProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndorserServer).ProcessProposal(ctx, req.(*SignedProposal))
	}
	return interceptor(ctx, in, info, handler)
}

var _Endorser_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Endorser",
	HandlerType: (*EndorserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessProposal",
			Handler:    _Endorser_ProcessProposal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "peer/peer.proto",
}