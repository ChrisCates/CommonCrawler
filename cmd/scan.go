package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"

	aurora "github.com/logrusorgru/aurora"
)

func scan(config Config) {
	fmt.Println(aurora.Green("Starting scanning..."))
	paths, err := os.Open(config.wetPaths)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(paths)
	index := 0

	for scanner.Scan() {
		uri := config.baseURI + scanner.Text()
		if index < config.start {
			index++
			continue
		} else if index >= config.stop {
			fmt.Println(aurora.Green("\nFinished scanning, you can review results in the output folders...\n"))
			break
		}

		filePath := path.Join(config.dataFolder, "wetfile_"+strconv.Itoa(index)+".wet.gz")

		fmt.Printf("\n  Download uri %s\n\t", uri)
		err := download(uri, filePath)
		if err != nil {
			fmt.Println(aurora.Red(fmt.Sprintf("\n  Download was not successfull: %s\n\t", err)))
			index++
			continue
		}

		fmt.Println(aurora.Green("\n  Download was successfull extracting:\n\t" + uri))

		err = extract(filePath)
		if err != nil {
			fmt.Println(aurora.Red(fmt.Sprintf("\n  Exctraction %s err: %s\n\t", filePath, err)))
			index++
			continue
		}

		fmt.Println(aurora.Green("\n  Finished extracting:\n\t" + uri))

		extractedPath := path.Join(config.dataFolder, "wetfile_"+strconv.Itoa(index)+".wet")
		scanPath := path.Join(config.matchFolder, "info."+strconv.Itoa(index)+".txt")

		err = analyzeFile(extractedPath, scanPath)

		if err != nil {
			fmt.Println(aurora.Red(fmt.Sprintf("\n  There was a problem analyzing, make sure to look into this file:\n\t%s\n", extractedPath)))
			fmt.Println(aurora.Red(fmt.Sprintf("\t  The error is: %s", err)))
			index++
			continue
		}

		fmt.Println(aurora.Green("\n  Finished analyzing:\n\t" + extractedPath))
		fmt.Println(aurora.Green("  Wrote results to" + scanPath))
		index++
	}
}