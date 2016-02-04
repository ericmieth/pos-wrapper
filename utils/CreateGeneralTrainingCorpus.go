package utils

import (
	"strings"
)

// this contains a general parsing for all taggers
func CreateGeneralTrainingCorpus(inputSentence string, tagMap map[string]string) string {

	inputWords := strings.Split(inputSentence, " ")
	trainingSentence := ""
	for _, inputWord := range inputWords {
		inputWordSlice := strings.Split(inputWord, "|")
		word := strings.TrimSpace(inputWordSlice[0])
		tag := tagMap[strings.TrimSpace(inputWordSlice[1])]
		taggedWord := word + "|" + tag
		trainingSentence = trainingSentence + taggedWord + "\n"
	}

	return trainingSentence

}
