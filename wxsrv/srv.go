package wxsrv

import (
	"GraduationDesign/msgtype"
	"GraduationDesign/wxclient"
	"crypto/sha1"
	"fmt"
	"io"
	"io/ioutil"
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

func getData(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://www.zhangleispace.club:8009/images" + r.RequestURI)
	if err != nil {
		log.Println("li chen hui")
		return
	}
	body, _ := ioutil.ReadAll(resp.Body)
	w.Write(body)
}

func wxHandle(w http.ResponseWriter, requestBody *msgtypetype.RequestBody) {
	var content string
	responseBody := make([]byte, 0)
	var err error
	if requestBody.MsgType == "text" {
		fmt.Println("user:", requestBody.FromUserName, "msg:", requestBody.Content)
		if requestBody.Content == "1" {
			content = "http://www.zhangleispace.club:8009/images/"
		} else if requestBody.Content == "2" {
			content = "15029236434"
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
	http.HandleFunc("/", wxclient.HomeHanler)
	http.HandleFunc("/admin", wxclient.AdminHandler)
	srvMux := http.NewServeMux()
	srvMux.HandleFunc("/", getData)
	srvMux.HandleFunc("/wx", procRequest)
	srvMux.HandleFunc("/getData", wxclient.GetData)
	go http.ListenAndServe(":8081", nil)
	http.ListenAndServe(":80", srvMux)
}
