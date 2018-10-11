package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

func download(uri string, path string) error {
	//check if file exists
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return fmt.Errorf("file \"%s\" has already been downloaded", path)
	}
	//create output file
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	//make a GET to the specified URL
	resp, err := http.Get(uri)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// check the server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	//redirect get responce to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func extract(path string) bool {

	fmt.Println("Unzipping", path)
	_, err := exec.Command("gunzip", path).CombinedOutput()
	return err == nil
}
