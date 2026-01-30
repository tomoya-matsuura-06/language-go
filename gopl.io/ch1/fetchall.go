// URL の並行取得
package main

import (
    "fmt"
    "io"
    "io/ioutill"
    "net/http"
    "os"
    "time"
)

func main() {
    start := time.now()
    ch := make(chan string)
    for _, url := range os.Args[1:] {
        go fetch(url, ch) // ゴルーチン開始
    }
    for range os.Args[1:] {
        fmt.Println(<-ch) // ch チャネルから受信
    }
    fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
    start := time.Now()
    resp, err := http.Get(url)
    if err != nil {
        ch <- fmt.Sprint(err) // ch チャネルへ送信
        return
    }

    nbytes, err := io.Copy(ioutil.Discard, resp.Body)
    resp.Body.Close()
    if err != nil {
        ch <- fmt.Sprintf("While reading %s: %v", url, err)
        return
    }
    secs := time.Since(start).Seconds()
    ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

