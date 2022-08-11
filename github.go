package third_login

import (
	"errors"
	"regexp"
)

type AuthGithub struct {
	*Auth
}

//获取登录地址
func (a *AuthGithub) Get_Rurl(state string) string {
	return "https://github.com/login/oauth/authorize?client_id=" + a.AppId + "&redirect_uri=" + a.RedirUri + "&state=" + state
}

//获取token
func (a *AuthGithub) Get_Token(code string) (string, error) {
	str, err := HttpGet("https://github.com/login/oauth/access_token?client_id=" + a.AppId + "&client_secret=" + a.AppSecret + "&code=" + code + "&redirect_uri=" + a.RedirUri)
	if err != nil {
		return "", err
	}

	ismatch, _ := regexp.MatchString("error", str)
	if ismatch {

		return "", errors.New(str)

	} else {
		re, _ := regexp.Compile("access_token=(.*)&scope")
		newres := re.FindStringSubmatch(str)
		if len(newres) >= 2 {
			return newres[1], nil
		}
		return "", nil
	}

}

//获取第三方用户信息
func (a *AuthGithub) Get_User_Info(access_token string) (string, error) {
	str, err := HttpGet("https://api.github.com/user?access_token=" + access_token)
	if err != nil {
		return "", err
	}
	return string(str), nil
}

//构造方法
func NewAuthGithub(config *Auth) *AuthGithub {
	return &AuthGithub{config}
}
