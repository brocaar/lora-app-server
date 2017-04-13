// Code generated by protoc-gen-go.
// source: channelList.proto
// DO NOT EDIT!

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	channelList.proto
	node.proto
	application.proto
	downlinkQueue.proto
	common.proto
	user.proto
	gateway.proto
	organization.proto

It has these top-level messages:
	CreateChannelListRequest
	CreateChannelListResponse
	UpdateChannelListRequest
	UpdateChannelListResponse
	GetChannelListRequest
	GetChannelListResponse
	ListChannelListRequest
	ListChannelListResponse
	DeleteChannelListRequest
	DeleteChannelListResponse
	CreateNodeRequest
	CreateNodeResponse
	GetNodeRequest
	GetNodeResponse
	DeleteNodeRequest
	DeleteNodeResponse
	ListNodeByApplicationIDRequest
	ListNodeResponse
	UpdateNodeRequest
	UpdateNodeResponse
	ActivateNodeRequest
	ActivateNodeResponse
	GetNodeActivationRequest
	GetNodeActivationResponse
	GetRandomDevAddrRequest
	GetRandomDevAddrResponse
	CreateApplicationRequest
	CreateApplicationResponse
	GetApplicationRequest
	GetApplicationResponse
	UpdateApplicationRequest
	UpdateApplicationResponse
	DeleteApplicationRequest
	DeleteApplicationResponse
	ListApplicationRequest
	ListApplicationResponse
	ListApplicationUsersRequest
	GetApplicationUserResponse
	ListApplicationUsersResponse
	AddApplicationUserRequest
	ApplicationUserRequest
	UpdateApplicationUserRequest
	EmptyApplicationUserResponse
	EnqueueDownlinkQueueItemRequest
	EnqueueDownlinkQueueItemResponse
	DeleteDownlinkQeueueItemRequest
	DeleteDownlinkQueueItemResponse
	DownlinkQueueItem
	ListDownlinkQueueItemsRequest
	ListDownlinkQueueItemsResponse
	ApplicationLink
	UserProfile
	ProfileRequest
	ProfileResponse
	LoginRequest
	LoginResponse
	ListUserRequest
	UserRequest
	AddUserResponse
	UserSettings
	UserInfo
	GetUserResponse
	AddUserRequest
	UpdateUserRequest
	ListUserResponse
	UserEmptyResponse
	UpdateUserPasswordRequest
	CreateGatewayRequest
	CreateGatewayResponse
	GetGatewayRequest
	GetGatewayResponse
	DeleteGatewayRequest
	DeleteGatewayResponse
	ListGatewayRequest
	ListGatewayItem
	ListGatewayResponse
	UpdateGatewayRequest
	UpdateGatewayResponse
	GatewayStats
	GetGatewayStatsRequest
	GetGatewayStatsResponse
	ListOrganizationRequest
	OrganizationRequest
	GetOrganizationResponse
	CreateOrganizationRequest
	CreateOrganizationResponse
	UpdateOrganizationRequest
	ListOrganizationResponse
	OrganizationEmptyResponse
	OrganizationUserRequest
	DeleteOrganizationUserRequest
	ListOrganizationUsersRequest
	GetOrganizationUserRequest
	GetOrganizationUserResponse
	ListOrganizationUsersResponse
*/
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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateChannelListRequest struct {
	Name     string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Channels []uint32 `protobuf:"varint,2,rep,packed,name=channels" json:"channels,omitempty"`
}

func (m *CreateChannelListRequest) Reset()                    { *m = CreateChannelListRequest{} }
func (m *CreateChannelListRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateChannelListRequest) ProtoMessage()               {}
func (*CreateChannelListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateChannelListRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateChannelListRequest) GetChannels() []uint32 {
	if m != nil {
		return m.Channels
	}
	return nil
}

