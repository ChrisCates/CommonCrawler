package crawler

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func (c *crawler) Download(uri string, path string) error {
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

func (c *crawler) Extract(path string) error {
	//get extracted file path
	_, fname := filepath.Split(path)
	ext := filepath.Ext(fname)
	extractedPath := path[:len(path)-len(ext)]
	//create extruction destination

	out, err := os.Create(extractedPath)
	if err != nil {
		return err
	}
	defer out.Close()

	//open gzip file
	fi, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fi.Close()
	//create gz reader
	fz, err := gzip.NewReader(fi)
	if err != nil {
		return err
	}
	defer fz.Close()

	//write extracted to file
	_, err = io.Copy(out, fz)
	if err != nil {
		return err
	}

	return nil

}
