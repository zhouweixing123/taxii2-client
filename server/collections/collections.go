package collections

import (
	"encoding/json"
	"fmt"
	"gitee.com/zhouweixing/taxii2.git/consts"
	"gitee.com/zhouweixing/taxii2.git/server/httpClient"
)

// GetCollectionsObject 结构体用于获取集合对象的请求
type GetCollectionsObject struct {
	Conn    *httpClient.ConnectionSession // HTTP连接会话对象
	Url     string                        // 请求的URL
	Header  map[string]string             // 请求头信息
	ApiRoot string                        // API根路径
}

// NewGetCollectionsObject 返回一个 GetCollectionsObject 的实例指针
// conn: HTTP连接会话对象
// url: 请求的URL
// header: 请求头信息
func NewGetCollectionsObject(conn *httpClient.ConnectionSession, url string, header map[string]string) *GetCollectionsObject {
	return &GetCollectionsObject{
		Conn:    conn,
		Header:  header,
		ApiRoot: url,
	}
}

// GetCollectionsObject 函数返回集合对象的ID和错误信息
func (g *GetCollectionsObject) GetCollectionsObject() (string, error) {
	// 解析URL
	parsedURL, err := g.Conn.ParseUrl()
	if err != nil {
		return ``, err
	}
	host := parsedURL.Host
	scheme := parsedURL.Scheme
	// 构造请求URL
	requestUrl := fmt.Sprintf(`%s://%s%scollections/`, scheme, host, g.ApiRoot)
	// 发起HTTP GET请求
	res, err := g.Conn.HttpGetRequest(requestUrl, g.Header, nil)
	var collection consts.Collections
	// 解析响应的JSON数据
	err = json.Unmarshal([]byte(res), &collection)
	if err != nil {
		return ``, err
	}
	// 返回第一个集合的ID
	return collection.Collections[0].Id, nil
}
