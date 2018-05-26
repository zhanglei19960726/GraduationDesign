package wxsrv

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//Access请求成功
type AccessTokenResponse struct {
	AccessToken string  `json:"access_token"`
	ExpiresIn   float64 `json:"expires_in"`
}

//请求获取Access失败
type AccessTokenErrorResponse struct {
	Error  float64
	Errmsg string
}

var (
	AppID               = "wxf4b1e3a9d5753984"
	AppSecret           = "c8981b2fc40b3ecc24f22dc644829099"
	accessTokenFetchUrl = "https://api.weixin.qq.com/cgi-bin/token"
)

//获取AccessToken
func fetchAccessToken(appID, appSecret, accessTokenFetchUrl string) (string, error) {
	requestLien := strings.Join([]string{accessTokenFetchUrl, "?grant_type=client_credential&appid=", appID, "&secret=", appSecret}, "")
	resp, err := http.Get(requestLien)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Println("发送get请求获取 atoken错误：", err)
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("发送get请求获取atoken读取返回body错误", err)
		return "", err
	}
	if bytes.Contains(body, []byte("access_token")) {
		atr := AccessTokenResponse{}
		err = json.Unmarshal(body, &atr)
		if err != nil {
			log.Println("发送get请求获取atoken返回数据json解析错误，", err)
			return "", err
		}
		return atr.AccessToken, nil
	} else {
		fmt.Println("发送get请求获取 微信返回 err")
		atr := AccessTokenErrorResponse{}
		err = json.Unmarshal(body, &atr)
		fmt.Println("发送get请求获取 微信返回的错误信息", atr)
		if err != nil {
			return "", err
		}
		return "", fmt.Errorf("%s", atr.Errmsg)
	}
}

func GetAndUpdateDBWxAToken() (string, error) {
	token, err := fetchAccessToken(AppID, AppSecret, accessTokenFetchUrl)
	if err != nil {
		log.Println("向微信服务器发送获取accessToken的get请求失败：", err)
		return "", err
	}
	return token, nil
}
