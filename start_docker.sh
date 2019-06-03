#!/bin/sh

protoc ./spellcheck/spellcheck.proto --go_out=plugins=grpc:.

docker build -t eudic-hunspell .
docker container stop hunspell-server
docker container rm hunspell-server
docker run -d -p 50052:50052 --name hunspell-server eudic-hunspell