package srv

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"time"
)

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

func value2CDATA(v string) string {
	return "<![CATA[" + v + "]]>"
}

//打包
func makeTextResponseBody(fromeUserName, toUserName, content string) ([]byte, error) {
	textResponseBody := &TextRequestBody{}
	textResponseBody.FromUserName = value2CDATA(fromeUserName)
	textResponseBody.ToUserName = value2CDATA(toUserName)
	textResponseBody.MsgType = value2CDATA("text")
	textResponseBody.Content = value2CDATA(content)
	textResponseBody.CreateTime = time.Duration(time.Now().Unix())
	return xml.Marshal(textResponseBody)
}
