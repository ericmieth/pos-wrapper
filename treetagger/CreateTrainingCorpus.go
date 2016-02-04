package treetagger

import (
	"../utils"
	"strings"
)

// stripping of annotations from a string
func CreateTrainingCorpus(inputSentence string, tagMap map[string]string) []byte {

	// do the general parsing
	trainingSentence := utils.CreateGeneralTrainingCorpus(inputSentence, tagMap)

	singleRows := strings.Split(trainingSentence, "\n")
	var validRows []byte
	for i := range singleRows {
		appendRow := strings.Replace(string(singleRows[i]), "|", "\t", -1) + "\n"
		validRows = append(validRows, []byte(appendRow)...)
	}

	returnValue := string(validRows)
	return []byte(returnValue)
}
