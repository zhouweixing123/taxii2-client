package getCollectionsData

import (
	"encoding/json"
	"fmt"
	"gitee.com/zhouweixing/taxii2.git/consts"
	"gitee.com/zhouweixing/taxii2.git/server/httpClient"
)

// GetCollectionsData 结构体用于获取集合数据
type GetCollectionsData struct {
	// http客户端连接会话
	Conn *httpClient.ConnectionSession
	// API根节点ID
	ApiRootId string
	// 集合ID
	CollectionId string
	// 请求头
	Header map[string]string
	// 请求参数
	Params map[string]interface{}
}

// NewGetCollectionsData 返回一个 GetCollectionsData 结构体的指针，该结构体包含了用于获取集合数据的相关信息
func NewGetCollectionsData(conn *httpClient.ConnectionSession, apiRootId, collectionId string, header map[string]string, params map[string]interface{}) *GetCollectionsData {
	return &GetCollectionsData{
		Conn:         conn,
		ApiRootId:    apiRootId,
		CollectionId: collectionId,
		Header:       header,
		Params:       params,
	}
}

func (g *GetCollectionsData) GetCollectionData() (*consts.GetCollectionDataRes, error) {
	// 解析url连接
	parsedURL, err := g.Conn.ParseUrl()
	if err != nil {
		return nil, err
	}
	// 获取ip和端口
	host := parsedURL.Host
	//  获取url连接的协议 http/https
	scheme := parsedURL.Scheme
	requestUrl := fmt.Sprintf(`%s://%s%scollections/%s/objects/`, scheme, host, g.ApiRootId, g.CollectionId)
	//fmt.Println(requestUrl)
	res, err := g.Conn.HttpGetRequest(requestUrl, g.Header, g.Params)
	//fmt.Println(res, err)
	if err != nil {
		return nil, err
	}
	var getCollectionDataRes consts.GetCollectionDataRes
	err = json.Unmarshal([]byte(res), &getCollectionDataRes)
	if err != nil {
		return nil, err
	}
	return &getCollectionDataRes, nil
}
