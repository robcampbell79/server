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

type Vampires []Vampire

func returnVampire(w http.ResponseWriter, r *http.Request) {
  vampire := Vampires{
    Vampire{Name: "Vlad Tepes", Country: "Romania", Died: "December 1476"},
    Vampire{Name: "Elizabeth Bathory", Country: "Hungary", Died: "August 1614"},
  }

  fmt.Println("Endpoint Hit: returnVampire")
  fmt.Println(vampire)

  json.NewEncoder(w).Encode(vampire)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, welcome to this index page.")
}

func main() {
  http.HandleFunc("/", handleIndex)
  http.HandleFunc("/vampire", returnVampire)

  http.ListenAndServe(":8088", nil)
}
