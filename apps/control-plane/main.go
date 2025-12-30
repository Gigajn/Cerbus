package main

import (
  "fmt"
  "net/http"
)

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintln(w, "ok")
  })
  http.ListenAndServe(":8080", mux)
}
