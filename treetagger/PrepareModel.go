package treetagger

import (
	"../utils"
	"log"
	"os"
)

func PrepareModel(dbName string, configMap map[string]string) {
	// open database
	db := utils.DbOpen(dbName, configMap)
	defer db.Close()

	log.Printf("creating text corpora")
	// get all sentences
	rows, err := db.Query("SELECT sentence FROM sentences_tagged")
	//rows, err := db.Query("SELECT sentence FROM sentences_tagged LIMIT 100000")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// define tmp files
	outputPath := "output/treetagger/" + dbName + "/"
	os.MkdirAll(outputPath, 0760)
	corpusUntaggedFileName := outputPath + "corpusUntagged"
	corpusTrainingFileName := outputPath + "corpusTraining"
	mappingFileName := "tagsets/mapping_ud17_stts"
	lexiconFileName := outputPath + "lexicon"

	// remove file if existing
	os.Remove(corpusUntaggedFileName)
	os.Remove(corpusTrainingFileName)

	corpusUntaggedFile, err := os.OpenFile(corpusUntaggedFileName,
		os.O_CREATE|os.O_RDWR|os.O_APPEND,
		0660)
	if err != nil {
		log.Fatal(err)
	}
	defer corpusUntaggedFile.Close()

	corpusTrainingFile, err := os.OpenFile(corpusTrainingFileName,
		os.O_CREATE|os.O_RDWR|os.O_APPEND,
		0660)
	if err != nil {
		log.Fatal(err)
	}

	defer corpusTrainingFile.Close()

	var tagMap map[string]string = utils.CreateTagMap(mappingFileName)
	var sentenceFromDB string

	for rows.Next() {
		rows.Scan(&sentenceFromDB)
		// create an untagged corpus (for validation)
		corpusUntaggedFile.Write(CreatePlainSentence(sentenceFromDB))

		// create the input file
		corpusTrainingFile.Write(CreateTrainingCorpus(sentenceFromDB, tagMap))
	}

	// create a lexicon
	log.Printf("creating lexicon")
	CreateLexicon(db, lexiconFileName)

}
