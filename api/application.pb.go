// Code generated by protoc-gen-go.
// source: application.proto
// DO NOT EDIT!

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CreateApplicationRequest struct {
	// Name of the application (must be unique).
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Description of the application.
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
}

func (m *CreateApplicationRequest) Reset()                    { *m = CreateApplicationRequest{} }
func (m *CreateApplicationRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateApplicationRequest) ProtoMessage()               {}
func (*CreateApplicationRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *CreateApplicationRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateApplicationRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type CreateApplicationResponse struct {
	// ID of the application that was created.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *CreateApplicationResponse) Reset()                    { *m = CreateApplicationResponse{} }
func (m *CreateApplicationResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateApplicationResponse) ProtoMessage()               {}
func (*CreateApplicationResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *CreateApplicationResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetApplicationRequest struct {
	// Name of the application.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetApplicationRequest) Reset()                    { *m = GetApplicationRequest{} }
func (m *GetApplicationRequest) String() string            { return proto.CompactTextString(m) }
func (*GetApplicationRequest) ProtoMessage()               {}
func (*GetApplicationRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *GetApplicationRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetApplicationResponse struct {
	// ID of the application.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// Name of the application.
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// Description of the application.
	Description string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
}

func (m *GetApplicationResponse) Reset()                    { *m = GetApplicationResponse{} }
func (m *GetApplicationResponse) String() string            { return proto.CompactTextString(m) }
func (*GetApplicationResponse) ProtoMessage()               {}
func (*GetApplicationResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

func (m *GetApplicationResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetApplicationResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetApplicationResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type UpdateApplicationRequest struct {
	// ID of the application to update.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// Name of the application (must be unique).
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// Description of the application.
	Description string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
}

func (m *UpdateApplicationRequest) Reset()                    { *m = UpdateApplicationRequest{} }
func (m *UpdateApplicationRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateApplicationRequest) ProtoMessage()               {}
func (*UpdateApplicationRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{4} }

func (m *UpdateApplicationRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateApplicationRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateApplicationRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type UpdateApplicationResponse struct {
}

func (m *UpdateApplicationResponse) Reset()                    { *m = UpdateApplicationResponse{} }
func (m *UpdateApplicationResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateApplicationResponse) ProtoMessage()               {}
func (*UpdateApplicationResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{5} }

type DeleteApplicationRequest struct {
	// ID of the application.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteApplicationRequest) Reset()                    { *m = DeleteApplicationRequest{} }
func (m *DeleteApplicationRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteApplicationRequest) ProtoMessage()               {}
func (*DeleteApplicationRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{6} }

func (m *DeleteApplicationRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteApplicationResponse struct {
}

func (m *DeleteApplicationResponse) Reset()                    { *m = DeleteApplicationResponse{} }
func (m *DeleteApplicationResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteApplicationResponse) ProtoMessage()               {}
func (*DeleteApplicationResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{7} }

type ListApplicationRequest struct {
	// Max number of applications to return in the result-test.
	Limit int64 `protobuf:"varint,1,opt,name=limit" json:"limit,omitempty"`
	// Offset in the result-set (for pagination).
	Offset int64 `protobuf:"varint,2,opt,name=offset" json:"offset,omitempty"`
}

func (m *ListApplicationRequest) Reset()                    { *m = ListApplicationRequest{} }
func (m *ListApplicationRequest) String() string            { return proto.CompactTextString(m) }
func (*ListApplicationRequest) ProtoMessage()               {}
func (*ListApplicationRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{8} }

func (m *ListApplicationRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListApplicationRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type ListApplicationResponse struct {
	// Total number of applications available within the result-set.
	TotalCount int64 `protobuf:"varint,1,opt,name=totalCount" json:"totalCount,omitempty"`
	// Applications within this result-set.
	Result []*GetApplicationResponse `protobuf:"bytes,2,rep,name=result" json:"result,omitempty"`
}

func (m *ListApplicationResponse) Reset()                    { *m = ListApplicationResponse{} }
func (m *ListApplicationResponse) String() string            { return proto.CompactTextString(m) }
func (*ListApplicationResponse) ProtoMessage()               {}
func (*ListApplicationResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{9} }

func (m *ListApplicationResponse) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *ListApplicationResponse) GetResult() []*GetApplicationResponse {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateApplicationRequest)(nil), "api.CreateApplicationRequest")
	proto.RegisterType((*CreateApplicationResponse)(nil), "api.CreateApplicationResponse")
	proto.RegisterType((*GetApplicationRequest)(nil), "api.GetApplicationRequest")
	proto.RegisterType((*GetApplicationResponse)(nil), "api.GetApplicationResponse")
	proto.RegisterType((*UpdateApplicationRequest)(nil), "api.UpdateApplicationRequest")
	proto.RegisterType((*UpdateApplicationResponse)(nil), "api.UpdateApplicationResponse")
	proto.RegisterType((*DeleteApplicationRequest)(nil), "api.DeleteApplicationRequest")
	proto.RegisterType((*DeleteApplicationResponse)(nil), "api.DeleteApplicationResponse")
	proto.RegisterType((*ListApplicationRequest)(nil), "api.ListApplicationRequest")
	proto.RegisterType((*ListApplicationResponse)(nil), "api.ListApplicationResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Application service

type ApplicationClient interface {
	// Create creates the given application.
	Create(ctx context.Context, in *CreateApplicationRequest, opts ...grpc.CallOption) (*CreateApplicationResponse, error)
	// Get returns the requested application.
	Get(ctx context.Context, in *GetApplicationRequest, opts ...grpc.CallOption) (*GetApplicationResponse, error)
	// Update updates the given application.
	Update(ctx context.Context, in *UpdateApplicationRequest, opts ...grpc.CallOption) (*UpdateApplicationResponse, error)
	// Delete deletes the given application.
	Delete(ctx context.Context, in *DeleteApplicationRequest, opts ...grpc.CallOption) (*DeleteApplicationResponse, error)
	// List lists the available applications.
	List(ctx context.Context, in *ListApplicationRequest, opts ...grpc.CallOption) (*ListApplicationResponse, error)
}

type applicationClient struct {
	cc *grpc.ClientConn
}

func NewApplicationClient(cc *grpc.ClientConn) ApplicationClient {
	return &applicationClient{cc}
}

func (c *applicationClient) Create(ctx context.Context, in *CreateApplicationRequest, opts ...grpc.CallOption) (*CreateApplicationResponse, error) {
	out := new(CreateApplicationResponse)
	err := grpc.Invoke(ctx, "/api.Application/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationClient) Get(ctx context.Context, in *GetApplicationRequest, opts ...grpc.CallOption) (*GetApplicationResponse, error) {
	out := new(GetApplicationResponse)
	err := grpc.Invoke(ctx, "/api.Application/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationClient) Update(ctx context.Context, in *UpdateApplicationRequest, opts ...grpc.CallOption) (*UpdateApplicationResponse, error) {
	out := new(UpdateApplicationResponse)
	err := grpc.Invoke(ctx, "/api.Application/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationClient) Delete(ctx context.Context, in *DeleteApplicationRequest, opts ...grpc.CallOption) (*DeleteApplicationResponse, error) {
	out := new(DeleteApplicationResponse)
	err := grpc.Invoke(ctx, "/api.Application/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationClient) List(ctx context.Context, in *ListApplicationRequest, opts ...grpc.CallOption) (*ListApplicationResponse, error) {
	out := new(ListApplicationResponse)
	err := grpc.Invoke(ctx, "/api.Application/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Application service

type ApplicationServer interface {
	// Create creates the given application.
	Create(context.Context, *CreateApplicationRequest) (*CreateApplicationResponse, error)
	// Get returns the requested application.
	Get(context.Context, *GetApplicationRequest) (*GetApplicationResponse, error)
	// Update updates the given application.
	Update(context.Context, *UpdateApplicationRequest) (*UpdateApplicationResponse, error)
	// Delete deletes the given application.
	Delete(context.Context, *DeleteApplicationRequest) (*DeleteApplicationResponse, error)
	// List lists the available applications.
	List(context.Context, *ListApplicationRequest) (*ListApplicationResponse, error)
}

func RegisterApplicationServer(s *grpc.Server, srv ApplicationServer) {
	s.RegisterService(&_Application_serviceDesc, srv)
}

func _Application_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Application/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServer).Create(ctx, req.(*CreateApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Application_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Application/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServer).Get(ctx, req.(*GetApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Application_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Application/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServer).Update(ctx, req.(*UpdateApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Application_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Application/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServer).Delete(ctx, req.(*DeleteApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Application_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Application/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServer).List(ctx, req.(*ListApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Application_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Application",
	HandlerType: (*ApplicationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Application_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Application_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Application_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Application_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Application_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "application.proto",
}

func init() { proto.RegisterFile("application.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 435 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x41, 0x6f, 0xda, 0x30,
	0x14, 0x16, 0x09, 0x8b, 0xb4, 0x87, 0x34, 0x09, 0x6f, 0xcb, 0x42, 0x60, 0x88, 0xe5, 0x32, 0xc4,
	0x24, 0x90, 0xe0, 0xb6, 0xdb, 0xc4, 0x34, 0x2e, 0x3b, 0x54, 0x91, 0x7a, 0xab, 0x50, 0x5d, 0x62,
	0x90, 0xab, 0x10, 0xbb, 0xb1, 0x39, 0x55, 0xbd, 0xf4, 0xd2, 0x1f, 0xd0, 0x9f, 0xd6, 0xbf, 0xd0,
	0x1f, 0x52, 0xc5, 0x76, 0x0b, 0x22, 0x76, 0xd5, 0x43, 0x6f, 0xd8, 0xef, 0xf3, 0xf7, 0xbd, 0xef,
	0xbd, 0x8f, 0x40, 0x1b, 0x73, 0x9e, 0xd3, 0x15, 0x96, 0x94, 0x15, 0x63, 0x5e, 0x32, 0xc9, 0x90,
	0x8f, 0x39, 0x8d, 0x7b, 0x1b, 0xc6, 0x36, 0x39, 0x99, 0x60, 0x4e, 0x27, 0xb8, 0x28, 0x98, 0x54,
	0x08, 0xa1, 0x21, 0xc9, 0x09, 0x44, 0xf3, 0x92, 0x60, 0x49, 0xfe, 0xec, 0x5f, 0xa7, 0xe4, 0x6a,
	0x47, 0x84, 0x44, 0x08, 0x9a, 0x05, 0xde, 0x92, 0xa8, 0x31, 0x68, 0x0c, 0x3f, 0xa6, 0xea, 0x37,
	0x1a, 0x40, 0x2b, 0x23, 0x62, 0x55, 0x52, 0x5e, 0x21, 0x23, 0x4f, 0x95, 0x0e, 0xaf, 0x92, 0x5f,
	0xd0, 0xb1, 0x30, 0x0a, 0xce, 0x0a, 0x41, 0xd0, 0x27, 0xf0, 0x68, 0xa6, 0x08, 0xfd, 0xd4, 0xa3,
	0x59, 0xf2, 0x13, 0xbe, 0x2e, 0x88, 0xb4, 0x68, 0x1f, 0x03, 0x97, 0x10, 0x1e, 0x03, 0xed, 0x94,
	0x2f, 0x5d, 0x7b, 0xee, 0xae, 0xfd, 0x7a, 0xd7, 0xe7, 0x10, 0x9d, 0xf2, 0xcc, 0x3e, 0x87, 0xf7,
	0x51, 0xe8, 0x42, 0xc7, 0xa2, 0xa0, 0x4d, 0x24, 0x23, 0x88, 0xfe, 0x92, 0x9c, 0xbc, 0x45, 0xbe,
	0x22, 0xb2, 0x60, 0x0d, 0xd1, 0x3f, 0x08, 0xff, 0x53, 0x61, 0x9b, 0xe8, 0x17, 0xf8, 0x90, 0xd3,
	0x2d, 0x95, 0x86, 0x49, 0x1f, 0x50, 0x08, 0x01, 0x5b, 0xaf, 0x05, 0x91, 0xca, 0x8d, 0x9f, 0x9a,
	0x53, 0x52, 0xc0, 0xb7, 0x1a, 0x8f, 0x19, 0x78, 0x1f, 0x40, 0x32, 0x89, 0xf3, 0x39, 0xdb, 0x15,
	0xcf, 0x6c, 0x07, 0x37, 0x68, 0x06, 0x41, 0x49, 0xc4, 0x2e, 0xaf, 0x28, 0xfd, 0x61, 0x6b, 0xda,
	0x1d, 0x63, 0x4e, 0xc7, 0xf6, 0xed, 0xa5, 0x06, 0x3a, 0xbd, 0x6b, 0x42, 0xeb, 0xa0, 0x8e, 0x08,
	0x04, 0x3a, 0x45, 0xe8, 0xbb, 0x7a, 0xee, 0x0a, 0x69, 0xdc, 0x77, 0x95, 0xcd, 0x40, 0x7a, 0xb7,
	0x0f, 0x8f, 0xf7, 0x5e, 0x98, 0xb4, 0xf5, 0x1f, 0x60, 0x8f, 0x10, 0xbf, 0x1b, 0x23, 0xb4, 0x04,
	0x7f, 0x41, 0x24, 0x8a, 0xad, 0x2d, 0x6a, 0x81, 0xd7, 0xda, 0x4f, 0xfa, 0x8a, 0x3d, 0x42, 0x61,
	0x8d, 0x7d, 0x72, 0x4d, 0xb3, 0x1b, 0x74, 0x09, 0x81, 0x5e, 0xba, 0xb1, 0xe1, 0xca, 0x98, 0xb1,
	0xe1, 0x0e, 0xc8, 0x0f, 0x25, 0xd4, 0x8d, 0x1d, 0x42, 0x95, 0x97, 0x0d, 0x04, 0x3a, 0x17, 0x46,
	0xcb, 0x15, 0x28, 0xa3, 0xe5, 0xce, 0x90, 0x31, 0x35, 0x72, 0x99, 0x3a, 0x83, 0x66, 0x95, 0x0d,
	0xa4, 0x27, 0x63, 0x8f, 0x5b, 0xdc, 0xb3, 0x17, 0x8d, 0x44, 0x47, 0x49, 0x7c, 0x46, 0xf5, 0xad,
	0x5c, 0x04, 0xea, 0xc3, 0x34, 0x7b, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xce, 0x78, 0x5f, 0x6f, 0xd0,
	0x04, 0x00, 0x00,
}
