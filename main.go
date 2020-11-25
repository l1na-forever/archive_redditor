package main

import (
	"errors"
	"fmt"
	"github.com/jzelinskie/geddit"
	"github.com/stvp/slug"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"
)

const (
	userAgent                = "github.com/l1na-forever/archive_redditor 0.1"
	newDirectoryMode         = 0755
	newFileMode              = 0644
	submissionDateLayout     = "20060102"
	submissionTitleSeparator = '-'
)

func init() {
	slug.Replacement = submissionTitleSeparator
}

func main() {
	if len(os.Args) < 3 {
		log.Fatal("USAGE: " + os.Args[0] + " <username> <output path>")
	}

	err := archive(os.Args[1], os.Args[2])
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

	for _, submission := range submissions {
		err = archiveSubmission(baseOutputPath, submission)
		if err != nil {
			// Log the error, but continue archiving other submissions
			log.Println(err)
		}
	}

	return nil
}

func archiveSubmission(baseOutputPath string, submission *geddit.Submission) error {
	// At the moment, the tool only supports archiving self text
	if !submission.IsSelf {
		return nil
	}

	outputPath := path.Join(baseOutputPath, submissionFilename(submission))

	// Don't overwrite existing posts
	if _, err := os.Stat(outputPath); err != nil {
		if os.IsExist(err) {
			return nil
		}
	}

	err := ioutil.WriteFile(outputPath, []byte(submission.Selftext), newFileMode)
	if err != nil {
		return errors.New(fmt.Sprintf("Error saving %s: %s", outputPath, err.Error()))
	}

	return nil
}

func submissionFilename(submission *geddit.Submission) string {
	dateCreated := time.Unix(int64(submission.DateCreated), 0)
	dateCreatedFormatted := dateCreated.Format(submissionDateLayout)
	return fmt.Sprintf("%s_%s_%s_%s.txt", dateCreatedFormatted, submission.ID, slug.Clean(submission.Author), slug.Clean(submission.Title))
}
