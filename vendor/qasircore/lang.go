package qasircore

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Lang struct {
	languange map[string]string
}

func (l *Lang) Get(key string) string {
	_, exists := l.languange[key]
	if !exists {
		return ""
	}
	return l.languange[key]
}

func (l *Lang) GenerateDataMappingLanguange(language string) {
	dir, _ := os.Getwd()

	pathDir := dir + "/lang/" + language

	files, err := ioutil.ReadDir(pathDir)
	if err != nil {
		log.Println(err)
	}

	dataLanguange := map[string]string{}

	for _, f := range files {
		filename := f.Name()
		raw, _ := ioutil.ReadFile(pathDir + "/" + filename)

		var content map[string]string

		json.Unmarshal(raw, &content)

		for k, v := range content {
			dataLanguange[k] = v
		}
	}

	l.languange = dataLanguange
}

func NewLang(languange string) *Lang {
	var lang Lang

	lang.GenerateDataMappingLanguange(languange)

	return &lang
}
