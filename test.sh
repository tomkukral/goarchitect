#!/bin/bash

go get
go get -u github.com/golang/dep/cmd/dep

dep ensure

go test -cover -race ./...
