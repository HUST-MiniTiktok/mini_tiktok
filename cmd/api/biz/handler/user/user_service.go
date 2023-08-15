// Code generated by hertz generator.

package user

import (
	"context"

	user "github.com/HUST-MiniTiktok/mini_tiktok/cmd/api/biz/model/user"
	rpc "github.com/HUST-MiniTiktok/mini_tiktok/cmd/api/biz/rpc"
	"github.com/HUST-MiniTiktok/mini_tiktok/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/codes"
)

// User .
// @router /douyin/user/ [GET]
func User(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.NewErrorResp(int32(codes.InvalidArgument), err.Error()))
		return
	}

	kitex_resp, err := rpc.User(ctx, utils.ToKitexUserRequest(&req))

	if err == nil {
		c.JSON(consts.StatusOK, utils.ToHertzUserResponse(kitex_resp))
	} else {
		c.JSON(consts.StatusOK, utils.NewErrorResp(int32(codes.Internal), err.Error()))
	}
}

// Register .
// @router /douyin/user/register/ [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserRegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.NewErrorResp(int32(codes.InvalidArgument), err.Error()))
		return
	}

	kitex_resp, err := rpc.Register(ctx, utils.ToKitexUserRegisterRequest(&req))

	if err == nil {
		c.JSON(consts.StatusOK, utils.ToHertzUserRegisterResponse(kitex_resp))
	} else {
		c.JSON(consts.StatusOK, utils.NewErrorResp(int32(codes.Internal), err.Error()))
	}
}

// Login .
// @router /douyin/user/login/ [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserLoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.NewErrorResp(int32(codes.InvalidArgument), err.Error()))
		return
	}

	kitex_resp, err := rpc.Login(ctx, utils.ToKitexUserLoginRequest(&req))

	if err == nil {
		c.JSON(consts.StatusOK, utils.ToHertzUserLoginResponse(kitex_resp))
	} else {
		c.JSON(consts.StatusOK, utils.NewErrorResp(int32(codes.Internal), err.Error()))
	}
}
