#!/bin/sh

if [ -z $1 ] || [ -z $2 ]; then
  echo "USAGE: ./make_book.sh <username> <filename.pdf / filename.epub>"
  exit 1
fi

if [ ! -f "archive_redditor" ]; then
  make
fi

if [ ! command -v pandoc &> /dev/null ]; then
  echo "Please install pandoc, as well as components for PDF support:"
  echo "  https://pandoc.org/installing.html"
  exit 1
fi

echo "Archiving u/$1 to $2"

tempdir=$(mktemp -d)
./archive_redditor -template templates/book_template.md "$1" "$tempdir"
pandoc $tempdir/*.txt -o "$2" --metadata title="u/$1" --toc --toc-depth=2
rm -rf "$tempdir"
