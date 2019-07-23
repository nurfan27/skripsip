package qasircore

import "qasircore/interfaces"

type Controller struct {
	Response   ApiResponse
	MerchantID string
	Subdomain  string
	Data       map[string]interface{}
	Cache      interfaces.CacheInterface
}
