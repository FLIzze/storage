package utils

import (
    "bufio"
    "io"
)

func Read(reader *bufio.Reader, len uint16) ([]byte, error) {
    tmp := make([]byte, len)
    if _, err := io.ReadFull(reader, tmp); err != nil {
        return nil, err
    }
    return tmp, nil
}

func Read32(reader *bufio.Reader, len uint32) ([]byte, error) {
    tmp := make([]byte, len)
    if _, err := io.ReadFull(reader, tmp); err != nil {
        return nil, err
    }
    return tmp, nil
}
