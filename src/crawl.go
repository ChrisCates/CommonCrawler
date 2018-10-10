package main

import (
	"fmt"
	"os"

	aurora "github.com/logrusorgru/aurora"
)

func main() {
	fmt.Println(aurora.Green("Getting configurations for Common Crawl Extractor..."))
	config := getConfiguration()

	fmt.Println(aurora.Blue("  Creating folder: " + config.dataFolder))
	os.Mkdir(config.dataFolder, 0740)
	fmt.Println(aurora.Blue("  Creating folder: " + config.matchFolder))
	os.Mkdir(config.matchFolder, 0740)

	fmt.Println(aurora.Green("Starting scanning..."))
	scan(config)
}
