package main

import (
    s "server/server"
)

func main() {
    server := &s.Server{
        Host: "localhost",
        Port: "6969",
    }

    server.Run()
}
