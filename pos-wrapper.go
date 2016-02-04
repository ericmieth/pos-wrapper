package main

import (
	"./rdrpostagger"
	"./treetagger"
	"./utils"
	"flag"
	//"fmt"
)

func main() {
	// evaluate command line flags
	taggerPtr := flag.String("tagger", "", "the tagger which should be used")
	dbNamePtr := flag.String("db-name", "", "name of database")
	// TODO: option to skip preparation of models

	flag.Parse()

	taggerName := *taggerPtr
	dbName := *dbNamePtr
	configMap := utils.ReturnConfig()

	if taggerName == "treetagger" {

		treetagger.PrepareModel(dbName, configMap)
		treetagger.CreateModel(dbName)
		//treetagger.EvaluateModel(dbName)

	} else if taggerName == "rdrpostagger" {
		rdrpostagger.PrepareModel(dbName, configMap)
		//TODO: this isn't implemented yet
		//rdrpostagger.CreateModel(dbName)
	}

}
