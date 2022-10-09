package main

import (
	"log"
	"path/filepath"
	"plans-renamer/components"
	"regexp"
)

func main() {
	args, err := components.NewArguments()
	if err != nil {
		log.Fatalln(err)
	}

	files, err := components.GetFilesFromDirectory(args.SourcePath)
	if err != nil {
		log.Fatalln(err)
	}

	rules, err := components.NewRules("rules.json")
	if err != nil {
		log.Fatalln(err)
	}

	for _, filePath := range files {
		filename := filepath.Base(filePath)
		for searchRegex, replaceRegex := range rules {
			regex := regexp.MustCompile(searchRegex)
			if regex.MatchString(filename) {
				newFilename := regex.ReplaceAllString(filename, replaceRegex)
				err = components.CopyFile(filePath, filepath.Clean(args.DestinationPath+"/"+args.Prefix+"_"+newFilename))
				break
				//if err != nil {
				//	log.Fatalln(err.Error())
				//}
			}
		}
	}
}
