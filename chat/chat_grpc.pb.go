// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package chat

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceClient interface {
	SayHello(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

var chatServiceSayHelloStreamDesc = &grpc.StreamDesc{
	StreamName: "SayHello",
}

func (c *chatServiceClient) SayHello(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/chat.ChatService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceService is the service API for ChatService service.
// Fields should be assigned to their respective handler implementations only before
// RegisterChatServiceService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type ChatServiceService struct {
	SayHello func(context.Context, *Message) (*Message, error)
}

func (s *ChatServiceService) sayHello(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/chat.ChatService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.SayHello(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

// RegisterChatServiceService registers a service implementation with a gRPC server.
func RegisterChatServiceService(s grpc.ServiceRegistrar, srv *ChatServiceService) {
	srvCopy := *srv
	if srvCopy.SayHello == nil {
		srvCopy.SayHello = func(context.Context, *Message) (*Message, error) {
			return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
		}
	}
	sd := grpc.ServiceDesc{
		ServiceName: "chat.ChatService",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "SayHello",
				Handler:    srvCopy.sayHello,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "chat.proto",
	}

	s.RegisterService(&sd, nil)
}
