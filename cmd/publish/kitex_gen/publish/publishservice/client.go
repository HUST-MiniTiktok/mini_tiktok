// Code generated by Kitex v0.7.0. DO NOT EDIT.

package publishservice

import (
	"context"
	publish "github.com/HUST-MiniTiktok/mini_tiktok/cmd/publish/kitex_gen/publish"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	PublishAction(ctx context.Context, request *publish.PublishActionRequest, callOptions ...callopt.Option) (r *publish.PublishActionResponse, err error)
	PublishList(ctx context.Context, request *publish.PublishListRequest, callOptions ...callopt.Option) (r *publish.PublishListResponse, err error)
	GetVideoById(ctx context.Context, request *publish.GetVideoByIdRequest, callOptions ...callopt.Option) (r *publish.GetVideoByIdResponse, err error)
	GetVideoByIdList(ctx context.Context, request *publish.GetVideoByIdListRequest, callOptions ...callopt.Option) (r *publish.GetVideoByIdListResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kPublishServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kPublishServiceClient struct {
	*kClient
}

func (p *kPublishServiceClient) PublishAction(ctx context.Context, request *publish.PublishActionRequest, callOptions ...callopt.Option) (r *publish.PublishActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PublishAction(ctx, request)
}

func (p *kPublishServiceClient) PublishList(ctx context.Context, request *publish.PublishListRequest, callOptions ...callopt.Option) (r *publish.PublishListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PublishList(ctx, request)
}

func (p *kPublishServiceClient) GetVideoById(ctx context.Context, request *publish.GetVideoByIdRequest, callOptions ...callopt.Option) (r *publish.GetVideoByIdResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetVideoById(ctx, request)
}

func (p *kPublishServiceClient) GetVideoByIdList(ctx context.Context, request *publish.GetVideoByIdListRequest, callOptions ...callopt.Option) (r *publish.GetVideoByIdListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetVideoByIdList(ctx, request)
}
