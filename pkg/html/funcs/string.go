package funcs

import (
	"html/template"
	"strings"
)

func Implode(a []string, sep string) string {
	return strings.Join(a, sep)
}

func Html(s string) template.HTML {
	return template.HTML(s)
}
