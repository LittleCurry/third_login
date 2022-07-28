package third_login

import (
	"encoding/json"
	"errors"
	"regexp"
	"strconv"
)

type AuthWb struct {
	*Auth
}

type AuthWbErrRes struct {
	Error            int    `json:"error_code"`
	ErrorDescription string `json:"error"`
}

type AuthWbSuccRes struct {
	AccessToken string `json:"access_token"`
	Openid      string `json:"uid"`
}

//获取登录地址
func (a *AuthWb) GetRurl(state string) string {
	return "https://api.weibo.com/oauth2/authorize?client_id=" + a.AppId + "&response_type=code&display=page&redirect_uri=" + a.RedirUri + "&state=" + state
}

//获取token
func (a *AuthWb) GetToken(code string) (*AuthWbSuccRes, error) {
	str, err := HttpPost("https://api.weibo.com/oauth2/access_token?client_id=" + a.AppId + "&client_secret=" + a.AppSecret + "&code=" + code + "&grant_type=authorization_code&redirect_uri=" + a.RedirUri)
	if err != nil {
		return nil, err
	}
	ismatch, _ := regexp.MatchString("error", str)
	if ismatch {

		p := &AuthWbErrRes{}
		err := json.Unmarshal([]byte(str), p)
		if err != nil {
			return nil, err
		}
		return nil, errors.New("Error:" + strconv.Itoa(p.Error) + " Error_description:" + p.ErrorDescription)

	} else {
		p := &AuthWbSuccRes{}
		err := json.Unmarshal([]byte(str), p)
		if err != nil {
			return nil, err
		}
		return p, nil
	}
}

//获取第三方用户信息
func (a *AuthWb) Get_User_Info(access_token string, openid string) (string, error) {
	return HttpGet("https://api.weibo.com/2/users/show.json?access_token=" + access_token + "&uid=" + openid)
}

//构造方法
func NewAuthWb(config *Auth) *AuthWb {
	return &AuthWb{config}
}
