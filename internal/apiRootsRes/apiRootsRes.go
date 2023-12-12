package apiRootsRes

import (
	"encoding/json"
	"gitee.com/zhouweixing/taxii2.git/consts"
	"gitee.com/zhouweixing/taxii2.git/internal/httpClient"
)

// ApiRoots 结构体用于管理API根路径的相关信息
type ApiRoots struct {
	Conn   *httpClient.ConnectionSession // http客户端连接会话
	Header map[string]string             // 请求头部信息
}

// NewApiRoots 函数用于创建一个ApiRoots对象。
// 参数conn是一个httpClient.ConnectionSession类型的指针，表示http客户端的连接会话。
// 参数header是一个map[string]string类型的对象，表示请求的头部信息。
// 返回值是一个指向新创建的ApiRoots对象的指针。
func NewApiRoots(conn *httpClient.ConnectionSession, header map[string]string) *ApiRoots {
	return &ApiRoots{
		Conn:   conn,
		Header: header,
	}
}

// GetApiRoot 获取ApiRoot字符串和错误信息
func (a *ApiRoots) GetApiRoot() (string, error) {
	// 向指定的 URL 发送 HTTP GET 请求，并获取响应结果和可能发生的错误
	res, err := a.Conn.HttpGetRequest(a.Conn.Url, a.Header, nil)
	if err != nil {
		return ``, err
	}
	// 定义变量 apiRootRes 为类型 ApiRootsRes
	var apiRootRes consts.ApiRootsRes
	// 将字节数组 res 反序列化为 apiRootRes 变量
	err = json.Unmarshal([]byte(res), &apiRootRes)
	// 如果反序列化过程中出现错误，则返回空字符串和错误信息
	if err != nil {
		return ``, err
	}
	// 返回 apiRootRes 结构体中的第一个 ApiRoot 元素和空错误
	return apiRootRes.ApiRoots[0], nil
}
