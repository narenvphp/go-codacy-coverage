#!/bin/bash

go install github.com/axw/gocov/gocov@latest
go install github.com/AlekSi/gocov-xml@latest

set -xe

mkdir -p -- "coverage"

go test -failfast -v ./tests/ --cover -coverprofile=./coverage/coverage.out

# Code Coverage + Convert to XML
echo "mode: set" > ./coverage/coverage.out && cat ./coverage/*.out | grep -v mode: | sort -r | \
awk '{if($1 != last) {print $0;last=$1}}' >> ./coverage/coverage.out
go tool cover -func=./coverage/coverage.out
gocov convert ./coverage/coverage.out | gocov-xml > ./coverage/coverage.xml
