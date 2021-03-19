package funcs

import (
	"strings"
)

func Implode(a []string, sep string) string {
	return strings.Join(a, sep)
}
