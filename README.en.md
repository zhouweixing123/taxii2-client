# taxii2

#### Introduction
taxii2 is a minimal client implementation for TAXII 2.X servers. It supports the following TAXII 2.X API services:
1. GET API Root Information
2. Get Collections

#### Software Architecture
Software architecture description:
1. Get API Root Information 
2. Get Collections

#### Installation
```shell
go get -u gitee.com/zhouweixing/taxii2.git
```

#### Usage Instructions
The TAXII client is intended to be used as a Go library. Currently, there is no command-line client available.
#### taxii2-client provides three classes:
1.  Server
2.  ApiRoot
3.  Collection

#### Example
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
	// Server interface: http://ip:port/taxii2/
	conn := httpClient.NewConnectionSession(`username`, `password`, `server interface`)
	apiRoots := apiRootsRes.NewApiRoots(conn, map[string]string{
		`User-Agent`: `taxii2-client/2.3.0`,
		`Accept`:     `application/taxii+json;version=2.1`,
	})
	// Get apiRootId
	apiRootId, err := apiRoots.GetApiRoot()
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get collectionID
	collection := collections.NewGetCollectionsObject(conn, apiRootId, map[string]string{
		`User-Agent`: `taxii2-client/2.3.0`,
		`Accept`:     `application/taxii+json;version=2.1`,
	})
	collectionId, err := collection.GetCollectionsObject()
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get data
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