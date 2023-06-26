package aliyunoss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type AliBucket struct {
	B *oss.Bucket
}

func LoadConfig(c Config) (AliBucket, error) {
	ab := AliBucket{}
	client, err := oss.New(c.Endpoint, c.AccessKeyId, c.AccessKeySecret)
	if err != nil {
		return ab, err
	}

	b, err := client.Bucket(c.BucketName)
	ab.B = b

	return ab, err
}
