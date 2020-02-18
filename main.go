package main

import (
	"github.com/ChrisCates/CommonCrawler/config"
	"github.com/ChrisCates/CommonCrawler/crawler"
	"github.com/fatih/color"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		color.Red("error reading CommonCrawler configuration: %s", err)
		return
	}

	dirs := []string{cfg.DataFolder, cfg.MatchFolder}
	if err = config.DirectorySetup(dirs); err != nil {
		color.Red("error while creating directories: %s", err)
		return
	}

	color.Green("Starting scanning...")
	if err = crawler.Scan(cfg); err != nil {
		color.Red("error while scanning: %s", err)
		return
	}

	color.Green("scanning was successful")
}
