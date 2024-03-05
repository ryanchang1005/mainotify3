package mainotify

import "github.com/maideax/mainotify/services/aws"

func NewAWSClient() *aws.AWSClient {
	return &aws.AWSClient{}
}
