package structs

import (
    "net"
)

type Server struct {
    host string
    port string
}

type Client struct {
    conn net.Conn
}

type Header struct {
    extensionType string
    fileName string
}
