#!/bin/bash

# documentation
# https://github.com/gin-gonic/gin/blob/master/docs/doc.md#build-with-json-replacement

#go run -tags=sonic main.go
#go build -tags=go_json -gcflags=-m main.go
go run -tags=go_json main.go

