# 分阶段编译，先编译出 Linux 下的 Go 包，然后再放入 alpine。这样可以使包更小
# 不使用交叉编译，会有动态编译失败的问题
FROM golang:1.12.5 AS build
RUN  mkdir -p /go/src \
  && mkdir -p /go/bin \
  && mkdir -p /go/pkg

# 先统一将需要编译的文件全部引入
COPY ./src /app/backend

# 排除开发机下编译的C++产物
RUN rm -rf /app/backend/binding/lib && rm -rf /app/backend/binding/obj

# 排除资源文件，后续会将资源文件放入部署环境
RUN rm -rf /app/backend/binding/include/dictionaries
ENV GOPATH=/go

# 由于 Go 可能读不出 vendor，因此将 vendor 内生成的文件放入默认 GOPATH
COPY ./src/vendor /go/src
WORKDIR /app/backend

# 编译 C++ Hunspell
RUN cd binding && make
# 静态编译，否则无法在 alpine 内运行
RUN go build -tags release --ldflags '-extldflags -static' -a -o /main .

# 部署环境
FROM alpine:latest
WORKDIR /main
COPY --from=build /main .
# 将 hunspell 资源文件引入
COPY ./src/binding/include/dictionaries /main/dictionaries
EXPOSE 50051
CMD ["./main"]