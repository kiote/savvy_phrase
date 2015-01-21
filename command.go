package main

import ( 
          "fmt"
  
          "os"
       )

func main() {
  args := os.Args[1:]

  first := args[0]
  second := args[1]
 
  fmt.Println("args 1,2:")
  fmt.Println(first)
  fmt.Println(second)
}
