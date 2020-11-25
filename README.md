archive_redditor
================

A simple utility to archive a redditor's submissions to disk (for all of the really good NoSleep stories you don't want to become bookmarked 404s).

## Installation

To install from source, first, [Install Go](https://golang.org/doc/install), and then run:

    go get github.com/l1na-forever/archive_redditor

**Binary releases can be found on the [Releases page](https://github.com/l1na-forever/archive_redditor/releases/).**

## Usage

To get help, run the command:

    archive_redditor -help

To archive a redditor's submissions, run the command:

    archive_redditor <username> <output directory path>

This will create the directory and populate it with text files of the following naming scheme:

    <year><month><date>_<post ID>_<username>_<post title slug>.txt

The contexts of the file will be the submission's selftext. Existing files won't be erased or overwritten.

**Using templates:**

By default, posts will be written to a plain text file consisting of a simple header, and the post's contents. This output format can be customized by passing in a `-template` argument:

    archive_redditor -template templates/book_template.md <username> <output directory path>

A template file might look like this:

    # {{.Title}}

    **{{.Author}}** - {{.Date}}

    {{.Selftext}}

    \newpage


Notably, each field is accessed with the syntax `{{.FieldName}}`. Other text is from the templating engine, and can be used to set directives (e.g., for TeX/Pandoc). See [the `templates` directory](https://github.com/l1na-forever/archive_redditor/tree/mainline/templates/) in this repository for examples. Templates are rendered using [Go's built-in templating facility](https://golang.org/pkg/text/template/); no sanitization is performed on the template (you should not run templates you don't trust). Available fields include:

* `Author`
* `Date`
* `NumComments`
* `Permalink`
* `Score`, `Ups`, `Downs`
* `Selftext`
* `SelftextHTML`
* `Subreddit`
* `Title`

The full list of available fields can be found [here](https://godoc.org/github.com/jzelinskie/geddit#Submission).

**Generating PDFs/E-books:**

The script `make_book.sh`, included in this repository, can be used to generate a single document. The document hands the user's posts to [Pandoc](https://pandoc.org/), which renders to one of its supported file formats (based on the output filename). To run the included script:

    ./make_book.sh <username> <filename>

To generate an E-book suitable for an E-reader, the ePub (`.epub`) format is usually the best bet. For Kindle devices, [Calibre](https://calibre-ebook.com/) includes a handy `ebook-convert` utility which can convert to the Kindle-compatible `.mobi` format (`pandoc` doesn't support outputting `.mobi` on its own):

    ebook-convert output.epub output.mobi

## Status

Presently, the utility saves only selftext posts hosted on reddit.com. Customization of the output file format is supported. In the future, support may be added for archiving additional types of posts. Support may also be added for archiving off-site links (for example, archiving a user's images hosted off-site).

## Licence

Copyright © 2020 Lina

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.