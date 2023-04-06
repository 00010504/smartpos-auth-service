// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: order_main.proto

package order_service

import (
	context "context"
	common "genproto/common"
	report_service "genproto/report_service"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderServiceClient interface {
	// order
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*GetOrderByIDResponse, error)
	GetOrderById(ctx context.Context, in *common.RequestID, opts ...grpc.CallOption) (*GetOrderByIDResponse, error)
	GetAllOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (*GetAllOrdersResponse, error)
	DeleteOrderById(ctx context.Context, in *common.RequestID, opts ...grpc.CallOption) (*common.ResponseID, error)
	// order item
	CreateOrderItem(ctx context.Context, in *CreateOrderItemRequest, opts ...grpc.CallOption) (*GetOrderByIDResponse, error)
	UpdateOrderItem(ctx context.Context, in *UpdateOrderItemRequest, opts ...grpc.CallOption) (*GetOrderByIDResponse, error)
	DeleteOrderItem(ctx context.Context, in *common.RequestID, opts ...grpc.CallOption) (*GetOrderByIDResponse, error)
	// order discount
	UpsertOrderDiscount(ctx context.Context, in *UpsertOrderDiscountRequest, opts ...grpc.CallOption) (*GetOrderByIDResponse, error)
	// order pays
	CreateOrderPays(ctx context.Context, in *CreateOrderPaysRequest, opts ...grpc.CallOption) (*CreateOrderPaysResponse, error)
	// Payment types For dashboardGet
	GetPaymentTypeAnalytics(ctx context.Context, in *report_service.GetDashboardAnalyticsReq, opts ...grpc.CallOption) (*GetPaymentTypeAnalyticsResponse, error)
	AddClientToOrder(ctx context.Context, in *AddClientToOrderRequest, opts ...grpc.CallOption) (*GetOrderByIDResponse, error)
	RemoveClientFromOrder(ctx context.Context, in *RemoveClientFromOrderRequest, opts ...grpc.CallOption) (*GetOrderByIDResponse, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*GetOrderByIDResponse, error) {
	out := new(GetOrderByIDResponse)
	err := c.cc.Invoke(ctx, "/OrderService/CreateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetOrderById(ctx context.Context, in *common.RequestID, opts ...grpc.CallOption) (*GetOrderByIDResponse, error) {
	out := new(GetOrderByIDResponse)
	err := c.cc.Invoke(ctx, "/OrderService/GetOrderById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetAllOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (*GetAllOrdersResponse, error) {
	out := new(GetAllOrdersResponse)
	err := c.cc.Invoke(ctx, "/OrderService/GetAllOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) DeleteOrderById(ctx context.Context, in *common.RequestID, opts ...grpc.CallOption) (*common.ResponseID, error) {
	out := new(common.ResponseID)
	err := c.cc.Invoke(ctx, "/OrderService/DeleteOrderById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) CreateOrderItem(ctx context.Context, in *CreateOrderItemRequest, opts ...grpc.CallOption) (*GetOrderByIDResponse, error) {
	out := new(GetOrderByIDResponse)
	err := c.cc.Invoke(ctx, "/OrderService/CreateOrderItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) UpdateOrderItem(ctx context.Context, in *UpdateOrderItemRequest, opts ...grpc.CallOption) (*GetOrderByIDResponse, error) {
	out := new(GetOrderByIDResponse)
	err := c.cc.Invoke(ctx, "/OrderService/UpdateOrderItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) DeleteOrderItem(ctx context.Context, in *common.RequestID, opts ...grpc.CallOption) (*GetOrderByIDResponse, error) {
	out := new(GetOrderByIDResponse)
	err := c.cc.Invoke(ctx, "/OrderService/DeleteOrderItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) UpsertOrderDiscount(ctx context.Context, in *UpsertOrderDiscountRequest, opts ...grpc.CallOption) (*GetOrderByIDResponse, error) {
	out := new(GetOrderByIDResponse)
	err := c.cc.Invoke(ctx, "/OrderService/UpsertOrderDiscount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) CreateOrderPays(ctx context.Context, in *CreateOrderPaysRequest, opts ...grpc.CallOption) (*CreateOrderPaysResponse, error) {
	out := new(CreateOrderPaysResponse)
	err := c.cc.Invoke(ctx, "/OrderService/CreateOrderPays", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetPaymentTypeAnalytics(ctx context.Context, in *report_service.GetDashboardAnalyticsReq, opts ...grpc.CallOption) (*GetPaymentTypeAnalyticsResponse, error) {
	out := new(GetPaymentTypeAnalyticsResponse)
	err := c.cc.Invoke(ctx, "/OrderService/GetPaymentTypeAnalytics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) AddClientToOrder(ctx context.Context, in *AddClientToOrderRequest, opts ...grpc.CallOption) (*GetOrderByIDResponse, error) {
	out := new(GetOrderByIDResponse)
	err := c.cc.Invoke(ctx, "/OrderService/AddClientToOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) RemoveClientFromOrder(ctx context.Context, in *RemoveClientFromOrderRequest, opts ...grpc.CallOption) (*GetOrderByIDResponse, error) {
	out := new(GetOrderByIDResponse)
	err := c.cc.Invoke(ctx, "/OrderService/RemoveClientFromOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
// All implementations should embed UnimplementedOrderServiceServer
// for forward compatibility
type OrderServiceServer interface {
	// order
	CreateOrder(context.Context, *CreateOrderRequest) (*GetOrderByIDResponse, error)
	GetOrderById(context.Context, *common.RequestID) (*GetOrderByIDResponse, error)
	GetAllOrders(context.Context, *GetOrdersRequest) (*GetAllOrdersResponse, error)
	DeleteOrderById(context.Context, *common.RequestID) (*common.ResponseID, error)
	// order item
	CreateOrderItem(context.Context, *CreateOrderItemRequest) (*GetOrderByIDResponse, error)
	UpdateOrderItem(context.Context, *UpdateOrderItemRequest) (*GetOrderByIDResponse, error)
	DeleteOrderItem(context.Context, *common.RequestID) (*GetOrderByIDResponse, error)
	// order discount
	UpsertOrderDiscount(context.Context, *UpsertOrderDiscountRequest) (*GetOrderByIDResponse, error)
	// order pays
	CreateOrderPays(context.Context, *CreateOrderPaysRequest) (*CreateOrderPaysResponse, error)
	// Payment types For dashboardGet
	GetPaymentTypeAnalytics(context.Context, *report_service.GetDashboardAnalyticsReq) (*GetPaymentTypeAnalyticsResponse, error)
	AddClientToOrder(context.Context, *AddClientToOrderRequest) (*GetOrderByIDResponse, error)
	RemoveClientFromOrder(context.Context, *RemoveClientFromOrderRequest) (*GetOrderByIDResponse, error)
}

// UnimplementedOrderServiceServer should be embedded to have forward compatible implementations.
type UnimplementedOrderServiceServer struct {
}

func (UnimplementedOrderServiceServer) CreateOrder(context.Context, *CreateOrderRequest) (*GetOrderByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedOrderServiceServer) GetOrderById(context.Context, *common.RequestID) (*GetOrderByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderById not implemented")
}
func (UnimplementedOrderServiceServer) GetAllOrders(context.Context, *GetOrdersRequest) (*GetAllOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllOrders not implemented")
}
func (UnimplementedOrderServiceServer) DeleteOrderById(context.Context, *common.RequestID) (*common.ResponseID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrderById not implemented")
}
func (UnimplementedOrderServiceServer) CreateOrderItem(context.Context, *CreateOrderItemRequest) (*GetOrderByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrderItem not implemented")
}
func (UnimplementedOrderServiceServer) UpdateOrderItem(context.Context, *UpdateOrderItemRequest) (*GetOrderByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrderItem not implemented")
}
func (UnimplementedOrderServiceServer) DeleteOrderItem(context.Context, *common.RequestID) (*GetOrderByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrderItem not implemented")
}
func (UnimplementedOrderServiceServer) UpsertOrderDiscount(context.Context, *UpsertOrderDiscountRequest) (*GetOrderByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertOrderDiscount not implemented")
}
func (UnimplementedOrderServiceServer) CreateOrderPays(context.Context, *CreateOrderPaysRequest) (*CreateOrderPaysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrderPays not implemented")
}
func (UnimplementedOrderServiceServer) GetPaymentTypeAnalytics(context.Context, *report_service.GetDashboardAnalyticsReq) (*GetPaymentTypeAnalyticsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPaymentTypeAnalytics not implemented")
}
func (UnimplementedOrderServiceServer) AddClientToOrder(context.Context, *AddClientToOrderRequest) (*GetOrderByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddClientToOrder not implemented")
}
func (UnimplementedOrderServiceServer) RemoveClientFromOrder(context.Context, *RemoveClientFromOrderRequest) (*GetOrderByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveClientFromOrder not implemented")
}

// UnsafeOrderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServiceServer will
// result in compilation errors.
type UnsafeOrderServiceServer interface {
	mustEmbedUnimplementedOrderServiceServer()
}

func RegisterOrderServiceServer(s grpc.ServiceRegistrar, srv OrderServiceServer) {
	s.RegisterService(&OrderService_ServiceDesc, srv)
}

func _OrderService_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/CreateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CreateOrder(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetOrderById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.RequestID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetOrderById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/GetOrderById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetOrderById(ctx, req.(*common.RequestID))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetAllOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetAllOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/GetAllOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetAllOrders(ctx, req.(*GetOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_DeleteOrderById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.RequestID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).DeleteOrderById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/DeleteOrderById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).DeleteOrderById(ctx, req.(*common.RequestID))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_CreateOrderItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CreateOrderItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/CreateOrderItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CreateOrderItem(ctx, req.(*CreateOrderItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_UpdateOrderItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrderItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).UpdateOrderItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/UpdateOrderItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).UpdateOrderItem(ctx, req.(*UpdateOrderItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_DeleteOrderItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.RequestID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).DeleteOrderItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/DeleteOrderItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).DeleteOrderItem(ctx, req.(*common.RequestID))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_UpsertOrderDiscount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertOrderDiscountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).UpsertOrderDiscount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/UpsertOrderDiscount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).UpsertOrderDiscount(ctx, req.(*UpsertOrderDiscountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_CreateOrderPays_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderPaysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CreateOrderPays(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/CreateOrderPays",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CreateOrderPays(ctx, req.(*CreateOrderPaysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetPaymentTypeAnalytics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(report_service.GetDashboardAnalyticsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetPaymentTypeAnalytics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/GetPaymentTypeAnalytics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetPaymentTypeAnalytics(ctx, req.(*report_service.GetDashboardAnalyticsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_AddClientToOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddClientToOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).AddClientToOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/AddClientToOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).AddClientToOrder(ctx, req.(*AddClientToOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_RemoveClientFromOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveClientFromOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).RemoveClientFromOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/RemoveClientFromOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).RemoveClientFromOrder(ctx, req.(*RemoveClientFromOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderService_ServiceDesc is the grpc.ServiceDesc for OrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _OrderService_CreateOrder_Handler,
		},
		{
			MethodName: "GetOrderById",
			Handler:    _OrderService_GetOrderById_Handler,
		},
		{
			MethodName: "GetAllOrders",
			Handler:    _OrderService_GetAllOrders_Handler,
		},
		{
			MethodName: "DeleteOrderById",
			Handler:    _OrderService_DeleteOrderById_Handler,
		},
		{
			MethodName: "CreateOrderItem",
			Handler:    _OrderService_CreateOrderItem_Handler,
		},
		{
			MethodName: "UpdateOrderItem",
			Handler:    _OrderService_UpdateOrderItem_Handler,
		},
		{
			MethodName: "DeleteOrderItem",
			Handler:    _OrderService_DeleteOrderItem_Handler,
		},
		{
			MethodName: "UpsertOrderDiscount",
			Handler:    _OrderService_UpsertOrderDiscount_Handler,
		},
		{
			MethodName: "CreateOrderPays",
			Handler:    _OrderService_CreateOrderPays_Handler,
		},
		{
			MethodName: "GetPaymentTypeAnalytics",
			Handler:    _OrderService_GetPaymentTypeAnalytics_Handler,
		},
		{
			MethodName: "AddClientToOrder",
			Handler:    _OrderService_AddClientToOrder_Handler,
		},
		{
			MethodName: "RemoveClientFromOrder",
			Handler:    _OrderService_RemoveClientFromOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order_main.proto",
}
