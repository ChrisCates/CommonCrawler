# Common Crawler

## A simple way to extract data from Common Crawl

This repository has been revitalized with better logging and organization of respective components for extraction.

[![Go Report Card](https://goreportcard.com/badge/github.com/ChrisCates/CommonCrawler)](https://goreportcard.com/report/github.com/ChrisCates/CommonCrawler)
[![Build Status](https://travis-ci.org/ChrisCates/CommonCrawler.svg?branch=master)](https://travis-ci.org/ChrisCates/CommonCrawler)

## Dependencies

```bash
go get -u github.com/logrusorgru/aurora # for colors
```

## Running + Configuring

1. To run the application:

```bash
# Will run the application
go run src/*.go
```

2. Configure desired outputs and paths in `src/config.go`:

```golang
// Config is the preset variables for your extractor
type Config struct {
    baseURI     string
    wetPaths    string
    dataFolder  string
    matchFolder string
    start       int
    stop        int
}

//Defaults
Config{
    start:       0,
    stop:        5,
    baseURI:     "https://commoncrawl.s3.amazonaws.com/",
    wetPaths:    path.Join(cwd, "wet.paths"),
    dataFolder:  path.Join(cwd, "/output/crawl-data"),
    matchFolder: path.Join(cwd, "/output/match-data"),
}
```

3. Useful scripts:

```bash
# clean.sh cleans up default output folders
sh clean.sh

# extract.sh runs the extracting application
sh extract.sh
```

### Additional Notes

* MIT Licensed :heart:

* Need help with things? Email hello@chriscates.ca
