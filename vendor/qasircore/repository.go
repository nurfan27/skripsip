package qasircore

import "github.com/gin-gonic/gin"

type Repository struct {
	C       *gin.Context
	Request HttpRequest
}

func (r *Repository) SetGinContext(c *gin.Context) {
	r.C = c
}

func (r *Repository) GetGinContext() *gin.Context {
	return r.C
}
