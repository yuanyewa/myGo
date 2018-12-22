package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		r, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stdout, "Fetch: %v\n", err)
			os.Exit(1)
		}
		// b, err := ioutil.ReadAll(r.Body)
		// r.Body.Close()
		// fmt.Printf("%s\n", b)
		_, err = io.Copy(os.Stdout, r.Body)
		if err != nil {
			fmt.Fprintf(os.Stdout, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Println(r.Status)
	}
}
