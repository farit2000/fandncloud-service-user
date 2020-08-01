// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/fandncloud-service-user/user.proto

package user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for UserService service

func NewUserServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for UserService service

type UserService interface {
	Create(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	Delete(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	Block(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	UnBlock(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	Update(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	GetById(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	GetUserGroups(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	GetUserPermissions(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Auth(ctx context.Context, in *User, opts ...client.CallOption) (*TokenResponse, error)
	Validate(ctx context.Context, in *AccessToken, opts ...client.CallOption) (*TokenResponse, error)
	Refresh(ctx context.Context, in *RefreshToken, opts ...client.CallOption) (*TokenResponse, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) Create(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UserService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Delete(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UserService.Delete", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Block(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UserService.Block", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UnBlock(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UserService.UnBlock", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Update(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UserService.Update", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetById(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UserService.GetById", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UserService.GetAll", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetUserGroups(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UserService.GetUserGroups", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetUserPermissions(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UserService.GetUserPermissions", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Auth(ctx context.Context, in *User, opts ...client.CallOption) (*TokenResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.Auth", in)
	out := new(TokenResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Validate(ctx context.Context, in *AccessToken, opts ...client.CallOption) (*TokenResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.Validate", in)
	out := new(TokenResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Refresh(ctx context.Context, in *RefreshToken, opts ...client.CallOption) (*TokenResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.Refresh", in)
	out := new(TokenResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	Create(context.Context, *User, *Response) error
	Delete(context.Context, *User, *Response) error
	Block(context.Context, *User, *Response) error
	UnBlock(context.Context, *User, *Response) error
	Update(context.Context, *User, *Response) error
	GetById(context.Context, *User, *Response) error
	GetAll(context.Context, *Request, *Response) error
	GetUserGroups(context.Context, *Request, *Response) error
	GetUserPermissions(context.Context, *Request, *Response) error
	Auth(context.Context, *User, *TokenResponse) error
	Validate(context.Context, *AccessToken, *TokenResponse) error
	Refresh(context.Context, *RefreshToken, *TokenResponse) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) error {
	type userService interface {
		Create(ctx context.Context, in *User, out *Response) error
		Delete(ctx context.Context, in *User, out *Response) error
		Block(ctx context.Context, in *User, out *Response) error
		UnBlock(ctx context.Context, in *User, out *Response) error
		Update(ctx context.Context, in *User, out *Response) error
		GetById(ctx context.Context, in *User, out *Response) error
		GetAll(ctx context.Context, in *Request, out *Response) error
		GetUserGroups(ctx context.Context, in *Request, out *Response) error
		GetUserPermissions(ctx context.Context, in *Request, out *Response) error
		Auth(ctx context.Context, in *User, out *TokenResponse) error
		Validate(ctx context.Context, in *AccessToken, out *TokenResponse) error
		Refresh(ctx context.Context, in *RefreshToken, out *TokenResponse) error
	}
	type UserService struct {
		userService
	}
	h := &userServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&UserService{h}, opts...))
}

type userServiceHandler struct {
	UserServiceHandler
}

func (h *userServiceHandler) Create(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.Create(ctx, in, out)
}

func (h *userServiceHandler) Delete(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.Delete(ctx, in, out)
}

func (h *userServiceHandler) Block(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.Block(ctx, in, out)
}

func (h *userServiceHandler) UnBlock(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.UnBlock(ctx, in, out)
}

func (h *userServiceHandler) Update(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.Update(ctx, in, out)
}

func (h *userServiceHandler) GetById(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.GetById(ctx, in, out)
}

func (h *userServiceHandler) GetAll(ctx context.Context, in *Request, out *Response) error {
	return h.UserServiceHandler.GetAll(ctx, in, out)
}

func (h *userServiceHandler) GetUserGroups(ctx context.Context, in *Request, out *Response) error {
	return h.UserServiceHandler.GetUserGroups(ctx, in, out)
}

func (h *userServiceHandler) GetUserPermissions(ctx context.Context, in *Request, out *Response) error {
	return h.UserServiceHandler.GetUserPermissions(ctx, in, out)
}

func (h *userServiceHandler) Auth(ctx context.Context, in *User, out *TokenResponse) error {
	return h.UserServiceHandler.Auth(ctx, in, out)
}

func (h *userServiceHandler) Validate(ctx context.Context, in *AccessToken, out *TokenResponse) error {
	return h.UserServiceHandler.Validate(ctx, in, out)
}

func (h *userServiceHandler) Refresh(ctx context.Context, in *RefreshToken, out *TokenResponse) error {
	return h.UserServiceHandler.Refresh(ctx, in, out)
}
