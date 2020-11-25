LDFLAGS = "-X main.Buildstamp=`date -u '+%Y-%m-%d_%I:%M:%S%p'` -X main.Githash=`git rev-parse HEAD`"
SRC := $(shell find ./ -type f -regex ".*\.go")
RELEASE_DIR = $(shell echo "releases/release_`git rev-parse HEAD`")

all: archive_redditor

release: target-darwin-amd64 target-linux-amd64 target-windows-amd64

test:
	go test -v .

archive_redditor:
	go build -v .

install:
	go install -v .

target-darwin-amd64: $(SRC)
	mkdir -p $(RELEASE_DIR)
	GOOS=darwin GOARCH=amd64 go build -o archive_redditor -ldflags $(LDFLAGS) .
	zip $(RELEASE_DIR)/archive_redditor_darwin_amd64.zip archive_redditor
	rm -f archive_redditor

target-linux-amd64: $(SRC)
	mkdir -p $(RELEASE_DIR)
	GOOS=linux GOARCH=amd64 go build -o archive_redditor -ldflags $(LDFLAGS) .
	zip $(RELEASE_DIR)/archive_redditor_linux_amd64.zip archive_redditor
	rm -f archive_redditor

target-windows-amd64: $(SRC)
	mkdir -p $(RELEASE_DIR)
	GOOS=windows GOARCH=amd64 go build -o archive_redditor.exe -ldflags $(LDFLAGS) .
	zip $(RELEASE_DIR)/archive_redditor_windows_amd64.zip archive_redditor.exe
	rm -f archive_redditor.exe

clean:
	rm -rf archive_redditor archive_redditor.exe $(RELEASE_DIR)

.PHONY: $(release) clean install test
