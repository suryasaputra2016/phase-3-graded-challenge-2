// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: service.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	BookService_Register_FullMethodName         = "/books.BookService/Register"
	BookService_Login_FullMethodName            = "/books.BookService/Login"
	BookService_BorrowABook_FullMethodName      = "/books.BookService/BorrowABook"
	BookService_ReturnABook_FullMethodName      = "/books.BookService/ReturnABook"
	BookService_ShowAllBook_FullMethodName      = "/books.BookService/ShowAllBook"
	BookService_ShowBorrowedBook_FullMethodName = "/books.BookService/ShowBorrowedBook"
)

// BookServiceClient is the client API for BookService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookServiceClient interface {
	Register(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	Login(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	BorrowABook(ctx context.Context, in *BookRequest, opts ...grpc.CallOption) (*BookResponse, error)
	ReturnABook(ctx context.Context, in *BookRequest, opts ...grpc.CallOption) (*BookResponse, error)
	ShowAllBook(ctx context.Context, in *BookRequest, opts ...grpc.CallOption) (*AllBookResponse, error)
	ShowBorrowedBook(ctx context.Context, in *BorrowedBookRequest, opts ...grpc.CallOption) (*AllBorrowedBookResponse, error)
}

type bookServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookServiceClient(cc grpc.ClientConnInterface) BookServiceClient {
	return &bookServiceClient{cc}
}

func (c *bookServiceClient) Register(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, BookService_Register_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) Login(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, BookService_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) BorrowABook(ctx context.Context, in *BookRequest, opts ...grpc.CallOption) (*BookResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BookResponse)
	err := c.cc.Invoke(ctx, BookService_BorrowABook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) ReturnABook(ctx context.Context, in *BookRequest, opts ...grpc.CallOption) (*BookResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BookResponse)
	err := c.cc.Invoke(ctx, BookService_ReturnABook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) ShowAllBook(ctx context.Context, in *BookRequest, opts ...grpc.CallOption) (*AllBookResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AllBookResponse)
	err := c.cc.Invoke(ctx, BookService_ShowAllBook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) ShowBorrowedBook(ctx context.Context, in *BorrowedBookRequest, opts ...grpc.CallOption) (*AllBorrowedBookResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AllBorrowedBookResponse)
	err := c.cc.Invoke(ctx, BookService_ShowBorrowedBook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookServiceServer is the server API for BookService service.
// All implementations must embed UnimplementedBookServiceServer
// for forward compatibility.
type BookServiceServer interface {
	Register(context.Context, *UserRequest) (*RegisterResponse, error)
	Login(context.Context, *UserRequest) (*LoginResponse, error)
	BorrowABook(context.Context, *BookRequest) (*BookResponse, error)
	ReturnABook(context.Context, *BookRequest) (*BookResponse, error)
	ShowAllBook(context.Context, *BookRequest) (*AllBookResponse, error)
	ShowBorrowedBook(context.Context, *BorrowedBookRequest) (*AllBorrowedBookResponse, error)
	mustEmbedUnimplementedBookServiceServer()
}

// UnimplementedBookServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBookServiceServer struct{}

func (UnimplementedBookServiceServer) Register(context.Context, *UserRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedBookServiceServer) Login(context.Context, *UserRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedBookServiceServer) BorrowABook(context.Context, *BookRequest) (*BookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BorrowABook not implemented")
}
func (UnimplementedBookServiceServer) ReturnABook(context.Context, *BookRequest) (*BookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReturnABook not implemented")
}
func (UnimplementedBookServiceServer) ShowAllBook(context.Context, *BookRequest) (*AllBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowAllBook not implemented")
}
func (UnimplementedBookServiceServer) ShowBorrowedBook(context.Context, *BorrowedBookRequest) (*AllBorrowedBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowBorrowedBook not implemented")
}
func (UnimplementedBookServiceServer) mustEmbedUnimplementedBookServiceServer() {}
func (UnimplementedBookServiceServer) testEmbeddedByValue()                     {}

// UnsafeBookServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookServiceServer will
// result in compilation errors.
type UnsafeBookServiceServer interface {
	mustEmbedUnimplementedBookServiceServer()
}

func RegisterBookServiceServer(s grpc.ServiceRegistrar, srv BookServiceServer) {
	// If the following call pancis, it indicates UnimplementedBookServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BookService_ServiceDesc, srv)
}

func _BookService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookService_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).Register(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).Login(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_BorrowABook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).BorrowABook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookService_BorrowABook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).BorrowABook(ctx, req.(*BookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_ReturnABook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).ReturnABook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookService_ReturnABook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).ReturnABook(ctx, req.(*BookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_ShowAllBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).ShowAllBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookService_ShowAllBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).ShowAllBook(ctx, req.(*BookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_ShowBorrowedBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BorrowedBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).ShowBorrowedBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookService_ShowBorrowedBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).ShowBorrowedBook(ctx, req.(*BorrowedBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BookService_ServiceDesc is the grpc.ServiceDesc for BookService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "books.BookService",
	HandlerType: (*BookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _BookService_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _BookService_Login_Handler,
		},
		{
			MethodName: "BorrowABook",
			Handler:    _BookService_BorrowABook_Handler,
		},
		{
			MethodName: "ReturnABook",
			Handler:    _BookService_ReturnABook_Handler,
		},
		{
			MethodName: "ShowAllBook",
			Handler:    _BookService_ShowAllBook_Handler,
		},
		{
			MethodName: "ShowBorrowedBook",
			Handler:    _BookService_ShowBorrowedBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
