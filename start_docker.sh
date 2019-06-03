#!/bin/sh
protoc ./src/spellcheck/spellcheck.proto --go_out=plugins=grpc:.
glide install

# 这里需要移除编译好的文件，因为最终是需要部署在 Linux 里的
cd src/binding
rm -rf lib
rm -rf obj
cd ..
glide install
cd ..

docker build -t eudic-hunspell .
docker container stop hunspell-server
docker container rm hunspell-server
docker run -d -p 50052:50052 --name hunspell-server eudic-hunspell