package wxsrv

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"time"
)

type CDATAText struct {
	Text string `xml:",innerxml"`
}

//创建菜单微信返回json格式
type MenErrorResponse struct {
	ErrorCode string
	ErrMsg    string
}

type msgBase struct {
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	Content      string
}

//请求普通消息格式
type RequestBody struct {
	XMLName xml.Name `xml:"xml"`
	msgBase
	MsgId    int
	Event    string
	EventKey string
}

type repMsgBase struct {
	ToUserName   CDATAText
	FromUserName CDATAText
	CreateTime   time.Duration
	MsgType      CDATAText
	Content      CDATAText
}

//响应普通消息格式
type TextReponseBody struct {
	XMLName xml.Name `xml:"xml"`
	repMsgBase
}

type ClickResponse struct {
	XMLName xml.Name `xml:"xml"`
	repMsgBase
	Event    CDATAText
	EventKey CDATAText
}

//解析微信客户端消息内容
func parseTextRequestBody(r *http.Request) (*RequestBody, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	requestBody := &RequestBody{}
	xml.Unmarshal(body, requestBody)
	return requestBody, nil
}

func value2CDATA(v string) CDATAText {
	return CDATAText{Text: "<![CDATA[" + v + "]]>"}
}

//打包服务器响应消息
func makeTextResponseBody(fromeUserName, toUserName, content string) ([]byte, error) {
	textResponseBody := &TextReponseBody{}
	textResponseBody.FromUserName = value2CDATA(fromeUserName)
	textResponseBody.ToUserName = value2CDATA(toUserName)
	textResponseBody.MsgType = value2CDATA("text")
	textResponseBody.Content = value2CDATA(content)
	textResponseBody.CreateTime = time.Duration(time.Now().Unix())
	return xml.Marshal(textResponseBody)
}
