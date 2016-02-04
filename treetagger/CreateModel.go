package treetagger

import (
	"log"
	"os"
	"syscall"
)

// execute exteral tool "train-tree-tagger" to create a model
func CreateModel(dbName string) {

	log.Printf("creating model")

	outputPath := "output/treetagger/" + dbName + "/"
	lexiconFileName := outputPath + "lexicon"
	tagsetFileName := "tagsets/ud17"
	corpusTrainingFileName := outputPath + "corpusTraining"

	modelFileName := outputPath + "model"
	treetaggerTrainingBinaryPath := "/opt/pos/tools/TreeTagger/bin/train-tree-tagger"

	execArgsTraining := []string{"-st PUNCT", lexiconFileName, tagsetFileName, corpusTrainingFileName, modelFileName}

	env := os.Environ()
	log.Printf("triggering ./train-tree-tagger, waiting for command to finish...")
	err := syscall.Exec(treetaggerTrainingBinaryPath, execArgsTraining, env)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Command finished with error: %v", err)

}
