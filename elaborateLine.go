package main

import (
	"fmt"
	"strconv"
	"strings"
)

func elaborateJson(json string) string {
	//( = ò
	//) = ç
	//, = £
	// //json//
	//extract out json
	jsonArray := strings.Split(json, "//")
	if len(jsonArray)%2 == 0 {
		panic("Json in argument formatting wrong, have you mismatched a //?")
	}
	//replace stuff in the json
	for i, doubleSlash := range jsonArray {
		if i%2 == 0 {
			continue
		}
		internalJson := replaceMap(doubleSlash, map[string]string{",": "£", " ": "", "\t": "", "(": "ò", ")": "ç", "//": "", "\r\n": "", "\n": ""})
		json = strings.ReplaceAll(json, doubleSlash, internalJson)
	}
	return json
}

// \w+(\(\(??[^(]*?\))
// test(rotation,3,/"ciao":{"cosa":["ksdjkfvn"\,"josdijf"]}/)
func elaborateCommand(command string, mp map[string]string) string {

	//for
	if command[0:strings.Index(command, "(")] == "for" {
		argumentGroup := command[strings.Index(command, "(")+1 : strings.LastIndex(command, ")")]
		argumentArray := strings.SplitN(argumentGroup, ",", 3)
		ret := ""
		forIndex := argumentArray[1]
		forChar := argumentArray[0]
		forBody := argumentArray[2]
		if forN, err := strconv.Atoi(forIndex); err == nil {
			for i := 0; i < forN; i++ {
				ret += strings.ReplaceAll(forBody, forChar, fmt.Sprintf("%d", i))
				if i < forN-1 {
					ret += ","
				}
			}
		} else {
			forArray := arrayStack[forIndex]
			for i, j := range forArray {
				ret += strings.ReplaceAll(forBody, forChar, fmt.Sprintf("%s", j))
				if i < len(forArray)-1 {
					ret += ","
				}
			}
		}

		return ret
	}
	if command[0:strings.Index(command, "(")] == "array" {
		argumentGroup := command[strings.Index(command, "(")+1 : strings.LastIndex(command, ")")]
		argumentArray := strings.Split(argumentGroup, ",")
		arrayName := argumentArray[0]
		wholeArray := argumentArray[1:]
		arrayStack[arrayName] = wholeArray
		return ""
	}
	//clean it up
	// //json//
	argumentGroup := command[strings.Index(command, "(")+1 : strings.LastIndex(command, ")")]
	//extract out json
	command = strings.ReplaceAll(command, argumentGroup, elaborateJson(argumentGroup))
	argumentGroup = command[strings.Index(command, "(")+1 : strings.LastIndex(command, ")")]
	argumentArray := strings.Split(argumentGroup, ",")
	commandName := command[0:strings.Index(command, "(")]
	//fmt.Sprintf("%s%d", commandName, strings.Count(commandReturn, "$"))
	ret := ""
	//handle case like version(), with no args
	if len(argumentArray[0]) > 0 {
		ret = mp[fmt.Sprintf("%s%d", commandName, len(argumentArray))]
	} else {
		ret = mp[fmt.Sprintf("%s%d", commandName, 0)]
		return ret
	}

	for i, argument := range argumentArray {
		ret = strings.ReplaceAll(ret, fmt.Sprintf("$%d", i+1), argument)
	}
	return ret
}
