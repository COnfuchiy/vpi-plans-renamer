package components

import (
	"errors"
	"flag"
	"os"
)

type Arguments struct {
	SourcePath      string
	DestinationPath string
	Prefix          string
}

func NewArguments() (Arguments, error) {
	args := Arguments{}
	flag.StringVar(&args.SourcePath, "src", "", "Path to documents root directory")
	flag.StringVar(&args.DestinationPath, "dest", "", "Path to documents destination directory")
	flag.StringVar(&args.Prefix, "prefix", "", "Document prefix")
	flag.Parse()
	err := args.checkField()
	if err != nil {
		return args, err
	}
	err = args.checkPaths()
	if err != nil {
		return args, err
	}
	return args, nil
}

func (args *Arguments) checkField() error {
	if args.SourcePath == "" {
		return errors.New("option src is required")
	}
	if args.Prefix == "" {
		return errors.New("option prefix is required")
	}
	if args.DestinationPath == "" {
		args.DestinationPath = args.SourcePath
	}
	return nil
}

func (args *Arguments) checkPaths() error {
	sourcePathFileInfo, err := os.Stat(args.SourcePath)
	if err != nil {
		return err
	}
	if !sourcePathFileInfo.IsDir() {
		return errors.New(args.SourcePath + " is not a directory")
	}

	destinationPathFileInfo, err := os.Stat(args.DestinationPath)
	if err != nil {
		return err
	}
	if !destinationPathFileInfo.IsDir() {
		return errors.New(args.DestinationPath + " is not a directory")
	}

	return nil
}
