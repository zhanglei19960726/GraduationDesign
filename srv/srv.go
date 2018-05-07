package srv

import (
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
	w.Header().Set("Content-type", "text/xml")
	r.ParseForm()
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
		fmt.Println("Wechat Service: Recv text msg [" + requestBody.Content.Text + "] from user [" + requestBody.FromUserName.Text + "]!")
		responseBody, err := makeTextResponseBody(requestBody.ToUserName.Text, requestBody.FromUserName.Text, "hello "+requestBody.Content.Text)
		if err != nil {
			log.Println(err.Error())
			return
		}
		fmt.Fprint(w, responseBody)
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
