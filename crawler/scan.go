package crawler

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"os"
	"path"
	"strconv"
)

func (c *crawler) Scan() error {
	paths, err := os.Open(c.config.WetPaths)

	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(paths)
	index := 0

	for scanner.Scan() {
		uri := c.config.BaseURI + scanner.Text()

		if index < c.config.Start {
			continue
		} else if index >= c.config.Stop {
			color.Green("\nFinished scanning, you can review results in the output folders...\n")
			break
		}

		index++

		filePath := path.Join(c.config.DataFolder, "wetfile_"+strconv.Itoa(index)+".wet.gz")

		fmt.Printf("\n  Download uri %s\n\t", uri)
		err := c.Download(uri, filePath)
		if err != nil {
			color.Red("\n  Download was not successful: %s\n\t", err)
			continue
		}

		color.Green("\n  Download was successful extracting:\n\t" + uri)

		err = c.Extract(filePath)
		if err != nil {
			color.Red("\n  Exctraction %s err: %s\n\t", filePath, err)
			continue
		}

		color.Green("\n  Finished extracting:\n\t" + uri)

		extractedPath := path.Join(c.config.DataFolder, "wetfile_"+strconv.Itoa(index)+".wet")
		scanPath := path.Join(c.config.MatchFolder, "info."+strconv.Itoa(index)+".txt")

		err = c.AnalyzeFile(extractedPath, scanPath)

		if err != nil {
			color.Red("\n  There was a problem analyzing, make sure to look into this file:\n\t%s\n", extractedPath)
			color.Red("\t  The error is: %s", err)
			continue
		}

		color.Green("\n  Finished analyzing: %s \n\t", extractedPath)
		color.Green("\t Wrote results to" + scanPath)

	}

	return nil
}
