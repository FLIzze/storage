package client

import (
    "net"
    "bufio"
    h "server/header"
)

type Client struct {
    Conn net.Conn
}

func (client *Client) HandleRequest() {
    defer client.Conn.Close()
    reader := bufio.NewReader(client.Conn)

    h.GetHeader(reader)

    client.Conn.Close()
}
