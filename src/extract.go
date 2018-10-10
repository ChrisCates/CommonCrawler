package main

import (
	"fmt"
	"os/exec"
)

func extract(uri string, path string) bool {
	fmt.Println("CURLing", uri)
	_, err := exec.Command("curl", uri, "-o", path).CombinedOutput()
	if err != nil {
		return false
	}

	fmt.Println("Unzipping", path)
	_, err = exec.Command("gunzip", path).CombinedOutput()
	if err != nil {
		return false
	}

	return true
}
