package main

import "fmt"

func main() {
  // x := 5
  // zero(&x)
  // fmt.Println(x)

  xPtr := new(int)
  one(xPtr)
  fmt.Println(*xPtr)
}

func zero(xPtr *int) {
  *xPtr = 0
}

func one(xPtr *int) {
  *xPtr = 1
}
