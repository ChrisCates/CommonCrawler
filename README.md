# Common Crawler

## 🕸 A simple and easy way to extract data from Common Crawl with little or no hassle.

![Go Version](https://img.shields.io/badge/Go-v1.12.4-blue.svg)
![License](https://img.shields.io/badge/license-MIT-blue.svg)
[![Build Status](https://travis-ci.org/ChrisCates/CommonCrawler.svg?branch=master)](https://travis-ci.org/ChrisCates/CommonCrawler)
[![Go Report Card](https://goreportcard.com/badge/github.com/ChrisCates/CommonCrawler)](https://goreportcard.com/report/github.com/ChrisCates/CommonCrawler)

## As a library

***This will be implemented soon, please review issues for Gitcoin bounties***

Install as a dependency:

```bash
go get https://github.com/ChrisCates/CommonCrawler
```

Access the library functions by `import`ing it:

```golang
import(
  cc "github.com/ChrisCates/CommonCrawler"
)

func main() {
  cc.scan()
  cc.download()
  cc.extract()
  // And so forth
}
```

## As a command line tool

***This will be implemented soon, please review issues for Gitcoin bounties***

Install from source:

```bash
go install  https://github.com/ChrisCates/CommonCrawler
```

Or you can curl from Github:

```bash
curl https://github.com/ChrisCates/CommonCrawler/raw/master/dist/commoncrawler -o commoncrawler
```

Then run as a binary:

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

### With Docker

```bash
docker build -t commoncrawler .
docker run commoncrawler
```

### Without Docker

```bash
go build -i -o ./dist/commoncrawler ./src/*.go
./dist/commoncrawler
```

Or you can run simply just run it.

```bash
go run src/*.go
```

### Resources

- MIT Licensed

- If people are interested or need it. I can create a documentation and tutorial page on https://commoncrawl.chriscates.ca

- You can post issues if they are valid, and, I could potentially fund them based on priority.
