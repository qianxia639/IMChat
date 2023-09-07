package bytesconv

import (
	"IMChat/utils"
	"bytes"
	"crypto/rand"
	"testing"
)

var testString = "This's test str..."
var testBytes = []byte(testString)

func rawBytesToString(b []byte) string {
	return string(b)
}

func rawStringTobytes(s string) []byte {
	return []byte(s)
}

func TestBytesToString(t *testing.T) {
	data := make([]byte, 1024)
	for i := 0; i < 100; i++ {
		rand.Read(data)
		if rawBytesToString(data) != BytesToString(data) {
			t.Fatal("don't match")
		}
	}
}

func TestStringToBytes(t *testing.T) {
	for i := 0; i < 100; i++ {
		str := utils.RandomString(6)
		if !bytes.Equal(rawStringTobytes(str), StringToBytes(str)) {
			t.Fatal("don't match")
		}
	}
}

func BenchmarkBytesConvBytesToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BytesToString(testBytes)
	}
}

func BenchmarkBytesConvBytesToStringRaw(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rawBytesToString(testBytes)
	}
}

func BenchmarkBytesConvStringToBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringToBytes(testString)
	}
}

func BenchmarkBytesConvStringToBytesRaw(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rawStringTobytes(testString)
	}
}
