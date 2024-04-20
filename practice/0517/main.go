package main

import (
	"reflect"
	"unsafe"
)

func main() {

	a := "hello"
	b := StringToBytes(a)
	b[0] = 'H'
}

func StringToBytes(str string) (bytes []byte) {
	ss := *(*reflect.StringHeader)(unsafe.Pointer(&str))
	bs := (*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	bs.Data = ss.Data
	bs.Len = ss.Len
	bs.Cap = ss.Len
	return bytes
}

// BytesToString 实现 []byte 转换成 string, 不需要额外的内存分配
func BytesToString(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}
