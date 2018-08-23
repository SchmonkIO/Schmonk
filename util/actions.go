package util

import (
	"strings"
)

func CheckAction(message []byte, action string) bool {
	searchString := "\"action\":\"" + action + "\""
	b := strings.Contains(string(message), searchString)
	return b
}
