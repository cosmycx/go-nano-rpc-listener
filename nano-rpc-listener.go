package main

import (
  "fmt"
  "net/http"
  "encoding/json"
)

type Block interface{}

func main() {

  http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request){

    var block Block
		json.NewDecoder(r.Body).Decode(&block)

    fmt.Println(block)
  })

  fmt.Println(http.ListenAndServe(":10077", nil))

}
