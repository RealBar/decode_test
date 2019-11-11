package utils

import (
	"github.com/google/uuid"
	"math/big"
	"strings"
)

func TruncateFilename(s string, limit int) string {
	r := []rune(s)
	runeLen := len(r)
	if runeLen > limit {
		i := strings.LastIndex(s, ".")
		if i != -1 && (len(r)-i-1 <= limit) {
			suffix := r[i+1:]
			r = append(r[:i-(runeLen-limit)], suffix...)
		} else {
			r = r[:limit]
		}
		return string(r)
	}
	return s
}

func GenerateUUID() string {
	var value big.Int
	value.SetString(strings.Replace(uuid.New().String(), "-", "", 4), 16)
	return value.Text(16)
}
