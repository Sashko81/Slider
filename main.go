package main

import (
    "net/http"
    "math/rand"
    "fmt"
    "strconv"
)

func serveTemplate(w http.ResponseWriter, r *http.Request) {
    var number = rand.Intn(4)+1
    fmt.Fprintf(w, strconv.Itoa(number))
}

func main() {
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))
    http.HandleFunc("/", serveTemplate)
    http.ListenAndServe(":5000", nil)
}