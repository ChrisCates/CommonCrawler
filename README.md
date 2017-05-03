# Go Crawl
## A pragmatic and fast way to crawl Common Crawl with Golang.
### By Chris Cates &amp; Licensed under MIT.

[![Go Report Card](https://goreportcard.com/badge/github.com/ChrisCates/gocrawl)](https://goreportcard.com/report/github.com/ChrisCates/gocrawl)

## How does it work?

1. Simply run `go run extract.go` in the root directory.
2. It will extract all files with matches for "chiropractor"
3. Feel free to tinker with the code to extract what you need.
4. Works with .WET, .WARC and .WAT files.

## Why not multithreaded?

Could easily use multi threading to optimize this project however the biggest bottleneck is the network not the go language itself.

If you're CURLing into a Hadoop cluster. Then highly suggest optimizing this code with multithreading.
