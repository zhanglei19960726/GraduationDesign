package wxsrv

import (
	"GraduationDesign/msgtype"
	"GraduationDesign/wxclient"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
	"strings"
)

const (
	token = "zhang"
)

func makeSignature(timestamp, nonce string) string {
	sl := []string{token, timestamp, nonce}
	sort.Strings(sl)
	s := sha1.New()
	io.WriteString(s, strings.Join(sl, ""))
	return fmt.Sprintf("%x", s.Sum(nil))

}

func validateUrl(w http.ResponseWriter, r *http.Request) bool {
	timestamp := strings.Join(r.Form["timestamp"], "")
	nonce := strings.Join(r.Form["nonce"], "")
	signatureGen := makeSignature(timestamp, nonce)
	signatureIn := strings.Join(r.Form["signature"], "")
	echostr := strings.Join(r.Form["echostr"], "")
	if signatureGen != signatureIn {
		return false
	}
	fmt.Fprintf(w, echostr)
	return true

}
func procRequest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//验证消息是否是微信消息
	if !validateUrl(w, r) {
		log.Println("Wechat Service: this http request is not from Wechat platform!")
		return

	}
	log.Println("Wechat Service: validateUrl Ok!")

	if r.Method == "POST" {
		requestBody, err := parseTextRequestBody(r)
		if err != nil {
			log.Println(err.Error())
			return
		}
		wxHandle(w, requestBody)
	}
}

func wxHandle(w http.ResponseWriter, requestBody *msgtypetype.RequestBody) {
	var content string
	responseBody := make([]byte, 0)
	var err error
	if requestBody.MsgType == "text" {
		fmt.Println("user:", requestBody.FromUserName, "msg:", requestBody.Content)
		if requestBody.Content == "1" {
			content = "ftp://140.143.14.180"
		} else if requestBody.Content == "2" {

		} else {
			content = "回复“1”获得课件下载地址\n回复“2”获得联系方式"
		}
	} else if requestBody.MsgType == "event" {
		if requestBody.Event == "subscribe" {
			wxclient.CreateWxMenu()
			content = "回复“1”获得课件下载地址\n回复“2”获得联系方式"
		}
	} else {
		content = "hello"
	}
	responseBody, err = makeTextResponseBody(requestBody.ToUserName, requestBody.FromUserName, content)
	if err != nil {
		log.Println("Wechat Service : makeTextResponseBody error:", err)
		return
	}
	w.Write(responseBody)

}

func Run() {
	log.Println("Wechat Service: Start!")
	clientMux := http.NewServeMux()
	clientMux.HandleFunc("/", wxclient.AdminHanler)
	srvMux := http.NewServeMux()
	srvMux.HandleFunc("/", procRequest)
	go http.ListenAndServe(":8081", clientMux)
	http.ListenAndServe(":80", srvMux)
}
