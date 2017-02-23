// Code generated by protoc-gen-go.
// source: node.proto
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

type CreateNodeRequest struct {
	// Hex encoded DevEUI.
	DevEUI string `protobuf:"bytes,1,opt,name=devEUI" json:"devEUI,omitempty"`
	// Hex encoded AppEUI.
	AppEUI string `protobuf:"bytes,2,opt,name=appEUI" json:"appEUI,omitempty"`
	// Hex encoded AppKey. When isABP is set to true, this field is not needed.
	AppKey string `protobuf:"bytes,3,opt,name=appKey" json:"appKey,omitempty"`
	// RX delay.
	RxDelay uint32 `protobuf:"varint,4,opt,name=rxDelay" json:"rxDelay,omitempty"`
	// RX1 data-rate offset.
	Rx1DROffset uint32 `protobuf:"varint,5,opt,name=rx1DROffset" json:"rx1DROffset,omitempty"`
	// Channel-list ID used for CFlist (see LoRaWAN regional parameters if this applies to your region).
	ChannelListID int64 `protobuf:"varint,6,opt,name=channelListID" json:"channelListID,omitempty"`
	// RX window to use.
	RxWindow RXWindow `protobuf:"varint,7,opt,name=rxWindow,enum=api.RXWindow" json:"rxWindow,omitempty"`
	// Data-rate to use for RX2.
	Rx2DR uint32 `protobuf:"varint,8,opt,name=rx2DR" json:"rx2DR,omitempty"`
	// Name of the node (if left blank, it will be set to the DevEUI).
	Name string `protobuf:"bytes,9,opt,name=name" json:"name,omitempty"`
	// Relax frame-counter mode is enabled.
	RelaxFCnt bool `protobuf:"varint,10,opt,name=relaxFCnt" json:"relaxFCnt,omitempty"`
	// Interval (in frames) in which the ADR engine may adapt the data-rate of the node (0 = disabled).
	AdrInterval uint32 `protobuf:"varint,11,opt,name=adrInterval" json:"adrInterval,omitempty"`
	// Installation-margin to use for ADR calculation.
	InstallationMargin float64 `protobuf:"fixed64,12,opt,name=installationMargin" json:"installationMargin,omitempty"`
	// ID of the application to which the node must be added.
	ApplicationID int64 `protobuf:"varint,13,opt,name=applicationID" json:"applicationID,omitempty"`
	// Description of the node.
	Description string `protobuf:"bytes,14,opt,name=description" json:"description,omitempty"`
	// Node is activated by ABP.
	IsABP bool `protobuf:"varint,15,opt,name=isABP" json:"isABP,omitempty"`
	// Node operates in Class-C.
	IsClassC bool `protobuf:"varint,16,opt,name=isClassC" json:"isClassC,omitempty"`
}

func (m *CreateNodeRequest) Reset()                    { *m = CreateNodeRequest{} }
func (m *CreateNodeRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateNodeRequest) ProtoMessage()               {}
func (*CreateNodeRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *CreateNodeRequest) GetDevEUI() string {
	if m != nil {
		return m.DevEUI
	}
	return ""
}

func (m *CreateNodeRequest) GetAppEUI() string {
	if m != nil {
		return m.AppEUI
	}
	return ""
}

func (m *CreateNodeRequest) GetAppKey() string {
	if m != nil {
		return m.AppKey
	}
	return ""
}

func (m *CreateNodeRequest) GetRxDelay() uint32 {
	if m != nil {
		return m.RxDelay
	}
	return 0
}

func (m *CreateNodeRequest) GetRx1DROffset() uint32 {
	if m != nil {
		return m.Rx1DROffset
	}
	return 0
}

func (m *CreateNodeRequest) GetChannelListID() int64 {
	if m != nil {
		return m.ChannelListID
	}
	return 0
}

func (m *CreateNodeRequest) GetRxWindow() RXWindow {
	if m != nil {
		return m.RxWindow
	}
	return RXWindow_RX1
}

func (m *CreateNodeRequest) GetRx2DR() uint32 {
	if m != nil {
		return m.Rx2DR
	}
	return 0
}

func (m *CreateNodeRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateNodeRequest) GetRelaxFCnt() bool {
	if m != nil {
		return m.RelaxFCnt
	}
	return false
}

func (m *CreateNodeRequest) GetAdrInterval() uint32 {
	if m != nil {
		return m.AdrInterval
	}
	return 0
}

func (m *CreateNodeRequest) GetInstallationMargin() float64 {
	if m != nil {
		return m.InstallationMargin
	}
	return 0
}

func (m *CreateNodeRequest) GetApplicationID() int64 {
	if m != nil {
		return m.ApplicationID
	}
	return 0
}

func (m *CreateNodeRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CreateNodeRequest) GetIsABP() bool {
	if m != nil {
		return m.IsABP
	}
	return false
}

func (m *CreateNodeRequest) GetIsClassC() bool {
	if m != nil {
		return m.IsClassC
	}
	return false
}

type CreateNodeResponse struct {
}

