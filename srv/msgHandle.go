package srv

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"time"
)

type CDATAText struct {
	Text string `xml:",innerxml"`
}

//微信收消息报文格式
type TextRequestBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	Content      string
	MsgId        int
}

//微信发消息报文格式
type TextReponseBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATAText
	FromUserName CDATAText
	CreateTime   time.Duration
	MsgType      CDATAText
	Content      CDATAText
}

//解析微信客户端消息内容
func parseTextRequestBody(r *http.Request) (*TextRequestBody, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	requestBody := &TextRequestBody{}
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
