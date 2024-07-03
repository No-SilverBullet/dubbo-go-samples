// Code generated by protoc-gen-triple. DO NOT EDIT.
//
// Source: proto/recommendationservice.proto
package hipstershop

import (
	"context"
)

import (
	"dubbo.apache.org/dubbo-go/v3"
	"dubbo.apache.org/dubbo-go/v3/client"
	"dubbo.apache.org/dubbo-go/v3/common"
	"dubbo.apache.org/dubbo-go/v3/common/constant"
	"dubbo.apache.org/dubbo-go/v3/protocol/triple/triple_protocol"
	"dubbo.apache.org/dubbo-go/v3/server"
)

// This is a compile-time assertion to ensure that this generated file and the Triple package
// are compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of Triple newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of Triple or updating the Triple
// version compiled into your binary.
const _ = triple_protocol.IsAtLeastVersion0_1_0

const (
	// RecommendationServiceName is the fully-qualified name of the RecommendationService service.
	RecommendationServiceName = "hipstershop.RecommendationService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// RecommendationServiceListRecommendationsProcedure is the fully-qualified name of the RecommendationService's ListRecommendations RPC.
	RecommendationServiceListRecommendationsProcedure = "/hipstershop.RecommendationService/ListRecommendations"
)
const (
	// ProductCatalogServiceName is the fully-qualified name of the ProductCatalogService service.
	ProductCatalogServiceName = "hipstershop.ProductCatalogService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ProductCatalogServiceListProductsProcedure is the fully-qualified name of the ProductCatalogService's ListProducts RPC.
	ProductCatalogServiceListProductsProcedure = "/hipstershop.ProductCatalogService/ListProducts"
	// ProductCatalogServiceGetProductProcedure is the fully-qualified name of the ProductCatalogService's GetProduct RPC.
	ProductCatalogServiceGetProductProcedure = "/hipstershop.ProductCatalogService/GetProduct"
	// ProductCatalogServiceSearchProductsProcedure is the fully-qualified name of the ProductCatalogService's SearchProducts RPC.
	ProductCatalogServiceSearchProductsProcedure = "/hipstershop.ProductCatalogService/SearchProducts"
)

var (
	_ RecommendationService = (*RecommendationServiceImpl)(nil)

	_ ProductCatalogService = (*ProductCatalogServiceImpl)(nil)
)

// RecommendationService is a client for the hipstershop.RecommendationService service.
type RecommendationService interface {
	ListRecommendations(ctx context.Context, req *ListRecommendationsRequest, opts ...client.CallOption) (*ListRecommendationsResponse, error)
}

// ProductCatalogService is a client for the hipstershop.ProductCatalogService service.
type ProductCatalogService interface {
	ListProducts(ctx context.Context, req *Empty, opts ...client.CallOption) (*ListProductsResponse, error)
	GetProduct(ctx context.Context, req *GetProductRequest, opts ...client.CallOption) (*Product, error)
	SearchProducts(ctx context.Context, req *SearchProductsRequest, opts ...client.CallOption) (*SearchProductsResponse, error)
}

// NewRecommendationService constructs a client for the hipstershop.RecommendationService service.
func NewRecommendationService(cli *client.Client, opts ...client.ReferenceOption) (RecommendationService, error) {
	conn, err := cli.DialWithInfo("hipstershop.RecommendationService", &RecommendationService_ClientInfo, opts...)
	if err != nil {
		return nil, err
	}
	return &RecommendationServiceImpl{
		conn: conn,
	}, nil
}

func SetConsumerRecommendationService(srv common.RPCService) {
	dubbo.SetConsumerServiceWithInfo(srv, &RecommendationService_ClientInfo)
}

// RecommendationServiceImpl implements RecommendationService.
type RecommendationServiceImpl struct {
	conn *client.Connection
}

func (c *RecommendationServiceImpl) ListRecommendations(ctx context.Context, req *ListRecommendationsRequest, opts ...client.CallOption) (*ListRecommendationsResponse, error) {
	resp := new(ListRecommendationsResponse)
	if err := c.conn.CallUnary(ctx, []interface{}{req}, resp, "ListRecommendations", opts...); err != nil {
		return nil, err
	}
	return resp, nil
}

// NewProductCatalogService constructs a client for the hipstershop.ProductCatalogService service.
func NewProductCatalogService(cli *client.Client, opts ...client.ReferenceOption) (ProductCatalogService, error) {
	conn, err := cli.DialWithInfo("hipstershop.ProductCatalogService", &ProductCatalogService_ClientInfo, opts...)
	if err != nil {
		return nil, err
	}
	return &ProductCatalogServiceImpl{
		conn: conn,
	}, nil
}

func SetConsumerProductCatalogService(srv common.RPCService) {
	dubbo.SetConsumerServiceWithInfo(srv, &ProductCatalogService_ClientInfo)
}

// ProductCatalogServiceImpl implements ProductCatalogService.
type ProductCatalogServiceImpl struct {
	conn *client.Connection
}

