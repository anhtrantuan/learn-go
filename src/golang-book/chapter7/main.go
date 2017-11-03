package main

import (
  "fmt"
  "os"
  "bufio"
)

func main() {
  xs := []float64{ 98, 93, 77, 82, 83 }
  fmt.Println(average(xs))
  fmt.Println()

  xs1 := []int{ 1, 2, 3 }
  fmt.Println(add(xs1...))
  fmt.Println()

  multiply := func(args ...int) int {
    total := 1
    for _, v := range args {
      total *= v
    }
    return total
  }
  fmt.Println(multiply(1, 2, 3, 4, 5))
  fmt.Println()

  nextOdd := makeOddGenerator()
  fmt.Println(nextOdd())
  fmt.Println(nextOdd())
  fmt.Println(nextOdd())
  fmt.Println()

  f, _ := os.Open("main.go")
  defer f.Close()
  scanner := bufio.NewScanner(f)
  scanner.Split(bufio.ScanLines)
  fmt.Println("======================= main.go =======================")
  for scanner.Scan() {
    fmt.Println(scanner.Text())
  }
  fmt.Println("=======================================================")
  fmt.Println()

  defer func() {
    str := recover()
    fmt.Println("PANIC:", str)
  }()
  panic("Error happened")
}

func average(xs []float64) (avg float64) {
  total := 0.0
  for _, v := range xs {
    total += v
  }
  avg = total / float64(len(xs))
  return
}

func add(args ...int) int {
  total := 0
  for _, v := range args {
    total += v
  }
  return total
}

func makeOddGenerator() func() uint {
  i := uint(1)
  return func() (ret uint) {
    ret = i
    i += 2
    return
  }
}
