
代码：<https://github.com/robberphex/kitex-examples/tree/sae-demo>

## 构建server
```shell
docker build . -f discovery/server/Dockerfile -t server:0.0.1
```

## 构建client

```shell
docker build . -f discovery/client/Dockerfile -t client:0.0.1
```

## 部署说明

通过环境变量的方式来接入Nacos，环境变量说明如下(<https://help.aliyun.com/document_detail/427558.html>)：


|     Key     |           说明           | 是否必须 |         默认值          |
|-------------|------------------------|------|----------------------|
| serverAddr  | Nacos的地址（MSE Nacos的地址） | 必须   | 127.0.0.1            |
| Port        | Nacos的端口               | 可选   | 8848                 |
| NamespaceId | 命名空间ID                 | 可选   | ""（空字符串，即Public命名空间） |

   