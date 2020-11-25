package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
	"time"

	"github.com/jzelinskie/geddit"
	"github.com/stvp/slug"
)

var (
	flagTemplatePath = flag.String("template", "",
		`(optional) path to template file
	Templates can be used to customize the format in which posts are archived.
	For example, the following template file could be used to generate a
	Markdown entry for each selftext:

	  # {{.Title}}
	  *Posted by {{.Author}} on {{.Date}} in r/{{.Subreddit}}*

	  {{.SelftextHTML}}

	A full list of available flags can be found here:
`)

	defaultTemplate = `{{.Title}}
Posted by {{.Author}} in r/{{.Subreddit}} on {{.Date}}

{{.Selftext}}
`
)

const (
	templateDateLayout     = "2006/01/02"
	newFileMode            = 0644
	filenameDateLayout     = "20060102"
	filenameTitleSeparator = '-'
)

func init() {
	slug.Replacement = filenameTitleSeparator
}

type SubmissionTemplateData struct {
	Date string
	*geddit.Submission
}

// Returns a record of the submission's data for use within the template's
// execution context.
func NewSubmissionTemplateData(submission *geddit.Submission) *SubmissionTemplateData {
	return &SubmissionTemplateData{
		Date:       formatSubmissionDate(submission, templateDateLayout),
		Submission: submission,
	}
}

// Returns the template to be used when archiving self posts. If a temmplate
// path argument was passed to the program, the path given is parsed and used
// as the template. Otherwise, a default template that outputs a basic header
// the post's contents is used.
func SubmissionTemplate() (*template.Template, error) {
	var templateContents string
	if *flagTemplatePath != "" {
		contents, err := ioutil.ReadFile(*flagTemplatePath)
		if err != nil {
			return nil, err
		}

		templateContents = string(contents)
	} else {
		templateContents = defaultTemplate
	}

	return template.New("").Parse(templateContents)
}

func ArchiveSubmission(path string, tmpl *template.Template, submission *geddit.Submission) error {
	// At the moment, the tool only supports archiving self text
	if !submission.IsSelf {
		return nil
	}

	// Don't overwrite existing posts
	if _, err := os.Stat(path); err != nil {
		if os.IsExist(err) {
			return nil
		}
	}

	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, newFileMode)
	if err != nil {
		return errors.New(fmt.Sprintf("Error creating %s: %s", path, err.Error()))
	}
	defer file.Close()

	return tmpl.Execute(file, NewSubmissionTemplateData(submission))
}

// Returns a filename uniquely and consistently identifying the given
// submission. Filenames are prefixed with the date, for easy chronological
// sorting.
func SubmissionFilename(submission *geddit.Submission) string {
	return fmt.Sprintf("%s_%s_%s_%s.txt",
		formatSubmissionDate(submission, filenameDateLayout),
		submission.ID,
		slug.Clean(submission.Author),
		slug.Clean(submission.Title))
}

// Returns a formatted string of a submission's creation date, formatted
// according to stdlib's "time" formatting scheme.
func formatSubmissionDate(submission *geddit.Submission, layout string) string {
	dateCreated := time.Unix(int64(submission.DateCreated), 0)
	return dateCreated.Format(layout)
}