func (m *CreateNodeResponse) Reset()                    { *m = CreateNodeResponse{} }
func (m *CreateNodeResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateNodeResponse) ProtoMessage()               {}
func (*CreateNodeResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type GetNodeRequest struct {
	// Hex encoded DevEUI of the node.
	DevEUI string `protobuf:"bytes,1,opt,name=devEUI" json:"devEUI,omitempty"`
}

func (m *GetNodeRequest) Reset()                    { *m = GetNodeRequest{} }
func (m *GetNodeRequest) String() string            { return proto.CompactTextString(m) }
func (*GetNodeRequest) ProtoMessage()               {}
func (*GetNodeRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *GetNodeRequest) GetDevEUI() string {
	if m != nil {
		return m.DevEUI
	}
	return ""
}

type GetNodeResponse struct {
	// Hex encoded DevEUI.
	DevEUI string `protobuf:"bytes,1,opt,name=devEUI" json:"devEUI,omitempty"`
	// Hex encoded AppEUI.
	AppEUI string `protobuf:"bytes,2,opt,name=appEUI" json:"appEUI,omitempty"`
	// Hex encoded AppKey.
	AppKey string `protobuf:"bytes,3,opt,name=appKey" json:"appKey,omitempty"`
	// RX delay.
	RxDelay uint32 `protobuf:"varint,4,opt,name=rxDelay" json:"rxDelay,omitempty"`
	// RX1 data-rate offset.
	Rx1DROffset uint32 `protobuf:"varint,5,opt,name=rx1DROffset" json:"rx1DROffset,omitempty"`
	// Channel-list ID used for CFlist (see LoRaWAN regional parameters if this applies to your region).
	ChannelListID int64 `protobuf:"varint,6,opt,name=channelListID" json:"channelListID,omitempty"`
	// RX window to use.
	RxWindow RXWindow `protobuf:"varint,7,opt,name=rxWindow,enum=api.RXWindow" json:"rxWindow,omitempty"`
	// Data-rate to use for RX2.
	Rx2DR uint32 `protobuf:"varint,8,opt,name=rx2DR" json:"rx2DR,omitempty"`
	// Name of the node.
	Name string `protobuf:"bytes,9,opt,name=name" json:"name,omitempty"`
	// Relax frame-counter mode is enabled.
	RelaxFCnt bool `protobuf:"varint,10,opt,name=relaxFCnt" json:"relaxFCnt,omitempty"`
	// Interval (in frames) in which the ADR engine may adapt the data-rate of the node (0 = disabled).
	AdrInterval uint32 `protobuf:"varint,11,opt,name=adrInterval" json:"adrInterval,omitempty"`
	// Installation-margin to use for ADR calculation.
	InstallationMargin float64 `protobuf:"fixed64,12,opt,name=installationMargin" json:"installationMargin,omitempty"`
	// Description of the node.
	Description string `protobuf:"bytes,13,opt,name=description" json:"description,omitempty"`
	// Node is activated by ABP.
	IsABP bool `protobuf:"varint,14,opt,name=isABP" json:"isABP,omitempty"`
	// ID of the application owning the node.
	ApplicationID int64 `protobuf:"varint,15,opt,name=applicationID" json:"applicationID,omitempty"`
	// Node operates in Class-C.
	IsClassC bool `protobuf:"varint,16,opt,name=isClassC" json:"isClassC,omitempty"`
}

func (m *GetNodeResponse) Reset()                    { *m = GetNodeResponse{} }
func (m *GetNodeResponse) String() string            { return proto.CompactTextString(m) }
func (*GetNodeResponse) ProtoMessage()               {}
func (*GetNodeResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *GetNodeResponse) GetDevEUI() string {
	if m != nil {
		return m.DevEUI
	}
	return ""
}

func (m *GetNodeResponse) GetAppEUI() string {
	if m != nil {
		return m.AppEUI
	}
	return ""
}

func (m *GetNodeResponse) GetAppKey() string {
	if m != nil {
		return m.AppKey
	}
	return ""
}

func (m *GetNodeResponse) GetRxDelay() uint32 {
	if m != nil {
		return m.RxDelay
	}
	return 0
}

func (m *GetNodeResponse) GetRx1DROffset() uint32 {
	if m != nil {
		return m.Rx1DROffset
	}
	return 0
}

func (m *GetNodeResponse) GetChannelListID() int64 {
	if m != nil {
		return m.ChannelListID
	}
	return 0
}

func (m *GetNodeResponse) GetRxWindow() RXWindow {
	if m != nil {
		return m.RxWindow
	}
	return RXWindow_RX1
}

func (m *GetNodeResponse) GetRx2DR() uint32 {
	if m != nil {
		return m.Rx2DR
	}
	return 0
}

func (m *GetNodeResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetNodeResponse) GetRelaxFCnt() bool {
	if m != nil {
		return m.RelaxFCnt
	}
	return false
}

func (m *GetNodeResponse) GetAdrInterval() uint32 {
	if m != nil {
		return m.AdrInterval
	}
	return 0
}

func (m *GetNodeResponse) GetInstallationMargin() float64 {
	if m != nil {
		return m.InstallationMargin
	}
	return 0
}

func (m *GetNodeResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *GetNodeResponse) GetIsABP() bool {
	if m != nil {
		return m.IsABP
	}
	return false
}

func (m *GetNodeResponse) GetApplicationID() int64 {
	if m != nil {
		return m.ApplicationID
	}
	return 0
}

func (m *GetNodeResponse) GetIsClassC() bool {
	if m != nil {
		return m.IsClassC
	}
	return false
}

type DeleteNodeRequest struct {
	// Hex encoded DevEUI of the node.
	DevEUI string `protobuf:"bytes,1,opt,name=devEUI" json:"devEUI,omitempty"`
}

func (m *DeleteNodeRequest) Reset()                    { *m = DeleteNodeRequest{} }
func (m *DeleteNodeRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteNodeRequest) ProtoMessage()               {}
func (*DeleteNodeRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *DeleteNodeRequest) GetDevEUI() string {
	if m != nil {
		return m.DevEUI
	}
	return ""
}

type DeleteNodeResponse struct {
}

func (m *DeleteNodeResponse) Reset()                    { *m = DeleteNodeResponse{} }
func (m *DeleteNodeResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteNodeResponse) ProtoMessage()               {}
func (*DeleteNodeResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

type ListNodeByApplicationIDRequest struct {
	// ID of the application for which to list the nodes.
	ApplicationID int64 `protobuf:"varint,3,opt,name=applicationID" json:"applicationID,omitempty"`
	// Max number of nodes to return in the result-set.
	Limit int64 `protobuf:"varint,1,opt,name=limit" json:"limit,omitempty"`
	// Offset of the result-set (for pagination).
	Offset int64 `protobuf:"varint,2,opt,name=offset" json:"offset,omitempty"`
}

func (m *ListNodeByApplicationIDRequest) Reset()                    { *m = ListNodeByApplicationIDRequest{} }
func (m *ListNodeByApplicationIDRequest) String() string            { return proto.CompactTextString(m) }
func (*ListNodeByApplicationIDRequest) ProtoMessage()               {}
func (*ListNodeByApplicationIDRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *ListNodeByApplicationIDRequest) GetApplicationID() int64 {
	if m != nil {
		return m.ApplicationID
	}
	return 0
}

func (m *ListNodeByApplicationIDRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListNodeByApplicationIDRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type ListNodeResponse struct {
	// Total number of nodes available within the result-set.
	TotalCount int64 `protobuf:"varint,1,opt,name=totalCount" json:"totalCount,omitempty"`
	// Nodes within this result-set.
	Result []*GetNodeResponse `protobuf:"bytes,2,rep,name=result" json:"result,omitempty"`
}

func (m *ListNodeResponse) Reset()                    { *m = ListNodeResponse{} }
func (m *ListNodeResponse) String() string            { return proto.CompactTextString(m) }
func (*ListNodeResponse) ProtoMessage()               {}
func (*ListNodeResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *ListNodeResponse) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *ListNodeResponse) GetResult() []*GetNodeResponse {
	if m != nil {
		return m.Result
	}
	return nil
}

type UpdateNodeRequest struct {
	// Hex encoded DevEUI of the node to update.
	DevEUI string `protobuf:"bytes,1,opt,name=devEUI" json:"devEUI,omitempty"`
	// Hex encoded AppEUI.
	AppEUI string `protobuf:"bytes,2,opt,name=appEUI" json:"appEUI,omitempty"`
	// Hex encoded AppKey.
	AppKey string `protobuf:"bytes,3,opt,name=appKey" json:"appKey,omitempty"`
	// RX delay.
	RxDelay uint32 `protobuf:"varint,4,opt,name=rxDelay" json:"rxDelay,omitempty"`
	// RX1 data-rate offset.
	Rx1DROffset uint32 `protobuf:"varint,5,opt,name=rx1DROffset" json:"rx1DROffset,omitempty"`
	// Channel-list ID used for CFlist (see LoRaWAN regional parameters if this applies to your region).
	ChannelListID int64 `protobuf:"varint,6,opt,name=channelListID" json:"channelListID,omitempty"`
	// RX window to use.
	RxWindow RXWindow `protobuf:"varint,7,opt,name=rxWindow,enum=api.RXWindow" json:"rxWindow,omitempty"`
	// Data-rate to use for RX2.
	Rx2DR uint32 `protobuf:"varint,8,opt,name=rx2DR" json:"rx2DR,omitempty"`
	// Name of the node (note that renaming the node affects its api endpoint)
	Name string `protobuf:"bytes,9,opt,name=name" json:"name,omitempty"`
	// Relax frame-counter mode is enabled.
	RelaxFCnt bool `protobuf:"varint,10,opt,name=relaxFCnt" json:"relaxFCnt,omitempty"`
	// Interval (in frames) in which the ADR engine may adapt the data-rate of the node (0 = disabled).
	AdrInterval uint32 `protobuf:"varint,11,opt,name=adrInterval" json:"adrInterval,omitempty"`
	// Installation-margin to use for ADR calculation.
	InstallationMargin float64 `protobuf:"fixed64,12,opt,name=installationMargin" json:"installationMargin,omitempty"`
	// ID of the application owning the node.
	ApplicationID int64 `protobuf:"varint,13,opt,name=applicationID" json:"applicationID,omitempty"`
	// Description of the node.
	Description string `protobuf:"bytes,14,opt,name=description" json:"description,omitempty"`
	// Node is activated by ABP.
	IsABP bool `protobuf:"varint,15,opt,name=isABP" json:"isABP,omitempty"`
	// Node operates in Class-C.
	IsClassC bool `protobuf:"varint,16,opt,name=isClassC" json:"isClassC,omitempty"`
}

func (m *UpdateNodeRequest) Reset()                    { *m = UpdateNodeRequest{} }
func (m *UpdateNodeRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateNodeRequest) ProtoMessage()               {}
func (*UpdateNodeRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *UpdateNodeRequest) GetDevEUI() string {
	if m != nil {
		return m.DevEUI
	}
	return ""
}

func (m *UpdateNodeRequest) GetAppEUI() string {
	if m != nil {
		return m.AppEUI
	}
	return ""
}

func (m *UpdateNodeRequest) GetAppKey() string {
	if m != nil {
		return m.AppKey
	}
	return ""
}

func (m *UpdateNodeRequest) GetRxDelay() uint32 {
	if m != nil {
		return m.RxDelay
	}
	return 0
}

func (m *UpdateNodeRequest) GetRx1DROffset() uint32 {
	if m != nil {
		return m.Rx1DROffset
	}
	return 0
}

func (m *UpdateNodeRequest) GetChannelListID() int64 {
	if m != nil {
		return m.ChannelListID
	}
	return 0
}

func (m *UpdateNodeRequest) GetRxWindow() RXWindow {
	if m != nil {
		return m.RxWindow
	}
	return RXWindow_RX1
}

func (m *UpdateNodeRequest) GetRx2DR() uint32 {
	if m != nil {
		return m.Rx2DR
	}
	return 0
}

func (m *UpdateNodeRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateNodeRequest) GetRelaxFCnt() bool {
	if m != nil {
		return m.RelaxFCnt
	}
	return false
}

func (m *UpdateNodeRequest) GetAdrInterval() uint32 {
	if m != nil {
		return m.AdrInterval
	}
	return 0
}

func (m *UpdateNodeRequest) GetInstallationMargin() float64 {
	if m != nil {
		return m.InstallationMargin
	}
	return 0
}

func (m *UpdateNodeRequest) GetApplicationID() int64 {
	if m != nil {
		return m.ApplicationID
	}
	return 0
}

func (m *UpdateNodeRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *UpdateNodeRequest) GetIsABP() bool {
	if m != nil {
		return m.IsABP
	}
	return false
}

func (m *UpdateNodeRequest) GetIsClassC() bool {
	if m != nil {
		return m.IsClassC
	}
	return false
}

type UpdateNodeResponse struct {
}

func (m *UpdateNodeResponse) Reset()                    { *m = UpdateNodeResponse{} }
func (m *UpdateNodeResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateNodeResponse) ProtoMessage()               {}
func (*UpdateNodeResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

type ActivateNodeRequest struct {
	// Hex encoded DevEUI of the node to activate.
	DevEUI string `protobuf:"bytes,1,opt,name=devEUI" json:"devEUI,omitempty"`
	// Hex encoded DevAddr.
	DevAddr string `protobuf:"bytes,2,opt,name=devAddr" json:"devAddr,omitempty"`
	// Hex encoded AppSKey.
	AppSKey string `protobuf:"bytes,3,opt,name=appSKey" json:"appSKey,omitempty"`
	// Hex encoded NwkSKey.
	NwkSKey string `protobuf:"bytes,4,opt,name=nwkSKey" json:"nwkSKey,omitempty"`
	// Uplink frame-counter.
	FCntUp uint32 `protobuf:"varint,5,opt,name=fCntUp" json:"fCntUp,omitempty"`
	// Downlink frame-counter.
	FCntDown uint32 `protobuf:"varint,6,opt,name=fCntDown" json:"fCntDown,omitempty"`
}

func (m *ActivateNodeRequest) Reset()                    { *m = ActivateNodeRequest{} }
func (m *ActivateNodeRequest) String() string            { return proto.CompactTextString(m) }
func (*ActivateNodeRequest) ProtoMessage()               {}
func (*ActivateNodeRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{10} }

func (m *ActivateNodeRequest) GetDevEUI() string {
	if m != nil {
		return m.DevEUI
	}
	return ""
}

func (m *ActivateNodeRequest) GetDevAddr() string {
	if m != nil {
		return m.DevAddr
	}
	return ""
}

func (m *ActivateNodeRequest) GetAppSKey() string {
	if m != nil {
		return m.AppSKey
	}
	return ""
}

func (m *ActivateNodeRequest) GetNwkSKey() string {
	if m != nil {
		return m.NwkSKey
	}
	return ""
}

func (m *ActivateNodeRequest) GetFCntUp() uint32 {
	if m != nil {
		return m.FCntUp
	}
	return 0
}

func (m *ActivateNodeRequest) GetFCntDown() uint32 {
	if m != nil {
		return m.FCntDown
	}
	return 0
}

type ActivateNodeResponse struct {
}

func (m *ActivateNodeResponse) Reset()                    { *m = ActivateNodeResponse{} }
func (m *ActivateNodeResponse) String() string            { return proto.CompactTextString(m) }
func (*ActivateNodeResponse) ProtoMessage()               {}
func (*ActivateNodeResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{11} }

type GetNodeActivationRequest struct {
	// Hex encoded DevEUI of the node.
	DevEUI string `protobuf:"bytes,1,opt,name=devEUI" json:"devEUI,omitempty"`
}

func (m *GetNodeActivationRequest) Reset()                    { *m = GetNodeActivationRequest{} }
func (m *GetNodeActivationRequest) String() string            { return proto.CompactTextString(m) }
func (*GetNodeActivationRequest) ProtoMessage()               {}
func (*GetNodeActivationRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{12} }

func (m *GetNodeActivationRequest) GetDevEUI() string {
	if m != nil {
		return m.DevEUI
	}
	return ""
}

type GetNodeActivationResponse struct {
	// Hex encoded DevAddr.
	DevAddr string `protobuf:"bytes,1,opt,name=devAddr" json:"devAddr,omitempty"`
	// Hex encoded AppSKey.
	AppSKey string `protobuf:"bytes,2,opt,name=appSKey" json:"appSKey,omitempty"`
	// Hex encoded NwkSKey.
	NwkSKey string `protobuf:"bytes,3,opt,name=nwkSKey" json:"nwkSKey,omitempty"`
	// Uplink frame-counter.
	FCntUp uint32 `protobuf:"varint,4,opt,name=fCntUp" json:"fCntUp,omitempty"`
	// Downlink frame-counter.
	FCntDown uint32 `protobuf:"varint,5,opt,name=fCntDown" json:"fCntDown,omitempty"`
}

func (m *GetNodeActivationResponse) Reset()                    { *m = GetNodeActivationResponse{} }
func (m *GetNodeActivationResponse) String() string            { return proto.CompactTextString(m) }
func (*GetNodeActivationResponse) ProtoMessage()               {}
func (*GetNodeActivationResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{13} }

func (m *GetNodeActivationResponse) GetDevAddr() string {
	if m != nil {
		return m.DevAddr
	}
	return ""
}

func (m *GetNodeActivationResponse) GetAppSKey() string {
	if m != nil {
		return m.AppSKey
	}
	return ""
}

func (m *GetNodeActivationResponse) GetNwkSKey() string {
	if m != nil {
		return m.NwkSKey
	}
	return ""
}

func (m *GetNodeActivationResponse) GetFCntUp() uint32 {
	if m != nil {
		return m.FCntUp
	}
	return 0
}

func (m *GetNodeActivationResponse) GetFCntDown() uint32 {
	if m != nil {
		return m.FCntDown
	}
	return 0
}

func init() {
	proto.RegisterType((*CreateNodeRequest)(nil), "api.CreateNodeRequest")
	proto.RegisterType((*CreateNodeResponse)(nil), "api.CreateNodeResponse")
	proto.RegisterType((*GetNodeRequest)(nil), "api.GetNodeRequest")
	proto.RegisterType((*GetNodeResponse)(nil), "api.GetNodeResponse")
	proto.RegisterType((*DeleteNodeRequest)(nil), "api.DeleteNodeRequest")
	proto.RegisterType((*DeleteNodeResponse)(nil), "api.DeleteNodeResponse")
	proto.RegisterType((*ListNodeByApplicationIDRequest)(nil), "api.ListNodeByApplicationIDRequest")
	proto.RegisterType((*ListNodeResponse)(nil), "api.ListNodeResponse")
	proto.RegisterType((*UpdateNodeRequest)(nil), "api.UpdateNodeRequest")
	proto.RegisterType((*UpdateNodeResponse)(nil), "api.UpdateNodeResponse")
	proto.RegisterType((*ActivateNodeRequest)(nil), "api.ActivateNodeRequest")
	proto.RegisterType((*ActivateNodeResponse)(nil), "api.ActivateNodeResponse")
	proto.RegisterType((*GetNodeActivationRequest)(nil), "api.GetNodeActivationRequest")
	proto.RegisterType((*GetNodeActivationResponse)(nil), "api.GetNodeActivationResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Node service

type NodeClient interface {
	// Create creates the given node.
	Create(ctx context.Context, in *CreateNodeRequest, opts ...grpc.CallOption) (*CreateNodeResponse, error)
	// Get returns the node for the requested name.
	Get(ctx context.Context, in *GetNodeRequest, opts ...grpc.CallOption) (*GetNodeResponse, error)
	// Delete deletes the node matching the given name.
	Delete(ctx context.Context, in *DeleteNodeRequest, opts ...grpc.CallOption) (*DeleteNodeResponse, error)
	// ListByApplicationID lists the nodes by the given application ID, sorted by the name of the node.
	ListByApplicationID(ctx context.Context, in *ListNodeByApplicationIDRequest, opts ...grpc.CallOption) (*ListNodeResponse, error)
	// Update updates the node matching the given name.
	Update(ctx context.Context, in *UpdateNodeRequest, opts ...grpc.CallOption) (*UpdateNodeResponse, error)
	// Activate (re)activates the node (only when ABP is set to true).
	Activate(ctx context.Context, in *ActivateNodeRequest, opts ...grpc.CallOption) (*ActivateNodeResponse, error)
	// GetActivation returns the current activation details of the node (OTAA and ABP).
	GetActivation(ctx context.Context, in *GetNodeActivationRequest, opts ...grpc.CallOption) (*GetNodeActivationResponse, error)
}

type nodeClient struct {
	cc *grpc.ClientConn
}

func NewNodeClient(cc *grpc.ClientConn) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) Create(ctx context.Context, in *CreateNodeRequest, opts ...grpc.CallOption) (*CreateNodeResponse, error) {
	out := new(CreateNodeResponse)
	err := grpc.Invoke(ctx, "/api.Node/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Get(ctx context.Context, in *GetNodeRequest, opts ...grpc.CallOption) (*GetNodeResponse, error) {
	out := new(GetNodeResponse)
	err := grpc.Invoke(ctx, "/api.Node/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Delete(ctx context.Context, in *DeleteNodeRequest, opts ...grpc.CallOption) (*DeleteNodeResponse, error) {
	out := new(DeleteNodeResponse)
	err := grpc.Invoke(ctx, "/api.Node/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) ListByApplicationID(ctx context.Context, in *ListNodeByApplicationIDRequest, opts ...grpc.CallOption) (*ListNodeResponse, error) {
	out := new(ListNodeResponse)
	err := grpc.Invoke(ctx, "/api.Node/ListByApplicationID", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Update(ctx context.Context, in *UpdateNodeRequest, opts ...grpc.CallOption) (*UpdateNodeResponse, error) {
	out := new(UpdateNodeResponse)
	err := grpc.Invoke(ctx, "/api.Node/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Activate(ctx context.Context, in *ActivateNodeRequest, opts ...grpc.CallOption) (*ActivateNodeResponse, error) {
	out := new(ActivateNodeResponse)
	err := grpc.Invoke(ctx, "/api.Node/Activate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) GetActivation(ctx context.Context, in *GetNodeActivationRequest, opts ...grpc.CallOption) (*GetNodeActivationResponse, error) {
	out := new(GetNodeActivationResponse)
	err := grpc.Invoke(ctx, "/api.Node/GetActivation", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Node service

type NodeServer interface {
	// Create creates the given node.
	Create(context.Context, *CreateNodeRequest) (*CreateNodeResponse, error)
	// Get returns the node for the requested name.
	Get(context.Context, *GetNodeRequest) (*GetNodeResponse, error)
	// Delete deletes the node matching the given name.
	Delete(context.Context, *DeleteNodeRequest) (*DeleteNodeResponse, error)
	// ListByApplicationID lists the nodes by the given application ID, sorted by the name of the node.
	ListByApplicationID(context.Context, *ListNodeByApplicationIDRequest) (*ListNodeResponse, error)
	// Update updates the node matching the given name.
	Update(context.Context, *UpdateNodeRequest) (*UpdateNodeResponse, error)
	// Activate (re)activates the node (only when ABP is set to true).
	Activate(context.Context, *ActivateNodeRequest) (*ActivateNodeResponse, error)
	// GetActivation returns the current activation details of the node (OTAA and ABP).
	GetActivation(context.Context, *GetNodeActivationRequest) (*GetNodeActivationResponse, error)
}

func RegisterNodeServer(s *grpc.Server, srv NodeServer) {
	s.RegisterService(&_Node_serviceDesc, srv)
}

func _Node_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Node/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Create(ctx, req.(*CreateNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Node/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Get(ctx, req.(*GetNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Node/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Delete(ctx, req.(*DeleteNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_ListByApplicationID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNodeByApplicationIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).ListByApplicationID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Node/ListByApplicationID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).ListByApplicationID(ctx, req.(*ListNodeByApplicationIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Node/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Update(ctx, req.(*UpdateNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Activate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivateNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Activate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Node/Activate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Activate(ctx, req.(*ActivateNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_GetActivation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodeActivationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).GetActivation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Node/GetActivation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).GetActivation(ctx, req.(*GetNodeActivationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Node_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Node",
	HandlerType: (*NodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Node_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Node_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Node_Delete_Handler,
		},
		{
			MethodName: "ListByApplicationID",
			Handler:    _Node_ListByApplicationID_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Node_Update_Handler,
		},
		{
			MethodName: "Activate",
			Handler:    _Node_Activate_Handler,
		},
		{
			MethodName: "GetActivation",
			Handler:    _Node_GetActivation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node.proto",
}

func init() { proto.RegisterFile("node.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 874 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x56, 0xdd, 0x72, 0xdb, 0x44,
	0x14, 0x1e, 0x45, 0xb6, 0xe3, 0x9c, 0xd4, 0x49, 0xbb, 0x71, 0xc2, 0x56, 0x14, 0x8f, 0x46, 0x65,
	0x40, 0x29, 0x4c, 0x3c, 0x98, 0x3b, 0xee, 0x5c, 0x1b, 0x32, 0x9e, 0x52, 0xe8, 0x88, 0x09, 0x3f,
	0x77, 0x2c, 0xd6, 0x26, 0x2c, 0xc8, 0xbb, 0x42, 0xda, 0x38, 0xf6, 0x74, 0x7a, 0xd3, 0x57, 0xe0,
	0x9a, 0x67, 0xe0, 0x21, 0x78, 0x04, 0x5e, 0x81, 0xa7, 0x60, 0xb8, 0x60, 0xf6, 0xc7, 0xb2, 0x1c,
	0xcb, 0x69, 0xae, 0xb8, 0x21, 0x77, 0x3e, 0xdf, 0x59, 0x9d, 0xef, 0x7c, 0xbb, 0xdf, 0x1e, 0x2f,
	0x00, 0x17, 0x31, 0x3d, 0x49, 0x33, 0x21, 0x05, 0x72, 0x49, 0xca, 0xbc, 0x47, 0x17, 0x42, 0x5c,
	0x24, 0xb4, 0x4b, 0x52, 0xd6, 0x25, 0x9c, 0x0b, 0x49, 0x24, 0x13, 0x3c, 0x37, 0x4b, 0xbc, 0x7b,
	0x63, 0x31, 0x99, 0x08, 0x6e, 0xa2, 0xe0, 0x1f, 0x17, 0x1e, 0x0c, 0x32, 0x4a, 0x24, 0xfd, 0x42,
	0xc4, 0x34, 0xa2, 0xbf, 0x5c, 0xd2, 0x5c, 0xa2, 0x23, 0x68, 0xc4, 0x74, 0xfa, 0xe9, 0xd9, 0x08,
	0x3b, 0xbe, 0x13, 0xee, 0x44, 0x36, 0x52, 0x38, 0x49, 0x53, 0x85, 0x6f, 0x19, 0xdc, 0x44, 0x16,
	0x7f, 0x46, 0xe7, 0xd8, 0x2d, 0xf0, 0x67, 0x74, 0x8e, 0x30, 0x6c, 0x67, 0xb3, 0x21, 0x4d, 0xc8,
	0x1c, 0xd7, 0x7c, 0x27, 0x6c, 0x45, 0x8b, 0x10, 0xf9, 0xb0, 0x9b, 0xcd, 0x3e, 0x1a, 0x46, 0x5f,
	0x9e, 0x9f, 0xe7, 0x54, 0xe2, 0xba, 0xce, 0x96, 0x21, 0xf4, 0x2e, 0xb4, 0xc6, 0x3f, 0x12, 0xce,
	0x69, 0xf2, 0x39, 0xcb, 0xe5, 0x68, 0x88, 0x1b, 0xbe, 0x13, 0xba, 0xd1, 0x2a, 0x88, 0x8e, 0xa1,
	0x99, 0xcd, 0xbe, 0x61, 0x3c, 0x16, 0x57, 0x78, 0xdb, 0x77, 0xc2, 0xbd, 0x5e, 0xeb, 0x84, 0xa4,
	0xec, 0x24, 0xfa, 0xd6, 0x80, 0x51, 0x91, 0x46, 0x6d, 0xa8, 0x67, 0xb3, 0xde, 0x30, 0xc2, 0x4d,
	0x4d, 0x66, 0x02, 0x84, 0xa0, 0xc6, 0xc9, 0x84, 0xe2, 0x1d, 0xdd, 0xb8, 0xfe, 0x8d, 0x1e, 0xc1,
	0x4e, 0x46, 0x13, 0x32, 0xfb, 0x6c, 0xc0, 0x25, 0x06, 0xdf, 0x09, 0x9b, 0xd1, 0x12, 0x50, 0xad,
	0x93, 0x38, 0x1b, 0x71, 0x49, 0xb3, 0x29, 0x49, 0xf0, 0xae, 0x69, 0xbd, 0x04, 0xa1, 0x13, 0x40,
	0x8c, 0xe7, 0x92, 0x24, 0x89, 0xde, 0xf9, 0xe7, 0x24, 0xbb, 0x60, 0x1c, 0xdf, 0xf3, 0x9d, 0xd0,
	0x89, 0x2a, 0x32, 0x4a, 0x2a, 0x49, 0xd3, 0x84, 0x8d, 0x35, 0x38, 0x1a, 0xe2, 0x96, 0x91, 0xba,
	0x02, 0x2a, 0xde, 0x98, 0xe6, 0xe3, 0x8c, 0xa5, 0x0a, 0xc0, 0x7b, 0xba, 0xe1, 0x32, 0xa4, 0x14,
	0xb2, 0xbc, 0xff, 0xf4, 0x05, 0xde, 0xd7, 0x3d, 0x9b, 0x00, 0x79, 0xd0, 0x64, 0xf9, 0x20, 0x21,
	0x79, 0x3e, 0xc0, 0xf7, 0x75, 0xa2, 0x88, 0x83, 0x36, 0xa0, 0xf2, 0xe9, 0xe7, 0xa9, 0xe0, 0x39,
	0x0d, 0x42, 0xd8, 0x3b, 0xa5, 0xf2, 0x16, 0x86, 0x08, 0xfe, 0x76, 0x61, 0xbf, 0x58, 0x6a, 0xbe,
	0xbe, 0x33, 0xcf, 0x7f, 0x6b, 0x9e, 0x6b, 0xb6, 0x68, 0xdd, 0x60, 0x8b, 0xbd, 0xb2, 0x2d, 0xd6,
	0x4c, 0xb7, 0x5f, 0x65, 0xba, 0x9b, 0xcc, 0xf3, 0x01, 0x3c, 0x18, 0xd2, 0x84, 0xde, 0x6a, 0x74,
	0x28, 0xa7, 0x95, 0x17, 0x5b, 0xa7, 0x49, 0xe8, 0xa8, 0xb3, 0x50, 0xd8, 0xd3, 0x79, 0xbf, 0xcc,
	0xbc, 0xa8, 0xb7, 0xd6, 0xa6, 0x5b, 0xd5, 0x66, 0x1b, 0xea, 0x09, 0x9b, 0x30, 0xa9, 0x49, 0xdd,
	0xc8, 0x04, 0xaa, 0x17, 0x61, 0x2c, 0xb2, 0xa5, 0x61, 0x1b, 0x05, 0xdf, 0xc3, 0xfd, 0x05, 0x6b,
	0xe1, 0xda, 0x0e, 0x80, 0x14, 0x92, 0x24, 0x03, 0x71, 0xc9, 0x17, 0x65, 0x4a, 0x08, 0xfa, 0x10,
	0x1a, 0x19, 0xcd, 0x2f, 0x13, 0x55, 0xcb, 0x0d, 0x77, 0x7b, 0x6d, 0xed, 0x94, 0x6b, 0xde, 0x8f,
	0xec, 0x1a, 0x3d, 0x56, 0xcf, 0xd2, 0xf8, 0x6e, 0xac, 0xfe, 0x6f, 0xc7, 0x6a, 0xf9, 0xf4, 0xad,
	0xd9, 0x7f, 0x77, 0xe0, 0xa0, 0x3f, 0x96, 0x6c, 0x7a, 0x4b, 0x5b, 0x60, 0xd8, 0x8e, 0xe9, 0xb4,
	0x1f, 0xc7, 0x99, 0xf5, 0xc5, 0x22, 0x54, 0x19, 0x92, 0xa6, 0x5f, 0x2d, 0x9d, 0xb1, 0x08, 0x55,
	0x86, 0x5f, 0xfd, 0xac, 0x33, 0x35, 0x93, 0xb1, 0xa1, 0x62, 0x39, 0x1f, 0x70, 0x79, 0x96, 0x5a,
	0x57, 0xd8, 0x48, 0xe9, 0x50, 0xbf, 0x86, 0xe2, 0x8a, 0x6b, 0x2f, 0xb4, 0xa2, 0x22, 0x0e, 0x8e,
	0xa0, 0xbd, 0xda, 0xb0, 0x55, 0xd2, 0x03, 0x6c, 0x9d, 0x6f, 0xd3, 0x4c, 0xf0, 0x37, 0x0d, 0x80,
	0xdf, 0x1c, 0x78, 0x58, 0xf1, 0x91, 0xbd, 0x7e, 0x25, 0xad, 0xce, 0x46, 0xad, 0x5b, 0x1b, 0xb5,
	0xba, 0x9b, 0xb4, 0xd6, 0x36, 0x6a, 0xad, 0xaf, 0x6a, 0xed, 0xfd, 0x51, 0x87, 0x9a, 0x6a, 0x0e,
	0xbd, 0x80, 0x86, 0xf9, 0x4f, 0x44, 0x47, 0xda, 0xf3, 0x6b, 0xcf, 0x23, 0xef, 0xad, 0x35, 0xdc,
	0xee, 0xcb, 0xe1, 0xeb, 0x3f, 0xff, 0xfa, 0x75, 0x6b, 0x3f, 0x00, 0xfd, 0xf6, 0x52, 0xef, 0xb2,
	0xfc, 0x13, 0xe7, 0x09, 0x7a, 0x0e, 0xee, 0x29, 0x95, 0xe8, 0x60, 0x75, 0x64, 0x98, 0x5a, 0x95,
	0x73, 0x24, 0x78, 0x5b, 0x17, 0x3a, 0x44, 0x07, 0xcb, 0x42, 0xdd, 0x97, 0x66, 0x23, 0x5f, 0xa1,
	0xaf, 0xa1, 0x61, 0x46, 0xa9, 0x6d, 0x70, 0x6d, 0x08, 0xdb, 0x06, 0x2b, 0xe6, 0xad, 0xad, 0xfb,
	0xa4, 0xb2, 0xee, 0x6b, 0x07, 0x0e, 0xd4, 0xfd, 0xbf, 0x36, 0x89, 0xd1, 0x63, 0x5d, 0xed, 0xe6,
	0x39, 0xed, 0x1d, 0xae, 0x2c, 0x2a, 0x08, 0xbb, 0x9a, 0xf0, 0x18, 0xbd, 0x6f, 0x5e, 0xa3, 0xcb,
	0x2f, 0xf3, 0xee, 0xcb, 0x95, 0x7b, 0xf8, 0xca, 0x74, 0x83, 0xbe, 0x83, 0x86, 0xb9, 0x3a, 0x56,
	0xdc, 0xda, 0x14, 0xb5, 0xe2, 0x2a, 0xee, 0x57, 0x47, 0x73, 0x61, 0xaf, 0x4a, 0x9c, 0x3a, 0x86,
	0x9f, 0xa0, 0xb9, 0x70, 0x33, 0xc2, 0xba, 0x48, 0xc5, 0x6d, 0xf4, 0x1e, 0x56, 0x64, 0x2c, 0xc1,
	0xb1, 0x26, 0x78, 0x1c, 0x74, 0x2a, 0x08, 0xba, 0xa4, 0x30, 0xb5, 0xe2, 0x9a, 0x42, 0xeb, 0x94,
	0xca, 0xa5, 0xd1, 0xd1, 0x3b, 0xe5, 0x73, 0x5e, 0xbb, 0x35, 0x5e, 0x67, 0x53, 0xda, 0x52, 0xbf,
	0xa7, 0xa9, 0x7d, 0xf4, 0x06, 0xea, 0x1f, 0x1a, 0xfa, 0x59, 0xff, 0xf1, 0xbf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x99, 0x0e, 0xde, 0xa8, 0x15, 0x0c, 0x00, 0x00,
}
