#FROM golang:latest
#
#COPY ./ $GOPATH/src/hunspellgo
#WORKDIR $GOPATH/src/hunspellgo
#
#RUN ls -la $GOPATH/src/hunspellgo/*
#

#
#RUN git config --global http.postBuffer 1048576000
#
#RUN go get -v ./
#RUN cd ./binding && make
#RUN go build .
#
#EXPOSE 50052
#ENTRYPOINT ["./hunspellgo"]

FROM golang:1.12.5 AS build
RUN  mkdir -p /go/src \
  && mkdir -p /go/bin \
  && mkdir -p /go/pkg
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH
RUN mkdir -p $GOPATH/src/app
COPY ./src $GOPATH/src/app
WORKDIR $GOPATH/src/app
RUN ls -al
RUN go build -o myapp .


FROM alpine:latest
WORKDIR /myapp
COPY --from=builder /myapp .
CMD ["./myapp"]