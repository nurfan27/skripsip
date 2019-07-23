package qasircore

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Config(filepath string) map[string]interface{} {
	configs := make(map[string]interface{})
	pathWithFilename := filepath + ".env"
	b, err := ioutil.ReadFile(pathWithFilename)
	if err != nil {
		fmt.Print(err)
	}

	str := string(b)

	ArrayLines := strings.Split(str, "\n")
	for _, val := range ArrayLines {
		val = strings.TrimSpace(val)
		if val != "" && !strings.Contains(val, "#") {
			explode := strings.Split(val, "=")
			configs[explode[0]] = explode[1]
			os.Setenv(explode[0], fmt.Sprint(explode[1]))
		}
	}

	return configs
}
