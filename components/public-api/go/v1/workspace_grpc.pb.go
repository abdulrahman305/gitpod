// Copyright (c) 2025 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: gitpod/v1/workspace.proto

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

// WorkspaceServiceClient is the client API for WorkspaceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkspaceServiceClient interface {
	// GetWorkspace returns a single workspace.
	//
	// +return NOT_FOUND User does not have access to a workspace with the given
	// ID +return NOT_FOUND Workspace does not exist
	GetWorkspace(ctx context.Context, in *GetWorkspaceRequest, opts ...grpc.CallOption) (*GetWorkspaceResponse, error)
	// WatchWorkspaceStatus watches the workspaces status changes
	//
	// workspace_id +return NOT_FOUND Workspace does not exist
	WatchWorkspaceStatus(ctx context.Context, in *WatchWorkspaceStatusRequest, opts ...grpc.CallOption) (WorkspaceService_WatchWorkspaceStatusClient, error)
	// ListWorkspaces returns a list of workspaces that match the query.
	ListWorkspaces(ctx context.Context, in *ListWorkspacesRequest, opts ...grpc.CallOption) (*ListWorkspacesResponse, error)
	// ListWorkspaceSessions returns a list of workspace sessions that match the
	ListWorkspaceSessions(ctx context.Context, in *ListWorkspaceSessionsRequest, opts ...grpc.CallOption) (*ListWorkspaceSessionsResponse, error)
	// CreateAndStartWorkspace creates a new workspace and starts it.
	CreateAndStartWorkspace(ctx context.Context, in *CreateAndStartWorkspaceRequest, opts ...grpc.CallOption) (*CreateAndStartWorkspaceResponse, error)
	// StartWorkspace starts an existing workspace.
	// If the specified workspace is not in stopped phase, this will return the
	// workspace as is.
	StartWorkspace(ctx context.Context, in *StartWorkspaceRequest, opts ...grpc.CallOption) (*StartWorkspaceResponse, error)
	// UpdateWorkspace updates the workspace.
	UpdateWorkspace(ctx context.Context, in *UpdateWorkspaceRequest, opts ...grpc.CallOption) (*UpdateWorkspaceResponse, error)
	// StopWorkspace stops a running workspace.
	StopWorkspace(ctx context.Context, in *StopWorkspaceRequest, opts ...grpc.CallOption) (*StopWorkspaceResponse, error)
	// DeleteWorkspace deletes a workspace.
	// When the workspace is running, it will be stopped as well.
	// Deleted workspaces cannot be started again.
	DeleteWorkspace(ctx context.Context, in *DeleteWorkspaceRequest, opts ...grpc.CallOption) (*DeleteWorkspaceResponse, error)
	// ListWorkspaceClasses enumerates all available workspace classes.
	ListWorkspaceClasses(ctx context.Context, in *ListWorkspaceClassesRequest, opts ...grpc.CallOption) (*ListWorkspaceClassesResponse, error)
	// ParseContextURL parses a context URL and returns the workspace metadata and
	// spec. Not implemented yet.
	ParseContextURL(ctx context.Context, in *ParseContextURLRequest, opts ...grpc.CallOption) (*ParseContextURLResponse, error)
	// GetWorkspaceDefaultImage returns the default workspace image of specified
	// workspace.
	GetWorkspaceDefaultImage(ctx context.Context, in *GetWorkspaceDefaultImageRequest, opts ...grpc.CallOption) (*GetWorkspaceDefaultImageResponse, error)
	// SendHeartBeat sends a heartbeat to activate the workspace
	SendHeartBeat(ctx context.Context, in *SendHeartBeatRequest, opts ...grpc.CallOption) (*SendHeartBeatResponse, error)
	// GetWorkspaceOwnerToken returns an owner token of workspace.
	GetWorkspaceOwnerToken(ctx context.Context, in *GetWorkspaceOwnerTokenRequest, opts ...grpc.CallOption) (*GetWorkspaceOwnerTokenResponse, error)
	// GetWorkspaceEditorCredentials returns an credentials that is used in editor
	// to encrypt and decrypt secrets
	GetWorkspaceEditorCredentials(ctx context.Context, in *GetWorkspaceEditorCredentialsRequest, opts ...grpc.CallOption) (*GetWorkspaceEditorCredentialsResponse, error)
	// CreateWorkspaceSnapshot creates a snapshot of the workspace that can be
	// shared with others.
	CreateWorkspaceSnapshot(ctx context.Context, in *CreateWorkspaceSnapshotRequest, opts ...grpc.CallOption) (*CreateWorkspaceSnapshotResponse, error)
	// WaitWorkspaceSnapshot waits for the snapshot to be available or failed.
	WaitForWorkspaceSnapshot(ctx context.Context, in *WaitForWorkspaceSnapshotRequest, opts ...grpc.CallOption) (*WaitForWorkspaceSnapshotResponse, error)
	// UpdateWorkspacePort updates the port of workspace.
	UpdateWorkspacePort(ctx context.Context, in *UpdateWorkspacePortRequest, opts ...grpc.CallOption) (*UpdateWorkspacePortResponse, error)
}

type workspaceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkspaceServiceClient(cc grpc.ClientConnInterface) WorkspaceServiceClient {
	return &workspaceServiceClient{cc}
}

func (c *workspaceServiceClient) GetWorkspace(ctx context.Context, in *GetWorkspaceRequest, opts ...grpc.CallOption) (*GetWorkspaceResponse, error) {
	out := new(GetWorkspaceResponse)
	err := c.cc.Invoke(ctx, "/gitpod.v1.WorkspaceService/GetWorkspace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) WatchWorkspaceStatus(ctx context.Context, in *WatchWorkspaceStatusRequest, opts ...grpc.CallOption) (WorkspaceService_WatchWorkspaceStatusClient, error) {
	stream, err := c.cc.NewStream(ctx, &WorkspaceService_ServiceDesc.Streams[0], "/gitpod.v1.WorkspaceService/WatchWorkspaceStatus", opts...)
	if err != nil {
		return nil, err
	}
	x := &workspaceServiceWatchWorkspaceStatusClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WorkspaceService_WatchWorkspaceStatusClient interface {
	Recv() (*WatchWorkspaceStatusResponse, error)
	grpc.ClientStream
}

type workspaceServiceWatchWorkspaceStatusClient struct {
	grpc.ClientStream
}

func (x *workspaceServiceWatchWorkspaceStatusClient) Recv() (*WatchWorkspaceStatusResponse, error) {
	m := new(WatchWorkspaceStatusResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *workspaceServiceClient) ListWorkspaces(ctx context.Context, in *ListWorkspacesRequest, opts ...grpc.CallOption) (*ListWorkspacesResponse, error) {
	out := new(ListWorkspacesResponse)
	err := c.cc.Invoke(ctx, "/gitpod.v1.WorkspaceService/ListWorkspaces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) ListWorkspaceSessions(ctx context.Context, in *ListWorkspaceSessionsRequest, opts ...grpc.CallOption) (*ListWorkspaceSessionsResponse, error) {
	out := new(ListWorkspaceSessionsResponse)
	err := c.cc.Invoke(ctx, "/gitpod.v1.WorkspaceService/ListWorkspaceSessions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) CreateAndStartWorkspace(ctx context.Context, in *CreateAndStartWorkspaceRequest, opts ...grpc.CallOption) (*CreateAndStartWorkspaceResponse, error) {
	out := new(CreateAndStartWorkspaceResponse)
	err := c.cc.Invoke(ctx, "/gitpod.v1.WorkspaceService/CreateAndStartWorkspace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) StartWorkspace(ctx context.Context, in *StartWorkspaceRequest, opts ...grpc.CallOption) (*StartWorkspaceResponse, error) {
	out := new(StartWorkspaceResponse)
	err := c.cc.Invoke(ctx, "/gitpod.v1.WorkspaceService/StartWorkspace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) UpdateWorkspace(ctx context.Context, in *UpdateWorkspaceRequest, opts ...grpc.CallOption) (*UpdateWorkspaceResponse, error) {
	out := new(UpdateWorkspaceResponse)
	err := c.cc.Invoke(ctx, "/gitpod.v1.WorkspaceService/UpdateWorkspace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) StopWorkspace(ctx context.Context, in *StopWorkspaceRequest, opts ...grpc.CallOption) (*StopWorkspaceResponse, error) {
	out := new(StopWorkspaceResponse)
	err := c.cc.Invoke(ctx, "/gitpod.v1.WorkspaceService/StopWorkspace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) DeleteWorkspace(ctx context.Context, in *DeleteWorkspaceRequest, opts ...grpc.CallOption) (*DeleteWorkspaceResponse, error) {
	out := new(DeleteWorkspaceResponse)
	err := c.cc.Invoke(ctx, "/gitpod.v1.WorkspaceService/DeleteWorkspace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) ListWorkspaceClasses(ctx context.Context, in *ListWorkspaceClassesRequest, opts ...grpc.CallOption) (*ListWorkspaceClassesResponse, error) {
	out := new(ListWorkspaceClassesResponse)
	err := c.cc.Invoke(ctx, "/gitpod.v1.WorkspaceService/ListWorkspaceClasses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) ParseContextURL(ctx context.Context, in *ParseContextURLRequest, opts ...grpc.CallOption) (*ParseContextURLResponse, error) {
	out := new(ParseContextURLResponse)
	err := c.cc.Invoke(ctx, "/gitpod.v1.WorkspaceService/ParseContextURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) GetWorkspaceDefaultImage(ctx context.Context, in *GetWorkspaceDefaultImageRequest, opts ...grpc.CallOption) (*GetWorkspaceDefaultImageResponse, error) {
	out := new(GetWorkspaceDefaultImageResponse)
	err := c.cc.Invoke(ctx, "/gitpod.v1.WorkspaceService/GetWorkspaceDefaultImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) SendHeartBeat(ctx context.Context, in *SendHeartBeatRequest, opts ...grpc.CallOption) (*SendHeartBeatResponse, error) {
	out := new(SendHeartBeatResponse)
	err := c.cc.Invoke(ctx, "/gitpod.v1.WorkspaceService/SendHeartBeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) GetWorkspaceOwnerToken(ctx context.Context, in *GetWorkspaceOwnerTokenRequest, opts ...grpc.CallOption) (*GetWorkspaceOwnerTokenResponse, error) {
	out := new(GetWorkspaceOwnerTokenResponse)
	err := c.cc.Invoke(ctx, "/gitpod.v1.WorkspaceService/GetWorkspaceOwnerToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) GetWorkspaceEditorCredentials(ctx context.Context, in *GetWorkspaceEditorCredentialsRequest, opts ...grpc.CallOption) (*GetWorkspaceEditorCredentialsResponse, error) {
	out := new(GetWorkspaceEditorCredentialsResponse)
	err := c.cc.Invoke(ctx, "/gitpod.v1.WorkspaceService/GetWorkspaceEditorCredentials", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) CreateWorkspaceSnapshot(ctx context.Context, in *CreateWorkspaceSnapshotRequest, opts ...grpc.CallOption) (*CreateWorkspaceSnapshotResponse, error) {
	out := new(CreateWorkspaceSnapshotResponse)
	err := c.cc.Invoke(ctx, "/gitpod.v1.WorkspaceService/CreateWorkspaceSnapshot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) WaitForWorkspaceSnapshot(ctx context.Context, in *WaitForWorkspaceSnapshotRequest, opts ...grpc.CallOption) (*WaitForWorkspaceSnapshotResponse, error) {
	out := new(WaitForWorkspaceSnapshotResponse)
	err := c.cc.Invoke(ctx, "/gitpod.v1.WorkspaceService/WaitForWorkspaceSnapshot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) UpdateWorkspacePort(ctx context.Context, in *UpdateWorkspacePortRequest, opts ...grpc.CallOption) (*UpdateWorkspacePortResponse, error) {
	out := new(UpdateWorkspacePortResponse)
	err := c.cc.Invoke(ctx, "/gitpod.v1.WorkspaceService/UpdateWorkspacePort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkspaceServiceServer is the server API for WorkspaceService service.
// All implementations must embed UnimplementedWorkspaceServiceServer
// for forward compatibility
type WorkspaceServiceServer interface {
	// GetWorkspace returns a single workspace.
	//
	// +return NOT_FOUND User does not have access to a workspace with the given
	// ID +return NOT_FOUND Workspace does not exist
	GetWorkspace(context.Context, *GetWorkspaceRequest) (*GetWorkspaceResponse, error)
	// WatchWorkspaceStatus watches the workspaces status changes
	//
	// workspace_id +return NOT_FOUND Workspace does not exist
	WatchWorkspaceStatus(*WatchWorkspaceStatusRequest, WorkspaceService_WatchWorkspaceStatusServer) error
	// ListWorkspaces returns a list of workspaces that match the query.
	ListWorkspaces(context.Context, *ListWorkspacesRequest) (*ListWorkspacesResponse, error)
	// ListWorkspaceSessions returns a list of workspace sessions that match the
	ListWorkspaceSessions(context.Context, *ListWorkspaceSessionsRequest) (*ListWorkspaceSessionsResponse, error)
	// CreateAndStartWorkspace creates a new workspace and starts it.
	CreateAndStartWorkspace(context.Context, *CreateAndStartWorkspaceRequest) (*CreateAndStartWorkspaceResponse, error)
	// StartWorkspace starts an existing workspace.
	// If the specified workspace is not in stopped phase, this will return the
	// workspace as is.
	StartWorkspace(context.Context, *StartWorkspaceRequest) (*StartWorkspaceResponse, error)
	// UpdateWorkspace updates the workspace.
	UpdateWorkspace(context.Context, *UpdateWorkspaceRequest) (*UpdateWorkspaceResponse, error)
	// StopWorkspace stops a running workspace.
	StopWorkspace(context.Context, *StopWorkspaceRequest) (*StopWorkspaceResponse, error)
	// DeleteWorkspace deletes a workspace.
	// When the workspace is running, it will be stopped as well.
	// Deleted workspaces cannot be started again.
	DeleteWorkspace(context.Context, *DeleteWorkspaceRequest) (*DeleteWorkspaceResponse, error)
	// ListWorkspaceClasses enumerates all available workspace classes.
	ListWorkspaceClasses(context.Context, *ListWorkspaceClassesRequest) (*ListWorkspaceClassesResponse, error)
	// ParseContextURL parses a context URL and returns the workspace metadata and
	// spec. Not implemented yet.
	ParseContextURL(context.Context, *ParseContextURLRequest) (*ParseContextURLResponse, error)
	// GetWorkspaceDefaultImage returns the default workspace image of specified
	// workspace.
	GetWorkspaceDefaultImage(context.Context, *GetWorkspaceDefaultImageRequest) (*GetWorkspaceDefaultImageResponse, error)
	// SendHeartBeat sends a heartbeat to activate the workspace
	SendHeartBeat(context.Context, *SendHeartBeatRequest) (*SendHeartBeatResponse, error)
	// GetWorkspaceOwnerToken returns an owner token of workspace.
	GetWorkspaceOwnerToken(context.Context, *GetWorkspaceOwnerTokenRequest) (*GetWorkspaceOwnerTokenResponse, error)
	// GetWorkspaceEditorCredentials returns an credentials that is used in editor
	// to encrypt and decrypt secrets
	GetWorkspaceEditorCredentials(context.Context, *GetWorkspaceEditorCredentialsRequest) (*GetWorkspaceEditorCredentialsResponse, error)
	// CreateWorkspaceSnapshot creates a snapshot of the workspace that can be
	// shared with others.
	CreateWorkspaceSnapshot(context.Context, *CreateWorkspaceSnapshotRequest) (*CreateWorkspaceSnapshotResponse, error)
	// WaitWorkspaceSnapshot waits for the snapshot to be available or failed.
	WaitForWorkspaceSnapshot(context.Context, *WaitForWorkspaceSnapshotRequest) (*WaitForWorkspaceSnapshotResponse, error)
	// UpdateWorkspacePort updates the port of workspace.
	UpdateWorkspacePort(context.Context, *UpdateWorkspacePortRequest) (*UpdateWorkspacePortResponse, error)
	mustEmbedUnimplementedWorkspaceServiceServer()
}

// UnimplementedWorkspaceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWorkspaceServiceServer struct {
}

func (UnimplementedWorkspaceServiceServer) GetWorkspace(context.Context, *GetWorkspaceRequest) (*GetWorkspaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkspace not implemented")
}
func (UnimplementedWorkspaceServiceServer) WatchWorkspaceStatus(*WatchWorkspaceStatusRequest, WorkspaceService_WatchWorkspaceStatusServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchWorkspaceStatus not implemented")
}
func (UnimplementedWorkspaceServiceServer) ListWorkspaces(context.Context, *ListWorkspacesRequest) (*ListWorkspacesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWorkspaces not implemented")
}
func (UnimplementedWorkspaceServiceServer) ListWorkspaceSessions(context.Context, *ListWorkspaceSessionsRequest) (*ListWorkspaceSessionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWorkspaceSessions not implemented")
}
func (UnimplementedWorkspaceServiceServer) CreateAndStartWorkspace(context.Context, *CreateAndStartWorkspaceRequest) (*CreateAndStartWorkspaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAndStartWorkspace not implemented")
}
func (UnimplementedWorkspaceServiceServer) StartWorkspace(context.Context, *StartWorkspaceRequest) (*StartWorkspaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartWorkspace not implemented")
}
func (UnimplementedWorkspaceServiceServer) UpdateWorkspace(context.Context, *UpdateWorkspaceRequest) (*UpdateWorkspaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWorkspace not implemented")
}
func (UnimplementedWorkspaceServiceServer) StopWorkspace(context.Context, *StopWorkspaceRequest) (*StopWorkspaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopWorkspace not implemented")
}
func (UnimplementedWorkspaceServiceServer) DeleteWorkspace(context.Context, *DeleteWorkspaceRequest) (*DeleteWorkspaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWorkspace not implemented")
}
func (UnimplementedWorkspaceServiceServer) ListWorkspaceClasses(context.Context, *ListWorkspaceClassesRequest) (*ListWorkspaceClassesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWorkspaceClasses not implemented")
}
func (UnimplementedWorkspaceServiceServer) ParseContextURL(context.Context, *ParseContextURLRequest) (*ParseContextURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ParseContextURL not implemented")
}
func (UnimplementedWorkspaceServiceServer) GetWorkspaceDefaultImage(context.Context, *GetWorkspaceDefaultImageRequest) (*GetWorkspaceDefaultImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkspaceDefaultImage not implemented")
}
func (UnimplementedWorkspaceServiceServer) SendHeartBeat(context.Context, *SendHeartBeatRequest) (*SendHeartBeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendHeartBeat not implemented")
}
func (UnimplementedWorkspaceServiceServer) GetWorkspaceOwnerToken(context.Context, *GetWorkspaceOwnerTokenRequest) (*GetWorkspaceOwnerTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkspaceOwnerToken not implemented")
}
func (UnimplementedWorkspaceServiceServer) GetWorkspaceEditorCredentials(context.Context, *GetWorkspaceEditorCredentialsRequest) (*GetWorkspaceEditorCredentialsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkspaceEditorCredentials not implemented")
}
func (UnimplementedWorkspaceServiceServer) CreateWorkspaceSnapshot(context.Context, *CreateWorkspaceSnapshotRequest) (*CreateWorkspaceSnapshotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWorkspaceSnapshot not implemented")
}
func (UnimplementedWorkspaceServiceServer) WaitForWorkspaceSnapshot(context.Context, *WaitForWorkspaceSnapshotRequest) (*WaitForWorkspaceSnapshotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WaitForWorkspaceSnapshot not implemented")
}
func (UnimplementedWorkspaceServiceServer) UpdateWorkspacePort(context.Context, *UpdateWorkspacePortRequest) (*UpdateWorkspacePortResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWorkspacePort not implemented")
}
func (UnimplementedWorkspaceServiceServer) mustEmbedUnimplementedWorkspaceServiceServer() {}

// UnsafeWorkspaceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkspaceServiceServer will
// result in compilation errors.
type UnsafeWorkspaceServiceServer interface {
	mustEmbedUnimplementedWorkspaceServiceServer()
}

func RegisterWorkspaceServiceServer(s grpc.ServiceRegistrar, srv WorkspaceServiceServer) {
	s.RegisterService(&WorkspaceService_ServiceDesc, srv)
}

func _WorkspaceService_GetWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkspaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).GetWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.v1.WorkspaceService/GetWorkspace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).GetWorkspace(ctx, req.(*GetWorkspaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_WatchWorkspaceStatus_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchWorkspaceStatusRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WorkspaceServiceServer).WatchWorkspaceStatus(m, &workspaceServiceWatchWorkspaceStatusServer{stream})
}

type WorkspaceService_WatchWorkspaceStatusServer interface {
	Send(*WatchWorkspaceStatusResponse) error
	grpc.ServerStream
}

type workspaceServiceWatchWorkspaceStatusServer struct {
	grpc.ServerStream
}

func (x *workspaceServiceWatchWorkspaceStatusServer) Send(m *WatchWorkspaceStatusResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _WorkspaceService_ListWorkspaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWorkspacesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).ListWorkspaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.v1.WorkspaceService/ListWorkspaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).ListWorkspaces(ctx, req.(*ListWorkspacesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_ListWorkspaceSessions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWorkspaceSessionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).ListWorkspaceSessions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.v1.WorkspaceService/ListWorkspaceSessions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).ListWorkspaceSessions(ctx, req.(*ListWorkspaceSessionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_CreateAndStartWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAndStartWorkspaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).CreateAndStartWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.v1.WorkspaceService/CreateAndStartWorkspace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).CreateAndStartWorkspace(ctx, req.(*CreateAndStartWorkspaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_StartWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartWorkspaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).StartWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.v1.WorkspaceService/StartWorkspace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).StartWorkspace(ctx, req.(*StartWorkspaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_UpdateWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWorkspaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).UpdateWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.v1.WorkspaceService/UpdateWorkspace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).UpdateWorkspace(ctx, req.(*UpdateWorkspaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_StopWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopWorkspaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).StopWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.v1.WorkspaceService/StopWorkspace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).StopWorkspace(ctx, req.(*StopWorkspaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_DeleteWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWorkspaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).DeleteWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.v1.WorkspaceService/DeleteWorkspace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).DeleteWorkspace(ctx, req.(*DeleteWorkspaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_ListWorkspaceClasses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWorkspaceClassesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).ListWorkspaceClasses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.v1.WorkspaceService/ListWorkspaceClasses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).ListWorkspaceClasses(ctx, req.(*ListWorkspaceClassesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_ParseContextURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParseContextURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).ParseContextURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.v1.WorkspaceService/ParseContextURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).ParseContextURL(ctx, req.(*ParseContextURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_GetWorkspaceDefaultImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkspaceDefaultImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).GetWorkspaceDefaultImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.v1.WorkspaceService/GetWorkspaceDefaultImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).GetWorkspaceDefaultImage(ctx, req.(*GetWorkspaceDefaultImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_SendHeartBeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendHeartBeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).SendHeartBeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.v1.WorkspaceService/SendHeartBeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).SendHeartBeat(ctx, req.(*SendHeartBeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_GetWorkspaceOwnerToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkspaceOwnerTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).GetWorkspaceOwnerToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.v1.WorkspaceService/GetWorkspaceOwnerToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).GetWorkspaceOwnerToken(ctx, req.(*GetWorkspaceOwnerTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_GetWorkspaceEditorCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkspaceEditorCredentialsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).GetWorkspaceEditorCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.v1.WorkspaceService/GetWorkspaceEditorCredentials",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).GetWorkspaceEditorCredentials(ctx, req.(*GetWorkspaceEditorCredentialsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_CreateWorkspaceSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWorkspaceSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).CreateWorkspaceSnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.v1.WorkspaceService/CreateWorkspaceSnapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).CreateWorkspaceSnapshot(ctx, req.(*CreateWorkspaceSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_WaitForWorkspaceSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WaitForWorkspaceSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).WaitForWorkspaceSnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.v1.WorkspaceService/WaitForWorkspaceSnapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).WaitForWorkspaceSnapshot(ctx, req.(*WaitForWorkspaceSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_UpdateWorkspacePort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWorkspacePortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).UpdateWorkspacePort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.v1.WorkspaceService/UpdateWorkspacePort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).UpdateWorkspacePort(ctx, req.(*UpdateWorkspacePortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WorkspaceService_ServiceDesc is the grpc.ServiceDesc for WorkspaceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WorkspaceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gitpod.v1.WorkspaceService",
	HandlerType: (*WorkspaceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWorkspace",
			Handler:    _WorkspaceService_GetWorkspace_Handler,
		},
		{
			MethodName: "ListWorkspaces",
			Handler:    _WorkspaceService_ListWorkspaces_Handler,
		},
		{
			MethodName: "ListWorkspaceSessions",
			Handler:    _WorkspaceService_ListWorkspaceSessions_Handler,
		},
		{
			MethodName: "CreateAndStartWorkspace",
			Handler:    _WorkspaceService_CreateAndStartWorkspace_Handler,
		},
		{
			MethodName: "StartWorkspace",
			Handler:    _WorkspaceService_StartWorkspace_Handler,
		},
		{
			MethodName: "UpdateWorkspace",
			Handler:    _WorkspaceService_UpdateWorkspace_Handler,
		},
		{
			MethodName: "StopWorkspace",
			Handler:    _WorkspaceService_StopWorkspace_Handler,
		},
		{
			MethodName: "DeleteWorkspace",
			Handler:    _WorkspaceService_DeleteWorkspace_Handler,
		},
		{
			MethodName: "ListWorkspaceClasses",
			Handler:    _WorkspaceService_ListWorkspaceClasses_Handler,
		},
		{
			MethodName: "ParseContextURL",
			Handler:    _WorkspaceService_ParseContextURL_Handler,
		},
		{
			MethodName: "GetWorkspaceDefaultImage",
			Handler:    _WorkspaceService_GetWorkspaceDefaultImage_Handler,
		},
		{
			MethodName: "SendHeartBeat",
			Handler:    _WorkspaceService_SendHeartBeat_Handler,
		},
		{
			MethodName: "GetWorkspaceOwnerToken",
			Handler:    _WorkspaceService_GetWorkspaceOwnerToken_Handler,
		},
		{
			MethodName: "GetWorkspaceEditorCredentials",
			Handler:    _WorkspaceService_GetWorkspaceEditorCredentials_Handler,
		},
		{
			MethodName: "CreateWorkspaceSnapshot",
			Handler:    _WorkspaceService_CreateWorkspaceSnapshot_Handler,
		},
		{
			MethodName: "WaitForWorkspaceSnapshot",
			Handler:    _WorkspaceService_WaitForWorkspaceSnapshot_Handler,
		},
		{
			MethodName: "UpdateWorkspacePort",
			Handler:    _WorkspaceService_UpdateWorkspacePort_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchWorkspaceStatus",
			Handler:       _WorkspaceService_WatchWorkspaceStatus_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "gitpod/v1/workspace.proto",
}
