package utils

import (
	"fmt"
	"strings"
)

func DesensitizationEmail(email string) (result string) {
	index := strings.Index(email, "@")

	result = fmt.Sprintf("%s****%s", email[0:2], email[index-2:])

	return

}
