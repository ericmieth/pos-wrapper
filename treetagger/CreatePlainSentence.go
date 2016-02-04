package treetagger

import (
	"regexp"
	"strings"
)

// stripping of annotations from a string
func CreatePlainSentence(inputSentence string) []byte {
	re := regexp.MustCompile("\\|[A-Z$.(,]*")
	plainSentence := re.ReplaceAllLiteralString(inputSentence, "")

	plainSentence = strings.Replace(plainSentence, " ", "\n", -1)

	return []byte(plainSentence)
}
