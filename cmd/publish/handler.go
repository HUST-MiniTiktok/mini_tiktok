package main

import (
	"context"

	service "github.com/HUST-MiniTiktok/mini_tiktok/cmd/publish/service"
	publish "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/publish"
)

// PublishServiceImpl implements the last service interface defined in the IDL.
type PublishServiceImpl struct{}

// PublishAction implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishAction(ctx context.Context, request *publish.PublishActionRequest) (resp *publish.PublishActionResponse, err error) {
	publish_service := service.NewPublishService(ctx)
	resp, err = publish_service.PublishAction(request)
	return
}

// PublishList implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishList(ctx context.Context, request *publish.PublishListRequest) (resp *publish.PublishListResponse, err error) {
	publish_service := service.NewPublishService(ctx)
	resp, err = publish_service.PublishList(request)
	return
}

// GetVideoById implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) GetVideoById(ctx context.Context, request *publish.GetVideoByIdRequest) (resp *publish.GetVideoByIdResponse, err error) {
	publish_service := service.NewPublishService(ctx)
	resp, err = publish_service.GetVideoById(request)
	return
}

// GetVideoByIdList implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) GetVideoByIdList(ctx context.Context, request *publish.GetVideoByIdListRequest) (resp *publish.GetVideoByIdListResponse, err error) {
	publish_service := service.NewPublishService(ctx)
	resp, err = publish_service.GetVideoByIdList(request)
	return
}

// GetPublishInfoByUserId implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) GetPublishInfoByUserId(ctx context.Context, request *publish.GetPublishInfoByUserIdRequest) (resp *publish.GetPublishInfoByUserIdResponse, err error) {
	publish_service := service.NewPublishService(ctx)
	resp, err = publish_service.GetPublishInfoByUserId(request)
	return
}
