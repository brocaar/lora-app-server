package api

import (
	log "github.com/Sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	pb "github.com/brocaar/lora-app-server/api"
	"github.com/brocaar/lora-app-server/internal/api/auth"
	"github.com/brocaar/lora-app-server/internal/common"
	"github.com/brocaar/lora-app-server/internal/storage"
	"github.com/brocaar/loraserver/api/ns"
	"github.com/brocaar/lorawan"
)

// NodeAPI exports the Node related functions.
type NodeAPI struct {
	ctx       common.Context
	validator auth.Validator
}

// NewNodeAPI creates a new NodeAPI.
func NewNodeAPI(ctx common.Context, validator auth.Validator) *NodeAPI {
	return &NodeAPI{
		ctx:       ctx,
		validator: validator,
	}
}

// Create creates the given Node.
func (a *NodeAPI) Create(ctx context.Context, req *pb.CreateNodeRequest) (*pb.CreateNodeResponse, error) {
	var appEUI, devEUI lorawan.EUI64
	var appKey lorawan.AES128Key

	if err := appEUI.UnmarshalText([]byte(req.AppEUI)); err != nil {
		return nil, grpc.Errorf(codes.InvalidArgument, err.Error())
	}
	if err := devEUI.UnmarshalText([]byte(req.DevEUI)); err != nil {
		return nil, grpc.Errorf(codes.InvalidArgument, err.Error())
	}
	if err := appKey.UnmarshalText([]byte(req.AppKey)); err != nil {
		return nil, grpc.Errorf(codes.InvalidArgument, err.Error())
	}

	if err := a.validator.Validate(ctx,
		auth.ValidateAPIMethod("Node.Create"),
		auth.ValidateApplicationID(req.ApplicationID),
		auth.ValidateNode(devEUI),
	); err != nil {
		return nil, grpc.Errorf(codes.Unauthenticated, "authentication failed: %s", err)
	}

	// if Name is "", set it to the DevEUI
	if req.Name == "" {
		req.Name = req.DevEUI
	}

	node := storage.Node{
		ApplicationID: req.ApplicationID,
		Name:          req.Name,
		Description:   req.Description,
		DevEUI:        devEUI,
		AppEUI:        appEUI,
		AppKey:        appKey,
		IsABP:         req.IsABP,
		IsClassC:      req.IsClassC,
		RelaxFCnt:     req.RelaxFCnt,

		RXDelay:     uint8(req.RxDelay),
		RX1DROffset: uint8(req.Rx1DROffset),
		RXWindow:    storage.RXWindow(req.RxWindow),
		RX2DR:       uint8(req.Rx2DR),

		ADRInterval:        req.AdrInterval,
		InstallationMargin: req.InstallationMargin,
	}
	if req.ChannelListID > 0 {
		node.ChannelListID = &req.ChannelListID
	}

	if err := storage.CreateNode(a.ctx.DB, node); err != nil {
		return nil, grpc.Errorf(codes.Unknown, err.Error())
	}

	return &pb.CreateNodeResponse{}, nil
}

// Get returns the Node for the given name.
func (a *NodeAPI) Get(ctx context.Context, req *pb.GetNodeRequest) (*pb.GetNodeResponse, error) {
	var eui lorawan.EUI64
	if err := eui.UnmarshalText([]byte(req.DevEUI)); err != nil {
		return nil, grpc.Errorf(codes.InvalidArgument, err.Error())
	}

	node, err := storage.GetNode(a.ctx.DB, eui)
	if err != nil {
		return nil, grpc.Errorf(codes.Unknown, err.Error())
	}

	if err := a.validator.Validate(ctx,
		auth.ValidateAPIMethod("Node.Get"),
		auth.ValidateApplicationID(node.ApplicationID),
		auth.ValidateNode(node.DevEUI),
	); err != nil {
		return nil, grpc.Errorf(codes.Unauthenticated, "authentication failed: %s", err)
	}

	resp := pb.GetNodeResponse{
		Name:               node.Name,
		Description:        node.Description,
		DevEUI:             node.DevEUI.String(),
		AppEUI:             node.AppEUI.String(),
		AppKey:             node.AppKey.String(),
		IsABP:              node.IsABP,
		IsClassC:           node.IsClassC,
		RxDelay:            uint32(node.RXDelay),
		Rx1DROffset:        uint32(node.RX1DROffset),
		RxWindow:           pb.RXWindow(node.RXWindow),
		Rx2DR:              uint32(node.RX2DR),
		RelaxFCnt:          node.RelaxFCnt,
		AdrInterval:        node.ADRInterval,
		InstallationMargin: node.InstallationMargin,
		ApplicationID:      node.ApplicationID,
	}

	if node.ChannelListID != nil {
		resp.ChannelListID = *node.ChannelListID
	}

	log.WithFields(log.Fields{
		"dev_eui":        node.DevEUI,
		"application_id": node.ApplicationID,
	}).Info("node created")

	return &resp, nil
}

