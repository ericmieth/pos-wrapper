package rdrpostagger

import (
	"regexp"
)

// stripping of annotations from a string
func CreatePlainSentence(inputSentence string) []byte {
	re := regexp.MustCompile("\\|[A-Z$.(,]*")
	plainSentence := re.ReplaceAllLiteralString(inputSentence, "")  + "\n"
	return []byte(plainSentence)
}
