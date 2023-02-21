// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: simfiny/proto/service/api/v1/service.proto

package gen

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

// FinancialServiceClient is the client API for FinancialService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FinancialServiceClient interface {
	//
	//ServiceHealth: Used to get the service health status of a service
	ServiceHealth(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*ServiceHealthResponse, error)
	//
	//ServiceReady: Used to deduce wether service is ready to serve request
	ServiceReady(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*ServiceReadyResponse, error)
	//
	//CreateAccount: Used to create an account record
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error)
	//
	//DeleteAccount: Used to perform a delete account request
	DeleteAccount(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	//
	//In order to connect a user's account with Plaid and thus their bank account,
	//you need to create a link token. This token is then used by the Plaid SDK in
	//the UI in order to provide an authentication flow for that user and their
	//bank
	GetNewLinkToken(ctx context.Context, in *PlaidLinkTokenRequest, opts ...grpc.CallOption) (*PlaidLinktTokenResponse, error)
	//
	//Once the user has completed authenticating their bank account using the
	//Plaid SDK, you will be given a public_token. This new token must be
	//exchanged with Plaid's API in order to allow simfiny to access the bank
	//data. The bank account should not be considered linked until this token is
	//successfully exchanged.
	PlaidTokenCallback(ctx context.Context, in *PlaidLinkTokenExchangeRequest, opts ...grpc.CallOption) (*PlaidLinkTokenExchangeResponse, error)
	//
	//Plaid links can be updated after they have been established. This can be
	//used as a method of re-authenticating a link if it ends up in an error state
	//(though sometimes a link can end up in an error state without indicating
	//[2]). This can also be used as a way to add additional accounts to a link if
	//those accounts were not originally granted access when the link was created.
	UpdatePlaidLink(ctx context.Context, in *PlaidLinkUpdateRequest, opts ...grpc.CallOption) (*PlaidLinkUpdateResponse, error)
	//
	//Sometimes you might need to manually trigger a resync with Plaid; this can
	//happen if there were issues where a webhook was not properly received. By
	//triggering a manual resync, transactions for the last 14 days and balances
	//for all bank accounts within the specified link will be checked.
	//
	//Links can be manually synced once every 30 minutes, an error will be
	//returned if you try to resync a link too quickly.
	//
	//NOTE - This will not send a "sync" request to Plaid. This will only retrieve
	//data already available via Plaid's API and update simfiny's data
	//accordingly.
	ManuallyResyncPlaidLink(ctx context.Context, in *PlaidManualLinkResyncRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
}

type financialServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFinancialServiceClient(cc grpc.ClientConnInterface) FinancialServiceClient {
	return &financialServiceClient{cc}
}

func (c *financialServiceClient) ServiceHealth(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*ServiceHealthResponse, error) {
	out := new(ServiceHealthResponse)
	err := c.cc.Invoke(ctx, "/endpoint.service.FinancialService/ServiceHealth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financialServiceClient) ServiceReady(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*ServiceReadyResponse, error) {
	out := new(ServiceReadyResponse)
	err := c.cc.Invoke(ctx, "/endpoint.service.FinancialService/ServiceReady", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financialServiceClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error) {
	out := new(CreateAccountResponse)
	err := c.cc.Invoke(ctx, "/endpoint.service.FinancialService/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financialServiceClient) DeleteAccount(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/endpoint.service.FinancialService/DeleteAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financialServiceClient) GetNewLinkToken(ctx context.Context, in *PlaidLinkTokenRequest, opts ...grpc.CallOption) (*PlaidLinktTokenResponse, error) {
	out := new(PlaidLinktTokenResponse)
	err := c.cc.Invoke(ctx, "/endpoint.service.FinancialService/GetNewLinkToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financialServiceClient) PlaidTokenCallback(ctx context.Context, in *PlaidLinkTokenExchangeRequest, opts ...grpc.CallOption) (*PlaidLinkTokenExchangeResponse, error) {
	out := new(PlaidLinkTokenExchangeResponse)
	err := c.cc.Invoke(ctx, "/endpoint.service.FinancialService/PlaidTokenCallback", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financialServiceClient) UpdatePlaidLink(ctx context.Context, in *PlaidLinkUpdateRequest, opts ...grpc.CallOption) (*PlaidLinkUpdateResponse, error) {
	out := new(PlaidLinkUpdateResponse)
	err := c.cc.Invoke(ctx, "/endpoint.service.FinancialService/UpdatePlaidLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financialServiceClient) ManuallyResyncPlaidLink(ctx context.Context, in *PlaidManualLinkResyncRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/endpoint.service.FinancialService/ManuallyResyncPlaidLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FinancialServiceServer is the server API for FinancialService service.
// All implementations must embed UnimplementedFinancialServiceServer
// for forward compatibility
type FinancialServiceServer interface {
	//
	//ServiceHealth: Used to get the service health status of a service
	ServiceHealth(context.Context, *EmptyRequest) (*ServiceHealthResponse, error)
	//
	//ServiceReady: Used to deduce wether service is ready to serve request
	ServiceReady(context.Context, *EmptyRequest) (*ServiceReadyResponse, error)
	//
	//CreateAccount: Used to create an account record
	CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error)
	//
	//DeleteAccount: Used to perform a delete account request
	DeleteAccount(context.Context, *DeleteAccountRequest) (*EmptyResponse, error)
	//
	//In order to connect a user's account with Plaid and thus their bank account,
	//you need to create a link token. This token is then used by the Plaid SDK in
	//the UI in order to provide an authentication flow for that user and their
	//bank
	GetNewLinkToken(context.Context, *PlaidLinkTokenRequest) (*PlaidLinktTokenResponse, error)
	//
	//Once the user has completed authenticating their bank account using the
	//Plaid SDK, you will be given a public_token. This new token must be
	//exchanged with Plaid's API in order to allow simfiny to access the bank
	//data. The bank account should not be considered linked until this token is
	//successfully exchanged.
	PlaidTokenCallback(context.Context, *PlaidLinkTokenExchangeRequest) (*PlaidLinkTokenExchangeResponse, error)
	//
	//Plaid links can be updated after they have been established. This can be
	//used as a method of re-authenticating a link if it ends up in an error state
	//(though sometimes a link can end up in an error state without indicating
	//[2]). This can also be used as a way to add additional accounts to a link if
	//those accounts were not originally granted access when the link was created.
	UpdatePlaidLink(context.Context, *PlaidLinkUpdateRequest) (*PlaidLinkUpdateResponse, error)
	//
	//Sometimes you might need to manually trigger a resync with Plaid; this can
	//happen if there were issues where a webhook was not properly received. By
	//triggering a manual resync, transactions for the last 14 days and balances
	//for all bank accounts within the specified link will be checked.
	//
	//Links can be manually synced once every 30 minutes, an error will be
	//returned if you try to resync a link too quickly.
	//
	//NOTE - This will not send a "sync" request to Plaid. This will only retrieve
	//data already available via Plaid's API and update simfiny's data
	//accordingly.
	ManuallyResyncPlaidLink(context.Context, *PlaidManualLinkResyncRequest) (*EmptyResponse, error)
	mustEmbedUnimplementedFinancialServiceServer()
}

// UnimplementedFinancialServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFinancialServiceServer struct {
}

func (UnimplementedFinancialServiceServer) ServiceHealth(context.Context, *EmptyRequest) (*ServiceHealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServiceHealth not implemented")
}
func (UnimplementedFinancialServiceServer) ServiceReady(context.Context, *EmptyRequest) (*ServiceReadyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServiceReady not implemented")
}
func (UnimplementedFinancialServiceServer) CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedFinancialServiceServer) DeleteAccount(context.Context, *DeleteAccountRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccount not implemented")
}
func (UnimplementedFinancialServiceServer) GetNewLinkToken(context.Context, *PlaidLinkTokenRequest) (*PlaidLinktTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNewLinkToken not implemented")
}
func (UnimplementedFinancialServiceServer) PlaidTokenCallback(context.Context, *PlaidLinkTokenExchangeRequest) (*PlaidLinkTokenExchangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaidTokenCallback not implemented")
}
func (UnimplementedFinancialServiceServer) UpdatePlaidLink(context.Context, *PlaidLinkUpdateRequest) (*PlaidLinkUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePlaidLink not implemented")
}
func (UnimplementedFinancialServiceServer) ManuallyResyncPlaidLink(context.Context, *PlaidManualLinkResyncRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ManuallyResyncPlaidLink not implemented")
}
func (UnimplementedFinancialServiceServer) mustEmbedUnimplementedFinancialServiceServer() {}

// UnsafeFinancialServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FinancialServiceServer will
// result in compilation errors.
type UnsafeFinancialServiceServer interface {
	mustEmbedUnimplementedFinancialServiceServer()
}

func RegisterFinancialServiceServer(s grpc.ServiceRegistrar, srv FinancialServiceServer) {
	s.RegisterService(&FinancialService_ServiceDesc, srv)
}

func _FinancialService_ServiceHealth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinancialServiceServer).ServiceHealth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/endpoint.service.FinancialService/ServiceHealth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinancialServiceServer).ServiceHealth(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinancialService_ServiceReady_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinancialServiceServer).ServiceReady(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/endpoint.service.FinancialService/ServiceReady",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinancialServiceServer).ServiceReady(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinancialService_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinancialServiceServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/endpoint.service.FinancialService/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinancialServiceServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinancialService_DeleteAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinancialServiceServer).DeleteAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/endpoint.service.FinancialService/DeleteAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinancialServiceServer).DeleteAccount(ctx, req.(*DeleteAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinancialService_GetNewLinkToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaidLinkTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinancialServiceServer).GetNewLinkToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/endpoint.service.FinancialService/GetNewLinkToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinancialServiceServer).GetNewLinkToken(ctx, req.(*PlaidLinkTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinancialService_PlaidTokenCallback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaidLinkTokenExchangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinancialServiceServer).PlaidTokenCallback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/endpoint.service.FinancialService/PlaidTokenCallback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinancialServiceServer).PlaidTokenCallback(ctx, req.(*PlaidLinkTokenExchangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinancialService_UpdatePlaidLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaidLinkUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinancialServiceServer).UpdatePlaidLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/endpoint.service.FinancialService/UpdatePlaidLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinancialServiceServer).UpdatePlaidLink(ctx, req.(*PlaidLinkUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinancialService_ManuallyResyncPlaidLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaidManualLinkResyncRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinancialServiceServer).ManuallyResyncPlaidLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/endpoint.service.FinancialService/ManuallyResyncPlaidLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinancialServiceServer).ManuallyResyncPlaidLink(ctx, req.(*PlaidManualLinkResyncRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FinancialService_ServiceDesc is the grpc.ServiceDesc for FinancialService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FinancialService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "endpoint.service.FinancialService",
	HandlerType: (*FinancialServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ServiceHealth",
			Handler:    _FinancialService_ServiceHealth_Handler,
		},
		{
			MethodName: "ServiceReady",
			Handler:    _FinancialService_ServiceReady_Handler,
		},
		{
			MethodName: "CreateAccount",
			Handler:    _FinancialService_CreateAccount_Handler,
		},
		{
			MethodName: "DeleteAccount",
			Handler:    _FinancialService_DeleteAccount_Handler,
		},
		{
			MethodName: "GetNewLinkToken",
			Handler:    _FinancialService_GetNewLinkToken_Handler,
		},
		{
			MethodName: "PlaidTokenCallback",
			Handler:    _FinancialService_PlaidTokenCallback_Handler,
		},
		{
			MethodName: "UpdatePlaidLink",
			Handler:    _FinancialService_UpdatePlaidLink_Handler,
		},
		{
			MethodName: "ManuallyResyncPlaidLink",
			Handler:    _FinancialService_ManuallyResyncPlaidLink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "simfiny/proto/service/api/v1/service.proto",
}
