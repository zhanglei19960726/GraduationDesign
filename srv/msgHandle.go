package srv

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"time"
)

//type CDATAText struct {
//	Text string `xml:",innerxml"`
//}

//微信报文格式
type TextRequestBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	Content      string
	MsgId        int
}

//解析报文内容
func parseTextRequestBody(r *http.Request) (*TextRequestBody, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	requestBody := &TextRequestBody{}
	xml.Unmarshal(body, requestBody)
	return requestBody, nil
}

//func value2CDATA(v string) CDATAText {
//	return CDATAText{Text: "<![CDATA[" + v + "]]>"}
//}

//打包
func makeTextResponseBody(fromeUserName, toUserName, content string) ([]byte, error) {
	textResponseBody := &TextRequestBody{}
	textResponseBody.FromUserName = fromeUserName
	textResponseBody.ToUserName = toUserName
	textResponseBody.MsgType = "text"
	textResponseBody.Content = content
	textResponseBody.CreateTime = time.Duration(time.Now().Unix())
	return xml.Marshal(textResponseBody)
}
