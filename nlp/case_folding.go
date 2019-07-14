package nlp

import (
	"log"
	"regexp"
)

// Proses penghilangan karakter selain huruf ‘a’ sampai ‘z’.
// Karakter seperti tanda ‘.’, ’,’ ,’:’

func CaseFolding(params string) string {
	// Make a Regex to say we only want letters and numbers
	reg, err := regexp.Compile("[^a-zA-Z0-9]*$")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(params, "")

	return processedString

}
