package main

import (
  // "bytes"
  // "container/list"
  // "crypto/sha1"
  "encoding/gob"
  // "errors"
  // "flag"
  "fmt"
  "hash/crc32"
  "io"
  "io/ioutil"
  // "math/rand"
  "net"
  "net/http"
  "net/rpc"
  // "os"
  // "path/filepath"
  // "sort"
  // "strings"
  "sync"
  "time"
)

type Person struct {
  Name string
  Age int
}

type ByName []Person
func (this ByName) Len() int {
  return len(this)
}
func (this ByName) Less(i, j int) bool {
  return this[i].Name < this[j].Name
}
func (this ByName) Swap(i, j int) {
  this[i], this[j] = this[j], this[i]
}

type ByAge []Person
func (this ByAge) Len() int {
  return len(this)
}
func (this ByAge) Less(i, j int) bool {
  return this[i].Age < this[j].Age
}
func (this ByAge) Swap(i, j int) {
  this[i], this[j] = this[j], this[i]
}

type Server struct {}
func (this *Server) Negate(i int64, reply *int64) error {
  *reply = -i
  return nil
}

func getHash(filename string) (uint32, error) {
  bs, err := ioutil.ReadFile(filename)
  if err != nil {
    return 0, err
  }
  h := crc32.NewIEEE()
  h.Write(bs)
  return h.Sum32(), nil
}

// func server() {
//   ln, err := net.Listen("tcp", ":9999")
//   if err != nil {
//     fmt.Println(err)
//     return
//   }
//   for {
//     c, err := ln.Accept()
//     if err != nil {
//       fmt.Println(err)
//       continue
//     }
//     go handleServerConnection(c)
//   }
// }

func server() {
  rpc.Register(new(Server))
  ln, err := net.Listen("tcp", ":9999")
  if err != nil {
    fmt.Println(err)
    return
  }
  for {
    c, err := ln.Accept()
    if err != nil {
      continue
    }
    go rpc.ServeConn(c)
  }
}

func handleServerConnection(c net.Conn) {
  var msg string
  err := gob.NewDecoder(c).Decode(&msg)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println("Received", msg)
  }

  c.Close()
}

// func client() {
//   c, err := net.Dial("tcp", "127.0.0.1:9999")
//   if err != nil {
//     fmt.Println(err)
//     return
//   }
//
//   msg := "Hello World"
//   fmt.Println("Sending", msg)
//   err = gob.NewEncoder(c).Encode(msg)
//   if err != nil {
//     fmt.Println(err)
//   }
//
//   c.Close()
// }

func client() {
  c, err := rpc.Dial("tcp", "127.0.0.1:9999")
  if err != nil {
    fmt.Println(err)
    return
  }
  var result int64
  err = c.Call("Server.Negate", int64(999), &result)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println("Server.Negate(999) =", result)
  }
}

func hello(res http.ResponseWriter, req *http.Request) {
  res.Header().Set(
    "Content-Type",
    "text/html",
  )
  io.WriteString(
    res,
    `
      <doctype html>
      <html>
        <head>
          <title>Hello World</title>
        </head>

        <body>
          Hello World!
        </body>
      </html>
    `,
  )
}

func main() {
  // fmt.Println(
  //   strings.Contains("test", "es"),
  //   strings.Count("test", "t"),
  //   strings.HasPrefix("test", "te"),
  //   strings.HasSuffix("test", "st"),
  //   strings.Index("test", "e"),
  //   strings.Join([]string{"a","b"}, "-"),
  //   strings.Repeat("a", 5),
  //   strings.Replace("aaaa", "a", "b", 2),
  //   strings.Split("a-b-c-d-e", "-"),
  //   strings.ToLower("TEST"),
  //   strings.ToUpper("test"),
  // )
  // arr := []byte("test")
  // str := string([]byte{'t','e','s','t'})
  // fmt.Println(arr, str)

  // var buf bytes.Buffer
  // buf.Write([]byte("test"))
  // fmt.Println(string(buf.Bytes()))



  // file, err := os.Open("test.txt")
  // if err != nil {
  //   fmt.Println("Error Occurred:", err)
  //   return
  // }
  // defer file.Close()
  //
  // stat, err := file.Stat()
  // if err != nil {
  //   fmt.Println("Error Occurred:", err)
  //   return
  // }
  //
  // bs := make([]byte, stat.Size())
  // _, err = file.Read(bs)
  // if err != nil {
  //   fmt.Println("Error Occurred:", err)
  //   return
  // }
  //
  // str := string(bs)
  // fmt.Println(str)

  // bs, err := ioutil.ReadFile("test.txt")
  // if err != nil {
  //   fmt.Println("Error Occurred:", err)
  //   return
  // }
  // str := string(bs)
  // fmt.Println(str)

  // file, err := os.Create("test1.txt")
  // if err != nil {
  //   fmt.Println("Error Occurred:", err)
  //   return
  // }
  // defer file.Close()
  //
  // file.WriteString("test")

  // dir, err := os.Open(".")
  // if err != nil {
  //   fmt.Println("Error Occurred:", err)
  //   return
  // }
  // defer dir.Close()
  //
  // fileInfos, err := dir.Readdir(-1)
  // if err != nil {
  //   fmt.Println("Error Occurred:", err)
  //   return
  // }
  // for _, fi := range fileInfos {
  //   fmt.Println(fi.Name())
  // }

  // filepath.Walk("..", func(path string, info os.FileInfo, err error) error {
  //   fmt.Println(path)
  //   return nil
  // })



  // err := errors.New("error message")
  // fmt.Println(err)



  // var x list.List
  // x.PushBack(1)
  // x.PushBack(2)
  // x.PushBack(3)
  //
  // for e := x.Front(); e != nil; e = e.Next() {
  //   fmt.Println(e.Value.(int))
  // }



  // kids := []Person{
  //   { "Jill", 9 },
  //   { "Jack", 10 },
  // }
  // sort.Sort(ByName(kids))
  // fmt.Println(kids)
  // sort.Sort(ByAge(kids))
  // fmt.Println(kids)



  // h := crc32.NewIEEE()
  // h.Write([]byte("test"))
  // v := h.Sum32()
  // fmt.Println(v)

  // h1, err := getHash("test1.txt")
  // if err != nil {
  //   return
  // }
  // h2, err := getHash("test2.txt")
  // if err != nil {
  //   return
  // }
  // fmt.Println(h1, h2, h1 == h2)

  // h := sha1.New()
  // h.Write([]byte("test"))
  // bs := h.Sum([]byte{})
  // fmt.Println(bs)



  // go server()
  // go client()
  //
  // var input string
  // fmt.Scanln(&input)

  // http.HandleFunc("/hello", hello)
  // http.Handle(
  //   "/assets/",
  //   http.StripPrefix(
  //     "/assets/",
  //     http.FileServer(http.Dir("assets")),
  //   ),
  // )
  // http.ListenAndServe(":9000", nil)



  // maxp := flag.Int("max", 6, "the max value")
  // flag.Parse()
  // fmt.Println(rand.Intn(*maxp))



  m := new(sync.Mutex)

  for i := 0; i < 10; i++ {
    go func(i int) {
      m.Lock()
      fmt.Println(i, "start")
      time.Sleep(time.Second)
      fmt.Println(i, "end")
      m.Unlock()
    }(i)
  }

  var input string
  fmt.Scanln(&input)
}
