package third_login

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//var (
//	AppId       = "" //应用ID
//	appSecret   = "" //应用密钥
//	redirectUri = "" //在API控制台页面配置的授权重定向URI, https://console.cloud.google.com/apis/credentials
//)
//const API_BASE            = 'https://www.googleapis.com/';
//const AUTHORIZE_URL       = 'https://accounts.google.com/o/oauth2/v2/auth';
//protected $AccessTokenURL = 'https://www.googleapis.com/oauth2/v4/token';

type AuthGoogle struct {
	*Auth
}

var authGoogle = &AuthGoogle{&Auth{}}

func InitGoogleLogin(app_id, secret, redirect_uri string) {
	authGoogle.AppId = app_id
	authGoogle.AppSecret = secret
	authGoogle.RedirUri = redirect_uri
}

func GetTokenFromGoogle(code string) string {
	//code 用户授权后客户端获取到的code
	reader := strings.NewReader(fmt.Sprintf("client_id=%s&client_secret=%s&redirect_uri=%s&grant_type=authorization_code&code=%s", authGoogle.AppId, authGoogle.AppSecret, authGoogle.RedirUri, code))
	res, err := http.Post("https://oauth2.googleapis.com/token", "application/x-www-form-urlencoded", reader)
	if err != nil {
		log.Println("res error:", err)
		return ""
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("body error:", err)
		return ""
	}
	log.Println("statuscode:", res.StatusCode, "\nhead[name]=", res.Header["Name"], "\nbody is ", string(body))
	return ""
}
