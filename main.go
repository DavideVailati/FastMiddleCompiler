package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var arrayStack map[string][]string

func submitLine(line string, cmdmap map[string]string) string {
	line = strings.ReplaceAll(line, ",_", "//,//")
	for strings.Contains(line, "(") {
		regex, _ := regexp.Compile("\\w+(\\(\\(??[^(]*?\\))")
		matches := regex.FindAllString(line, -1)
		for _, k := range matches {
			line = strings.ReplaceAll(line, k, elaborateCommand(k, cmdmap))
		}
	}
	final := replaceMap(line, map[string]string{"£": ",", "ò": "(", "ç": ")", "//": ""})
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, []byte(final), "", "  ")
	if err != nil {
		fmt.Printf("\n" + final)
		panic(err)
	}

	return prettyJSON.String()
}

func main() {
	arrayStack = make(map[string][]string)
	input := fileToStr("input.fmc")
	fb := bufio.NewScanner(strings.NewReader(input))
	fb.Scan()
	filePath := replaceMap(fb.Text()[1:], map[string]string{
		"bp/": fileToStr("rootPath.txt") + "0. Behaviour Pack-BP/",
		"rp/": fileToStr("rootPath.txt") + "1. Resource Pack-RP/",
	})
	f, _ := os.Create(filePath)
	toWrite := ""
	cmdmap := loadMap("./")

	for fb.Scan() {
		if strings.Contains(fb.Text(), "$") {

			_, err := f.WriteString(submitLine(toWrite, cmdmap))
			if err != nil {
				panic(err)
			}
			err = f.Close()
			if err != nil {
				panic(err)
			}
			toWrite = ""
			filePath = replaceMap(fb.Text()[1:], map[string]string{
				"bp/": fileToStr("rootPath.txt") + "0. Behaviour Pack-BP/",
				"rp/": fileToStr("rootPath.txt") + "1. Resource Pack-RP/",
			})
			f, _ = os.Create(filePath)
		} else {
			toWrite += replaceMap(fb.Text(), map[string]string{" ": "", "\t": ""})
		}
	}

	_, err := f.WriteString(submitLine(toWrite, cmdmap))
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}
}
