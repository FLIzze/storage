package main

import (
    s "server/server"
)

func main() {
    server := &s.Server{
        Host: "192.168.1.25",
        Port: "6969",
    }

    server.Run()
}
