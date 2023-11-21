// Code generated by protoc-gen-triple. DO NOT EDIT.
//
// Source: greet.proto
package greettriple

import (
	"context"
)

import (
	"dubbo.apache.org/dubbo-go/v3/client"
	"dubbo.apache.org/dubbo-go/v3/common/constant"
	"dubbo.apache.org/dubbo-go/v3/protocol/triple/triple_protocol"
	"dubbo.apache.org/dubbo-go/v3/server"
)

import (
	proto "github.com/apache/dubbo-go-samples/helloworld/proto"
)

// This is a compile-time assertion to ensure that this generated file and the Triple package
// are compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of Triple newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of Triple or updating the Triple
// version compiled into your binary.
const _ = triple_protocol.IsAtLeastVersion0_1_0

const (
	// GreetServiceName is the fully-qualified name of the GreetService service.
	GreetServiceName = "greet.GreetService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// GreetServiceGreetProcedure is the fully-qualified name of the GreetService's Greet RPC.
	GreetServiceGreetProcedure = "/greet.GreetService/Greet"
)

var (
	_ GreetService = (*GreetServiceImpl)(nil)
)

// GreetService is a client for the greet.GreetService service.
type GreetService interface {
	Greet(ctx context.Context, req *proto.GreetRequest, opts ...client.CallOption) (*proto.GreetResponse, error)
}

// NewGreetService constructs a client for the greet.GreetService service.
func NewGreetService(cli *client.Client, opts ...client.ReferenceOption) (GreetService, error) {
	group, version, err := cli.Init(&GreetService_ClientInfo, opts...)
	if err != nil {
		return nil, err
	}

	return &GreetServiceImpl{
		cli:     cli,
		group:   group,
		version: version,
	}, nil
}

// GreetServiceImpl implements GreetService.
type GreetServiceImpl struct {
	cli     *client.Client
	group   string
	version string
}

func (c *GreetServiceImpl) Greet(ctx context.Context, req *proto.GreetRequest, opts ...client.CallOption) (*proto.GreetResponse, error) {
	opts = appendGroupVersion(opts, c)
	resp := new(proto.GreetResponse)
	if err := c.cli.CallUnary(ctx, req, resp, "greet.GreetService", "Greet", opts...); err != nil {
		return nil, err
	}
	return resp, nil
}

func appendGroupVersion(opts []client.CallOption, c *GreetServiceImpl) []client.CallOption {
	opts = append(opts, client.WithCallGroup(c.group))
	opts = append(opts, client.WithCallVersion(c.version))
	return opts
}

var GreetService_ClientInfo = client.ClientInfo{
	InterfaceName: "greet.GreetService",
	MethodNames:   []string{"Greet"},
	ClientInjectFunc: func(dubboCliRaw interface{}, cli *client.Client) {
		dubboCli := dubboCliRaw.(GreetServiceImpl)
		dubboCli.cli = cli
	},
}

// GreetServiceHandler is an implementation of the greet.GreetService service.
type GreetServiceHandler interface {
	Greet(context.Context, *proto.GreetRequest) (*proto.GreetResponse, error)
}

func RegisterGreetServiceHandler(srv *server.Server, hdlr GreetServiceHandler, opts ...server.ServiceOption) error {
	return srv.Register(hdlr, &GreetService_ServiceInfo, opts...)
}

var GreetService_ServiceInfo = server.ServiceInfo{
	InterfaceName: "greet.GreetService",
	ServiceType:   (*GreetServiceHandler)(nil),
	Methods: []server.MethodInfo{
		{
			Name: "Greet",
			Type: constant.CallUnary,
			ReqInitFunc: func() interface{} {
				return new(proto.GreetRequest)
			},
			MethodFunc: func(ctx context.Context, args []interface{}, handler interface{}) (interface{}, error) {
				req := args[0].(*proto.GreetRequest)
				res, err := handler.(GreetServiceHandler).Greet(ctx, req)
				if err != nil {
					return nil, err
				}
				return triple_protocol.NewResponse(res), nil
			},
		},
	},
}
