package main

import (
    "net"
    "fmt"
    "log"
    "bufio"
    "os"
    "time"
    "path/filepath"
)

var SAVEPATH = "./storage/"
var SAVETYPE = "jpg"

type Server struct {
    host string
    port string
}

type Client struct {
    conn net.Conn
}

func (client *Client) handleRequest() {
    reader := bufio.NewReader(client.conn)

    var wholeFileContent []byte

    // header := make([]byte, 20)

    // byte, err := reader.ReadByte()
    // if err != nil {
    //     fmt.Println("Error")
    //     return
    // }
    //
    // fmt.Println(byte)

    for {
        message, err := reader.ReadBytes('\n')
        if err != nil {
            fmt.Println(err)
            break
        }

        wholeFileContent = append(wholeFileContent, message...)
        client.conn.Write([]byte("Message received.\n"))
    }

     

    // saveFile(wholeFileContent)
    // fmt.Println(wholeFileContent)
    client.conn.Close()
}

func (server *Server) Run() {
    listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", server.host, server.port))
    if err != nil {
        log.Fatal(err)
    }

    defer listener.Close()

    fmt.Printf("Server running on : %s/%s\n", server.host, server.port)

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Fatal(err)
        }

        client := &Client{
            conn: conn,
        }

        go client.handleRequest()
    }
}

func saveFile(fileContent []byte) {
    fileName := fmt.Sprintf("image_%d.%s", time.Now().UnixNano(), SAVETYPE)
    filePath := filepath.Join(SAVEPATH, fileName)

    f, err := os.Create(filePath)
    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    w := bufio.NewWriter(f)

    _, err = w.Write(fileContent)
    if err != nil {
        log.Fatal(err)
    }

    w.Flush()

    fmt.Printf("File saved as: %s\n", filePath)
}

func main() {
    server := &Server{
        host: "localhost",
        port: "6969",
    }

    server.Run()
}
