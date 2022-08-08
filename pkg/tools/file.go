package tools

import (
	// "fmt"
	"bytes"
	"os"
)


// handle log
func SaveToLocalFile(filename string, data []byte) {
	//line := []byte("\n")
	// if exist and create
	file,err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)

	if err != nil && os.IsNotExist(err) {
		// println("file is not exist")//需要加[]
		start := []byte("[")
		end := []byte("]")
		newdata := BytesCombine(start,data,end)

		file, _ = os.Create(filename)
		file.Write(newdata)
		file.Seek(-1,2)
	}else {
		//println("file exist and append")
		//fmt.Println(file.Seek(0, 2))
		start := []byte("[")
		end := []byte("]")
		newdata := BytesCombine(start,data,end)
		file.Seek(-1,2)
		file.Write(newdata)
	}
	file.Close()
}

// BytesCombine 多个[]byte数组合并成一个[]byte
func BytesCombine(pBytes ...[]byte) []byte {
	len := len(pBytes)
	s := make([][]byte, len)
	for index := 0; index < len; index++ {
		s[index] = pBytes[index]
	}
	sep := []byte("")
	return bytes.Join(s, sep)
}