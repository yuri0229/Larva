## Larva
Larva是一个基于 gin + gorm + grpc 的微服务框架。
>名字来源于pc游戏星际争霸中的幼虫（larva），虫族大部分单位的孵化都必须依靠幼虫来进行。

## 介绍
做为一名golang的初学者和爱好者，在边学习边实践过程中对常用模块进行了整理，用最简单平时的代码构建此项目。Larva是业余学习的作品，欢迎提出宝贵意见。

demo分为服务模块(service)和接口模块(api)，api模块基于gin框架，用来处理客户端请求，通过grpc转发给服务模块。

dao默认集成了gorm用来访问mysql，可根据需求添加其它访问驱动。

两个模块内部结构一致，可以分成两个项目单独部署，这里为了方便演示我放到了一个项目中。

## 使用

- grpc服务模块启动方式：
<pre>
cd app/service/cmd
go run main.go
</pre>

- http接口模块启动方式：
<pre>
cd app/api/cmd
go run main.go
</pre>

http://127.0.0.1:9800/detail?id=1&timestamp=1619768437&sign=847430163260510a07e59e9ce288efac

- 测试表数据

<pre>
CREATE DATABASE demo DEFAULT CHARACTER SET utf8mb4;

CREATE TABLE 'article' (
 'id' int(10) unsigned NOT NULL AUTO_INCREMENT,
 'title' varchar(30) NOT NULL DEFAULT '',
 PRIMARY KEY ('id')
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


INSERT INTO 'article' ('id', 'title') VALUES
(1, '测试数据');
</pre>

## 目录结构
<details>
<summary>展开查看</summary>
<pre><code>.
├── http
│   ├── http.go         http接口
├── grpc
│   ├── client.go       grpc客户端
│   ├── grpc.go         grpc接口
│   ├── demo.proto      protobuf用例
├── cmd
│   ├── conf.toml       配置文件
│   ├── main.go         运行入口
├── internal
│   ├── conf
│   │   ├── conf.go     配置文件解析
│   ├── dao
│   │   ├── dao.go      提供mysql,redis连接
│   ├── model
│   │   ├── model.go    定义数据实体
│   ├── server
│   │   ├── grpc
│   │   │   ├── server.go   grpc服务初始化
│   │   ├── http
│   │   │   ├── server.go   http服务初始化
│   ├── service
│   │   ├── service.go  业务逻辑
</code></pre>
</details>

## 更新记录
2021.5.25
>- 新增日志基础库 pkg/log
>- 新增签名验证基础库 pkg/verify

2021.4.29
>- 在http和grpc接口新增定义了interface，业务逻辑(service)负责实现这个接口。
>- grpc客户端从http移动到了dao。
>- demo分成api和service两个模块。

2021.4.18
>- larav 开发版 gin + gorm + grpc