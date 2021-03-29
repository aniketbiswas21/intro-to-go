package utils

import (
	"strings"
)

func MakeExcited(s string) string {
	k := strings.ToUpper(s) + "!"
	return k
}
