#!/bin/bash

GOOS=linux GOARCH=amd64 go build -o dmpr . 
ls -lh dmpr