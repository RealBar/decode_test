#!/usr/bin/env zsh

gofmt -w $(find . -name "*.go" | grep -v vendor/)
goimports -w $(find . -name "*.go" | grep -v vendor/)