// ListByApplicationID returns a list of nodes (given an application id, limit and offset).
func (a *NodeAPI) ListByApplicationID(ctx context.Context, req *pb.ListNodeByApplicationIDRequest) (*pb.ListNodeResponse, error) {
	if err := a.validator.Validate(ctx,
		auth.ValidateAPIMethod("Node.List"),
		auth.ValidateApplicationID(req.ApplicationID),
	); err != nil {
		return nil, grpc.Errorf(codes.Unauthenticated, "authentication failed: %s", err)
	}

	nodes, err := storage.GetNodesForApplicationID(a.ctx.DB, req.ApplicationID, int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, err.Error())
	}
	count, err := storage.GetNodesCountForApplicationID(a.ctx.DB, req.ApplicationID)
	return a.returnList(count, nodes)
}

// Update updates the node matching the given name.
func (a *NodeAPI) Update(ctx context.Context, req *pb.UpdateNodeRequest) (*pb.UpdateNodeResponse, error) {
	var appEUI, devEUI lorawan.EUI64
	var appKey lorawan.AES128Key

	if err := appEUI.UnmarshalText([]byte(req.AppEUI)); err != nil {
		return nil, grpc.Errorf(codes.InvalidArgument, err.Error())
	}
	if err := devEUI.UnmarshalText([]byte(req.DevEUI)); err != nil {
		return nil, grpc.Errorf(codes.InvalidArgument, err.Error())
	}
	if err := appKey.UnmarshalText([]byte(req.AppKey)); err != nil {
		return nil, grpc.Errorf(codes.InvalidArgument, err.Error())
	}

	node, err := storage.GetNode(a.ctx.DB, devEUI)
	if err != nil {
		return nil, grpc.Errorf(codes.Unknown, err.Error())
	}

	if err := a.validator.Validate(ctx,
		auth.ValidateAPIMethod("Node.Update"),
		auth.ValidateApplicationID(node.ApplicationID),
		auth.ValidateApplicationID(req.ApplicationID), // in case of an application ID update
		auth.ValidateNode(node.DevEUI),
	); err != nil {
		return nil, grpc.Errorf(codes.Unauthenticated, "authentication failed: %s", err)
	}

	node.Name = req.Name
	node.Description = req.Description
	node.AppEUI = appEUI
	node.AppKey = appKey
	node.IsABP = req.IsABP
	node.IsClassC = req.IsClassC
	node.RXDelay = uint8(req.RxDelay)
	node.RX1DROffset = uint8(req.Rx1DROffset)
	node.RXWindow = storage.RXWindow(req.RxWindow)
	node.RX2DR = uint8(req.Rx2DR)
	node.RelaxFCnt = req.RelaxFCnt
	node.ADRInterval = req.AdrInterval
	node.InstallationMargin = req.InstallationMargin
	node.ApplicationID = req.ApplicationID
	if req.ChannelListID > 0 {
		node.ChannelListID = &req.ChannelListID
	} else {
		node.ChannelListID = nil
	}

	if err := storage.UpdateNode(a.ctx.DB, node); err != nil {
		return nil, grpc.Errorf(codes.Unknown, err.Error())
	}

	log.WithFields(log.Fields{
		"dev_eui":        node.DevEUI,
		"application_id": node.ApplicationID,
	}).Info("node updated")

	return &pb.UpdateNodeResponse{}, nil
}

