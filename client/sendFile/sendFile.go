package sendFile

import (
	"log"
    "encoding/binary"
	"mime/multipart"
	"net"
    "fmt"
	"strings"
)

const (
     host = "localhost";
     port = "6969";
)

func SendFileData(file *multipart.FileHeader) {
    fmt.Println("Sending file")

    if file == nil {
        log.Fatal("File is nil")
    }

    conn := connectToServer()
    header := createHeader(file) 
    body := createBody(file)

    _, err := conn.Write(header)
    if err != nil {
        fmt.Printf("Error sending header: %v\n", err)
        return
    }

    _, err = conn.Write(body)
    if err != nil {
        fmt.Printf("Error sending body: %v\n", err)
        return
    }

    fmt.Println("File sent")
}

func createHeader(file *multipart.FileHeader) []byte {
    spliited_file_name := strings.Split(file.Filename, ".")
    file_extension := spliited_file_name[1]
    file_name := spliited_file_name[0]

    file_content_len := file.Size

    file_extension_len := len(file_extension)
    file_name_len := len(file_name)

    totalHeaderLen := 2 + file_extension_len +
                    2 + file_name_len +
                    4 

    header := make([]byte, totalHeaderLen)

    cursor := 0

    binary.BigEndian.PutUint16(header[cursor:], uint16(file_extension_len))
    cursor += 2
    copy(header[cursor:], []byte(file_extension))
    cursor += file_extension_len

    binary.BigEndian.PutUint16(header[cursor:], uint16(file_name_len))
    cursor += 2
    copy(header[cursor:], []byte(file_name))
    cursor += file_name_len

    binary.BigEndian.PutUint32(header[cursor:], uint32(file_content_len))

    return header
}

func createBody(file *multipart.FileHeader) []byte {
    file_content, err := file.Open()
    if err != nil {
        log.Fatal("Could not open file", err)
    }

    defer file_content.Close()

    file_content_bytes := make([]byte, file.Size)

    _, err = file_content.Read(file_content_bytes)
    if err != nil {
        log.Fatal("Could not read file", err)
    }

    return file_content_bytes
}

func connectToServer() net.Conn {
    fmt.Println("Connecting to server on " + host + ":" + port)

    conn, err := net.Dial("tcp", host + ":" + port)
    if err != nil {
        log.Fatal("Could not connect to server", err)
    }

    return conn
}
