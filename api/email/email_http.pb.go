// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.0
// - protoc             v4.23.3
// source: email/email.proto

package email

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

const OperationEmailSendOrderConfirmation = "/api.email.Email/SendOrderConfirmation"

type EmailHTTPServer interface {
	SendOrderConfirmation(context.Context, *SendOrderConfirmationRequest) (*Empty, error)
}

func RegisterEmailHTTPServer(s *http.Server, srv EmailHTTPServer) {
	r := s.Route("/")
	r.POST("/email/order-confirmation", _Email_SendOrderConfirmation0_HTTP_Handler(srv))
}

func _Email_SendOrderConfirmation0_HTTP_Handler(srv EmailHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SendOrderConfirmationRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationEmailSendOrderConfirmation)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SendOrderConfirmation(ctx, req.(*SendOrderConfirmationRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Empty)
		return ctx.Result(200, reply)
	}
}

type EmailHTTPClient interface {
	SendOrderConfirmation(ctx context.Context, req *SendOrderConfirmationRequest, opts ...http.CallOption) (rsp *Empty, err error)
}

type EmailHTTPClientImpl struct {
	cc *http.Client
}

func NewEmailHTTPClient(client *http.Client) EmailHTTPClient {
	return &EmailHTTPClientImpl{client}
}

func (c *EmailHTTPClientImpl) SendOrderConfirmation(ctx context.Context, in *SendOrderConfirmationRequest, opts ...http.CallOption) (*Empty, error) {
	var out Empty
	pattern := "/email/order-confirmation"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationEmailSendOrderConfirmation))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
