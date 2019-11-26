package main

import (
    "fmt"

    "encoding/json"
    "net/http"
)

type Message struct {
    Id      string `json:"id"`
    Message string `json:"message"`
}

func HelloWorld(res http.ResponseWriter, req *http.Request) {
    message := Message{"1", "Hello world"}


    js, err := json.Marshal(message)

    if err != nil {
        http.Error(res, err.Error(), http.StatusInternalServerError)
        return
    }

    res.Header().Set("Content-Type", "application/json")
    res.Write(js)
}

func Healthz(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(res, "OK")
}

func main() {
    http.HandleFunc("/", HelloWorld)
    http.HandleFunc("/healthz", Healthz)
    http.ListenAndServe(":3000", nil)
}
