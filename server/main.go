package main

import (
  "net/http"
  "encoding/json"
  "os"
  "time"
  "fmt"
  "io/ioutil"
  //"reflect"
  //"log"
  "github.com/gorilla/mux"
  "github.com/gorilla/handlers"
  "github.com/dgrijalva/jwt-go"
  "github.com/auth0/go-jwt-middleware"
)

func main() {
  r := mux.NewRouter()

  //r.Methods("OPTIONS").HandlerFunc(HandleCors)
  //r.Methods("GET").HandlerFunc(HandleCors)

  r.Handle("/", http.FileServer(http.Dir("./views/")))

  //Status of server
  r.Handle("/status", StatusHandler).Methods("GET")

  //Dummy data
  //r.Handle("/dogs", jwtMiddleware.Handler(DogsHandler)).Methods("GET")

  r.Handle("/dogs", DogsHandler).Methods("GET")

  r.Handle("/dogs", PreflightDogsHandler).Methods("OPTIONS")

  //oauth
  r.Handle("/get-token", GetTokenHandler).Methods("POST")



  //fmt.Println(reflect.TypeOf(jwtMiddleware.Handler(DogsHandler)))

  http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, r))
}

//var HandleCors = func

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

  if origin := r.Header.Get("Origin"); origin != "" {
    w.Header().Set("Access-Control-Allow-Origin", origin)
  }

  fmt.Println(r.Header.Get("Authorization"))

  tokenString := r.Header.Get("Authorization")

  token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

    /*if _, ok := token.Method.(jwt.SigningMethodHS256); !ok {
      return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
    }*/

    return signingKey, nil

  })

  if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
    fmt.Println(claims["google"])

    data, _ := json.Marshal(dogs)

    w.Write([]byte(data))

  } else {
    fmt.Println(err)
    w.Write([]byte("Not Authorized"))
    //w.Write([]byte(err))
  }
})

var PreflightDogsHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  if origin := r.Header.Get("Origin"); origin != "" {
    w.Header().Set("Access-Control-Allow-Origin", origin)
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    w.Header().Set("Access-Control-Allow-Headers",
            "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
    }
})

var signingKey = []byte("secret")

var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

  if origin := r.Header.Get("Origin"); origin != "" {
    w.Header().Set("Access-Control-Allow-Origin", origin)
  }

  //decoder := json.NewDecoder(r.Body.Get())

  body, err := ioutil.ReadAll(r.Body)
  if err != nil {
    panic(err)
  }

  fmt.Println(string(body))
  //body := r.Body

  //fmt.Println(body)

  //setup token
  token := jwt.New(jwt.SigningMethodHS256)

  claims := token.Claims.(jwt.MapClaims)

  claims["google"] = string(body)
  claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

  //sign the token
  tokenString, _ := token.SignedString(signingKey)


  //w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Write([]byte(tokenString))
  //w.Write([]byte("Test working"))
})

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
  ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
    fmt.Println("here")
    return signingKey, nil
  },
  SigningMethod: jwt.SigningMethodHS256,
})