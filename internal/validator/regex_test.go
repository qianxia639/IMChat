package validator

import (
	"testing"

	"github.com/dlclark/regexp2"
)

func TestRegexp(t *testing.T) {
	reg, _ := regexp2.Compile("^(?=.*[A-Za-z])(?=.*\\d|.*[\\W_])[A-Za-z\\d\\W_]{8,16}$", 0)
	// _, err := reg.MatchString(req.Password)

	testCases := []string{
		"qqqqq",
		"qianxia123",
		"qianxia",
		"_qianxia",
		"qianxia?!",
		"qianxia?!@",
	}

	for _, s := range testCases {
		t.Logf("%s ----------> %t\n", s, matchRegexp2(reg, s))
	}
}

func matchRegexp2(reg *regexp2.Regexp, str string) bool {
	matched, _ := reg.MatchString(str)
	return matched
}
