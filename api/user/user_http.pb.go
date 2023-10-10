// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.0
// - protoc             v4.23.3
// source: user/user.proto

package user

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

const OperationUserCreateUser = "/boutiqueshop.User/CreateUser"
const OperationUserGetUser = "/boutiqueshop.User/GetUser"
const OperationUserGetUserByUsername = "/boutiqueshop.User/GetUserByUsername"
const OperationUserSave = "/boutiqueshop.User/Save"
const OperationUserVerifyPassword = "/boutiqueshop.User/VerifyPassword"

type UserHTTPServer interface {
	CreateUser(context.Context, *CreateUserReq) (*CreateUserReply, error)
	GetUser(context.Context, *GetUserReq) (*GetUserReply, error)
	GetUserByUsername(context.Context, *GetUserByUsernameReq) (*GetUserByUsernameReply, error)
	Save(context.Context, *SaveUserReq) (*SaveUserReply, error)
	VerifyPassword(context.Context, *VerifyPasswordReq) (*VerifyPasswordReply, error)
}

func RegisterUserHTTPServer(s *http.Server, srv UserHTTPServer) {
	r := s.Route("/")
	r.GET("/user/{id}", _User_GetUser0_HTTP_Handler(srv))
	r.GET("/user/{username}", _User_GetUserByUsername0_HTTP_Handler(srv))
	r.PUT("/user", _User_Save0_HTTP_Handler(srv))
	r.POST("/user", _User_CreateUser0_HTTP_Handler(srv))
	r.POST("/user/verify", _User_VerifyPassword0_HTTP_Handler(srv))
}

func _User_GetUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserGetUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUser(ctx, req.(*GetUserReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserReply)
		return ctx.Result(200, reply)
	}
}

func _User_GetUserByUsername0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserByUsernameReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserGetUserByUsername)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserByUsername(ctx, req.(*GetUserByUsernameReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserByUsernameReply)
		return ctx.Result(200, reply)
	}
}

func _User_Save0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SaveUserReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserSave)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Save(ctx, req.(*SaveUserReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SaveUserReply)
		return ctx.Result(200, reply)
	}
}

func _User_CreateUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateUserReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserCreateUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateUser(ctx, req.(*CreateUserReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateUserReply)
		return ctx.Result(200, reply)
	}
}

func _User_VerifyPassword0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in VerifyPasswordReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserVerifyPassword)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.VerifyPassword(ctx, req.(*VerifyPasswordReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*VerifyPasswordReply)
		return ctx.Result(200, reply)
	}
}

type UserHTTPClient interface {
	CreateUser(ctx context.Context, req *CreateUserReq, opts ...http.CallOption) (rsp *CreateUserReply, err error)
	GetUser(ctx context.Context, req *GetUserReq, opts ...http.CallOption) (rsp *GetUserReply, err error)
	GetUserByUsername(ctx context.Context, req *GetUserByUsernameReq, opts ...http.CallOption) (rsp *GetUserByUsernameReply, err error)
	Save(ctx context.Context, req *SaveUserReq, opts ...http.CallOption) (rsp *SaveUserReply, err error)
	VerifyPassword(ctx context.Context, req *VerifyPasswordReq, opts ...http.CallOption) (rsp *VerifyPasswordReply, err error)
}

type UserHTTPClientImpl struct {
	cc *http.Client
}

func NewUserHTTPClient(client *http.Client) UserHTTPClient {
	return &UserHTTPClientImpl{client}
}

func (c *UserHTTPClientImpl) CreateUser(ctx context.Context, in *CreateUserReq, opts ...http.CallOption) (*CreateUserReply, error) {
	var out CreateUserReply
	pattern := "/user"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserCreateUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) GetUser(ctx context.Context, in *GetUserReq, opts ...http.CallOption) (*GetUserReply, error) {
	var out GetUserReply
	pattern := "/user/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserGetUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) GetUserByUsername(ctx context.Context, in *GetUserByUsernameReq, opts ...http.CallOption) (*GetUserByUsernameReply, error) {
	var out GetUserByUsernameReply
	pattern := "/user/{username}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserGetUserByUsername))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) Save(ctx context.Context, in *SaveUserReq, opts ...http.CallOption) (*SaveUserReply, error) {
	var out SaveUserReply
	pattern := "/user"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserSave))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) VerifyPassword(ctx context.Context, in *VerifyPasswordReq, opts ...http.CallOption) (*VerifyPasswordReply, error) {
	var out VerifyPasswordReply
	pattern := "/user/verify"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserVerifyPassword))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
