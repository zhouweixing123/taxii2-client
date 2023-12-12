package main

import (
	"encoding/json"
	"fmt"
	"gitee.com/zhouweixing/taxii2.git/internal/apiRootsRes"
	"gitee.com/zhouweixing/taxii2.git/internal/collections"
	"gitee.com/zhouweixing/taxii2.git/internal/getCollectionsData"
	"gitee.com/zhouweixing/taxii2.git/internal/httpClient"
	"time"
)

type taxii struct {
	conn         *httpClient.ConnectionSession
	apiRootId    string
	collectionId string
}

func NewTaxii(user, pass, url string) *taxii {
	conn := httpClient.NewConnectionSession(user, pass, url)
	return &taxii{
		conn: conn,
	}
}

func (t *taxii) GetCollectionData() {
	err := t.getApiRoot()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = t.getCollectionId()
	if err != nil {
		fmt.Println(err)
		return
	}
	params, err := t.getParams()
	if err != nil {
		fmt.Println(err)
		return
	}

	i := 0
	next := ``
	for {
		if next != `` {
			params[`next`] = []interface{}{next}
		}
		collection := getCollectionsData.NewGetCollectionsData(t.conn, t.apiRootId, t.collectionId, t.getHeader(), params)
		res, err := collection.GetCollectionData()
		fmt.Println(fmt.Sprintf(`%d页的数量: `, i), len(res.Objects))
		fmt.Println(err)
		c, _ := json.Marshal(res)
		fmt.Println(fmt.Sprintf(`%d页的数据: `, i), string(c))
		time.Sleep(1 * time.Second)
		if res.Next == `` {
			break
		}
		next = res.Next
		i++
	}
}
func (t *taxii) getApiRoot() error {
	apiRoot := apiRootsRes.NewApiRoots(t.conn, t.getHeader())
	apiRootId, err := apiRoot.GetApiRoot()
	if err != nil {
		return err
	}
	t.apiRootId = apiRootId
	return nil
}
func (t *taxii) getCollectionId() error {
	collection := collections.NewGetCollectionsObject(t.conn, t.apiRootId, t.getHeader())
	collectionId, err := collection.GetCollectionsObject()
	if err != nil {
		return err
	}
	t.collectionId = collectionId
	return nil
}

func (t *taxii) getHeader() map[string]string {
	return map[string]string{
		`User-Agent`: `taxii2-client/2.3.0`,
		`Accept`:     `application/taxii+json;version=2.1`,
	}
}

func (t *taxii) getParams() (map[string]interface{}, error) {
	startTime := `2023-01-01 00:00:00`
	parsedTime, err := time.Parse(`2006-01-02 15:04:05`, startTime)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return map[string]interface{}{
		`limit`:       []interface{}{10},
		`added_after`: []interface{}{parsedTime},
		`type`:        []interface{}{`indicator`},
	}, nil
}

func main() {
	client := NewTaxii(`xx`, `xx`, `http://xxxx.xxx.xxx.xxx:xxxx/taxii2/`)
	defer client.conn.Close()
	err := client.getApiRoot()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = client.getCollectionId()
	if err != nil {
		fmt.Println(err)
		return
	}
	client.GetCollectionData()
}
