package WxPlatUtil

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

//请求失败
type AccessTokenErrorResponse struct {
	Error  float64
	Errmsg string
}

var (
	appID               = "wxd2329671b539d3c8"
	appSecret           = "c90858154797ced838e0d828b68b653b"
	accessTokenFetchUrl = "https://api.weixin.qq.com/cgi-bin/token"
	AccessToken         string
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

func GetAndUpdateDBWxAToken() error {
	token, err := fetchAccessToken(appID, appSecret, accessTokenFetchUrl)
	if err != nil {
		log.Println("向微信服务器发送获取accessToken的get请求失败：", err)
		return err
	}
	AccessToken = token
	return nil
}
