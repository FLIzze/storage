package server

import (
    "net"
    "fmt"
    "log"
    c "server/client"
)

type Server struct {
    Host string
    Port string
}

func (server *Server) Run() {
    listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", server.Host, server.Port))
    if err != nil {
        log.Fatal(err)
    }

    defer listener.Close()

    fmt.Printf("Server running on: %s/%s\n", server.Host, server.Port)

    for {
        conn, err := listener.Accept()
        if err != nil {
            // log.Fatal(err)
            fmt.Println(err);
            return;
        }

        client := &c.Client{
            Conn: conn,
        }

        go client.HandleRequest()
    }
}

