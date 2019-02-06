# Khadijah

[![Build Statue](https://travis-ci.org//chengyumeng/khadijah.svg?branch=master)](https://travis-ci.org/chengyumeng/khadijah)
[![Build Release](https://img.shields.io/github/release/chengyumeng/Khadijah.svg)](https://github.com/chengyumeng/khadijah/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/chengyumeng/khadijah)](https://goreportcard.com/report/github.com/chengyumeng/khadijah)

Khadijah 是一款基于奇虎 360 开源的 Wayne 的 kubernetes 命令行管理工具。

### 为什么使用 Khadijah ?


## Feature

- 支持了多种 Wayne OpenAPI 的接口调用
- 支持了跨部门、跨项目的资源信息查询
- 支持了 Kubernetes 原生信息的查询
- 支持了基于 Wayne 的权限管理下的从命令行进入 Kubernetes Pod 容器的远程 Shell

## 快速安装

```bash
go get -u github.com/chengyumeng/khadijah
```

```bash
khadijah config set --apikey=example --websocketurl=ws://127.0.0.1:8080 --baseurl=http://127.0.0.1:4200
# 必须指定 baseurl，该值为 Wayne 服务的 http(s)地址
# websocketurl 是用于远程 Shell 的 Wayne 地址，可以按需指定
# apikey 为需要调用 Wayne OpenAPI 需要指定的值（khadijah query 下所有操作依赖于这个值）
```

```bash
khadijah login -uadmin -padmin
```

```bash
khadijah get namespace
Name: admin Email:admin@360.cn

+----+--------------+-----------------+-------------------------------+-------------------------------+
| ID |     NAME     |      USER       |          CREATETIME           |          UPDATETIME           |
+----+--------------+-----------------+-------------------------------+-------------------------------+
| 27 | argo         | admin           | 2018-12-26 11:01:08 +0800 CST | 2018-12-26 11:01:08 +0800 CST |
+----+--------------+-----------------+-------------------------------+-------------------------------+
```

```bash
khadijah get deployment -n=infra
+-----+----------------------------------------+------------+-----+-----------+-------------+-------------------------------+
| ID  |                  NAME                  |    TYPE    | APP | NAMESPACE |    USER     |          CREATETIME           |
+-----+----------------------------------------+------------+-----+-----------+-------------+-------------------------------+
|   1 | infra-wayne                            | deployment |     |           | admin       | 2018-05-30 16:59:59 +0800 CST |
|   2 | infra-test                             | deployment |     |           | admin       | 2018-05-30 17:20:22 +0800 CST |
|   3 | infra-cpu                              | deployment |     |           | admin       | 2018-05-31 10:36:21 +0800 CST |
+-----+----------------------------------------+------------+-----+-----------+-------------+-------------------------------+
```

```bash
khadijah describe deployment move-num1 -n=infra -o=pretty
+-----------+-----------+---------+--------------------------------+------------------------------+----------+--------------------------------+---------------------------+
|   NAME    | NAMESPACE | CLUSTER |             LABELS             |            CONTAINERS        | REPLICAS |            MESSAGE             |           PODS            |
+-----------+-----------+---------+--------------------------------+------------------------------+----------+--------------------------------+---------------------------+
| move-num1 | infra     | K8S-1   | app:move-num1 qihoo-app:move   | php:docker.hub/php:1.1.0     | 1/1      | 2018-12-10 19:03:00 +0800      | move-num1-945d9577f-wnvtm |
|           |           |         | qihoo-ns:infra                 | nginx:docker.hub/nginx:1.1.0 |          | CST:Deployment has minimum     |                           |
|           |           |         |                                | qconf:docker.hub/agent:0.1.0 |          | availability. 2018-12-11       |                           |
|           |           |         |                                |                              |          | 14:52:19 +0800 CST:ReplicaSet  |                           |
|           |           |         |                                |                              |          | "move-num1-945d9577f" has      |                           |
|           |           |         |                                |                              |          | successfully progressed.       |                           |
+-----------+-----------+---------+--------------------------------+------------------------------+----------+--------------------------------+---------------------------+
```

```bash
khadijah query getpodinfo -c=K8S-1 --label=app=move-num1
```
```json
{
  "code": 200,
  "pods": [
    {
      "podIp": "172.17.4.44",
      "labels": {
        "app": "move-num1",
        "pod-template-hash": "501851339",
        "qihoo-app": "move",
        "qihoo-ns": "infra"
      }
    }
  ]
}
```