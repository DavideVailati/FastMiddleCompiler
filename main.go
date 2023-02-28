package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
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

	return final
}
func getPath(in string) string {
	filePath := replaceMap(in, map[string]string{
		"bp/": settings["root"] + settings["bp"],
		"rp/": settings["root"] + settings["rp"],
	})
	return filePath
}

var settings map[string]string

func PrettyString(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "  "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}

func main() {
	arrayStack = make(map[string][]string)
	input := "start()\n" + fileToStr("input.fmc")
	fb := bufio.NewScanner(strings.NewReader(input))
	fb.Scan()
	FSettings, _ := os.Open("settings.json")
	FBSettings, _ := io.ReadAll(FSettings)
	settings = make(map[string]string)
	err := json.Unmarshal(FBSettings, &settings)
	if err != nil {
		panic(err)
	}
	cmdmap := loadMap("./")
	finalTotal := ""

	for fb.Scan() {
		finalTotal += replaceMap(fb.Text(), map[string]string{
			" ":  "",
			"\t": "",
		})
	}
	elaboratedTotal := submitLine(finalTotal, cmdmap)
	splitTotal := replaceMap(elaboratedTotal, map[string]string{
		"$":      "\n$",
		",$":     "\n$",
		".json,": ".json",
		".json":  ".json\n",
		"[!]":    " ",
	})
	finalSplit := strings.Split(splitTotal, "$")[1:]
	for _, k := range finalSplit {
		k = replaceMap(k, map[string]string{
			"\n": "",
			"\r": "",
		})
		midSplit := strings.Split(k, ".json")
		filePath := midSplit[0] + ".json"
		filePath = getPath(filePath)
		finalFile, err := os.Create(filePath)
		if err != nil {
			panic(err)
		}
		toWrite, err := PrettyString(midSplit[1])
		if err != nil {
			finalFile.WriteString(midSplit[1])
			panic(err)
		}
		finalFile.WriteString(toWrite)
		err = finalFile.Close()
		if err != nil {
			panic(err)
		}
	}
}
