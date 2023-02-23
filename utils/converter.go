package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type ConvertConfig struct {
	ConvertTo string
	Content   string
	OutputDir string
}

func ConvertFile(cfg ConvertConfig) {
	switch cfg.ConvertTo {
	case "plaintext":
		ConvertToPlainText(cfg)
	case "json":
		ConvertToJson(cfg)
	}
}

func ConvertToJson(cfg ConvertConfig) {
	result := make([]map[string]string, 0)

	lines := strings.Split(cfg.Content, "\n")
	for _, line := range lines {
		row := map[string]string{}
		for idx, word := range strings.Split(line, " ") {
			key := fmt.Sprintf("Column %v", idx)
			row[key] = word
		}
		result = append(result, row)
	}

	bytes, err := json.Marshal(result)
	if err != nil {
		log.Fatalf("Error marshal to json | err : %v", err)
	}

	err = os.WriteFile(cfg.OutputDir, bytes, os.ModePerm)
	if err != nil {
		log.Fatalf("Error write file | err : %v", err.Error())
	}
}

func ConvertToPlainText(cfg ConvertConfig) {
	err := os.WriteFile(cfg.OutputDir, []byte(cfg.Content), os.ModePerm)
	if err != nil {
		log.Fatalf("Error write file | err : %v", err.Error())
	}
}
