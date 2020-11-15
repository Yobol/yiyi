# yiyi

A web site that integrates various utilities using golang toolkits.

## 技术选型

- Go v1.15.3
- beego v1.12.0
  - bee v1.12.0
    - `go get github.com/beego/bee@v1.12.0`
  - 使用 `bee new yiyi` 创建 beego 项目，其版本与 bee 版本保持一致；
    - 可以在 go.mod 中更改 beego 版本，然后使用 `go mod install` 安装依赖；
  - 使用 `bee run` 命令启动 beego 项目。

