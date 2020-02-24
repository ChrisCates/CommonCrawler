package main

import (
	"github.com/ChrisCates/CommonCrawler/crawler"
	"github.com/fatih/color"
	"os"
)

func main() {
	color.Green("Getting configurations for Common Crawl Extractor...")
	cfg, err := crawler.ReadConfig()
	if err != nil {
		color.Red("error reading CommonCrawler config: %s", err)
		os.Exit(1)
	}

	if err = crawler.DirectorySetup(cfg.DataFolder, cfg.MatchFolder); err != nil {
		color.Red("error setting directories: %s", err)
		os.Exit(2)
	}

	cc := crawler.NewCrawler(cfg)
	color.Green("Starting scanning...")
	if err = cc.Scan(); err != nil {
		color.Red("error occurred while scanning: %s", err)
		os.Exit(3)
	}

	color.Green("finished scanning...")
}
