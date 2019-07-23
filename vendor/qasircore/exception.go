package qasircore

import "reflect"

type Exception struct {
	Response                 Response
	MappingExceptionFunction map[string]interface{}
}

/**
 * @brief      this function for set error exception
 * @param      exceptionname  The exceptionname
 * @return     object
 */
func (this *Exception) ErrorException(exceptionname string) Response {
	reflectFunction := reflect.ValueOf(this.MappingExceptionFunction[exceptionname])
	params := []reflect.Value{}
	t := reflectFunction.Call(params)
	resData := t[0].Interface().(Response)

	return resData
}
