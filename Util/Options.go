package util

import (
	"encoding/json"
	"log"
	"os"
)

var Options ProgramOptions

type ProgramOptions struct {
	DatabaseURL  string `json:"databaseURL"`
	DatabaseName string `json:"databaseName"`
}

func init() {
	data, err := os.ReadFile("options.json")
	if err != nil {
		log.Println("Failed to load options:", err.Error())
	}
	log.Println("Options OK")

	json.Unmarshal(data, &Options)
}
