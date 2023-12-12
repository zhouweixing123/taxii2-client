package httpClient

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var connectSession *ConnectionSession

// ConnectionSession 与服务器的连接回话
type ConnectionSession struct {
	session *http.Client // 执行http请求
	user    string       // 用于身份验证的用户名
	pass    string       // 用于身份验证的密码
	url     string       // 请求的url
}

// NewConnectionSession 创建一个新的连接会话
// 如果已经存在连接会话，则返回现有的连接会话，如果没有则创建一个新的连接会话
func NewConnectionSession(user, pass, url string) *ConnectionSession {
	if connectSession != nil {
		return connectSession
	}
	return &ConnectionSession{
		session: &http.Client{}, // 初始化http客户端
		user:    user,           // 用户名
		pass:    pass,           // 密码
		url:     url,            // 请求的url
	}
}

// HttpGetRequest 发起一个http的Get请求，并返回响应主题作为字符串
// requestUrl: 请求链接
// headers: 请求头
// params: 请求参数
func (c *ConnectionSession) HttpGetRequest(requestUrl string, headers map[string]string, params map[string]interface{}) (string, error) {
	// 创建一个合并后的请求头，用于实际发送请求
	mergeHeaders := make(map[string]string)
	for k, v := range headers {
		mergeHeaders[k] = v
	}
	// 将参数组装成查询字符串
	queryParams := c.filterToQueryParams(params)
	// 发送http的Get请求
	resp, err := c.sendRequest(`GET`, requestUrl, mergeHeaders, queryParams)
	defer resp.Body.Close()
	// 确保在函数结束时关闭响应体
	if err != nil {
		return ``, err
	}
	// 检查http响应状态码，如果不是200，则返回Http错误
	if resp.StatusCode != http.StatusOK {
		return ``, fmt.Errorf(`http Request Error: StatusCode: %d`, resp.StatusCode)
	}
	// 读取并返回响应体
	body, err := io.ReadAll(resp.Body)
	if !errors.Is(err, nil) {
		return ``, err
	}
	return string(body), nil
}

// sendRequest 发送Http请求到指定的URL
// method 请求方法,
// requestUrl 请求的url
// headers 参数包含请求头信息
// params 参数包含的查询参数
func (c *ConnectionSession) sendRequest(method string, requestUrl string, headers map[string]string, params map[string]string) (*http.Response, error) {
	// 创建一个新的http请求
	req, err := http.NewRequest(method, requestUrl, nil)
	if err != nil {
		return nil, err
	}
	// 使用Basic Auth进行身份认证
	req.SetBasicAuth(c.user, c.pass)
	// 设置请求头细心你
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	// 设置查询参数
	query := req.URL.Query()
	for k, v := range params {
		query.Add(k, v)
	}
	req.URL.RawQuery = query.Encode()
	// 执行http请求并返回响应
	return c.session.Do(req)
}

// filterToQueryParams 将传入的参数进行过滤并转换为查询参数
// filterMap 参数包含过滤器的键值对，用于构建查询参数
// return map[string]string,包含了用于查询的参数
func (c *ConnectionSession) filterToQueryParams(filterMap map[string]interface{}) map[string]string {
	// 创建空洞额查询参数映射
	queryParams := make(map[string]string)

	// 遍历过滤器参数
	for kwarg, value := range filterMap {
		switch arglist := value.(type) {
		case []interface{}:
			if len(arglist) == 0 {
				// 如果参数列表为空，继续下一次循环
				continue
			}
			// 根据不同的参数类型构建查询参数
			switch kwarg {
			case "version":
				// 处理版本号参数
				var versions []string
				for _, val := range arglist {
					versions = append(versions, c.ensureDatetimeToString(val).(string))
				}
				queryParams["match[version]"] = strings.Join(versions, ",")
			case "added_after":
				if len(arglist) > 1 {
					// Handle error condition in Go for multiple values
					continue
				}
				queryParams["added_after"] = c.ensureDatetimeToString(arglist[0]).(string)
			case "limit":
				// 每次获取多少条数据
				queryParams["limit"] = fmt.Sprintf("%v", arglist[0])
			case "next":
				// 处理下一批数据
				var nextValues []string
				for _, val := range arglist {
					nextValues = append(nextValues, fmt.Sprintf("%v", val))
				}
				queryParams["next"] = strings.Join(nextValues, ",")
			default:
				// 处理其他的参数
				var matchValues []string
				for _, val := range arglist {
					matchValues = append(matchValues, fmt.Sprintf("%v", val))
				}
				queryParams["match["+kwarg+"]"] = strings.Join(matchValues, ",")
			}
		}
	}
	return queryParams
}

// formatDatetime 格式化给定的事件为taxii的时间
// dttm 需要格式化的事件对象
// return 格式化后的时间字符串
func (c *ConnectionSession) formatDatetime(dttm time.Time) string {
	myTime := time.Date(dttm.Year(), dttm.Month(), dttm.Day(), dttm.Hour(), dttm.Minute(), dttm.Second(), dttm.Nanosecond(), time.UTC)
	formattedTime := myTime.Format("2006-01-02T15:04:05.000000Z")
	return formattedTime
}

func (c *ConnectionSession) ensureDatetimeToString(maybeDttm interface{}) interface{} {
	dttm, ok := maybeDttm.(time.Time)
	if ok {
		return c.formatDatetime(dttm)
	}
	return maybeDttm
}

func (c *ConnectionSession) ParseUrl() (*url.URL, error) {
	parseUrl, err := url.Parse(c.url)
	if err != nil {
		return nil, err
	}
	return parseUrl, nil
}
