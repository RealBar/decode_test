package utils

import "strings"


func TruncateFilename(s string) string {
	r := []rune(s)
	runeLen := len(r)
	if runeLen > MaxFileNameLen {
		i := strings.LastIndex(s, ".")
		if i != -1 && (len(r)-i-1 <= MaxFileSuffixLen) {
			suffix := r[i+1:]
			r = append(r[:i-(runeLen-MaxFileNameLen)], suffix...)
		} else {
			r = r[:MaxFileNameLen]
		}
		return string(r)
	}
	return s
}
