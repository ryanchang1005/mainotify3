package mainotify

import "github.com/ryanchang1005/mainotify2/services/aws"

func NewAWSClient() *aws.AWSClient {
	return &aws.AWSClient{}
}
