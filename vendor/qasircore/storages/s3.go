package storages

import (
	"fmt"
	"net/http"

	"gopkg.in/amz.v1/aws"
	"gopkg.in/amz.v1/s3"
)

type S3 struct {
	accessKey  string
	secretKey  string
	bucketName string
	awsregion  string
}

func (this *S3) SettingAllConfiguration(config map[string]interface{}) {
	this.SetAccessKey(fmt.Sprint(config["awsAccessKey"]))
	this.SetSecretKey(fmt.Sprint(config["awsSecretKey"]))
	this.SetAwsRegion(fmt.Sprint(config["awsRegion"]))
	this.SetBucketName(fmt.Sprint(config["s3Bucket"]))
}

func (this *S3) SetAccessKey(accessKey string) {
	this.accessKey = accessKey
}

func (this *S3) SetAwsRegion(awsRegion string) {
	this.awsregion = awsRegion
}

func (this *S3) SetSecretKey(secretkey string) {
	this.secretKey = secretkey
}

func (this *S3) SetBucketName(bucketName string) {
	this.bucketName = bucketName
}

func (this *S3) Get(path string) ([]byte, error) {
	AWSAuth := aws.Auth{
		AccessKey: this.accessKey, // change this to yours
		SecretKey: this.secretKey,
	}

	AWSRegion := aws.Regions[this.awsregion]

	connection := s3.New(AWSAuth, AWSRegion)

	bucket := connection.Bucket(this.bucketName)

	return bucket.Get(path)
}

func (this *S3) Upload(pathUpload string, bytes []byte) {
	AWSAuth := aws.Auth{
		AccessKey: this.accessKey, // change this to yours
		SecretKey: this.secretKey,
	}

	AWSRegion := aws.Regions[this.awsregion]

	connection := s3.New(AWSAuth, AWSRegion)

	bucket := connection.Bucket(this.bucketName)
	fileType := http.DetectContentType(bytes)
	err := bucket.Put(pathUpload, bytes, fileType, s3.ACL("public-read"))

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
