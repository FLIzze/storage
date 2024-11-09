package header

import (
    "bufio"
    "encoding/binary"
    "fmt"
    "os"
    "log"
    u "server/utils"
)

const (
    SAVEPATH = "./storage/"
    EXTENSION_LEN = 2
    FILE_NAME_LEN = 2
    DATA_LEN = 4
)

type Header struct {
    ExtensionType string
    FileName string
}

func (header *Header) saveFile(fileContent []byte) {
    filePath := SAVEPATH + header.FileName + "." + header.ExtensionType

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


func GetHeader(reader *bufio.Reader) {
    header := &Header{}

    extLenBytes, err := u.Read(reader, EXTENSION_LEN)
    if err != nil {
        fmt.Printf("Error reading extension length: %s\n", err)
        return 
    }
    extLen := binary.BigEndian.Uint16(extLenBytes)

    extNameBytes, err := u.Read(reader, extLen)
    if err != nil {
        fmt.Printf("Error reading extension name: %s\n", err)
        return 
    }
    header.ExtensionType = string(extNameBytes)

    fNameLenBytes, err := u.Read(reader, FILE_NAME_LEN)
    if err != nil {
        fmt.Printf("Error reading file name length: %s\n", err)
        return 
    }
    fNameLen := binary.BigEndian.Uint16(fNameLenBytes)

    fNameBytes, err := u.Read(reader, fNameLen)
    if err != nil {
        fmt.Printf("Error reading file name: %s\n", err)
        return 
    }
    header.FileName = string(fNameBytes)

    dLenBytes, err := u.Read(reader, DATA_LEN)
    if err != nil {
        fmt.Printf("Error reading data length: %s\n", err)
        return 
    }
    dataLen := binary.BigEndian.Uint32(dLenBytes)

    fileContent, err := u.Read32(reader, dataLen)
    if err != nil {
        fmt.Printf("Error reading file content: %s\n", err)
        return 
    }

    fmt.Printf("Header: %s, %s, %d\n", header.ExtensionType, header.FileName, dataLen)
    header.saveFile(fileContent)
}
