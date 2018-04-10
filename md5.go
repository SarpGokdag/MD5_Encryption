package main

import (
  "fmt"
  "io"
  "crypto/md5"
)

func main(){
  h := md5.New()
  io.WriteString(h, " mesaj ")
  fmt.Printf("%X", h.Sum(nil))
  fmt.Printf("\n")
}
