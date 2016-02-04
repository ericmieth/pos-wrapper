package utils

import (
	"encoding/csv"
	//"fmt"
	"log"
	"os"
	"strings"
)

func CreateTagMap(mappingFileName string) map[string]string {
	mappingFile, err := os.OpenFile(mappingFileName, os.O_RDONLY, 0660)
	if err != nil {
		log.Fatal(err)
	}
	defer mappingFile.Close()

	r := csv.NewReader(mappingFile)
	r.FieldsPerRecord = 2
	r.Comma = '\t'

	lines, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	tagMap := make(map[string]string)

	for _, line := range lines {
		tagMap[line[0]] = strings.TrimSpace(line[1])
	}

	return tagMap
}
