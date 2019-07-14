package app

import (
	"fmt"
	"os"
)

func GenerateWordList() {
	f, err := os.Create("./nlp/wordlist.txt")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	var d []string

	temp := fmt.Println()

	d = append(d, temp...)

	for _, v := range d {
		fmt.Fprintln(f, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file written successfully")
}
