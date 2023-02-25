package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func fileToStr(path string) string {
	f, _ := os.Open(path)
	fb, _ := io.ReadAll(f)
	input := string(fb)
	return input
}

func loadMap(path string) map[string]string {

	ret := make(map[string]string)
	files, err := os.ReadDir(path + "load/")
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if file.IsDir() {
			fmt.Printf("Error! %s is a directory...\n", file.Name())
			return nil
		}
		wholeFile := fileToStr(path + "load/" + file.Name())
		wholeFile = replaceMap(wholeFile, map[string]string{" ": "", "\r": "", "\n": "", "\t": ""})

		sections := strings.Split(wholeFile[2:], "!!")

		for _, section := range sections {
			commandWhole := strings.SplitN(section, ":", 2)
			commandName := commandWhole[0]
			commandReturn := commandWhole[1]
			x := strings.Count(commandReturn, "$")

			nArgs := 0
			for i := 0; i < x; i++ {
				if strings.Count(commandReturn, fmt.Sprintf("$%d", i+1)) > 0 {
					nArgs = i + 1

				}
			}

			if nArgs > 0 {
				ret[fmt.Sprintf("%s%d", commandName, nArgs)] = formatInputLoad(commandReturn)
			} else {
				ret[fmt.Sprintf("%s%d", commandName, 0)] = formatInputLoad(commandReturn)
			}

		}
	}

	return ret
}
