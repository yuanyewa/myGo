package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
func fetch(url string, ch chan<- string) {
	start := time.Now()
	r, e := http.Get(url)
	if e != nil {
		ch <- fmt.Sprint(e)
		return
	}
	n, e := io.Copy(ioutil.Discard, r.Body)
	r.Body.Close()
	if e != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, e)
		return
	}
	ch <- fmt.Sprintf("%.2fs %7d %s", time.Since(start).Seconds(), n, url)
}
