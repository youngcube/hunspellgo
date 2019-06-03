## HunspellGo Binding

#### macOS 环境搭建

- 官网运行安装最新版 `Go`，注意不要使用 `brew` 安装，否则会有许多奇怪问题 https://golang.org/dl/

- 将 \$GOPATH 设置为 `~/go`

- 本项目严重依赖翻墙工具，否则基本不可编译。因此请务必先开启必要的命令行翻墙工具，使命令行能够顺利翻墙。可参考以下命令：

```
git config http.proxy 'http://127.0.0.1:8118'
git config https.proxy 'http://127.0.0.1:8118'
```

> 确定命令行内已是翻墙状态后，运行本项目内的 `dependence.sh`。内置一些必要的插件。该命令行可重复执行，如果翻墙工具不稳定可多执行几次
> 由于 `VSCode` 不方便翻墙，因此独立写一份脚本。如果使用 `GoLand` 开发，也需要运行此脚本
> 如果发现执行后出现 `There is no tracking information for the current branch.`，请记下该包，然后进入 `~/go/src` 内删除这个包文件夹，然后重新执行一次 `dependence.sh`
> 等待成功后，再打开 `VSCode`/`GoLand`

- 本项目包含一些`VSCode`推荐插件，同时`VSCode`本身也会推荐`Go`插件，请全部安装，以便使用`gRPC`内的 `proto`高亮；但是，更推荐使用 `GoLand` 开发，体验相比 `VSCode` 更好

- 如果需要直接编译运行，请运行 `build.sh`。该命令行内包含了必要的 `Make` 命令

- 如果需要部署线上 `Docker`，请运行 `start_docker.sh`。如果不带参数，默认不执行 `protoc` 和 `glide install`：即不会重新安装依赖，也不会重新生成`gRPC`需要的接口文件`spellcheck.pb.go`。如果带参数 `-i`，那么就会安装依赖


#### Benchmark

- 本项目基于 `C++` 版 `Hunspell` 构建，实测性能相比`C++`版本，有大约 `10~20` 微秒的性能损耗

- `gRPC` 在直接发送消息的测试中（`Server Go` => `Client Go`），大约需要 `300~500` 微秒的时间

- `C#`版 `gRPC` 在直接发送消息的测试中（`Server Go` => `Client C#`），大约需要 `1200~2000` 微秒的时间

- 注意首次调用时耗时较长，可能会需要 `5000~15000` 微秒左右，但第二次开始就会下降到正常水平