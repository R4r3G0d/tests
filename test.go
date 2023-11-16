package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("X-XSS-Protection", "0")

    messages, ok := r.URL.Query()["message"]
    if !ok {
       messages = []string{"hello, world"}
    }
    fmt.Fprintf(w, "<html><p>%v</p></html>", messages[0])
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
