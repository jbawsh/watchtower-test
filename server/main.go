package main

import (
  "net/http"
  "encoding/json"
  "os"
  "github.com/gorilla/mux"
  "github.com/gorilla/handlers"
)

func main() {
  r := mux.NewRouter()

  r.Handle("/", http.FileServer(http.Dir("./views/")))

  //Status of server
  r.Handle("/status", StatusHandler).Methods("GET")

  //Dummy data
  r.Handle("/dogs", DogsHandler).Methods("GET")

  //oauth
  r.Handle("/get-token", GetTokenHandler).Methods("POST")

  http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, r))

}

var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
  w.Write([]byte("Not Implemented"))
})

type Dog struct {
    Id int
    Name string
    Description string
}

var dogs = []Dog{
  Dog{Id: 1, Name: "Husky", Description : "Sheds a lot"},
  Dog{Id: 2, Name: "Border Collie", Description : "Super smart"},
  Dog{Id: 3, Name: "Golden Retriever", Description : "Very friendly"},
  Dog{Id: 4, Name: "German Shepard", Description: "Very protective"},
}

var StatusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
  w.Write([]byte("API is running"))
})

var DogsHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

  data, _ := json.Marshal(dogs)

  w.Header().Set("Content-Type", "application/json")
  w.Write([]byte(data))
})

var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)

  os.Stdout(vars)

  w.Header().Set("Content-Type", "application/json")
  w.Write([]byte("Test working"))
})