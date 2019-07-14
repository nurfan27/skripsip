package nlp

import (
	"log"
)

func WordList(params map[string]string) string {
	log.Println("wordlist : ")
	for index, _ := range params {
		log.Println(index)
	}

	return ""
}
