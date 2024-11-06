package main

import (
    s "server/server"
)

func main() {
    server := &s.Server{
        Host: "localHost",
        Port: "6969",
    }

    server.Run()
}
