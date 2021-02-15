#!/bin/bash

go clean --cache && go test -v -cover microservices/authentication/...
go build -o authentication/authsvc authentication/main.go