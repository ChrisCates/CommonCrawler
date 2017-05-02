package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	paths, _ := os.Open("wet.paths")
	pathScanner := bufio.NewScanner(paths)

	os.Mkdir("crawl-data", 0755)
	os.Mkdir("match-data", 0755)

	pathIndex := 0

	for pathScanner.Scan() {
		if pathIndex < 2 {
			pathIndex++
			continue
		}
		path := pathScanner.Text()
		newPath := "wetfile_" + strconv.Itoa(pathIndex) + ".wet.gz"
		fmt.Println("CURLing", newPath)
		_, err := exec.Command("curl", "-o", "crawl-data/"+newPath, "https://commoncrawl.s3.amazonaws.com/"+path).CombinedOutput()
		pathIndex++
		if err != nil {
			fmt.Println("Problem CURLing", newPath)
			continue
		}
		fmt.Println("Unzipping", newPath)
		_, err = exec.Command("gunzip", "crawl-data/"+newPath).CombinedOutput()
		if err != nil {
			fmt.Println("Problem Unzipping", newPath)
			continue
		}

		fileList, _ := ioutil.ReadDir("crawl-data")
		for _, file := range fileList {
			fmt.Println("Reading", file.Name())
			f, _ := os.Open("crawl-data/" + file.Name())
			scanner := bufio.NewScanner(f)

			data, _ := os.Create("match-data/" + file.Name())
			matches := 0
			warcData := ""

			for scanner.Scan() {
				text := scanner.Text()
				if strings.Contains(text, "chiropractor") {
					matches = matches + 1
				}
				if strings.Contains(text, "WARC") {
					warcData = warcData + text + "\n"
				}
				if strings.Contains(text, "WARC/1.0") && matches > 0 {
					warcData = warcData + "MATCHES: " + strconv.Itoa(matches)
					fmt.Println("Found", matches, "matches")
					data.WriteString(warcData)
					matches = 0
				}
				if strings.Contains(text, "WARC/1.0") {
					warcData = ""
				}
			}

			f.Close()
			data.Close()
			os.Remove("crawl-data/" + file.Name())
		}

		os.Remove("crawl-data/" + newPath)

	}
}
