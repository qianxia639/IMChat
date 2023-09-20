package validator

import (
	"regexp"
	"testing"
)

func TestRegexp(t *testing.T) {
	var isBool = regexp.MustCompile("^[a-zA-Z0-9_?!]{6,20}$").MatchString

	testCases := []string{
		"qqqqq",
		"qianxia123",
		"qianxia",
		"_qianxia",
		"qianxia?!",
		"qianxia?!@",
	}

	for _, s := range testCases {
		t.Logf("%s ----------> %t\n", s, isBool(s))
	}
}
