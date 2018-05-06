package srv

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"time"
)

//微信请求报文格式
type TextRequestBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	Content      string
	MsgId        int
}

func parseTextRequestBody(r *http.Request) (*TextRequestBody, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	requestBody := &TextRequestBody{}
	xml.Unmarshal(body, requestBody)
	return requestBody, nil
}
