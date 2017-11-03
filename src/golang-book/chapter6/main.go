package main

import "fmt"

func main() {
  // Array
  // arr := [5]float64{ 1, 2, 3, 4, 5 }
  // var total float64
  // for i, value := range(arr) {
  //   fmt.Println("i =", i)
  //   total += value
  // }
  // fmt.Println(total / float64(len(arr)))




  // Slice
  // slice := make([]float64, 5, 10)
  // fmt.Println(len(slice))

  // slice1 := []int{ 1, 2, 3 }
  // slice2 := append(slice1, 4, 5)
  // fmt.Println(slice1, slice2)

  // slice1 := []int{ 1, 2, 3 }
  // slice2 := make([]int, 2)
  // copy(slice2, slice1)
  // fmt.Println(slice1, slice2)




  // Map
  m := map[string]int{
    "a": 1,
    "b": 2,
  }
  m["c"] = 3
  fmt.Println(m)

  delete(m, "c")
  fmt.Println(m)

  if v, ok := m["b"]; ok {
    fmt.Println(v, ok)
  }
}
