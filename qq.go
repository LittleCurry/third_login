package third_login

import (
	"encoding/json"
	"errors"
	"regexp"
	"strconv"
)

type AuthQQ struct {
	*Auth
}
type AuthQqErrRes struct {
	Error            int    `json:"error"`
	ErrorDescription string `json:"error_description"`
}
type AuthQQMe struct {
	ClientId string `json:"client_id"`
	OpenID   string `json:"openid"`
}

//获取登录地址
func (a *AuthQQ) GetRurl(state string) string {
	return "https://graph.qq.com/oauth2.0/authorize?response_type=code&client_id=" + a.AppId + "&redirect_uri=" + a.RedirUri + "&state=" + state
}

//获取token
func (a *AuthQQ) GetToken(code string) (string, error) {

	str, err := HttpGet("https://graph.qq.com/oauth2.0/token?grant_type=authorization_code&client_id=" + a.AppId + "&client_secret=" + a.AppSecret + "&code=" + code + "&redirect_uri=" + a.RedirUri)
	if err != nil {
		return "", err
	}

	ismatch, _ := regexp.MatchString("error", str)
	if ismatch {
		re, _ := regexp.Compile("({.*})")
		newres := re.FindStringSubmatch(str)
		errstr := newres[0]
		p := &AuthQqErrRes{}
		err := json.Unmarshal([]byte(errstr), p)
		if err != nil {
			return "", err
		}
		return "", errors.New("Error:" + strconv.Itoa(p.Error) + " Error_description:" + p.ErrorDescription)

	} else {
		re, _ := regexp.Compile("access_token=(.*)&expires_in")
		newres := re.FindStringSubmatch(str)
		if len(newres) >= 2 {
			return newres[1], nil
		}
		return "", nil
	}

}

//获取第三方id
func (a *AuthQQ) Get_Me(access_token string) (*AuthQQMe, error) {

	str, err := HttpGet("https://graph.qq.com/oauth2.0/me?access_token=" + access_token)
	if err != nil {
		return nil, err
	}
	ismatch, _ := regexp.MatchString("error", str)
	if ismatch {
		re, _ := regexp.Compile("({.*})")
		newres := re.FindStringSubmatch(str)
		errstr := newres[0]
		p := &AuthQqErrRes{}
		err := json.Unmarshal([]byte(errstr), p)
		if err != nil {
			return nil, err
		}
		return nil, errors.New("Error:" + strconv.Itoa(p.Error) + " Error_description:" + p.ErrorDescription)

	} else {
		re, _ := regexp.Compile("({.*})")
		newres := re.FindStringSubmatch(str)
		errstr := newres[0]
		p := &AuthQQMe{}
		err := json.Unmarshal([]byte(errstr), p)
		if err != nil {
			return nil, err
		}
		return p, nil
	}
}

//获取第三方用户信息
func (a *AuthQQ) GetUserInfo(access_token string, openid string) (string, error) {
	return HttpGet("https://graph.qq.com/user/get_user_info?access_token=" + access_token + "&oauth_consumer_key=" + a.AppId + "&openid=" + openid)
}

//构造方法
func NewAuthQQ(config *Auth) *AuthQQ {
	return &AuthQQ{config}
}
