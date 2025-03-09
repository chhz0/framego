# fg-cli - 基于cobra封装后的命令行工具

fg-cli 是一个用以快速构建命令行应用程序的包，在设计上主要参考了`hugo`和`iam`的命令行实现，如果你希望获取更强大的支持，请直接使用`cobra`,
设计的目的仅为通过抽象`Commander`、`Flager`接口，实现接口的组合，启动一个命令行应用程序，并使其能够快速应用在服务框架中

## Features

- 应用命令行框架 √
- 命令行参数解析 √
- 配置文件解析 √

## Example
查看本项目的[example/fg-cli](https://github.com/chhz0/framego/tree/main/example/fg-cli)

## Optimization

1. flag绑定
2. 支持配置文件解析
3. options 使用

## Inspiration

- [cobra](https://github.com/spf13/cobra)
- [simplecobra](https://github.com/bep/simplecobra)
- [hugo](https://github.com/gohugoio/hugo)
- [iam/pkg/app](https://github.com/marmotedu/iam/tree/master/pkg/app)