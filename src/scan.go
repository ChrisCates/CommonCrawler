package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"

	aurora "github.com/logrusorgru/aurora"
)

func scan(config Config) {
	paths, err := os.Open(config.wetPaths)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(paths)
	index := 0

	for scanner.Scan() {
		uri := config.baseURI + scanner.Text()

		if index < config.start {
			continue
		} else if index >= config.stop {
			fmt.Println(aurora.Green("\nFinished scanning, you can review results in the output folders...\n"))
			break
		}

		index++

		filePath := path.Join(config.dataFolder, "wetfile_"+strconv.Itoa(index)+".wet.gz")
		extracted := true

		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			extracted = extract(uri, filePath)
		} else {
			fmt.Println(aurora.Magenta("\n  " + uri + " has already been downloaded"))
		}

		if extracted == true {
			fmt.Println(aurora.Green("\n  Finished extracting:\n\t" + uri))
		} else {
			fmt.Println(aurora.Red("\n  There was a problem extracting: " + uri))
			fmt.Println(aurora.Red("  Make sure to look into this file: " + filePath))
		}

		extractedPath := path.Join(config.dataFolder, "wetfile_"+strconv.Itoa(index)+".wet")
		scanPath := path.Join(config.matchFolder, "info."+strconv.Itoa(index)+".txt")
		analyzed := analyze(extractedPath, scanPath)

		if analyzed == true {
			fmt.Println(aurora.Green("\n  Finished analyzing:\n\t" + extractedPath))
			fmt.Println(aurora.Green("  Wrote results to" + scanPath))
		} else {
			fmt.Println(aurora.Red("\n  There was a problem analyzing, make sure to look into this file:\n\t" + extractedPath))
		}
	}
}