// Delete deletes the node matching the given name.
func (a *NodeAPI) Delete(ctx context.Context, req *pb.DeleteNodeRequest) (*pb.DeleteNodeResponse, error) {
	var eui lorawan.EUI64
	if err := eui.UnmarshalText([]byte(req.DevEUI)); err != nil {
		return nil, grpc.Errorf(codes.InvalidArgument, err.Error())
	}

	node, err := storage.GetNode(a.ctx.DB, eui)
	if err != nil {
		return nil, grpc.Errorf(codes.Unknown, err.Error())
	}

	if err := a.validator.Validate(ctx,
		auth.ValidateAPIMethod("Node.Delete"),
		auth.ValidateApplicationID(node.ApplicationID),
		auth.ValidateNode(node.DevEUI),
	); err != nil {
		return nil, grpc.Errorf(codes.Unauthenticated, "authentication failed: %s", err)
	}

	if err := storage.DeleteNode(a.ctx.DB, node.DevEUI); err != nil {
		return nil, grpc.Errorf(codes.Unknown, err.Error())
	}

	// try to delete the node-session
	_, _ = a.ctx.NetworkServer.DeleteNodeSession(context.Background(), &ns.DeleteNodeSessionRequest{
		DevEUI: node.DevEUI[:],
	})

	log.WithFields(log.Fields{
		"dev_eui":        node.DevEUI,
		"application_id": node.ApplicationID,
	}).Info("node deleted")

	return &pb.DeleteNodeResponse{}, nil
}

// Activate activates the node (ABP only).
func (a *NodeAPI) Activate(ctx context.Context, req *pb.ActivateNodeRequest) (*pb.ActivateNodeResponse, error) {
	var devAddr lorawan.DevAddr
	var devEUI lorawan.EUI64
	var appSKey, nwkSKey lorawan.AES128Key

	if err := devAddr.UnmarshalText([]byte(req.DevAddr)); err != nil {
		return nil, grpc.Errorf(codes.InvalidArgument, "devAddr: %s", err)
	}
	if err := devEUI.UnmarshalText([]byte(req.DevEUI)); err != nil {
		return nil, grpc.Errorf(codes.InvalidArgument, "devEUI: %s", err)
	}
	if err := appSKey.UnmarshalText([]byte(req.AppSKey)); err != nil {
		return nil, grpc.Errorf(codes.InvalidArgument, "appSKey: %s", err)
	}
	if err := nwkSKey.UnmarshalText([]byte(req.NwkSKey)); err != nil {
		return nil, grpc.Errorf(codes.InvalidArgument, "nwkSKey: %s", err)
	}

	node, err := storage.GetNode(a.ctx.DB, devEUI)
	if err != nil {
		return nil, grpc.Errorf(codes.Unknown, err.Error())
	}

	if err := a.validator.Validate(ctx,
		auth.ValidateAPIMethod("Node.Activate"),
		auth.ValidateNode(node.DevEUI),
		auth.ValidateApplicationID(node.ApplicationID),
	); err != nil {
		return nil, grpc.Errorf(codes.Unauthenticated, "authentication failed: %s", err)
	}

	if !node.IsABP {
		return nil, grpc.Errorf(codes.FailedPrecondition, "node must be an ABP node")
	}

	// try to remove an existing node-session.
	// TODO: refactor once https://github.com/brocaar/loraserver/pull/124 is in place?
	// so that we can call something like SaveNodeSession which will either
	// create or update an existing node-session
	_, _ = a.ctx.NetworkServer.DeleteNodeSession(context.Background(), &ns.DeleteNodeSessionRequest{
		DevEUI: node.DevEUI[:],
	})

	cFList, err := storage.GetCFListForNode(a.ctx.DB, node)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, err.Error())
	}

	createNSReq := ns.CreateNodeSessionRequest{
		DevAddr:            devAddr[:],
		AppEUI:             node.AppEUI[:],
		DevEUI:             node.DevEUI[:],
		NwkSKey:            nwkSKey[:],
		FCntUp:             req.FCntUp,
		FCntDown:           req.FCntDown,
		RxDelay:            uint32(node.RXDelay),
		Rx1DROffset:        uint32(node.RX1DROffset),
		RxWindow:           ns.RXWindow(node.RXWindow),
		Rx2DR:              uint32(node.RX2DR),
		RelaxFCnt:          node.RelaxFCnt,
		AdrInterval:        node.ADRInterval,
		InstallationMargin: node.InstallationMargin,
	}
	if cFList != nil {
		for _, freq := range cFList {
			createNSReq.CFList = append(createNSReq.CFList, freq)
		}
	}

	_, err = a.ctx.NetworkServer.CreateNodeSession(context.Background(), &createNSReq)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, err.Error())
	}

	node.AppSKey = appSKey
	node.DevAddr = devAddr
	node.NwkSKey = nwkSKey

	if err = storage.UpdateNode(a.ctx.DB, node); err != nil {
		return nil, grpc.Errorf(codes.Internal, err.Error())
	}

	log.WithFields(log.Fields{
		"dev_addr":       devAddr,
		"dev_eui":        node.DevEUI,
		"application_id": node.ApplicationID,
	}).Info("node activated")

	return &pb.ActivateNodeResponse{}, nil
}

