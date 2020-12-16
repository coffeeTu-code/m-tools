COMMIT_HASH=$(shell git rev-parse --verify HEAD | cut -c 1-8)
BUILD_DATE=$(shell date +%Y-%m-%d_%H:%M:%S%z)
GIT_TAG=$(shell git describe --tags)
GIT_AUTHOR=$(shell git show -s --format=%an)
SHELL:=/bin/bash

exec=go build -ldflags "-X main.GitTag=$(GIT_TAG) -X main.BuildTime=$(BUILD_DATE) -X main.GitCommit=$(COMMIT_HASH) -X main.GitAuthor=$(GIT_AUTHOR)" -o bin/$(1)/$(1) cmd/$(1)/main.go

all: build

build: mod helloworld spider greeter

mod:
	go mod download && go mod tidy

spider:
	$(call exec,m-spider)

helloworld:
	$(call exec,hello-world)

greeter:
	$(call exec,greeter-sample)

.PHONY : clean

clean:
	-rm -rf bin/*