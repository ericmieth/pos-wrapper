package utils

import (
	"encoding/json"
	"log"
	"os"
)

type config struct {
	DbUser                    string
	DbPass                    string
	DbHost                    string
	DbPort                    string
	TrainTreeTaggerBinaryPath string
}

func ReturnConfig() map[string]string {

	file, err := os.Open("config")
	decoder := json.NewDecoder(file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	configuration := config{}
	err = decoder.Decode(&configuration)
	if err != nil {
		log.Fatal(err)
	}

	configMap := make(map[string]string)

	configMap["dbUser"] = configuration.DbUser
	configMap["dbPass"] = configuration.DbPass
	configMap["dbHost"] = configuration.DbHost
	configMap["dbPort"] = configuration.DbPort
	configMap["TrainTreeTaggerBinaryPath"] = configuration.TrainTreeTaggerBinaryPath

	return configMap
}
