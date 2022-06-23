/*
 * @Author: alexander.huang
 * @Date:   2022-06-23 20:14:50
 * @Last Modified by: alexander.huang
 * @Last Modified time: 2022-06-23 20:14:50
 */
package utils

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	awscfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const (
	endpoint   = ""
	accesskey  = ""
	secretkey  = ""
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
