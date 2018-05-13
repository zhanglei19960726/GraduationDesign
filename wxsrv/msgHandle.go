package wxsrv

import (
	"GraduationDesign/msgtype"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"time"
)

//解析微信客户端消息内容
func parseTextRequestBody(r *http.Request) (*msgtypetype.RequestBody, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	requestBody := &msgtypetype.RequestBody{}
	xml.Unmarshal(body, requestBody)
	return requestBody, nil
}

func value2CDATA(v string) msgtypetype.CDATAText {
	return msgtypetype.CDATAText{Text: "<![CDATA[" + v + "]]>"}
}

//打包服务器响应消息
func makeTextResponseBody(fromeUserName, toUserName, content string) ([]byte, error) {
	textResponseBody := &msgtypetype.TextReponseBody{}
	textResponseBody.FromUserName = value2CDATA(fromeUserName)
	textResponseBody.ToUserName = value2CDATA(toUserName)
	textResponseBody.MsgType = value2CDATA("text")
	textResponseBody.Content = value2CDATA(content)
	textResponseBody.CreateTime = time.Duration(time.Now().Unix())
	return xml.Marshal(textResponseBody)
}
