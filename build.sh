#!/bin/sh
protoc ./src/spellcheck/spellcheck.proto --go_out=plugins=grpc:.

cd src/binding
rm -rf lib
rm -rf obj
make

cd ..
glide install
