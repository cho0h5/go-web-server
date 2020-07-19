package main

import (
    "fmt"
    "net"
    "log"
    "os"
)

func main() {
    fmt.Println("Start Server")

    startServer()
}

func startServer() {
    l, err := net.Listen("tcp", ":8080")
    if err != nil {
      log.Fatal(err)
    }
    defer l.Close()

    for {
        conn, err := l.Accept()
        if err != nil {
          log.Println(err)
          continue
        }
        go handler(conn)
    }
}

func handler(conn net.Conn) {
    header := "HTTP/1.1 200 OK\r\n\r\n"
    conn.Write([]byte(header))

    f, err := os.Open("static/index.html")
    if err != nil {
        log.Println(err)
    }

    b:= make([]byte, 1024)
    _, err = f.Read(b)
    if err != nil {
        log.Println(err)
    }
    conn.Write(b)

    conn.Close()
}
