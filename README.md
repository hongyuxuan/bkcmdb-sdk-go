# 蓝鲸CMDB Golang SDK

## 接口列表
- 获取模型分组
- 创建模型
- 删除模型
- 获取模型属性分组
- 创建模型属性分组
- 删除模型属性分组
- 获取模型属性
- 创建模型属性
- 删除模型属性
- 获取模型关联关系
- 创建模型关联关系

## 用法
1.引入包
```go
import (
	bkcmdb "github.com/hongyuxuan/bkcmdb-sdk-go"
)
```
2.新建连接，并传入bkcmdb的用户参数、连接地址
```go
client := bkcmdb.NewClient(
  bkcmdb.WithBkUser("admin"),
  bkcmdb.WithSupplier("0"),
  bkcmdb.WithBaseUrl("http://bkcmdb_host:8080"))
```
3.示例：获取模型列表
```go
res, err := client.Classification().ListObject(context.Background())
if err != nil {
  log.Fatal(err)
}
fmt.Println(res)
```

更多示例详见test。