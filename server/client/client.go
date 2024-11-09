package client

import (
    "net"
    "bufio"
    "fmt"
    h "server/header"
)

type Client struct {
    Conn net.Conn
}

func (client *Client) HandleRequest() {
    fmt.Println("Handling request")

    defer client.Conn.Close()
    reader := bufio.NewReader(client.Conn)

    h.GetHeader(reader)

    client.Conn.Close()
}
