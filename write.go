package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"unsafe"
)

type version_json struct {
	Version string `json:"version"`
	Data    `json:"data"`
}

type Data [][]string

var v = &version_json{}

func getFileMd5(filename string) string {
	// 文件全路径名
	path := fmt.Sprintf("%s", filename)
	pFile, err := os.Open(path)
	if err != nil {
		fmt.Errorf("打开文件失败，filename=%v, err=%v", filename, err)
		return ""
	}
	defer pFile.Close()
	md5h := md5.New()
	io.Copy(md5h, pFile)

	return hex.EncodeToString(md5h.Sum(nil))
}

func writeVersion(version, writename, filename string) {
	v.Data = append(v.Data, []string{writename, getFileMd5(filename)})
}

func Admission(version string) {
	v = new(version_json)
	v.Version = version
}

func Do(filename string) {
	byt, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filename, byt, 0644)
	if err != nil {
		panic(err)
	}
}

func StringBytes(s string) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: stringHeader.Data,
		Cap:  stringHeader.Len,
		Len:  stringHeader.Len,
	}))
}
