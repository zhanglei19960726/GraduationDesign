package wxsrv

import (
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

func wxHandle(w http.ResponseWriter, requestBody *RequestBody) {
	if requestBody.MsgType == "text" {
		fmt.Println("user:", requestBody.FromUserName, "msg:", requestBody.Content)
		responseBody, err := makeTextResponseBody(requestBody.ToUserName, requestBody.FromUserName, "hello")
		if err != nil {
			log.Println("Wechat Service : makeTextResponseBody error:", err)
			return
		}
		w.Write(responseBody)
	} else if requestBody.MsgType == "event" {
		if requestBody.Event == "subscribe" {
			wxclient.CreateWxMenu()
			responseBody, err := makeTextResponseBody(requestBody.ToUserName, requestBody.FromUserName, "hello")
			if err != nil {
				log.Println("Wechat Service : makeTextResponseBody error:", err)
				return
			}
			w.Write(responseBody)
		}
	}

}

func Run() {
	log.Println("Wechat Service: Start!")
	http.HandleFunc("/", procRequest)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("Wechat Service: ListenAndServe failed, ", err)

	}
	log.Println("Wechat Service: Stop!")

}
