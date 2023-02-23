package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/mrizalr/mysooltan_test/utils"
)

func WriteHelper() {
	fmt.Println(`
LOG CONVERTER

usage : logconverter [<args>] [-t] [-o] [-h]

<args>	The directory of input (.log) files (required)
-t	Set type for convert result (json/plaintext) | default: plaintext
-o	The directory of convert result (required)
-h	Help

logconverter /lib/vars/error.log -t json -o /convert/logs/error.json
	`)
}

func main() {
	if len(os.Args) == 1 || os.Args[1] == "-h" {
		WriteHelper()
		return
	}

	logFileDir := os.Args[1]

	err := utils.CheckDirExists(logFileDir)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("%s isn't exists", logFileDir)
		}
		log.Fatalf("Error checking input directory | err : %v", err.Error())
	}

	convertFormat := "plaintext"
	outputDir := ""

	for i := 2; i < len(os.Args); i++ {
		if os.Args[i] == "-t" {
			convertFormat = strings.ToLower(os.Args[i+1])
			i++
			continue
		}

		if os.Args[i] == "-o" {
			outputDir = strings.ToLower(os.Args[i+1])
			i++
			continue
		}
	}

	if outputDir == "" {
		log.Fatalf("Error output dir is empty")
	}

	dir, _ := filepath.Split(outputDir)
	err = utils.CheckDirExists(dir)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.Mkdir(dir, os.ModePerm); err != nil {
				log.Fatalf("Error when make directory | err : %v", err.Error())
			}
		} else {
			log.Fatalf("Error checking output directory | err : %v", err.Error())
		}
	}

	content := utils.ReadLogFile(logFileDir)

	convertCfg := utils.ConvertConfig{
		ConvertTo: convertFormat,
		Content:   content,
		OutputDir: outputDir,
	}

	utils.ConvertFile(convertCfg)
}
