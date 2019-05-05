# Common Crawler

## 🕸 A simple and easy way to extract data from Common Crawl with little or no hassle.

![Go Version](https://img.shields.io/badge/Go-v1.1.0-blue.svg)
![License](https://img.shields.io/badge/license-MIT-blue.svg)
[![Build Status](https://travis-ci.org/ChrisCates/CommonCrawler.svg?branch=master)](https://travis-ci.org/ChrisCates/CommonCrawler)
[![Go Report Card](https://goreportcard.com/badge/github.com/ChrisCates/CommonCrawler)](https://goreportcard.com/report/github.com/ChrisCates/CommonCrawler)


## As a command line tool

***This will be implemented soon, please review issues for Gitcoin bounties***

```bash
# Output help
commoncrawler --help

# Specify configuration
commoncrawler --base-uri https://commoncrawl.s3.amazonaws.com/
commoncrawler --wet-paths wet.paths
commoncrawler --data-folder output/crawl-data
commoncrawler --start 0
commoncrawler --stop 5 # -1 will loop through all wet files from wet.paths

# Start crawling the web
commoncrawler start --stop -1
```

## Compilation and Configuration

### Installing dependencies

```bash
go get github.com/logrusorgru/aurora
```

### Running with docker

```bash
docker build -t commoncrawler .
docker run -it commoncrawler
```

### Downloading data with the application

First configure the type of data you want to extract.

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

Then you can simply just build and run it as an executable.

```bash
go build src/*.go
go install src/*.go
```

Or you can run simply just run it.

```bash
go run src/*.go
```

### Resources

- MIT Licensed

- If people are interested or need it. I can create a documentation and tutorial page on https://commoncrawl.chriscates.ca

- You can post issues if they are valid, and, I could potentially fund them based on priority.
