// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: chat/v1/chat.proto

package chatv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/x0y14/jackal/gen/chat/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// ChatServiceName is the fully-qualified name of the ChatService service.
	ChatServiceName = "chat.v1.ChatService"
)

// ChatServiceClient is a client for the chat.v1.ChatService service.
type ChatServiceClient interface {
	CreateUser(context.Context, *connect_go.Request[v1.CreateUserRequest]) (*connect_go.Response[v1.CreateUserResponse], error)
	GetUser(context.Context, *connect_go.Request[v1.GetUserRequest]) (*connect_go.Response[v1.GetUserResponse], error)
	SendMessage(context.Context, *connect_go.Request[v1.SendMessageRequest]) (*connect_go.Response[v1.SendMessageResponse], error)
}

// NewChatServiceClient constructs a client for the chat.v1.ChatService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewChatServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) ChatServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &chatServiceClient{
		createUser: connect_go.NewClient[v1.CreateUserRequest, v1.CreateUserResponse](
			httpClient,
			baseURL+"/chat.v1.ChatService/CreateUser",
			opts...,
		),
		getUser: connect_go.NewClient[v1.GetUserRequest, v1.GetUserResponse](
			httpClient,
			baseURL+"/chat.v1.ChatService/GetUser",
			opts...,
		),
		sendMessage: connect_go.NewClient[v1.SendMessageRequest, v1.SendMessageResponse](
			httpClient,
			baseURL+"/chat.v1.ChatService/SendMessage",
			opts...,
		),
	}
}

// chatServiceClient implements ChatServiceClient.
type chatServiceClient struct {
	createUser  *connect_go.Client[v1.CreateUserRequest, v1.CreateUserResponse]
	getUser     *connect_go.Client[v1.GetUserRequest, v1.GetUserResponse]
	sendMessage *connect_go.Client[v1.SendMessageRequest, v1.SendMessageResponse]
}

// CreateUser calls chat.v1.ChatService.CreateUser.
func (c *chatServiceClient) CreateUser(ctx context.Context, req *connect_go.Request[v1.CreateUserRequest]) (*connect_go.Response[v1.CreateUserResponse], error) {
	return c.createUser.CallUnary(ctx, req)
}

// GetUser calls chat.v1.ChatService.GetUser.
func (c *chatServiceClient) GetUser(ctx context.Context, req *connect_go.Request[v1.GetUserRequest]) (*connect_go.Response[v1.GetUserResponse], error) {
	return c.getUser.CallUnary(ctx, req)
}

// SendMessage calls chat.v1.ChatService.SendMessage.
func (c *chatServiceClient) SendMessage(ctx context.Context, req *connect_go.Request[v1.SendMessageRequest]) (*connect_go.Response[v1.SendMessageResponse], error) {
	return c.sendMessage.CallUnary(ctx, req)
}

// ChatServiceHandler is an implementation of the chat.v1.ChatService service.
type ChatServiceHandler interface {
	CreateUser(context.Context, *connect_go.Request[v1.CreateUserRequest]) (*connect_go.Response[v1.CreateUserResponse], error)
	GetUser(context.Context, *connect_go.Request[v1.GetUserRequest]) (*connect_go.Response[v1.GetUserResponse], error)
	SendMessage(context.Context, *connect_go.Request[v1.SendMessageRequest]) (*connect_go.Response[v1.SendMessageResponse], error)
}

// NewChatServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewChatServiceHandler(svc ChatServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/chat.v1.ChatService/CreateUser", connect_go.NewUnaryHandler(
		"/chat.v1.ChatService/CreateUser",
		svc.CreateUser,
		opts...,
	))
	mux.Handle("/chat.v1.ChatService/GetUser", connect_go.NewUnaryHandler(
		"/chat.v1.ChatService/GetUser",
		svc.GetUser,
		opts...,
	))
	mux.Handle("/chat.v1.ChatService/SendMessage", connect_go.NewUnaryHandler(
		"/chat.v1.ChatService/SendMessage",
		svc.SendMessage,
		opts...,
	))
	return "/chat.v1.ChatService/", mux
}

// UnimplementedChatServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedChatServiceHandler struct{}

func (UnimplementedChatServiceHandler) CreateUser(context.Context, *connect_go.Request[v1.CreateUserRequest]) (*connect_go.Response[v1.CreateUserResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("chat.v1.ChatService.CreateUser is not implemented"))
}

func (UnimplementedChatServiceHandler) GetUser(context.Context, *connect_go.Request[v1.GetUserRequest]) (*connect_go.Response[v1.GetUserResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("chat.v1.ChatService.GetUser is not implemented"))
}

func (UnimplementedChatServiceHandler) SendMessage(context.Context, *connect_go.Request[v1.SendMessageRequest]) (*connect_go.Response[v1.SendMessageResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("chat.v1.ChatService.SendMessage is not implemented"))
}
