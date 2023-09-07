package bytesconv

import "unsafe"

// BytesToString 在不分配内存的情况下将 []byte 转换为 string
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// StringToBytes 在不分配内存的情况下将 string 转换为 []byte
func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}