type CreateChannelListResponse struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *CreateChannelListResponse) Reset()                    { *m = CreateChannelListResponse{} }
func (m *CreateChannelListResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateChannelListResponse) ProtoMessage()               {}
func (*CreateChannelListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateChannelListResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type UpdateChannelListRequest struct {
	Id       int64    `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name     string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Channels []uint32 `protobuf:"varint,3,rep,packed,name=channels" json:"channels,omitempty"`
}

func (m *UpdateChannelListRequest) Reset()                    { *m = UpdateChannelListRequest{} }
func (m *UpdateChannelListRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateChannelListRequest) ProtoMessage()               {}
func (*UpdateChannelListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UpdateChannelListRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateChannelListRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateChannelListRequest) GetChannels() []uint32 {
	if m != nil {
		return m.Channels
	}
	return nil
}

type UpdateChannelListResponse struct {
}

func (m *UpdateChannelListResponse) Reset()                    { *m = UpdateChannelListResponse{} }
func (m *UpdateChannelListResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateChannelListResponse) ProtoMessage()               {}
func (*UpdateChannelListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type GetChannelListRequest struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetChannelListRequest) Reset()                    { *m = GetChannelListRequest{} }
func (m *GetChannelListRequest) String() string            { return proto.CompactTextString(m) }
func (*GetChannelListRequest) ProtoMessage()               {}
func (*GetChannelListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetChannelListRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetChannelListResponse struct {
	Id       int64    `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name     string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Channels []uint32 `protobuf:"varint,3,rep,packed,name=channels" json:"channels,omitempty"`
}

func (m *GetChannelListResponse) Reset()                    { *m = GetChannelListResponse{} }
func (m *GetChannelListResponse) String() string            { return proto.CompactTextString(m) }
func (*GetChannelListResponse) ProtoMessage()               {}
func (*GetChannelListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GetChannelListResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetChannelListResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetChannelListResponse) GetChannels() []uint32 {
	if m != nil {
		return m.Channels
	}
	return nil
}

type ListChannelListRequest struct {
	Limit  int64 `protobuf:"varint,1,opt,name=limit" json:"limit,omitempty"`
	Offset int64 `protobuf:"varint,2,opt,name=offset" json:"offset,omitempty"`
}

func (m *ListChannelListRequest) Reset()                    { *m = ListChannelListRequest{} }
func (m *ListChannelListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListChannelListRequest) ProtoMessage()               {}
func (*ListChannelListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ListChannelListRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListChannelListRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type ListChannelListResponse struct {
	TotalCount int64                     `protobuf:"varint,1,opt,name=totalCount" json:"totalCount,omitempty"`
	Result     []*GetChannelListResponse `protobuf:"bytes,2,rep,name=result" json:"result,omitempty"`
}

func (m *ListChannelListResponse) Reset()                    { *m = ListChannelListResponse{} }
func (m *ListChannelListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListChannelListResponse) ProtoMessage()               {}
func (*ListChannelListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ListChannelListResponse) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *ListChannelListResponse) GetResult() []*GetChannelListResponse {
	if m != nil {
		return m.Result
	}
	return nil
}

type DeleteChannelListRequest struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteChannelListRequest) Reset()                    { *m = DeleteChannelListRequest{} }
func (m *DeleteChannelListRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteChannelListRequest) ProtoMessage()               {}
func (*DeleteChannelListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *DeleteChannelListRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteChannelListResponse struct {
}

func (m *DeleteChannelListResponse) Reset()                    { *m = DeleteChannelListResponse{} }
func (m *DeleteChannelListResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteChannelListResponse) ProtoMessage()               {}
func (*DeleteChannelListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func init() {
	proto.RegisterType((*CreateChannelListRequest)(nil), "api.CreateChannelListRequest")
	proto.RegisterType((*CreateChannelListResponse)(nil), "api.CreateChannelListResponse")
	proto.RegisterType((*UpdateChannelListRequest)(nil), "api.UpdateChannelListRequest")
	proto.RegisterType((*UpdateChannelListResponse)(nil), "api.UpdateChannelListResponse")
	proto.RegisterType((*GetChannelListRequest)(nil), "api.GetChannelListRequest")
	proto.RegisterType((*GetChannelListResponse)(nil), "api.GetChannelListResponse")
	proto.RegisterType((*ListChannelListRequest)(nil), "api.ListChannelListRequest")
	proto.RegisterType((*ListChannelListResponse)(nil), "api.ListChannelListResponse")
	proto.RegisterType((*DeleteChannelListRequest)(nil), "api.DeleteChannelListRequest")
	proto.RegisterType((*DeleteChannelListResponse)(nil), "api.DeleteChannelListResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ChannelList service

type ChannelListClient interface {
	// Create creates the given channel-list.
	Create(ctx context.Context, in *CreateChannelListRequest, opts ...grpc.CallOption) (*CreateChannelListResponse, error)
	// Update updates the given channel-list.
	Update(ctx context.Context, in *UpdateChannelListRequest, opts ...grpc.CallOption) (*UpdateChannelListResponse, error)
	// Get returns the channel-list matching the given id.
	Get(ctx context.Context, in *GetChannelListRequest, opts ...grpc.CallOption) (*GetChannelListResponse, error)
	// List lists the channel-lists given an offset and limit.
	List(ctx context.Context, in *ListChannelListRequest, opts ...grpc.CallOption) (*ListChannelListResponse, error)
	// Delete deletes the channel-list matching the given id.
	Delete(ctx context.Context, in *DeleteChannelListRequest, opts ...grpc.CallOption) (*DeleteChannelListResponse, error)
}

type channelListClient struct {
	cc *grpc.ClientConn
}

func NewChannelListClient(cc *grpc.ClientConn) ChannelListClient {
	return &channelListClient{cc}
}

func (c *channelListClient) Create(ctx context.Context, in *CreateChannelListRequest, opts ...grpc.CallOption) (*CreateChannelListResponse, error) {
	out := new(CreateChannelListResponse)
	err := grpc.Invoke(ctx, "/api.ChannelList/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelListClient) Update(ctx context.Context, in *UpdateChannelListRequest, opts ...grpc.CallOption) (*UpdateChannelListResponse, error) {
	out := new(UpdateChannelListResponse)
	err := grpc.Invoke(ctx, "/api.ChannelList/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelListClient) Get(ctx context.Context, in *GetChannelListRequest, opts ...grpc.CallOption) (*GetChannelListResponse, error) {
	out := new(GetChannelListResponse)
	err := grpc.Invoke(ctx, "/api.ChannelList/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelListClient) List(ctx context.Context, in *ListChannelListRequest, opts ...grpc.CallOption) (*ListChannelListResponse, error) {
	out := new(ListChannelListResponse)
	err := grpc.Invoke(ctx, "/api.ChannelList/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelListClient) Delete(ctx context.Context, in *DeleteChannelListRequest, opts ...grpc.CallOption) (*DeleteChannelListResponse, error) {
	out := new(DeleteChannelListResponse)
	err := grpc.Invoke(ctx, "/api.ChannelList/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ChannelList service

type ChannelListServer interface {
	// Create creates the given channel-list.
	Create(context.Context, *CreateChannelListRequest) (*CreateChannelListResponse, error)
	// Update updates the given channel-list.
	Update(context.Context, *UpdateChannelListRequest) (*UpdateChannelListResponse, error)
	// Get returns the channel-list matching the given id.
	Get(context.Context, *GetChannelListRequest) (*GetChannelListResponse, error)
	// List lists the channel-lists given an offset and limit.
	List(context.Context, *ListChannelListRequest) (*ListChannelListResponse, error)
	// Delete deletes the channel-list matching the given id.
	Delete(context.Context, *DeleteChannelListRequest) (*DeleteChannelListResponse, error)
}

func RegisterChannelListServer(s *grpc.Server, srv ChannelListServer) {
	s.RegisterService(&_ChannelList_serviceDesc, srv)
}

func _ChannelList_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChannelListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelListServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ChannelList/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelListServer).Create(ctx, req.(*CreateChannelListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChannelList_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateChannelListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelListServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ChannelList/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelListServer).Update(ctx, req.(*UpdateChannelListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChannelList_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChannelListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelListServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ChannelList/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelListServer).Get(ctx, req.(*GetChannelListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChannelList_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListChannelListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelListServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ChannelList/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelListServer).List(ctx, req.(*ListChannelListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChannelList_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteChannelListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelListServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ChannelList/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelListServer).Delete(ctx, req.(*DeleteChannelListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChannelList_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.ChannelList",
	HandlerType: (*ChannelListServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ChannelList_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ChannelList_Update_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ChannelList_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ChannelList_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ChannelList_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "channelList.proto",
}

func init() { proto.RegisterFile("channelList.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x94, 0xcf, 0xaa, 0x1a, 0x31,
	0x14, 0xc6, 0x71, 0xc6, 0x0e, 0xed, 0x91, 0x96, 0xf6, 0x50, 0x75, 0x1c, 0xff, 0x20, 0xd9, 0x54,
	0x2c, 0x28, 0xe8, 0xae, 0x5b, 0x4b, 0x85, 0xd2, 0xd5, 0x40, 0xa1, 0xb4, 0x74, 0x91, 0x76, 0xa2,
	0x0d, 0x8c, 0xc9, 0xd4, 0xc4, 0x55, 0xe9, 0xc6, 0x57, 0xb8, 0x8f, 0x76, 0x5f, 0xe1, 0x3e, 0xc8,
	0xc5, 0x24, 0x7a, 0x65, 0xcc, 0xdc, 0x3f, 0x3b, 0x93, 0x1c, 0xbf, 0xdf, 0x77, 0xce, 0xf9, 0x14,
	0xde, 0xfc, 0xfe, 0x43, 0x85, 0x60, 0xf9, 0x17, 0xae, 0xf4, 0xa4, 0xd8, 0x4a, 0x2d, 0x31, 0xa4,
	0x05, 0x4f, 0x7a, 0x6b, 0x29, 0xd7, 0x39, 0x9b, 0xd2, 0x82, 0x4f, 0xa9, 0x10, 0x52, 0x53, 0xcd,
	0xa5, 0x50, 0xb6, 0x84, 0x7c, 0x86, 0x78, 0xb1, 0x65, 0x54, 0xb3, 0xc5, 0xdd, 0xb7, 0x53, 0xf6,
	0x77, 0xc7, 0x94, 0x46, 0x84, 0xba, 0xa0, 0x1b, 0x16, 0xd7, 0x86, 0xb5, 0xd1, 0x8b, 0xd4, 0x7c,
	0xc6, 0x04, 0x9e, 0x3b, 0x8e, 0x8a, 0x83, 0x61, 0x38, 0x7a, 0x99, 0x9e, 0xce, 0xe4, 0x3d, 0x74,
	0x3c, 0x5a, 0xaa, 0x90, 0x42, 0x31, 0x7c, 0x05, 0x01, 0xcf, 0x8c, 0x54, 0x98, 0x06, 0x3c, 0x23,
	0xdf, 0x21, 0xfe, 0x5a, 0x64, 0x7e, 0x70, 0xa9, 0xf6, 0x64, 0x24, 0xa8, 0x30, 0x12, 0x96, 0x8c,
	0x74, 0xa1, 0xe3, 0xd1, 0xb6, 0x46, 0xc8, 0x3b, 0x68, 0x2e, 0x99, 0x7e, 0x98, 0x4a, 0xbe, 0x41,
	0xab, 0x5c, 0xe8, 0xef, 0xe5, 0xc9, 0xfe, 0x3e, 0x41, 0xeb, 0xa0, 0xe7, 0xf1, 0xf0, 0x16, 0x9e,
	0xe5, 0x7c, 0xc3, 0xb5, 0x13, 0xb7, 0x07, 0x6c, 0x41, 0x24, 0x57, 0x2b, 0xc5, 0xb4, 0x21, 0x84,
	0xa9, 0x3b, 0x11, 0x01, 0xed, 0x0b, 0x1d, 0x67, 0x71, 0x00, 0xa0, 0xa5, 0xa6, 0xf9, 0x42, 0xee,
	0xc4, 0x51, 0xed, 0xec, 0x06, 0xe7, 0x10, 0x6d, 0x99, 0xda, 0xe5, 0xda, 0x6c, 0xb1, 0x31, 0xeb,
	0x4e, 0x68, 0xc1, 0x27, 0xfe, 0x7e, 0x53, 0x57, 0x4a, 0xc6, 0x10, 0x7f, 0x64, 0x39, 0x7b, 0xcc,
	0xce, 0x0e, 0x3b, 0xf0, 0xd4, 0x5a, 0xc1, 0xd9, 0xbe, 0x0e, 0x8d, 0xb3, 0x7b, 0xcc, 0x20, 0xb2,
	0xc9, 0xc1, 0xbe, 0xf1, 0x51, 0x15, 0xc9, 0x64, 0x50, 0xf5, 0xec, 0x96, 0xdb, 0xdd, 0x5f, 0xdf,
	0x5c, 0x05, 0x4d, 0xf2, 0xda, 0xc4, 0xfd, 0xec, 0x17, 0xf1, 0xa1, 0x36, 0x46, 0x0e, 0x91, 0x8d,
	0x85, 0xa3, 0x54, 0xe5, 0xcf, 0x51, 0xaa, 0x23, 0x34, 0x34, 0x94, 0x24, 0x69, 0x96, 0x29, 0xd3,
	0x7f, 0x3c, 0xfb, 0x7f, 0x40, 0xfd, 0x84, 0x70, 0xc9, 0x34, 0x26, 0xde, 0xa9, 0x5a, 0xc8, 0x7d,
	0x13, 0x27, 0x7d, 0x43, 0x68, 0xa3, 0x9f, 0x80, 0x3f, 0xa0, 0x6e, 0xe6, 0x66, 0x35, 0xfc, 0x59,
	0x4a, 0x7a, 0xfe, 0x47, 0x47, 0x88, 0x0d, 0x01, 0xf1, 0x62, 0x52, 0xb8, 0x82, 0xc8, 0x6e, 0xce,
	0x8d, 0xa9, 0x6a, 0xe5, 0x6e, 0x4c, 0x95, 0x5b, 0x3e, 0x36, 0x31, 0xf6, 0x37, 0xf1, 0x2b, 0x32,
	0xff, 0x40, 0xf3, 0xdb, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa2, 0xf7, 0xf2, 0x08, 0xb9, 0x04, 0x00,
	0x00,
}
