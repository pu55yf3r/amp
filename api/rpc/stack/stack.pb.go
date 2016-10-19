// Code generated by protoc-gen-go.
// source: github.com/appcelerator/amp/api/rpc/stack/stack.proto
// DO NOT EDIT!

/*
Package stack is a generated protocol buffer package.

It is generated from these files:
	github.com/appcelerator/amp/api/rpc/stack/stack.proto

It has these top-level messages:
	UpRequest
	UpReply
	StackRequest
	RemoveRequest
	StackReply
	ListRequest
	ListReply
	StackInfo
	StackID
	CustomNetwork
	IdList
	NetworkSpec
	NetworkIPAM
	NetworkIPAMConfig
	Stack
*/
package stack

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import service "github.com/appcelerator/amp/api/rpc/service"

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

type StackState int32

const (
	StackState_Stopped     StackState = 0
	StackState_Starting    StackState = 1
	StackState_Running     StackState = 2
	StackState_Redeploying StackState = 3
)

var StackState_name = map[int32]string{
	0: "Stopped",
	1: "Starting",
	2: "Running",
	3: "Redeploying",
}
var StackState_value = map[string]int32{
	"Stopped":     0,
	"Starting":    1,
	"Running":     2,
	"Redeploying": 3,
}

func (x StackState) String() string {
	return proto.EnumName(StackState_name, int32(x))
}
func (StackState) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// struct for up request function
type UpRequest struct {
	StackName string `protobuf:"bytes,1,opt,name=stack_name,json=stackName" json:"stack_name,omitempty"`
	Stackfile string `protobuf:"bytes,2,opt,name=stackfile" json:"stackfile,omitempty"`
}

