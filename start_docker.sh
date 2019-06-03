#!/bin/sh

protoc ./src/spellcheck/spellcheck.proto --go_out=plugins=grpc:.

cd src/binding
rm -rf lib
rm -rf obj
cd ../..

docker build -t eudic-hunspell .
docker container stop hunspell-server
docker container rm hunspell-server
docker run -d -p 50052:50052 --name hunspell-server eudic-hunspell