package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	aurora "github.com/logrusorgru/aurora"
)

func analyzeFile(filePath string, path string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := os.Create(path)
	if err != nil {
		return err
	}
	defer data.Close()

	matched := make(chan matchedWarc)

	go analyze("ninja", file, matched)

	for m := range matched {
		fmt.Println(aurora.Blue("\t  Found " + strconv.Itoa(m.matches) + " matches in this wet file..."))
		data.WriteString(fmt.Sprintf("%sMATCHES: %d", m.warcData, m.matches))
	}

	return nil
}

type matchedWarc struct {
	matches  int
	warcData string
}

func analyze(searchFor string, in io.Reader, matched chan matchedWarc) {
	defer close(matched)

	scanner := bufio.NewScanner(in)
	var warcDataBuilder strings.Builder
	matches := 0

	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, searchFor) {
			matches = matches + 1
		}

		if strings.Contains(text, "WARC") {
			warcDataBuilder.WriteString(text)
			warcDataBuilder.WriteRune('\n')
		}

		if strings.Contains(text, "WARC/1.0") {
			if matches > 0 {
				matched <- matchedWarc{matches, warcDataBuilder.String()}
			}
			warcDataBuilder.Reset()
			matches = 0
		}
	}
}
