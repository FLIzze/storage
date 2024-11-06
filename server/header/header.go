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


func GetHeader(reader *bufio.Reader) error {
    header := &Header{}

    extLenBytes, err := u.Read(reader, EXTENSION_LEN)
    if err != nil {
        return err
    }
    extLen := binary.BigEndian.Uint16(extLenBytes)
    fmt.Println(extLen)

    extNameBytes, err := u.Read(reader, extLen)
    if err != nil {
        return err
    }
    header.ExtensionType = string(extNameBytes)
    fmt.Println(string(extNameBytes))

    fNameLenBytes, err := u.Read(reader, FILE_NAME_LEN)
    if err != nil {
        return err
    }
    fNameLen := binary.BigEndian.Uint16(fNameLenBytes)
    fmt.Println(fNameLen)

    fNameBytes, err := u.Read(reader, fNameLen)
    if err != nil {
        return err
    }
    header.FileName = string(fNameBytes)
    fmt.Println(string(fNameBytes))

    dLenBytes, err := u.Read(reader, DATA_LEN)
    if err != nil {
        return err
    }
    dataLen := binary.BigEndian.Uint32(dLenBytes)
    fmt.Println(dataLen)

    fileContent, err := u.Read32(reader, dataLen)
    if err != nil {
        return err
    }

    header.saveFile(fileContent)
    return nil
}
