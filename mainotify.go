package mainotify

import "github.com/ryanchang1005/mainotify3/services/aws"

func NewAWSClient() *aws.AWSClient {
	return &aws.AWSClient{}
}
