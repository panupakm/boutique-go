// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.0
// - protoc             v4.23.3
// source: checkout/checkout.proto

package checkout

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationCheckoutPlaceOrder = "/boutiqueshop.Checkout/PlaceOrder"

type CheckoutHTTPServer interface {
	PlaceOrder(context.Context, *PlaceOrderRequest) (*PlaceOrderResponse, error)
}

func RegisterCheckoutHTTPServer(s *http.Server, srv CheckoutHTTPServer) {
	r := s.Route("/")
	r.POST("/checkout", _Checkout_PlaceOrder0_HTTP_Handler(srv))
}

func _Checkout_PlaceOrder0_HTTP_Handler(srv CheckoutHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PlaceOrderRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCheckoutPlaceOrder)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PlaceOrder(ctx, req.(*PlaceOrderRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PlaceOrderResponse)
		return ctx.Result(200, reply)
	}
}

type CheckoutHTTPClient interface {
	PlaceOrder(ctx context.Context, req *PlaceOrderRequest, opts ...http.CallOption) (rsp *PlaceOrderResponse, err error)
}

type CheckoutHTTPClientImpl struct {
	cc *http.Client
}

func NewCheckoutHTTPClient(client *http.Client) CheckoutHTTPClient {
	return &CheckoutHTTPClientImpl{client}
}

func (c *CheckoutHTTPClientImpl) PlaceOrder(ctx context.Context, in *PlaceOrderRequest, opts ...http.CallOption) (*PlaceOrderResponse, error) {
	var out PlaceOrderResponse
	pattern := "/checkout"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCheckoutPlaceOrder))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
