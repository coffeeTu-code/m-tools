#!/usr/bin/env bash

protoc --version

# grpc
protoc -I . --go_out=paths=source_relative,plugins=grpc:./go/ *.proto

