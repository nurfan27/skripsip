package interfaces

type ResponseApi interface {
	GetStatus() int
	GetHttpStatus() int
	GetData() map[string]interface{}
}
