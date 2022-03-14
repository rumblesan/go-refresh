package server

import (
  "fmt"
  "encoding/json"
  "log"
  "rumblesan.com/go-refresh/vector"
  "net/http"
)

func Start () {

  http.HandleFunc("/", HelloHandler)
  http.HandleFunc("/vector", RandomVectorHandler)

  log.Fatal(http.ListenAndServe(":8080", nil))

}

func HelloHandler (w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hey!\n")
}

func RandomVectorHandler (w http.ResponseWriter, r *http.Request) {

  v := vector.RandomVector()
  data, _ := json.Marshal(v)
  fmt.Fprintf(w, string(data))

}
