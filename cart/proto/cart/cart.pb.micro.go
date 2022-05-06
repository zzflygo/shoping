// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: cart.proto

package cart

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Cart service

func NewCartEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Cart service

type CartService interface {
	AddCart(ctx context.Context, in *CartInfo, opts ...client.CallOption) (*Response, error)
	ClearCart(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*Response, error)
	Incr(ctx context.Context, in *ChangeNum, opts ...client.CallOption) (*Response, error)
	Decr(ctx context.Context, in *ChangeNum, opts ...client.CallOption) (*Response, error)
	GetAll(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*CartAllResponse, error)
	DeleteByID(ctx context.Context, in *DeleteByIdRequest, opts ...client.CallOption) (*Response, error)
}

type cartService struct {
	c    client.Client
	name string
}

func NewCartService(name string, c client.Client) CartService {
	return &cartService{
		c:    c,
		name: name,
	}
}

func (c *cartService) AddCart(ctx context.Context, in *CartInfo, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Cart.AddCart", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartService) ClearCart(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Cart.ClearCart", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartService) Incr(ctx context.Context, in *ChangeNum, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Cart.Incr", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartService) Decr(ctx context.Context, in *ChangeNum, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Cart.Decr", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartService) GetAll(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*CartAllResponse, error) {
	req := c.c.NewRequest(c.name, "Cart.GetAll", in)
	out := new(CartAllResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartService) DeleteByID(ctx context.Context, in *DeleteByIdRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Cart.DeleteByID", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Cart service

type CartHandler interface {
	AddCart(context.Context, *CartInfo, *Response) error
	ClearCart(context.Context, *UserRequest, *Response) error
	Incr(context.Context, *ChangeNum, *Response) error
	Decr(context.Context, *ChangeNum, *Response) error
	GetAll(context.Context, *UserRequest, *CartAllResponse) error
	DeleteByID(context.Context, *DeleteByIdRequest, *Response) error
}

func RegisterCartHandler(s server.Server, hdlr CartHandler, opts ...server.HandlerOption) error {
	type cart interface {
		AddCart(ctx context.Context, in *CartInfo, out *Response) error
		ClearCart(ctx context.Context, in *UserRequest, out *Response) error
		Incr(ctx context.Context, in *ChangeNum, out *Response) error
		Decr(ctx context.Context, in *ChangeNum, out *Response) error
		GetAll(ctx context.Context, in *UserRequest, out *CartAllResponse) error
		DeleteByID(ctx context.Context, in *DeleteByIdRequest, out *Response) error
	}
	type Cart struct {
		cart
	}
	h := &cartHandler{hdlr}
	return s.Handle(s.NewHandler(&Cart{h}, opts...))
}

type cartHandler struct {
	CartHandler
}

func (h *cartHandler) AddCart(ctx context.Context, in *CartInfo, out *Response) error {
	return h.CartHandler.AddCart(ctx, in, out)
}

func (h *cartHandler) ClearCart(ctx context.Context, in *UserRequest, out *Response) error {
	return h.CartHandler.ClearCart(ctx, in, out)
}

func (h *cartHandler) Incr(ctx context.Context, in *ChangeNum, out *Response) error {
	return h.CartHandler.Incr(ctx, in, out)
}

func (h *cartHandler) Decr(ctx context.Context, in *ChangeNum, out *Response) error {
	return h.CartHandler.Decr(ctx, in, out)
}

func (h *cartHandler) GetAll(ctx context.Context, in *UserRequest, out *CartAllResponse) error {
	return h.CartHandler.GetAll(ctx, in, out)
}

func (h *cartHandler) DeleteByID(ctx context.Context, in *DeleteByIdRequest, out *Response) error {
	return h.CartHandler.DeleteByID(ctx, in, out)
}