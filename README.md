# GoginL(isegopher)
https://github.com/IseEkko

首先感谢Singo
此web框架是在Singo的基础上进行改进的一个框架
## 登录
在这个框架中，登录使用的是jwt认证。
在里面我们可以很轻松的改变登录产生token的逻辑。和token包含的东西

## 不足的地方
在里面对返回数据格式的封装不是很友好，后期我会对这个框架进行改进

## 后续方面
 1. 返回响应的封装
 2. 对上传的增加，集成 在本地上传文件和第三方的对象存储
 3. 邮箱认证的集成
 。。。。。。。
 后续我将持续的更新谢谢大家的支持
 本人联系方式：794776085
## Go Mod

本项目使用[Go Mod](https://github.com/golang/go/wiki/Modules)管理依赖。

```shell
go mod init go-crud
export GOPROXY=http://mirrors.aliyun.com/goproxy/
go run main.go // 自动安装
```

## 运行

```shell
go run main.go
```

项目运行后启动在3000端口（可以修改，参考gin文档)
## Singo文档
