package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomString(t *testing.T) {
	str1 := RandomString(0)
	require.Empty(t, str1)

	str2 := RandomString(6)
	require.NotEmpty(t, str2)
	require.Len(t, str2, 6)
}

func TestRandomStringWithString(t *testing.T) {
	str1 := RandomStringWithString("", 0)
	require.Empty(t, str1)

	str2 := RandomStringWithString("", 6)
	require.NotEmpty(t, str2)
	require.Len(t, str2, 6)

	testString := "This is test data..."

	str3 := RandomStringWithString(testString, 0)
	require.Empty(t, str3)

	str4 := RandomStringWithString(testString, 6)
	require.NotEmpty(t, str4)
	require.Len(t, str4, 6)
}

func TestGender(t *testing.T) {
	m := make(map[int64]int)
	for i := 0; i < 1000; i++ {
		num := RandomInt(0, 2)
		m[num]++
	}

	t.Logf("map[int64]int: %+v\n", m)
}
