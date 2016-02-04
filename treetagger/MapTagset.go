package treetagger

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
)

func MapTagsets(corpusTrainingFileName string, tagMap map[string]string) {

	corpusTrainingFile, err := os.OpenFile(corpusTrainingFileName, os.O_RDONLY, 0660)
	if err != nil {
		log.Fatal(err)
	}
	defer corpusTrainingFile.Close()

	mappingFileTmpName := corpusTrainingFileName + ".tmp"
	os.Remove(mappingFileTmpName)
	mappingFileTmp, err := os.OpenFile(mappingFileTmpName,
		os.O_CREATE|os.O_RDWR|os.O_APPEND,
		0660)
	if err != nil {
		log.Fatal(err)
	}
	defer mappingFileTmp.Close()

	r := csv.NewReader(corpusTrainingFile)
	r.FieldsPerRecord = 2
	r.Comma = '\t'

	lines, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var preparedRow string

	for _, line := range lines {
		line[1] = tagMap[strings.TrimSpace(line[1])]

		_, err = mappingFileTmp.Write([]byte(preparedRow))
		if err != nil {
			log.Fatal(err)
		}

	}

	os.Rename(mappingFileTmpName, corpusTrainingFileName)
}
