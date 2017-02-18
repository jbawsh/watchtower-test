package main

import (
  "net/http"
  "github.com/gorilla/mux"
)

func main() {
  r := mux.NewRouter()

  r.Handle("/", http.FileServer(http.Dir("./views/")))

  http.ListenAndServe(":3000", r)
}