package qasircore

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

/**
 * @brief      this function for handling validations
 * @param      c         gin.Context
 * @param      rules     The rules
 * @param      messages  The messages
 * @return     message errors
 */
func Validations(c *gin.Context, rules govalidator.MapData, messages govalidator.MapData) map[string]interface{} {
	var errors []string
	data := make(map[string]interface{}, 0)

	LoggingRequest(c.Request)
	opts := govalidator.Options{
		Request:         c.Request, // request object
		Rules:           rules,     // rules map
		Messages:        messages,  // custom message map (Optional)
		RequiredDefault: true,      // all the field to be pass the rules
	}
	v := govalidator.New(opts)
	validations := v.Validate()

	// check validations is nil
	if len(validations) != 0 {

		validation := govalidator.Options{
			Request: c.Request,
			Rules:   rules,
			Data:    &data,
		}

		vd := govalidator.New(validation)
		e := vd.ValidateJSON()
		if len(e) != 0 {
			// mapping if any error
			for _, data := range validations {
				errors = append(errors, data...)
			}

			// return data
			return map[string]interface{}{"errors": errors}
		}

	}

	return nil
}
