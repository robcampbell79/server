package main

import(
  "encoding/json"
  "net/http"
  "fmt"
)

type Vampire struct {
  Name string `json:"Name"`
  Country string `json:"Country"`
  Died string `json:"Died"`
}

type Dracula []Vampire

func returnVampire(w http.ResponseWriter, r *http.Request) {
  dracula := Dracula{
    Vampire{Name: "Vlad Tepes", Country: "Romaina", Died: "December 1476"},
  }

  fmt.Println("Endpoint Hit: returnVampire")

  json.NewEncoder(w).Encode(dracula)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, welcome to this index page.")
}

func main() {
  http.HandleFunc("/", handleIndex)
  http.HandleFunc("/vampire", returnVampire)

  http.ListenAndServe(":8088", nil)
}
