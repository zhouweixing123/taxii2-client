# taxii2

#### 介绍
taxii2是针对TAXII2.X服务器的最小化客户端实现。它支持一下TAXII2.XApi服务:


#### 软件架构
软件架构说明
1. Get API Root Information
2. Get Collections

#### 安装教程
```shell
go get -u gitee.com/zhouweixing/taxii2.git
```

#### 使用说明
TAXII 客户端旨在作为 Go 库使用。目前暂无命令行客户端。
#### taxii2-client 提供了三个类：
1.  Server(服务器)
2.  ApiRoot(Api根)
3.  Collection(集合)

#### 示例
```go
package main

import (
	"encoding/json"
	"fmt"
	"gitee.com/zhouweixing/taxii2.git/server/apiRootsRes"
	"gitee.com/zhouweixing/taxii2.git/server/collections"
	"gitee.com/zhouweixing/taxii2.git/server/getCollectionsData"
	"gitee.com/zhouweixing/taxii2.git/server/httpClient"
	"time"
)

func main() {
	// 服务器接口: http://ip:端口/taxii2/
	conn := httpClient.NewConnectionSession(`用户名`, `密码`, `服务器接口`)
	apiRoots := apiRootsRes.NewApiRoots(conn, map[string]string{
		`User-Agent`: `taxii2-client/2.3.0`,
		`Accept`:     `application/taxii+json;version=2.1`,
	})
	// 获取apiRootId
	apiRootId, err := apiRoots.GetApiRoot()
	if err != nil {
		fmt.Println(err)
		return
	}
	// 获取集合ID
	collection := collections.NewGetCollectionsObject(conn, apiRootId, map[string]string{
		`User-Agent`: `taxii2-client/2.3.0`,
		`Accept`:     `application/taxii+json;version=2.1`,
	})
	collectionId, err := collection.GetCollectionsObject()
	if err != nil {
		fmt.Println(err)
		return
	}
	// 获取数据
	startTime := `2023-01-01 00:00:00`
	parsedTime, err := time.Parse(`2006-01-02 15:04:05`, startTime)
	if err != nil {
		fmt.Println(err)
		return
	}
	params := map[string]interface{}{
		`limit`:       []interface{}{10},
		`added_after`: []interface{}{parsedTime},
		`type`:        []interface{}{`indicator`},
	}
	i := 0
	next := ``
	for {
		if next != `` {
			params[`next`] = []interface{}{next}
		}
		collection := getCollectionsData.NewGetCollectionsData(conn, apiRootId, collectionId, map[string]string{
			`User-Agent`: `taxii2-client/2.3.0`,
			`Accept`:     `application/taxii+json;version=2.1`,
		}, params)
		res, err := collection.GetCollectionData()
		if err != nil {
			fmt.Println(err)
			return
		}
		c, _ := json.Marshal(res)
		fmt.Println(fmt.Sprintf(`%s`, c))
		if res.Next == `` {
			break
		}
		next = res.Next
		i++
	}
}
```