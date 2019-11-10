package utils

import (
	"decode_test/pkg/app"
	"strings"
)


func TruncateFilename(s string) string {
	r := []rune(s)
	runeLen := len(r)
	if runeLen > app.MaxFileNameLen {
		i := strings.LastIndex(s, ".")
		if i != -1 && (len(r)-i-1 <= app.MaxFileSuffixLen) {
			suffix := r[i+1:]
			r = append(r[:i-(runeLen-app.MaxFileNameLen)], suffix...)
		} else {
			r = r[:app.MaxFileNameLen]
		}
		return string(r)
	}
	return s
}
