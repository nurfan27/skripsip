package qasircore

import (
	"github.com/gin-gonic/gin"
)

type HttpResponse struct {
	c *gin.Context
}

func (hres *HttpResponse) Json(httpStatus int, data interface{}) {
	hres.c.JSON(httpStatus, data)
	return
}

func (hres *HttpResponse) Xml(httpStatus int, data map[string]interface{}) {
	res := gin.H{}

	for k, v := range data {
		res[k] = v
	}

	hres.c.XML(httpStatus, res)
	return
}
