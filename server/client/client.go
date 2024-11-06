package client

import (
    "net"
    "fmt"
    "bufio"
    h "server/header"
)

type Client struct {
    Conn net.Conn
}

func (client *Client) HandleRequest() {
    fmt.Println("\nNew message")
    defer client.Conn.Close()
    reader := bufio.NewReader(client.Conn)

    h.GetHeader(reader)

    client.Conn.Close()
}
