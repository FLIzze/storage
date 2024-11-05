package main

import (
    "net"
    "fmt"
    "log"
    "bufio"
    "os"
    "encoding/binary"
    "io"
)

const (
    SAVEPATH = "./storage/"
    EXTENSION_LEN = 2
    FILE_NAME_LEN = 2
    DATA_LEN = 4
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

func (client *Client) handleRequest() {
    fmt.Println("\nNew message")
    defer client.conn.Close()
    reader := bufio.NewReader(client.conn)

    getHeader(reader)

    client.conn.Close()
}

func (server *Server) Run() {
    listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", server.host, server.port))
    if err != nil {
        log.Fatal(err)
    }

    defer listener.Close()

    fmt.Printf("Server running on: %s/%s\n", server.host, server.port)

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

func (header *Header) saveFile(fileContent []byte) {
    filePath := SAVEPATH + header.fileName + "." + header.extensionType

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

func read(reader *bufio.Reader, len uint16) ([]byte, error) {
    tmp := make([]byte, len)
    if _, err := io.ReadFull(reader, tmp); err != nil {
        return nil, err
    }
    return tmp, nil
}

func read32(reader *bufio.Reader, len uint32) ([]byte, error) {
    tmp := make([]byte, len)
    if _, err := io.ReadFull(reader, tmp); err != nil {
        return nil, err
    }
    return tmp, nil
}

func getHeader(reader *bufio.Reader) error {
    header := &Header{}

    extLenBytes, err := read(reader, EXTENSION_LEN)
    if err != nil {
        return err
    }
    extLen := binary.BigEndian.Uint16(extLenBytes)
    fmt.Println(extLen)

    extNameBytes, err := read(reader, extLen)
    if err != nil {
        return err
    }
    header.extensionType = string(extNameBytes)
    fmt.Println(string(extNameBytes))

    fNameLenBytes, err := read(reader, FILE_NAME_LEN)
    if err != nil {
        return err
    }
    fNameLen := binary.BigEndian.Uint16(fNameLenBytes)
    fmt.Println(fNameLen)

    fNameBytes, err := read(reader, fNameLen)
    if err != nil {
        return err
    }
    header.fileName = string(fNameBytes)
    fmt.Println(string(fNameBytes))

    dLenBytes, err := read(reader, DATA_LEN)
    if err != nil {
        return err
    }
    dataLen := binary.BigEndian.Uint32(dLenBytes)
    fmt.Println(dataLen)

    fileContent, err := read32(reader, dataLen)
    if err != nil {
        return err
    }

    header.saveFile(fileContent)
    return nil
}

func main() {
    server := &Server{
        host: "localhost",
        port: "6969",
    }

    server.Run()
}
