// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: comment/v1/comment.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CommentClient is the client API for Comment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommentClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateReply, error)
	GetComment(ctx context.Context, in *GetCommentRequest, opts ...grpc.CallOption) (*GetCommentReply, error)
	Like(ctx context.Context, in *LikeCommentRequest, opts ...grpc.CallOption) (*LikeCommentReply, error)
}

type commentClient struct {
	cc grpc.ClientConnInterface
}

func NewCommentClient(cc grpc.ClientConnInterface) CommentClient {
	return &commentClient{cc}
}

func (c *commentClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateReply, error) {
	out := new(CreateReply)
	err := c.cc.Invoke(ctx, "/comment.v1.Comment/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) GetComment(ctx context.Context, in *GetCommentRequest, opts ...grpc.CallOption) (*GetCommentReply, error) {
	out := new(GetCommentReply)
	err := c.cc.Invoke(ctx, "/comment.v1.Comment/GetComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) Like(ctx context.Context, in *LikeCommentRequest, opts ...grpc.CallOption) (*LikeCommentReply, error) {
	out := new(LikeCommentReply)
	err := c.cc.Invoke(ctx, "/comment.v1.Comment/Like", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentServer is the server API for Comment service.
// All implementations must embed UnimplementedCommentServer
// for forward compatibility
type CommentServer interface {
	Create(context.Context, *CreateRequest) (*CreateReply, error)
	GetComment(context.Context, *GetCommentRequest) (*GetCommentReply, error)
	Like(context.Context, *LikeCommentRequest) (*LikeCommentReply, error)
	mustEmbedUnimplementedCommentServer()
}

// UnimplementedCommentServer must be embedded to have forward compatible implementations.
type UnimplementedCommentServer struct {
}

func (UnimplementedCommentServer) Create(context.Context, *CreateRequest) (*CreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCommentServer) GetComment(context.Context, *GetCommentRequest) (*GetCommentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComment not implemented")
}
func (UnimplementedCommentServer) Like(context.Context, *LikeCommentRequest) (*LikeCommentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Like not implemented")
}
func (UnimplementedCommentServer) mustEmbedUnimplementedCommentServer() {}

// UnsafeCommentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommentServer will
// result in compilation errors.
type UnsafeCommentServer interface {
	mustEmbedUnimplementedCommentServer()
}

func RegisterCommentServer(s grpc.ServiceRegistrar, srv CommentServer) {
	s.RegisterService(&Comment_ServiceDesc, srv)
}

func _Comment_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comment.v1.Comment/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_GetComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).GetComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comment.v1.Comment/GetComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).GetComment(ctx, req.(*GetCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_Like_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikeCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).Like(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comment.v1.Comment/Like",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).Like(ctx, req.(*LikeCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Comment_ServiceDesc is the grpc.ServiceDesc for Comment service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Comment_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "comment.v1.Comment",
	HandlerType: (*CommentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Comment_Create_Handler,
		},
		{
			MethodName: "GetComment",
			Handler:    _Comment_GetComment_Handler,
		},
		{
			MethodName: "Like",
			Handler:    _Comment_Like_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comment/v1/comment.proto",
}
