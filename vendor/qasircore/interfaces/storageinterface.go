package interfaces

type StorageInterface interface {
	Put(pathUpload string, filename string, bytes []byte) bool
}