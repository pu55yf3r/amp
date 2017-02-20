// Code generated by protoc-gen-go.
// source: github.com/appcelerator/amp/api/rpc/account/account.proto
// DO NOT EDIT!

/*
Package account is a generated protocol buffer package.

It is generated from these files:
	github.com/appcelerator/amp/api/rpc/account/account.proto

It has these top-level messages:
	SignUpRequest
	VerificationRequest
	LogInRequest
	PasswordResetRequest
	PasswordSetRequest
	PasswordChangeRequest
	ForgotLoginRequest
*/
package account

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SignUpRequest struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
}

func (m *SignUpRequest) Reset()                    { *m = SignUpRequest{} }
func (m *SignUpRequest) String() string            { return proto.CompactTextString(m) }
func (*SignUpRequest) ProtoMessage()               {}
func (*SignUpRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SignUpRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SignUpRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *SignUpRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type VerificationRequest struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *VerificationRequest) Reset()                    { *m = VerificationRequest{} }
func (m *VerificationRequest) String() string            { return proto.CompactTextString(m) }
func (*VerificationRequest) ProtoMessage()               {}
func (*VerificationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *VerificationRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type LogInRequest struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *LogInRequest) Reset()                    { *m = LogInRequest{} }
func (m *LogInRequest) String() string            { return proto.CompactTextString(m) }
func (*LogInRequest) ProtoMessage()               {}
func (*LogInRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LogInRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LogInRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type PasswordResetRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *PasswordResetRequest) Reset()                    { *m = PasswordResetRequest{} }
func (m *PasswordResetRequest) String() string            { return proto.CompactTextString(m) }
func (*PasswordResetRequest) ProtoMessage()               {}
func (*PasswordResetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PasswordResetRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type PasswordSetRequest struct {
	Token    string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *PasswordSetRequest) Reset()                    { *m = PasswordSetRequest{} }
func (m *PasswordSetRequest) String() string            { return proto.CompactTextString(m) }
func (*PasswordSetRequest) ProtoMessage()               {}
func (*PasswordSetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PasswordSetRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *PasswordSetRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type PasswordChangeRequest struct {
	Name             string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	ExistingPassword string `protobuf:"bytes,2,opt,name=existingPassword" json:"existingPassword,omitempty"`
	NewPassword      string `protobuf:"bytes,3,opt,name=newPassword" json:"newPassword,omitempty"`
}

func (m *PasswordChangeRequest) Reset()                    { *m = PasswordChangeRequest{} }
func (m *PasswordChangeRequest) String() string            { return proto.CompactTextString(m) }
func (*PasswordChangeRequest) ProtoMessage()               {}
func (*PasswordChangeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PasswordChangeRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PasswordChangeRequest) GetExistingPassword() string {
	if m != nil {
		return m.ExistingPassword
	}
	return ""
}

func (m *PasswordChangeRequest) GetNewPassword() string {
	if m != nil {
		return m.NewPassword
	}
	return ""
}

type ForgotLoginRequest struct {
	Email string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
}

func (m *ForgotLoginRequest) Reset()                    { *m = ForgotLoginRequest{} }
func (m *ForgotLoginRequest) String() string            { return proto.CompactTextString(m) }
func (*ForgotLoginRequest) ProtoMessage()               {}
func (*ForgotLoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ForgotLoginRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func init() {
	proto.RegisterType((*SignUpRequest)(nil), "account.SignUpRequest")
	proto.RegisterType((*VerificationRequest)(nil), "account.VerificationRequest")
	proto.RegisterType((*LogInRequest)(nil), "account.LogInRequest")
	proto.RegisterType((*PasswordResetRequest)(nil), "account.PasswordResetRequest")
	proto.RegisterType((*PasswordSetRequest)(nil), "account.PasswordSetRequest")
	proto.RegisterType((*PasswordChangeRequest)(nil), "account.PasswordChangeRequest")
	proto.RegisterType((*ForgotLoginRequest)(nil), "account.ForgotLoginRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Account service

type AccountClient interface {
	SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	Verify(ctx context.Context, in *VerificationRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	Login(ctx context.Context, in *LogInRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	PasswordReset(ctx context.Context, in *PasswordResetRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	PasswordSet(ctx context.Context, in *PasswordSetRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	PasswordChange(ctx context.Context, in *PasswordChangeRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	ForgotLogin(ctx context.Context, in *ForgotLoginRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
}

type accountClient struct {
	cc *grpc.ClientConn
}

func NewAccountClient(cc *grpc.ClientConn) AccountClient {
	return &accountClient{cc}
}

func (c *accountClient) SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/account.Account/SignUp", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Verify(ctx context.Context, in *VerificationRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/account.Account/Verify", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Login(ctx context.Context, in *LogInRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/account.Account/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) PasswordReset(ctx context.Context, in *PasswordResetRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/account.Account/PasswordReset", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) PasswordSet(ctx context.Context, in *PasswordSetRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/account.Account/PasswordSet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) PasswordChange(ctx context.Context, in *PasswordChangeRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/account.Account/PasswordChange", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) ForgotLogin(ctx context.Context, in *ForgotLoginRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/account.Account/ForgotLogin", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Account service

type AccountServer interface {
	SignUp(context.Context, *SignUpRequest) (*google_protobuf1.Empty, error)
	Verify(context.Context, *VerificationRequest) (*google_protobuf1.Empty, error)
	Login(context.Context, *LogInRequest) (*google_protobuf1.Empty, error)
	PasswordReset(context.Context, *PasswordResetRequest) (*google_protobuf1.Empty, error)
	PasswordSet(context.Context, *PasswordSetRequest) (*google_protobuf1.Empty, error)
	PasswordChange(context.Context, *PasswordChangeRequest) (*google_protobuf1.Empty, error)
	ForgotLogin(context.Context, *ForgotLoginRequest) (*google_protobuf1.Empty, error)
}

func RegisterAccountServer(s *grpc.Server, srv AccountServer) {
	s.RegisterService(&_Account_serviceDesc, srv)
}

func _Account_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.Account/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).SignUp(ctx, req.(*SignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.Account/Verify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).Verify(ctx, req.(*VerificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.Account/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).Login(ctx, req.(*LogInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_PasswordReset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswordResetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).PasswordReset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.Account/PasswordReset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).PasswordReset(ctx, req.(*PasswordResetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_PasswordSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswordSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).PasswordSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.Account/PasswordSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).PasswordSet(ctx, req.(*PasswordSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_PasswordChange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswordChangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).PasswordChange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.Account/PasswordChange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).PasswordChange(ctx, req.(*PasswordChangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_ForgotLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForgotLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).ForgotLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.Account/ForgotLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).ForgotLogin(ctx, req.(*ForgotLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Account_serviceDesc = grpc.ServiceDesc{
	ServiceName: "account.Account",
	HandlerType: (*AccountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUp",
			Handler:    _Account_SignUp_Handler,
		},
		{
			MethodName: "Verify",
			Handler:    _Account_Verify_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Account_Login_Handler,
		},
		{
			MethodName: "PasswordReset",
			Handler:    _Account_PasswordReset_Handler,
		},
		{
			MethodName: "PasswordSet",
			Handler:    _Account_PasswordSet_Handler,
		},
		{
			MethodName: "PasswordChange",
			Handler:    _Account_PasswordChange_Handler,
		},
		{
			MethodName: "ForgotLogin",
			Handler:    _Account_ForgotLogin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/appcelerator/amp/api/rpc/account/account.proto",
}

func init() {
	proto.RegisterFile("github.com/appcelerator/amp/api/rpc/account/account.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 521 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x92, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0x87, 0xd5, 0x41, 0x3b, 0x78, 0xcb, 0x10, 0x98, 0xae, 0xab, 0xba, 0x0e, 0x8a, 0x05, 0x08,
	0x65, 0x28, 0x16, 0x70, 0x62, 0x07, 0x24, 0x84, 0x98, 0x84, 0xb4, 0xc3, 0xd4, 0x32, 0x38, 0x71,
	0x70, 0x83, 0xeb, 0x59, 0x34, 0xb6, 0x49, 0x9c, 0x8e, 0x0a, 0xed, 0xc2, 0x57, 0xe0, 0xa3, 0x71,
	0xe5, 0xc8, 0x07, 0x41, 0xb1, 0x9b, 0xd4, 0x51, 0xff, 0x20, 0x38, 0x25, 0xaf, 0xfd, 0xe6, 0x79,
	0xec, 0xbc, 0x3f, 0x78, 0xc1, 0x85, 0x39, 0xcf, 0x46, 0x61, 0xa4, 0x62, 0x42, 0xb5, 0x8e, 0xd8,
	0x84, 0x25, 0xd4, 0xa8, 0x84, 0xd0, 0x58, 0x13, 0xaa, 0x05, 0x49, 0x74, 0x44, 0x68, 0x14, 0xa9,
	0x4c, 0x9a, 0xe2, 0x19, 0xea, 0x44, 0x19, 0x85, 0xb6, 0xe7, 0x65, 0xb7, 0xc7, 0x95, 0xe2, 0x13,
	0x66, 0xdb, 0xa9, 0x94, 0xca, 0x50, 0x23, 0x94, 0x4c, 0x5d, 0x5b, 0x77, 0x7f, 0xbe, 0x6b, 0xab,
	0x51, 0x36, 0x26, 0x2c, 0xd6, 0x66, 0xe6, 0x36, 0xf1, 0x19, 0xec, 0x0c, 0x05, 0x97, 0x67, 0x7a,
	0xc0, 0xbe, 0x64, 0x2c, 0x35, 0x08, 0xc1, 0x55, 0x49, 0x63, 0xd6, 0xa9, 0xf5, 0x6b, 0x8f, 0xaf,
	0x0f, 0xec, 0x3b, 0xea, 0xc2, 0x35, 0x4d, 0xd3, 0xf4, 0x42, 0x25, 0x9f, 0x3a, 0x5b, 0x76, 0xbd,
	0xac, 0x51, 0x0b, 0xea, 0x2c, 0xa6, 0x62, 0xd2, 0xb9, 0x62, 0x37, 0x5c, 0x81, 0x0f, 0xe1, 0xce,
	0x7b, 0x96, 0x88, 0xb1, 0x88, 0xec, 0x51, 0x0a, 0x78, 0x0b, 0xea, 0x46, 0x7d, 0x66, 0x72, 0x4e,
	0x77, 0x05, 0x7e, 0x09, 0x37, 0x4e, 0x14, 0x7f, 0x2b, 0xff, 0xf3, 0x08, 0x38, 0x80, 0xd6, 0xe9,
	0xfc, 0x7d, 0xc0, 0x52, 0x66, 0x36, 0x70, 0xf0, 0x31, 0xa0, 0xa2, 0x77, 0xb8, 0xe8, 0x5c, 0x79,
	0xae, 0x8d, 0xce, 0x19, 0xec, 0x16, 0x9c, 0xd7, 0xe7, 0x54, 0x72, 0xb6, 0xe9, 0xf0, 0x01, 0xdc,
	0x62, 0x5f, 0x45, 0x6a, 0x84, 0xe4, 0xa7, 0x55, 0xe0, 0xd2, 0x3a, 0xea, 0x43, 0x53, 0xb2, 0x8b,
	0xb2, 0xcd, 0xfd, 0x55, 0x7f, 0x09, 0x07, 0x80, 0x8e, 0x55, 0xc2, 0x95, 0x39, 0x51, 0x5c, 0xf8,
	0xbf, 0xd6, 0xcd, 0xa1, 0xe6, 0xcd, 0xe1, 0xd9, 0xaf, 0x3a, 0x6c, 0xbf, 0x72, 0x29, 0x41, 0x1f,
	0xa0, 0xe1, 0x46, 0x8d, 0xda, 0x61, 0x11, 0xa4, 0xca, 0xec, 0xbb, 0xed, 0xd0, 0x45, 0x25, 0x2c,
	0xa2, 0x12, 0xbe, 0xc9, 0xa3, 0x82, 0x0f, 0xbe, 0xff, 0xfc, 0xfd, 0x63, 0x6b, 0x0f, 0x23, 0x32,
	0x7d, 0x5a, 0x66, 0x31, 0x15, 0x5c, 0x66, 0xfa, 0xa8, 0x16, 0xa0, 0x8f, 0xd0, 0xb0, 0xc3, 0x9e,
	0xa1, 0x5e, 0x09, 0x5e, 0x31, 0xfd, 0x7f, 0xc3, 0x4f, 0x2d, 0x31, 0xc7, 0xbf, 0x83, 0xba, 0xbd,
	0x29, 0xda, 0x2d, 0xe9, 0x7e, 0x5c, 0xd6, 0x62, 0x7b, 0x16, 0xdb, 0xc6, 0xb7, 0x7d, 0xec, 0x24,
	0x27, 0xe5, 0xd4, 0x14, 0x76, 0x2a, 0xa1, 0x41, 0x07, 0x25, 0x7d, 0x55, 0x98, 0xd6, 0x5a, 0x0e,
	0xad, 0xe5, 0x21, 0xee, 0xfb, 0x96, 0x6f, 0xf9, 0xd8, 0x2f, 0x89, 0xf6, 0x41, 0xb9, 0x74, 0x0c,
	0x4d, 0x2f, 0x7d, 0x68, 0x7f, 0x49, 0x39, 0xfc, 0xbb, 0x10, 0x5b, 0x61, 0x0f, 0xef, 0xf9, 0x42,
	0xbd, 0xf8, 0x3e, 0xf7, 0x4c, 0xe1, 0x66, 0x35, 0x9d, 0xe8, 0xee, 0x92, 0xaa, 0x12, 0xdb, 0xb5,
	0xb6, 0x27, 0xd6, 0xf6, 0x08, 0xdf, 0xdf, 0x70, 0x3d, 0x47, 0xca, 0xbd, 0x12, 0x9a, 0x5e, 0x34,
	0xbd, 0xfb, 0x2d, 0x07, 0x76, 0xad, 0x31, 0xb0, 0xc6, 0x07, 0xf8, 0x5e, 0xc5, 0x68, 0xe3, 0x7c,
	0x49, 0xc6, 0x0b, 0xce, 0x51, 0x2d, 0x18, 0x35, 0xec, 0xb7, 0xcf, 0xff, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x98, 0x95, 0xe7, 0xa8, 0x45, 0x05, 0x00, 0x00,
}