func (a *NodeAPI) GetActivation(ctx context.Context, req *pb.GetNodeActivationRequest) (*pb.GetNodeActivationResponse, error) {
	var devAddr lorawan.DevAddr
	var devEUI lorawan.EUI64
	var nwkSKey lorawan.AES128Key

	if err := devEUI.UnmarshalText([]byte(req.DevEUI)); err != nil {
		return nil, grpc.Errorf(codes.InvalidArgument, "devEUI: %s", err)
	}

	node, err := storage.GetNode(a.ctx.DB, devEUI)
	if err != nil {
		return nil, grpc.Errorf(codes.Unknown, err.Error())
	}

	if err := a.validator.Validate(ctx,
		auth.ValidateAPIMethod("Node.GetActivation"),
		auth.ValidateNode(node.DevEUI),
		auth.ValidateApplicationID(node.ApplicationID),
	); err != nil {
		return nil, grpc.Errorf(codes.Unauthenticated, "authentication failed: %s", err)
	}

	ns, err := a.ctx.NetworkServer.GetNodeSession(context.Background(), &ns.GetNodeSessionRequest{
		DevEUI: node.DevEUI[:],
	})
	if err != nil {
		return nil, err
	}

	copy(devAddr[:], ns.DevAddr)
	copy(nwkSKey[:], ns.NwkSKey)

	return &pb.GetNodeActivationResponse{
		DevAddr:  devAddr.String(),
		AppSKey:  node.AppSKey.String(),
		NwkSKey:  nwkSKey.String(),
		FCntUp:   ns.FCntUp,
		FCntDown: ns.FCntDown,
	}, nil
}

func (a *NodeAPI) returnList(count int, nodes []storage.Node) (*pb.ListNodeResponse, error) {
	resp := pb.ListNodeResponse{
		TotalCount: int64(count),
	}
	for _, node := range nodes {
		item := pb.GetNodeResponse{
			Name:               node.Name,
			Description:        node.Description,
			DevEUI:             node.DevEUI.String(),
			AppEUI:             node.AppEUI.String(),
			AppKey:             node.AppKey.String(),
			IsABP:              node.IsABP,
			IsClassC:           node.IsClassC,
			RxDelay:            uint32(node.RXDelay),
			Rx1DROffset:        uint32(node.RX1DROffset),
			RxWindow:           pb.RXWindow(node.RXWindow),
			Rx2DR:              uint32(node.RX2DR),
			RelaxFCnt:          node.RelaxFCnt,
			AdrInterval:        node.ADRInterval,
			InstallationMargin: node.InstallationMargin,
			ApplicationID:      node.ApplicationID,
		}

		if node.ChannelListID != nil {
			item.ChannelListID = *node.ChannelListID
		}

		resp.Result = append(resp.Result, &item)
	}
	return &resp, nil
}
