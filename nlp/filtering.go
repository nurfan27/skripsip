package nlp

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func WordList(params map[string]string) []string {
	log.Println("wordlist : ")

	var wordkey []string

	lines, err := readLines("wordlist.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	for index, _ := range params {
		for _, line := range lines {
			// fmt.Println(i, line)
			log.Println(index, line)
			if strings.TrimSpace(index) == strings.TrimSpace(line) {
				log.Println("sama bro ")
				wordkey = append(wordkey, index)
			}
		}

	}

	return wordkey
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
