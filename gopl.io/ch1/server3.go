// http リクエストを返す
package main

import (
    "fmt"
    "log"
    "net/http"
    "sync"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
    for k, v := range r.Header {
         fmt.FPrintf(w, "Header[%q] = %q\n", k, v)
    }
    fmt.FPrintf(w, "Host = %q\n", r.Host)
    fmt.FPrintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
    if err := r.ParseForm(); err != nil {
        log.Print(err)
    }
    for k, v := range r.Form {
        fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
    }
}

