// Code generated by hertz generator.

package feed

import (
	"context"

	feed "github.com/HUST-MiniTiktok/mini_tiktok/biz/model/feed"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// GetFeed .
// @router /douyin/feed/ [GET]
func GetFeed(ctx context.Context, c *app.RequestContext) {
	var err error
	var req feed.FeedRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(feed.FeedResponse)

	c.JSON(consts.StatusOK, resp)
}
