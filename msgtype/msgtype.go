package msgtypetype

import (
	"encoding/xml"
	"os"
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

//图文素材格式

type Articles struct {
	Title            string `json:"title"`              //标题
	ThumbMediaId     string `json:"thumb_media_id"`     //图文消息的封面图片素材ID
	Author           string `json:"author"`             //作者
	Digest           string `json:"digest"`             //图文消息摘要
	ShowCoverPic     int    `json:"show_cover_pic"`     //是否显示封面
	Content          string `json:"content"`            //图文消息的具体内容
	ContentSourceUrl string `json:"content_source_url"` //图文消息的原文地址，即点击“阅读原文”后的URL
}

type ArticlesReq struct {
	Articles []Articles `json:"articles"`
}

type ArticlesResp struct {
	MediaId string
}

type Material struct {
	File     *os.File
	Filename string
}

type MaterialReq struct {
	Media    []byte `json:"media"`
	Filename string `json:"filename"`
}
