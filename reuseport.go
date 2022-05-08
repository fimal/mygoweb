package main

import (
  "fmt"
  "html"
  "net/http"
  "os"
  "github.com/kavu/go_reuseport"
)

func main() {
  listener, err := reuseport.Listen("tcp", "8080")
  if err != nil {
    panic(err)
  }
  defer listener.Close()

  server := &http.Server{}
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Println(os.Getgid())
    fmt.Fprintf(w, "Hello, %q\n", html.EscapeString(r.URL.Path))
  })

  panic(server.Serve(listener))
}
