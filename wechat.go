package third_login

import (
	"encoding/json"
	"errors"
	"regexp"
	"strconv"
)

type AuthWx struct {
	*Auth
}

type AuthWxErrRes struct {
	Error            int    `json:"errcode"`
	ErrorDescription string `json:"errmsg"`
}

type AuthWxSuccRes struct {
	AccessToken string `json:"access_token"`
	Openid      string `json:"openid"`
}

//获取登录地址
func (a *AuthWx) GetRurl(state string) string {
	return "https://open.weixin.qq.com/connect/qrconnect?appid=" + a.AppId + "&redirect_uri=" + a.RedirUri + "&response_type=code&scope=snsapi_login&state=" + state
}

//获取token
func (a *AuthWx) GetToken(code string) (*AuthWxSuccRes, error) {

	str, err := HttpPost("https://api.weixin.qq.com/sns/oauth2/access_token?appid=" + a.AppId + "&secret=" + a.AppSecret + "&code=" + code + "&grant_type=authorization_code")
	if err != nil {
		return nil, err
	}

	ismatch, _ := regexp.MatchString("errcode", str)
	if ismatch {

		p := &AuthWxErrRes{}
		err := json.Unmarshal([]byte(str), p)
		if err != nil {
			return nil, err
		}
		return nil, errors.New("Error:" + strconv.Itoa(p.Error) + " Error_description:" + p.ErrorDescription)

	} else {

		p := &AuthWxSuccRes{}
		err := json.Unmarshal([]byte(str), p)
		if err != nil {
			return nil, err
		}
		return p, nil
	}
}

//获取第三方用户信息
func (a *AuthWx) Get_User_Info(access_token string, openid string) (string, error) {
	return HttpGet("https://api.weixin.qq.com/sns/userinfo?access_token=" + access_token + "&openid=" + openid)
}

//构造方法
func NewAuthWx(config *Auth) *AuthWx {
	return &AuthWx{config}
}
