package qasircore

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

/**
 * @brief      Creates a file logging.
 * @return     os files
 */
func CreateFileLogging() *os.File {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	current_date := getCurrentDate()
	f, err := os.OpenFile(dir+"/logs/"+current_date+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("error opening file: %v", err)
	}

	return f
}

/**
 * @brief      this functioin for logging data error
 * @param      error_data  The error data
 * @return     object
 */
func LoggingError(error_data error) {
	f := CreateFileLogging()
	log.SetOutput(f)
	gin.DefaultWriter = io.MultiWriter(f)
	log.Println("Error Data => %s", error_data)
	defer f.Close()
}

/**
 * @brief      create logging for requests
 * @param      request_data  The request data
 */
func LoggingRequest(request_data *http.Request) {
	f := CreateFileLogging()

	log.SetOutput(f)
	gin.DefaultWriter = io.MultiWriter(f)
	log.Println("Request Data => %s", request_data)
	defer f.Close()
}
