package cmd

import (
	"os"
	"path"
)

// Config is the preset variables for your extractor
type Config struct {
	baseURI     string
	wetPaths    string
	dataFolder  string
	matchFolder string
	start       int
	stop        int
}

func getConfiguration() Config {
	cwd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	return Config{
		start:       0,
		stop:        1,
		baseURI:     "https://commoncrawl.s3.amazonaws.com/",
		wetPaths:    path.Join(cwd, "wet.paths"),
		dataFolder:  path.Join(cwd, "/output/crawl-data"),
		matchFolder: path.Join(cwd, "/output/match-data"),
	}
}
