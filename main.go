package main

import (
    "fmt"
    "net"
    "log"
)

func main() {
    fmt.Println("Start Server");

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
    data := "HTTP/1.1 200 OK\n\nhello"
    conn.Write([]byte(data))

    conn.Close()
}
