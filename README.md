archive_redditor
================

A simple utility to archive a redditor's submissions to disk (for all of the really good NoSleep stories you don't want to become bookmarked 404s).

## Installation

To install from source, first, [Install Go](https://golang.org/doc/install), and then run:

    go get github.com/l1na-forever/archive_redditor

**Binary releases can be found on the [Releases page](https://github.com/l1na-forever/archive_redditor/releases/).**

## Usage

To archive a redditor's submissions, run the command:

    archive_redditor <username> <output directory path>

This will create the directory and populate it with text files of the following naming scheme:

    <year><month><date>_<post ID>_<username>_<post title slug>.txt

The contexts of the file will be the submission's selftext. Existing files won't be erased or overwritten.

## Status

Presently, the utility saves only selftext posts hosted on reddit.com. In the future, support may be added for archiving additional types of posts. Support may also be added for archiving off-site links (for example, archiving a user's images hosted off-site).

## Licence

Copyright © 2020 Lina

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.