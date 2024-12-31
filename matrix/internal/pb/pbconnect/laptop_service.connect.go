// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: laptop_service.proto

package pbconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	pb "matrix/internal/pb"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// LaptopServiceName is the fully-qualified name of the LaptopService service.
	LaptopServiceName = "pb.LaptopService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// LaptopServiceCreateLaptopProcedure is the fully-qualified name of the LaptopService's
	// CreateLaptop RPC.
	LaptopServiceCreateLaptopProcedure = "/pb.LaptopService/CreateLaptop"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	laptopServiceServiceDescriptor            = pb.File_laptop_service_proto.Services().ByName("LaptopService")
	laptopServiceCreateLaptopMethodDescriptor = laptopServiceServiceDescriptor.Methods().ByName("CreateLaptop")
)

// LaptopServiceClient is a client for the pb.LaptopService service.
type LaptopServiceClient interface {
	CreateLaptop(context.Context, *connect.Request[pb.CreateLaptopRequest]) (*connect.Response[pb.CreateLaptopResponse], error)
}

// NewLaptopServiceClient constructs a client for the pb.LaptopService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewLaptopServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) LaptopServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &laptopServiceClient{
		createLaptop: connect.NewClient[pb.CreateLaptopRequest, pb.CreateLaptopResponse](
			httpClient,
			baseURL+LaptopServiceCreateLaptopProcedure,
			connect.WithSchema(laptopServiceCreateLaptopMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// laptopServiceClient implements LaptopServiceClient.
type laptopServiceClient struct {
	createLaptop *connect.Client[pb.CreateLaptopRequest, pb.CreateLaptopResponse]
}

// CreateLaptop calls pb.LaptopService.CreateLaptop.
func (c *laptopServiceClient) CreateLaptop(ctx context.Context, req *connect.Request[pb.CreateLaptopRequest]) (*connect.Response[pb.CreateLaptopResponse], error) {
	return c.createLaptop.CallUnary(ctx, req)
}

// LaptopServiceHandler is an implementation of the pb.LaptopService service.
type LaptopServiceHandler interface {
	CreateLaptop(context.Context, *connect.Request[pb.CreateLaptopRequest]) (*connect.Response[pb.CreateLaptopResponse], error)
}

// NewLaptopServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewLaptopServiceHandler(svc LaptopServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	laptopServiceCreateLaptopHandler := connect.NewUnaryHandler(
		LaptopServiceCreateLaptopProcedure,
		svc.CreateLaptop,
		connect.WithSchema(laptopServiceCreateLaptopMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/pb.LaptopService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case LaptopServiceCreateLaptopProcedure:
			laptopServiceCreateLaptopHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedLaptopServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedLaptopServiceHandler struct{}

func (UnimplementedLaptopServiceHandler) CreateLaptop(context.Context, *connect.Request[pb.CreateLaptopRequest]) (*connect.Response[pb.CreateLaptopResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("pb.LaptopService.CreateLaptop is not implemented"))
}
