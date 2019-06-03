#!/bin/sh
if [ "$1" == "-i" ]; then
protoc ./src/spellcheck/spellcheck.proto --go_out=plugins=grpc:.
fi

cd src
if [ "$1" == "-i" ]; then
glide install
fi
cd ..

docker build -t eudic-hunspell .
docker container stop hunspell-server
docker container rm hunspell-server
docker run -d -p 50051:50051 --name hunspell-server eudic-hunspell