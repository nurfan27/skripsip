package nlp

import (
	"fmt"
	"log"

	"gopkg.in/jdkato/prose.v2"
)

var (
	token = make(map[string]string)
)

// Proses mengurai kalimat menjadi kata-kata yang menyusunnya.
func Tokenizing(params string) map[string]string {
	// load default configuration
	doc, err := prose.NewDocument(params)
	if err != nil {
		log.Fatal(err)
	}

	// Iterasi doc's:
	for _, val := range doc.Tokens() {
		// bisa ambil tag dan label
		fmt.Println(val.Text, val.Tag)
		token[val.Text] = val.Tag
	}

	return token
}
