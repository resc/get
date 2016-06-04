package main

import (
	"fmt"
	"io"
	"net/url"
	"path/filepath"
	//"io"
	"net/http"
	"os"

	"github.com/spf13/hugo"
)

func main() {
	if len(os.Args) != 2 {
		printHelp()
	}

	rawUrl := os.Args[1]
	statusCode := download(rawUrl)
	os.Exit(resp.StatusCode)
}

func printError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		os.Exit(1)
	}
}

func download(url string) int {
	url, err := url.Parse(rawUrl)
	if err != nil {
		printError(fmt.Errorf("Invalid url '%s': %s", rawUrl, err))
	}

	resp, err := http.Get(url.String())
	if err != nil {
		printError(err)
		return -1
	}

	defer resp.Body.Close()
	if 200 <= resp.StatusCode && resp.StatusCode < 300 {
		io.Copy(os.Stdout, io.LimitReader(resp.Body, 4*1024*1024))
	} else {
		fmt.Fprintf(os.Stderr, "%s\n", resp.Status)
	}
	return resp.StatusCode
}

func printHelp() {

	exe := os.Args[0]
	_, file := filepath.Split(exe)
	fmt.Fprintf(os.Stderr, "usage:\n %s [url]\n", file)
}