func (m *UpRequest) Reset()                    { *m = UpRequest{} }
func (m *UpRequest) String() string            { return proto.CompactTextString(m) }
func (*UpRequest) ProtoMessage()               {}
func (*UpRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// struct for up reply function
type UpReply struct {
	StackId string `protobuf:"bytes,1,opt,name=stack_id,json=stackId" json:"stack_id,omitempty"`
}

func (m *UpReply) Reset()                    { *m = UpReply{} }
func (m *UpReply) String() string            { return proto.CompactTextString(m) }
func (*UpReply) ProtoMessage()               {}
func (*UpReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// struct for start and stop request functions
type StackRequest struct {
	StackIdent string `protobuf:"bytes,1,opt,name=stack_ident,json=stackIdent" json:"stack_ident,omitempty"`
}

func (m *StackRequest) Reset()                    { *m = StackRequest{} }
func (m *StackRequest) String() string            { return proto.CompactTextString(m) }
func (*StackRequest) ProtoMessage()               {}
func (*StackRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// struct for remove request function
type RemoveRequest struct {
	StackIdent string `protobuf:"bytes,1,opt,name=stack_ident,json=stackIdent" json:"stack_ident,omitempty"`
	Force      bool   `protobuf:"varint,2,opt,name=force" json:"force,omitempty"`
}

func (m *RemoveRequest) Reset()                    { *m = RemoveRequest{} }
func (m *RemoveRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoveRequest) ProtoMessage()               {}
func (*RemoveRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// struct for start, stop and remove reply functions
type StackReply struct {
	StackId string `protobuf:"bytes,1,opt,name=stack_id,json=stackId" json:"stack_id,omitempty"`
}

func (m *StackReply) Reset()                    { *m = StackReply{} }
func (m *StackReply) String() string            { return proto.CompactTextString(m) }
func (*StackReply) ProtoMessage()               {}
func (*StackReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// struct for list request function
type ListRequest struct {
	All   bool  `protobuf:"varint,1,opt,name=all" json:"all,omitempty"`
	Limit int64 `protobuf:"varint,2,opt,name=limit" json:"limit,omitempty"`
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

// struct for list reply function
type ListReply struct {
	List []*StackInfo `protobuf:"bytes,1,rep,name=list" json:"list,omitempty"`
}

func (m *ListReply) Reset()                    { *m = ListReply{} }
func (m *ListReply) String() string            { return proto.CompactTextString(m) }
func (*ListReply) ProtoMessage()               {}
func (*ListReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ListReply) GetList() []*StackInfo {
	if m != nil {
		return m.List
	}
	return nil
}

// struct part of ListReply Struct
type StackInfo struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Id    string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	State string `protobuf:"bytes,3,opt,name=state" json:"state,omitempty"`
}

func (m *StackInfo) Reset()                    { *m = StackInfo{} }
func (m *StackInfo) String() string            { return proto.CompactTextString(m) }
func (*StackInfo) ProtoMessage()               {}
func (*StackInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

// struct to store Stack id in ETCD
type StackID struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *StackID) Reset()                    { *m = StackID{} }
func (m *StackID) String() string            { return proto.CompactTextString(m) }
func (*StackID) ProtoMessage()               {}
func (*StackID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

// struct to store network info in ETCD
type CustomNetwork struct {
	Id          string       `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	OwnerNumber int32        `protobuf:"varint,2,opt,name=owner_number,json=ownerNumber" json:"owner_number,omitempty"`
	Data        *NetworkSpec `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
}

func (m *CustomNetwork) Reset()                    { *m = CustomNetwork{} }
func (m *CustomNetwork) String() string            { return proto.CompactTextString(m) }
func (*CustomNetwork) ProtoMessage()               {}
func (*CustomNetwork) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *CustomNetwork) GetData() *NetworkSpec {
	if m != nil {
		return m.Data
	}
	return nil
}

// struct to store service id list in ETCD
type IdList struct {
	List []string `protobuf:"bytes,1,rep,name=list" json:"list,omitempty"`
}

func (m *IdList) Reset()                    { *m = IdList{} }
func (m *IdList) String() string            { return proto.CompactTextString(m) }
func (*IdList) ProtoMessage()               {}
func (*IdList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type NetworkSpec struct {
	Name       string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Driver     string            `protobuf:"bytes,2,opt,name=driver" json:"driver,omitempty"`
	EnableIpv6 bool              `protobuf:"varint,3,opt,name=enable_ipv6,json=enableIpv6" json:"enable_ipv6,omitempty"`
	Ipam       *NetworkIPAM      `protobuf:"bytes,4,opt,name=ipam" json:"ipam,omitempty"`
	Internal   bool              `protobuf:"varint,5,opt,name=internal" json:"internal,omitempty"`
	Options    map[string]string `protobuf:"bytes,6,rep,name=options" json:"options,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Labels     map[string]string `protobuf:"bytes,7,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	External   string            `protobuf:"bytes,8,opt,name=external" json:"external,omitempty"`
}

func (m *NetworkSpec) Reset()                    { *m = NetworkSpec{} }
func (m *NetworkSpec) String() string            { return proto.CompactTextString(m) }
func (*NetworkSpec) ProtoMessage()               {}
func (*NetworkSpec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *NetworkSpec) GetIpam() *NetworkIPAM {
	if m != nil {
		return m.Ipam
	}
	return nil
}

func (m *NetworkSpec) GetOptions() map[string]string {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *NetworkSpec) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type NetworkIPAM struct {
	Driver  string               `protobuf:"bytes,1,opt,name=driver" json:"driver,omitempty"`
	Options map[string]string    `protobuf:"bytes,2,rep,name=options" json:"options,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Config  []*NetworkIPAMConfig `protobuf:"bytes,3,rep,name=config" json:"config,omitempty"`
}

func (m *NetworkIPAM) Reset()                    { *m = NetworkIPAM{} }
func (m *NetworkIPAM) String() string            { return proto.CompactTextString(m) }
func (*NetworkIPAM) ProtoMessage()               {}
func (*NetworkIPAM) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *NetworkIPAM) GetOptions() map[string]string {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *NetworkIPAM) GetConfig() []*NetworkIPAMConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

type NetworkIPAMConfig struct {
	Subnet     string            `protobuf:"bytes,1,opt,name=subnet" json:"subnet,omitempty"`
	IpRange    string            `protobuf:"bytes,2,opt,name=ip_range,json=ipRange" json:"ip_range,omitempty"`
	Gateway    string            `protobuf:"bytes,3,opt,name=gateway" json:"gateway,omitempty"`
	AuxAddress map[string]string `protobuf:"bytes,4,rep,name=aux_address,json=auxAddress" json:"aux_address,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *NetworkIPAMConfig) Reset()                    { *m = NetworkIPAMConfig{} }
func (m *NetworkIPAMConfig) String() string            { return proto.CompactTextString(m) }
func (*NetworkIPAMConfig) ProtoMessage()               {}
func (*NetworkIPAMConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *NetworkIPAMConfig) GetAuxAddress() map[string]string {
	if m != nil {
		return m.AuxAddress
	}
	return nil
}

// Stack struct
type Stack struct {
	Name     string                 `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Id       string                 `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Services []*service.ServiceSpec `protobuf:"bytes,3,rep,name=services" json:"services,omitempty"`
	Networks []*NetworkSpec         `protobuf:"bytes,4,rep,name=networks" json:"networks,omitempty"`
	IsPublic bool                   `protobuf:"varint,5,opt,name=is_public,json=isPublic" json:"is_public,omitempty"`
}

func (m *Stack) Reset()                    { *m = Stack{} }
func (m *Stack) String() string            { return proto.CompactTextString(m) }
func (*Stack) ProtoMessage()               {}
func (*Stack) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *Stack) GetServices() []*service.ServiceSpec {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *Stack) GetNetworks() []*NetworkSpec {
	if m != nil {
		return m.Networks
	}
	return nil
}

func init() {
	proto.RegisterType((*UpRequest)(nil), "stack.UpRequest")
	proto.RegisterType((*UpReply)(nil), "stack.UpReply")
	proto.RegisterType((*StackRequest)(nil), "stack.StackRequest")
	proto.RegisterType((*RemoveRequest)(nil), "stack.removeRequest")
	proto.RegisterType((*StackReply)(nil), "stack.StackReply")
	proto.RegisterType((*ListRequest)(nil), "stack.ListRequest")
	proto.RegisterType((*ListReply)(nil), "stack.ListReply")
	proto.RegisterType((*StackInfo)(nil), "stack.StackInfo")
	proto.RegisterType((*StackID)(nil), "stack.StackID")
	proto.RegisterType((*CustomNetwork)(nil), "stack.CustomNetwork")
	proto.RegisterType((*IdList)(nil), "stack.IdList")
	proto.RegisterType((*NetworkSpec)(nil), "stack.NetworkSpec")
	proto.RegisterType((*NetworkIPAM)(nil), "stack.NetworkIPAM")
	proto.RegisterType((*NetworkIPAMConfig)(nil), "stack.NetworkIPAMConfig")
	proto.RegisterType((*Stack)(nil), "stack.Stack")
	proto.RegisterEnum("stack.StackState", StackState_name, StackState_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for StackService service

type StackServiceClient interface {
	Up(ctx context.Context, in *UpRequest, opts ...grpc.CallOption) (*UpReply, error)
	Start(ctx context.Context, in *StackRequest, opts ...grpc.CallOption) (*StackReply, error)
	Stop(ctx context.Context, in *StackRequest, opts ...grpc.CallOption) (*StackReply, error)
	Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*StackReply, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error)
}

type stackServiceClient struct {
	cc *grpc.ClientConn
}

func NewStackServiceClient(cc *grpc.ClientConn) StackServiceClient {
	return &stackServiceClient{cc}
}

func (c *stackServiceClient) Up(ctx context.Context, in *UpRequest, opts ...grpc.CallOption) (*UpReply, error) {
	out := new(UpReply)
	err := grpc.Invoke(ctx, "/stack.StackService/Up", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stackServiceClient) Start(ctx context.Context, in *StackRequest, opts ...grpc.CallOption) (*StackReply, error) {
	out := new(StackReply)
	err := grpc.Invoke(ctx, "/stack.StackService/Start", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stackServiceClient) Stop(ctx context.Context, in *StackRequest, opts ...grpc.CallOption) (*StackReply, error) {
	out := new(StackReply)
	err := grpc.Invoke(ctx, "/stack.StackService/Stop", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stackServiceClient) Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*StackReply, error) {
	out := new(StackReply)
	err := grpc.Invoke(ctx, "/stack.StackService/Remove", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stackServiceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error) {
	out := new(ListReply)
	err := grpc.Invoke(ctx, "/stack.StackService/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for StackService service

type StackServiceServer interface {
	Up(context.Context, *UpRequest) (*UpReply, error)
	Start(context.Context, *StackRequest) (*StackReply, error)
	Stop(context.Context, *StackRequest) (*StackReply, error)
	Remove(context.Context, *RemoveRequest) (*StackReply, error)
	List(context.Context, *ListRequest) (*ListReply, error)
}

func RegisterStackServiceServer(s *grpc.Server, srv StackServiceServer) {
	s.RegisterService(&_StackService_serviceDesc, srv)
}

func _StackService_Up_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StackServiceServer).Up(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stack.StackService/Up",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StackServiceServer).Up(ctx, req.(*UpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StackService_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StackServiceServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stack.StackService/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StackServiceServer).Start(ctx, req.(*StackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StackService_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StackServiceServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stack.StackService/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StackServiceServer).Stop(ctx, req.(*StackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StackService_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StackServiceServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stack.StackService/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StackServiceServer).Remove(ctx, req.(*RemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StackService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StackServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stack.StackService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StackServiceServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StackService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stack.StackService",
	HandlerType: (*StackServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Up",
			Handler:    _StackService_Up_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _StackService_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _StackService_Stop_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _StackService_Remove_Handler,
		},
		{
			MethodName: "List",
			Handler:    _StackService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() {
	proto.RegisterFile("github.com/appcelerator/amp/api/rpc/stack/stack.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 869 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x55, 0xdd, 0x6e, 0xdc, 0x44,
	0x14, 0xc6, 0xfb, 0x6b, 0x1f, 0xa7, 0xed, 0x76, 0x88, 0x90, 0xbb, 0x14, 0x0a, 0x56, 0x55, 0x22,
	0x84, 0xbc, 0x6d, 0x50, 0x23, 0x52, 0x89, 0x8b, 0xa8, 0x04, 0xb1, 0x52, 0x09, 0xd5, 0xac, 0x7a,
	0xbd, 0x9a, 0xb5, 0x27, 0xcb, 0x10, 0x7b, 0x3c, 0x8c, 0xc7, 0x9b, 0xec, 0xfb, 0x70, 0xc5, 0xc3,
	0xf0, 0x20, 0x3c, 0x01, 0x97, 0x68, 0x7e, 0x6c, 0x39, 0xcd, 0x56, 0x24, 0xe2, 0x66, 0x77, 0xce,
	0x39, 0xdf, 0x77, 0xfe, 0xc6, 0x73, 0x0e, 0xbc, 0x5c, 0x33, 0xf5, 0x6b, 0xbd, 0x4a, 0xd2, 0xb2,
	0x98, 0x11, 0x21, 0x52, 0x9a, 0x53, 0x49, 0x54, 0x29, 0x67, 0xa4, 0x10, 0x33, 0x22, 0xd8, 0x4c,
	0x8a, 0x74, 0x56, 0x29, 0x92, 0x5e, 0xd8, 0xdf, 0x44, 0xc8, 0x52, 0x95, 0x68, 0x68, 0x84, 0xe9,
	0xf1, 0xad, 0xd8, 0x54, 0x6e, 0x58, 0x4a, 0x9b, 0x7f, 0xeb, 0x21, 0xfe, 0x09, 0x82, 0x77, 0x02,
	0xd3, 0xdf, 0x6b, 0x5a, 0x29, 0xf4, 0x19, 0x80, 0x71, 0xb8, 0xe4, 0xa4, 0xa0, 0x91, 0xf7, 0x85,
	0x77, 0x10, 0xe0, 0xc0, 0x68, 0xce, 0x48, 0x41, 0xd1, 0x63, 0xb0, 0xc2, 0x39, 0xcb, 0x69, 0xd4,
	0xeb, 0x58, 0xb5, 0x22, 0x7e, 0x0a, 0x63, 0xed, 0x49, 0xe4, 0x5b, 0xf4, 0x08, 0x7c, 0xeb, 0x87,
	0x65, 0xce, 0xcb, 0xd8, 0xc8, 0xf3, 0x2c, 0x9e, 0xc1, 0xde, 0x42, 0x1f, 0x9b, 0x90, 0x4f, 0x20,
	0x6c, 0xa0, 0x94, 0x2b, 0x87, 0x06, 0x87, 0xa6, 0x5c, 0xc5, 0x3f, 0xc2, 0x3d, 0x49, 0x8b, 0x72,
	0x43, 0x6f, 0xcb, 0x40, 0xfb, 0x30, 0x3c, 0x2f, 0x65, 0x6a, 0x53, 0xf4, 0xb1, 0x15, 0xe2, 0xaf,
	0x00, 0x5c, 0xe0, 0xff, 0xc8, 0xf0, 0x25, 0x84, 0x6f, 0x58, 0xa5, 0x9a, 0x70, 0x13, 0xe8, 0x93,
	0x3c, 0x37, 0x20, 0x1f, 0xeb, 0xa3, 0xf6, 0x9f, 0xb3, 0x82, 0x29, 0xe3, 0xbf, 0x8f, 0xad, 0x10,
	0xbf, 0x80, 0xc0, 0xd2, 0xb4, 0xfb, 0xa7, 0x30, 0xc8, 0x59, 0xa5, 0x93, 0xeb, 0x1f, 0x84, 0x87,
	0x93, 0xc4, 0xde, 0x99, 0x89, 0x3f, 0xe7, 0xe7, 0x25, 0x36, 0xd6, 0xf8, 0x14, 0x82, 0x56, 0x85,
	0x10, 0x0c, 0x3a, 0x5d, 0x37, 0x67, 0x74, 0x1f, 0x7a, 0x2c, 0x73, 0x9d, 0xee, 0xb1, 0x4c, 0x47,
	0xae, 0x14, 0x51, 0x34, 0xea, 0x1b, 0x95, 0x15, 0xe2, 0x47, 0x30, 0xb6, 0x6e, 0x7e, 0x70, 0x04,
	0xaf, 0x21, 0xc4, 0xbf, 0xc1, 0xbd, 0xd7, 0x75, 0xa5, 0xca, 0xe2, 0x8c, 0xaa, 0xcb, 0x52, 0x5e,
	0xbc, 0x0f, 0x40, 0x5f, 0xc2, 0x5e, 0x79, 0xc9, 0xa9, 0x5c, 0xf2, 0xba, 0x58, 0x51, 0x69, 0x62,
	0x0d, 0x71, 0x68, 0x74, 0x67, 0x46, 0x85, 0x9e, 0xc1, 0x20, 0x23, 0x8a, 0x98, 0x98, 0xe1, 0x21,
	0x72, 0xb5, 0x38, 0x87, 0x0b, 0x41, 0x53, 0x6c, 0xec, 0xf1, 0x63, 0x18, 0xcd, 0x33, 0xdd, 0x02,
	0x5d, 0x4a, 0x5b, 0x7d, 0xe0, 0x6a, 0xfd, 0xa3, 0x0f, 0x61, 0x87, 0xb3, 0xb3, 0xdc, 0x4f, 0x60,
	0x94, 0x49, 0xb6, 0x71, 0x69, 0x04, 0xd8, 0x49, 0xfa, 0xc6, 0x29, 0x27, 0xab, 0x9c, 0x2e, 0x99,
	0xd8, 0x1c, 0x99, 0x44, 0x7c, 0x0c, 0x56, 0x35, 0x17, 0x9b, 0x23, 0x9d, 0x22, 0x13, 0xa4, 0x88,
	0x06, 0xbb, 0x52, 0x9c, 0xbf, 0x3d, 0xf9, 0x19, 0x1b, 0x3b, 0x9a, 0x82, 0xcf, 0xb8, 0xa2, 0x92,
	0x93, 0x3c, 0x1a, 0x1a, 0x2f, 0xad, 0x8c, 0x8e, 0x61, 0x5c, 0x0a, 0xc5, 0x4a, 0x5e, 0x45, 0x23,
	0x73, 0x6b, 0x4f, 0x6e, 0x56, 0x9a, 0xfc, 0x62, 0x11, 0xa7, 0x5c, 0xc9, 0x2d, 0x6e, 0xf0, 0xe8,
	0x08, 0x46, 0x39, 0x59, 0xd1, 0xbc, 0x8a, 0xc6, 0x86, 0xf9, 0xf9, 0x0e, 0xe6, 0x1b, 0x03, 0xb0,
	0x44, 0x87, 0xd6, 0xe9, 0xd0, 0x2b, 0x97, 0x8e, 0x6f, 0x2a, 0x6e, 0xe5, 0xe9, 0x2b, 0xd8, 0xeb,
	0x06, 0xd3, 0x9f, 0xe1, 0x05, 0xdd, 0xba, 0x76, 0xe9, 0xa3, 0xfe, 0x18, 0x36, 0x24, 0xaf, 0x9b,
	0x97, 0x68, 0x85, 0x57, 0xbd, 0xef, 0xbc, 0xe9, 0x31, 0x84, 0x9d, 0x70, 0x77, 0xa1, 0xc6, 0x7f,
	0x79, 0xed, 0x35, 0xe9, 0xbe, 0x75, 0xae, 0xc4, 0xbb, 0x76, 0x25, 0x9d, 0x6e, 0xf5, 0x76, 0x75,
	0x4b, 0x93, 0x3f, 0xd0, 0xad, 0xe7, 0x30, 0x4a, 0x4b, 0x7e, 0xce, 0xd6, 0x51, 0xdf, 0x30, 0xa3,
	0x9b, 0xcc, 0xd7, 0xc6, 0x8e, 0x1d, 0xee, 0xff, 0xf4, 0x22, 0xfe, 0xdb, 0x83, 0x87, 0x37, 0x3c,
	0xeb, 0xb2, 0xaa, 0x7a, 0xc5, 0x69, 0x33, 0x3e, 0x9c, 0xa4, 0xc7, 0x02, 0x13, 0x4b, 0x49, 0xf8,
	0xba, 0x71, 0x35, 0x66, 0x02, 0x6b, 0x11, 0x45, 0x30, 0x5e, 0x13, 0x45, 0x2f, 0xc9, 0xd6, 0xbd,
	0xbe, 0x46, 0x44, 0x73, 0x08, 0x49, 0x7d, 0xb5, 0x24, 0x59, 0x26, 0x69, 0x55, 0x45, 0x03, 0x53,
	0xd5, 0xc1, 0x87, 0xaa, 0x4a, 0x4e, 0xea, 0xab, 0x13, 0x0b, 0xb5, 0x8d, 0x01, 0xd2, 0x2a, 0xa6,
	0xdf, 0xc3, 0x83, 0xf7, 0xcc, 0x77, 0x2a, 0xf6, 0x4f, 0x0f, 0x86, 0x66, 0x14, 0xdc, 0x6a, 0x9a,
	0x3c, 0x07, 0xdf, 0xed, 0x82, 0xca, 0x5d, 0xc5, 0x7e, 0xd2, 0x2c, 0x87, 0x85, 0xfd, 0x37, 0xcf,
	0xbb, 0x45, 0xa1, 0x04, 0x7c, 0x6e, 0xeb, 0x69, 0xca, 0xdc, 0x35, 0x0e, 0x5a, 0x0c, 0xfa, 0x14,
	0x02, 0x56, 0x2d, 0x45, 0xbd, 0xca, 0x59, 0xda, 0x3e, 0xb8, 0xea, 0xad, 0x91, 0xbf, 0x3e, 0x75,
	0x03, 0x79, 0xa1, 0x87, 0x18, 0x0a, 0xf5, 0x10, 0x2b, 0x85, 0xa0, 0xd9, 0xe4, 0x23, 0xb4, 0x07,
	0xfe, 0x42, 0x11, 0xa9, 0x18, 0x5f, 0x4f, 0x3c, 0x6d, 0xc2, 0x35, 0xe7, 0x5a, 0xe8, 0xa1, 0x07,
	0x10, 0x62, 0x9a, 0x51, 0x91, 0x97, 0x5b, 0xad, 0xe8, 0x1f, 0xfe, 0xe3, 0xb9, 0x8d, 0xe2, 0x52,
	0x46, 0xcf, 0xa0, 0xf7, 0x4e, 0xa0, 0x66, 0xe6, 0xb6, 0xcb, 0x6d, 0x7a, 0xbf, 0xa3, 0xd1, 0x33,
	0x7a, 0x66, 0x7a, 0x25, 0x15, 0xfa, 0xb8, 0x3b, 0x9e, 0x1b, 0xf4, 0xc3, 0xeb, 0x4a, 0x4d, 0x48,
	0x60, 0xa0, 0x53, 0xbc, 0x35, 0xfe, 0x05, 0x8c, 0xb0, 0xd9, 0x5c, 0x68, 0xdf, 0x19, 0xaf, 0x2d,
	0xb2, 0x5d, 0x94, 0x6f, 0x60, 0x60, 0x27, 0xa8, 0x33, 0x75, 0x16, 0xd1, 0x74, 0x72, 0x4d, 0x27,
	0xf2, 0xed, 0x6a, 0x64, 0x56, 0xf8, 0xb7, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x46, 0xce, 0x91,
	0x57, 0x3d, 0x08, 0x00, 0x00,
}