func (c *ProductCatalogServiceImpl) ListProducts(ctx context.Context, req *Empty, opts ...client.CallOption) (*ListProductsResponse, error) {
	resp := new(ListProductsResponse)
	if err := c.conn.CallUnary(ctx, []interface{}{req}, resp, "ListProducts", opts...); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ProductCatalogServiceImpl) GetProduct(ctx context.Context, req *GetProductRequest, opts ...client.CallOption) (*Product, error) {
	resp := new(Product)
	if err := c.conn.CallUnary(ctx, []interface{}{req}, resp, "GetProduct", opts...); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ProductCatalogServiceImpl) SearchProducts(ctx context.Context, req *SearchProductsRequest, opts ...client.CallOption) (*SearchProductsResponse, error) {
	resp := new(SearchProductsResponse)
	if err := c.conn.CallUnary(ctx, []interface{}{req}, resp, "SearchProducts", opts...); err != nil {
		return nil, err
	}
	return resp, nil
}

var RecommendationService_ClientInfo = client.ClientInfo{
	InterfaceName: "hipstershop.RecommendationService",
	MethodNames:   []string{"ListRecommendations"},
	ConnectionInjectFunc: func(dubboCliRaw interface{}, conn *client.Connection) {
		dubboCli := dubboCliRaw.(*RecommendationServiceImpl)
		dubboCli.conn = conn
	},
}
var ProductCatalogService_ClientInfo = client.ClientInfo{
	InterfaceName: "hipstershop.ProductCatalogService",
	MethodNames:   []string{"ListProducts", "GetProduct", "SearchProducts"},
	ConnectionInjectFunc: func(dubboCliRaw interface{}, conn *client.Connection) {
		dubboCli := dubboCliRaw.(*ProductCatalogServiceImpl)
		dubboCli.conn = conn
	},
}

// RecommendationServiceHandler is an implementation of the hipstershop.RecommendationService service.
type RecommendationServiceHandler interface {
	ListRecommendations(context.Context, *ListRecommendationsRequest) (*ListRecommendationsResponse, error)
}

func RegisterRecommendationServiceHandler(srv *server.Server, hdlr RecommendationServiceHandler, opts ...server.ServiceOption) error {
	return srv.Register(hdlr, &RecommendationService_ServiceInfo, opts...)
}

func SetProviderRecommendationService(srv common.RPCService) {
	dubbo.SetProviderServiceWithInfo(srv, &RecommendationService_ServiceInfo)
}

// ProductCatalogServiceHandler is an implementation of the hipstershop.ProductCatalogService service.
type ProductCatalogServiceHandler interface {
	ListProducts(context.Context, *Empty) (*ListProductsResponse, error)
	GetProduct(context.Context, *GetProductRequest) (*Product, error)
	SearchProducts(context.Context, *SearchProductsRequest) (*SearchProductsResponse, error)
}

func RegisterProductCatalogServiceHandler(srv *server.Server, hdlr ProductCatalogServiceHandler, opts ...server.ServiceOption) error {
	return srv.Register(hdlr, &ProductCatalogService_ServiceInfo, opts...)
}

func SetProviderProductCatalogService(srv common.RPCService) {
	dubbo.SetProviderServiceWithInfo(srv, &ProductCatalogService_ServiceInfo)
}

var RecommendationService_ServiceInfo = server.ServiceInfo{
	InterfaceName: "hipstershop.RecommendationService",
	ServiceType:   (*RecommendationServiceHandler)(nil),
	Methods: []server.MethodInfo{
		{
			Name: "ListRecommendations",
			Type: constant.CallUnary,
			ReqInitFunc: func() interface{} {
				return new(ListRecommendationsRequest)
			},
			MethodFunc: func(ctx context.Context, args []interface{}, handler interface{}) (interface{}, error) {
				req := args[0].(*ListRecommendationsRequest)
				res, err := handler.(RecommendationServiceHandler).ListRecommendations(ctx, req)
				if err != nil {
					return nil, err
				}
				return triple_protocol.NewResponse(res), nil
			},
		},
	},
}
var ProductCatalogService_ServiceInfo = server.ServiceInfo{
	InterfaceName: "hipstershop.ProductCatalogService",
	ServiceType:   (*ProductCatalogServiceHandler)(nil),
	Methods: []server.MethodInfo{
		{
			Name: "ListProducts",
			Type: constant.CallUnary,
			ReqInitFunc: func() interface{} {
				return new(Empty)
			},
			MethodFunc: func(ctx context.Context, args []interface{}, handler interface{}) (interface{}, error) {
				req := args[0].(*Empty)
				res, err := handler.(ProductCatalogServiceHandler).ListProducts(ctx, req)
				if err != nil {
					return nil, err
				}
				return triple_protocol.NewResponse(res), nil
			},
		},
		{
			Name: "GetProduct",
			Type: constant.CallUnary,
			ReqInitFunc: func() interface{} {
				return new(GetProductRequest)
			},
			MethodFunc: func(ctx context.Context, args []interface{}, handler interface{}) (interface{}, error) {
				req := args[0].(*GetProductRequest)
				res, err := handler.(ProductCatalogServiceHandler).GetProduct(ctx, req)
				if err != nil {
					return nil, err
				}
				return triple_protocol.NewResponse(res), nil
			},
		},
		{
			Name: "SearchProducts",
			Type: constant.CallUnary,
			ReqInitFunc: func() interface{} {
				return new(SearchProductsRequest)
			},
			MethodFunc: func(ctx context.Context, args []interface{}, handler interface{}) (interface{}, error) {
				req := args[0].(*SearchProductsRequest)
				res, err := handler.(ProductCatalogServiceHandler).SearchProducts(ctx, req)
				if err != nil {
					return nil, err
				}
				return triple_protocol.NewResponse(res), nil
			},
		},
	},
}
