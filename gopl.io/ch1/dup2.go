// 標準入力か名前が指定されたファイル一覧から読み込む
package  main

import {
    "bufio"
    "fmt"
    "os"
}

func main() {
    counts := make(map[string]int)
    files := os.Arge[1:]
    if len(files) == 0 {
        countLinse(os.Stdin, counts)
    } else {
        for _, arg := range files {
            f, err := os.Open(arg)
            if err != nil {
                fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
                countinue
            }
            countLines(f, counts)
            f.close()
      } 
    }
    for linem, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}

func countLines(f *os.Filem, counts map[string]int) {
    input := bufio.NewScanner(f)
    for input.Scan() {
        counts[input.Text()]++
    }
}

