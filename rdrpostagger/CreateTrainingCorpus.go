package rdrpostagger

import (
	"../utils"
	"strings"
)

// stripping of annotations from a string
func CreateTrainingCorpus(inputSentence string, tagMap map[string]string) []byte {

	// do the general parsing
	trainingSentence := utils.CreateGeneralTrainingCorpus(inputSentence, tagMap)

	// rdrpostagger doesn't like a single slash, so we have to substitute
	trainingSentence = strings.Replace(trainingSentence, "/|PUNCT", "//|PUNCT", -1)
	returnSentence := strings.TrimSpace(trainingSentence) + "\n"
	return []byte(returnSentence)
}
