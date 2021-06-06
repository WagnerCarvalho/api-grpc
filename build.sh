#!/bin/bash

go clean --cache && go test -v -cover api-grpc/...
go build -o authentication/authsvc-local authentication/main.go
go build -o api/apisvc-local api/main.go

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o authentication/authsvc authentication/main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api/apisvc api/main.go