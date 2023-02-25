package main

import (
	"fmt"
	"strings"
)

func replaceMap(input string, mp map[string]string) string {
	output := input
	for k, v := range mp {
		output = strings.ReplaceAll(output, k, v)
	}
	return output
}

func formatInputLoad(input string) string {
	//in fmc files, use \(\) instead of () for json, normal for function, thanks future dave, please do it though you lazy bastard <3
	ret := fmt.Sprintf("//%s//", input)
	return elaborateJson(ret)
}
