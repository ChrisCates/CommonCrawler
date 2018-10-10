package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	aurora "github.com/logrusorgru/aurora"
)

func analyze(filePath string, path string) bool {
	file, err := os.Open(filePath)
	if err != nil {
		return false
	}

	scanner := bufio.NewScanner(file)

	data, err := os.Create(path)
	if err != nil {
		return false
	}

	matches := 0
	warcData := ""

	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, "ninja") {
			matches = matches + 1
		}
		if strings.Contains(text, "WARC") {
			warcData = warcData + text + "\n"
		}
		if strings.Contains(text, "WARC/1.0") && matches > 0 {
			warcData = warcData + "MATCHES: " + strconv.Itoa(matches)
			fmt.Println(aurora.Blue("\t  Found " + strconv.Itoa(matches) + " matches in this wet file..."))
			data.WriteString(warcData)
			matches = 0
		}
		if strings.Contains(text, "WARC/1.0") {
			warcData = ""
		}
	}

	file.Close()
	data.Close()

	return true
}
