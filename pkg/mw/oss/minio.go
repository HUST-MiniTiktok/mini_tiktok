package oss

import (
	"context"

	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/conf"
	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/mw/redis"
	"github.com/cloudwego/kitex/pkg/klog"
	minio "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	OSSClient   *minio.Client
	RDClient    *redis.RDClient
	OSSEndpoint string
	err         error
)

func init() {
	ctx := context.Background()

	OSSEndpoint = conf.GetConf().GetString("oss.endpoint")
	OSSClient, err = minio.New(OSSEndpoint, &minio.Options{
		Creds: credentials.NewStaticV4(conf.GetConf().GetString("oss.accesskey"), conf.GetConf().GetString("oss.secretkey"), ""),
	})

	if err != nil {
		klog.Fatalf("init minio client failed: %v", err)
	}

	CreateBucket(ctx, conf.GetConf().GetString("oss.videobucket"))
	CreateBucket(ctx, conf.GetConf().GetString("oss.imagebucket"))
	LoadDefaultImageData(ctx, conf.GetConf().GetString("oss.imagebucket"))

	RDClient = redis.NewRDClient(conf.GetConf().GetInt("db.redis.id.oss"))
}
