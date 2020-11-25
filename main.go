package main

import (
	"flag"
	"fmt"
	"github.com/jzelinskie/geddit"
	"log"
	"os"
	"path"
)

const (
	userAgent        = "github.com/l1na-forever/archive_redditor 0.1"
	newDirectoryMode = 0755
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [optional flags] <username> <output path>\n\nOptional flags:\n", os.Args[0])
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		fmt.Printf("USAGE: %s <username> <output path>\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	err := archive(args[0], args[1])
	if err != nil {
		log.Fatal(err)
	}
}

func archive(username, baseOutputPath string) error {
	session := geddit.NewSession(userAgent)
	submissions, err := session.RedditorSubmissions(username, geddit.ListingOptions{})
	if err != nil {
		return err
	}

	err = os.MkdirAll(baseOutputPath, newDirectoryMode)
	if err != nil {
		return err
	}

	tmpl, err := SubmissionTemplate()
	if err != nil {
		return err
	}

	for _, submission := range submissions {
		path := path.Join(baseOutputPath, SubmissionFilename(submission))
		err = ArchiveSubmission(path, tmpl, submission)
		if err != nil {
			// Log the error, but continue archiving other submissions
			log.Println(err)
		}
	}

	return nil
}
