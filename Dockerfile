#FROM golang:latest
#
#COPY ./ $GOPATH/src/hunspellgo
#WORKDIR $GOPATH/src/hunspellgo
#
#RUN ls -la $GOPATH/src/hunspellgo/*
#
##ENV HTTP_PROXY http://127.0.0.1:8118
##ENV HTTPS_PROXY http://127.0.0.1:8118
#
#RUN git config --global http.postBuffer 1048576000
#
#RUN go get -v ./
#RUN cd ./binding && make
#RUN go build .
#
#EXPOSE 50052
#ENTRYPOINT ["./hunspellgo"]

FROM golang:1.12.5-alpine3.9
ADD hunspellgo /
EXPOSE 50052
CMD ["/hunspellgo"]