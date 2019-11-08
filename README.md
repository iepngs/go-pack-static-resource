# go-pack-static-resource
go打包静态资源和html示例

#### 打包静态html文件所需包：

```shell
$ go get github.com/jessevdk/go-assets-builder
```

#### 打包所有html文件为一个go文件
```shell
# linux系统终端下：
$ go-assets-builder templates -o integrate/assets.go -p demo -s /templates/

# windows的powershell下：
go-assets-builder templates -o="integrate/assets.go" -p="demo" -s="/templates/"
```
> 注： 需要把GOPATH下的bin文件夹（可以执行go env查看具体位置）加入到环境中不然无法识别`go-assets-builder`等命令；

#### 打包静态文件所需包：

```shell
$ go get github.com/jteeuwen/go-bindata/...
$ go get github.com/elazarl/go-bindata-assetfs/...
```

#### 打包静态文件为go文件
```shell
$ go-bindata -o="integrate/bindata.go" -pkg="demo" resources/... config/...
```

#### 相关文档：

https://github.com/gin-gonic/examples/tree/master/assets-in-binary


#### 静态资源打包也可以用statik:

https://github.com/rakyll/statik
