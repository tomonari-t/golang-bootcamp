package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func formatUrl(url string) string {
	if strings.HasPrefix(url, "http://") {
		return url
	}

	return "http://" + url
}

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(formatUrl(url))
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s:  %v\n", url, err)
			os.Exit(1)
		}
		resp.Body.Close()
	}
}
