## Larva
Larva是一个基于 gin + gorm + grpc 的开发框架。
>名字来源于PC游戏星际争霸中的幼虫（larva），虫族大部分单位的孵化都必须依靠幼虫来进行。

## 目的
做为一名golang的萌新，在边学习边实践过程中对常用模块进行了整理，用最简单平时的代码构建此项目。Larva是业余学习的作品，欢迎提出宝贵意见。

## 使用
默认内置一个demo接口用来取文章详情：http://127.0.0.1:9800/?id=1

使用rpc取详情数据的示例：http://127.0.0.1:9800/grpc?id=1

rpc服务端示例见api/grpc.go，服务端口9801

article表数据

<pre>
CREATE TABLE 'article' (
 'id' int(10) unsigned NOT NULL AUTO_INCREMENT,
 'title' varchar(30) NOT NULL DEFAULT '',
 PRIMARY KEY ('id')
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4


INSERT INTO 'article' ('id', 'title') VALUES
(1, '测试数据');
</pre>

## 目录结构
<details>
<summary>展开查看</summary>
<pre><code>.
├── app
│   ├── http.go         定义http协议接口
│   ├── grpc.go         定义grpc协议接口
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
│   │   ├── service.go  基于业务逻辑的数据处理
</code></pre>
</details>