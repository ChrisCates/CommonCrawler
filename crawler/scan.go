package crawler

import (
	"bufio"
	"github.com/ChrisCates/CommonCrawler/types"
	"github.com/fatih/color"
	"log"
	"os"
	"path"
	"strconv"
)

func Scan(cfg *types.Config) error {
	paths, err := os.Open(cfg.WetPaths)

	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(paths)
	index := 0

	for scanner.Scan() {
		uri := cfg.BaseURI + scanner.Text()

		if index < cfg.Start {
			continue
		} else if index >= cfg.Stop {
			color.Green("\nFinished scanning, you can review results in the output folders...\n")
			break
		}

		index++

		filePath := path.Join(cfg.DataFolder, "wetfile_"+strconv.Itoa(index)+".wet.gz")

		log.Printf("\n  Download uri %s\n\t", uri)
		err := download(uri, filePath)
		if err != nil {
			color.Red("\n  error downloading file: %s\n\t", err)
			continue
		}

		color.Green("\n  Download was successful. \n extracting:\n\t" + uri)

		err = extract(filePath)
		if err != nil {
			color.Red("\n  error while extracting %s: %s\n\t", filePath, err)
			continue
		}

		color.Green("\n  Finished extracting:\n\t" + uri)

		extractedPath := path.Join(cfg.DataFolder, "wetfile_"+strconv.Itoa(index)+".wet")
		scanPath := path.Join(cfg.MatchFolder, "info."+strconv.Itoa(index)+".txt")

		err = analyzeFile(extractedPath, scanPath)

		if err != nil {
			color.Red("\n  There was a problem analyzing, make sure to look into this file:\n\t%s\n", extractedPath)
			color.Red("\t  The error is: %s", err)
			continue
		}

		color.Green("\n  Finished analyzing:\n\t" + extractedPath)
		color.Green("  Wrote results to" + scanPath)
	}

	return nil
}
