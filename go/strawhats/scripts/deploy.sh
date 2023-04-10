#!/bin/bash

mkdir dist
GOOS=linux GOARCH=amd64 go build -o ./dist/strawhats ./server/main/main.go
cp -r ./static ./dist/static
scp -r ./dist ubuntu@strawhat.chrisclegg.com:~/current
ssh ubuntu@strawhat.chrisclegg.com