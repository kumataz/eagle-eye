package tools

import (
    "bytes"
    "encoding/binary"
    // "fmt"
)

func IntToBytes(n int) []uint8 {
    data := int64(n)
    bytebuf := bytes.NewBuffer([]byte{})
    binary.Write(bytebuf, binary.BigEndian, data)
    return bytebuf.Bytes()
}

func BytesToInt(bys []uint8) int {
    bytebuff := bytes.NewBuffer(bys)
    var data int64
    binary.Read(bytebuff, binary.BigEndian, &data)
    return int(data)
}
