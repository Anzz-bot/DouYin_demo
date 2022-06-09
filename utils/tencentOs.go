package utils

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	awscfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const (
	endpoint   = "https://cos.ap-nanjing.myqcloud.com"
	accesskey  = "AKIDQfeMnxqSoxDCuLVNyIL9gwvd4E6i37uR"
	secretkey  = "HH0SdW8TX4PWO1OVMxINCm8E4ez4ta1D"
	bucketName = "douyin-demo"
	region     = "ap-nanjing"
)

func NewS3Client() *s3.Client {
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL:           endpoint,
			SigningRegion: region,
		}, nil
	}) // 将endpoint设置为自己的endpoint
	cfg, err := awscfg.LoadDefaultConfig(context.TODO(), awscfg.WithEndpointResolverWithOptions(customResolver))
	if err != nil {
		panic(err)
	}
	cfg.Region = region
	cfg.Credentials = aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(accesskey, secretkey, ""))
	// 设定accesskey和secretkey
	client := s3.NewFromConfig(cfg)
	// 创建s3客户端
	return client
}
