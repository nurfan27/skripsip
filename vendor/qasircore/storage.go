package qasircore

import (
	"fmt"
	"qasircore/storages"
)

type Storage struct {
	driver           string
	config           map[string]interface{}
	prefixFolderPath string
	awsS3            storages.S3
}

func (this *Storage) SetDriver(driver string) {
	this.driver = driver
	if this.driver == "s3" {
		var s3 storages.S3
		s3.SettingAllConfiguration(this.config)
		this.awsS3 = s3
	}
}

func (this *Storage) SetConfig(configuration map[string]interface{}) {
	this.config = configuration
	this.prefixFolderPath = fmt.Sprint(configuration["prefixFolderPath"])
	this.SetDriver(fmt.Sprint(configuration["filesystemDriver"]))
}

func (this *Storage) Put(pathUpload string, filename string, bytes []byte) bool {
	if this.driver == "s3" {
		pathUpload = this.prefixFolderPath + "/" + pathUpload + "/" + filename
		this.awsS3.Upload(pathUpload, bytes)
	} else if this.driver == "local" {
		_ = storages.CreateImageLocal(bytes, filename, pathUpload)
	}
	return true
}

func (this *Storage) Get(path string) ([]byte, error) {
	return this.awsS3.Get(path)
}

func NewStorage(config map[string]interface{}) *Storage {
	var storages Storage
	storages.SetConfig(config)
	return &storages
}
