#!/bin/bash

go clean -cache ./...
go build -o twitter-bot -v ./...