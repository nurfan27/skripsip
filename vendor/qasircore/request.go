package qasircore

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type HttpRequest interface {
	GetBody() map[string]interface{}
	GetBindingJSON() interface{}
	GetContextGin() *gin.Context
	Input(key string) string
	Header(key string) string
	getResponse() HttpResponse
}

type Request struct {
	Res         HttpResponse
	c           *gin.Context
	BindingJSON interface{}
	Body        map[string]interface{}
}

func (ccr *Request) Header(key string) string {
	return ccr.c.Request.Header.Get(key)
}

func (ccr *Request) getResponse() HttpResponse {
	return ccr.Res
}

func (ccr *Request) GetContextGin() *gin.Context {
	return ccr.c
}

func (ccr *Request) SetBody(body map[string]interface{}) {
	ccr.Body = body
}

func (ccr *Request) GetBody() map[string]interface{} {
	return ccr.Body
}

func (ccr *Request) GetBindingJSON() interface{} {
	return ccr.BindingJSON
}

func (ccr *Request) Rules() map[string][]string {
	return nil
}

func (ccr *Request) Message() map[string][]string {
	return nil
}

func (ccr *Request) Authorize() bool {
	return true
}

func (ccr *Request) SetContextGin(c *gin.Context) {
	ccr.Res.c = c
	ccr.c = c
}

func (ccr *Request) Input(key string) string {
	return ccr.c.DefaultPostForm(key, ccr.c.DefaultQuery(key, ccr.c.Param(key)))
}

func (ccr *Request) Validation(rules govalidator.MapData, messages govalidator.MapData) []string {
	var errors []string

	opts := govalidator.Options{
		Request:         ccr.c.Request, // request object
		Rules:           rules,         // rules map
		Messages:        messages,      // custom message map (Optional)
		RequiredDefault: true,          // all the field to be pass the rules
	}
	v := govalidator.New(opts)
	validations := v.Validate()

	if ccr.BindingJSON != nil {
		validation := govalidator.Options{
			Request: ccr.c.Request,
			Rules:   rules,
			Data:    ccr.BindingJSON,
		}

		vd := govalidator.New(validation)
		e := vd.ValidateJSON()
		if len(e) != 0 {
			for _, data := range e {
				errors = append(errors, data...)
			}
			return errors
		} else {
			return nil
		}
	}

	if len(validations) != 0 {
		for _, data := range validations {
			errors = append(errors, data...)
		}
		return errors
	}

	return errors
}
