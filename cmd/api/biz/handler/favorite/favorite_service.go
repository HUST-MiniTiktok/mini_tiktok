// Code generated by hertz generator.

package favorite

import (
	"context"

	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/api/biz/client"
	favorite "github.com/HUST-MiniTiktok/mini_tiktok/cmd/api/biz/model/favorite"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/errno"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/utils/conv"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// FavoriteAction .
// @router /douyin/favorite/action/ [POST]
func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req favorite.FavoriteActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, conv.ToHertzBaseResponse(errno.ParamErr))
		return
	}

	kitex_resp, err := client.FavoriteRPC.FavoriteAction(ctx, conv.ToKitexFavoriteActionRequest(&req))

	if err == nil {
		c.JSON(consts.StatusOK, conv.ToHertzFavoriteActionResponse(kitex_resp))
	} else {
		c.JSON(consts.StatusOK, conv.ToHertzBaseResponse(err))
	}
}

// FavoriteList .
// @router /douyin/favorite/list/ [GET]
func FavoriteList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req favorite.FavoriteListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, conv.ToHertzBaseResponse(errno.ParamErr))
		return
	}

	kitex_resp, err := client.FavoriteRPC.FavoriteList(ctx, conv.ToKitexFavoriteListRequest(&req))

	if err == nil {
		c.JSON(consts.StatusOK, conv.ToHertzFavoriteListResponse(kitex_resp))
	} else {
		c.JSON(consts.StatusOK, conv.ToHertzBaseResponse(err))
	}
}
