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
	ch := make(chan string) //TODO ちょっと何言ってるかわからない
	for i, url := range os.Args[1:] {
		go fetch(url, ch, i) //start a goroutine
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) //receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string, i int) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) //send to channel ch
		return
	}
	b, _ := io.ReadAll(resp.Body)
	ioutil.WriteFile("./files", b, 0777)

	nbytes, err := io.Copy(ioutil.Discard, resp.Body) // ToDO
	// memo ioutil.Discard はとりあえず書き込んでbyte 数を題したいときに使うと思う
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
