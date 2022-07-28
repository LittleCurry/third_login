# third_login

> Golang第三方登录服务端(开发中, 还没上线使用)

### firebase

> 类似国内友盟, 可以在firebase 里直接集成 github,google,facebook,twitter...登录
[申请流程参考](https://www.freesion.com/article/66471214400)

```golang
package main

import (
	"third_login"
	"github.com/valyala/fasthttp"
)

func main() {
	third_login.InitFirebase("./configs/firebase.json") // 参考申请流程获取 firebase.json
	r := NewRouter()
	r.Use(Middleware)
	r.Add("/firebase/login", FirebaseCheckIdToken)
}

func FirebaseCheckIdToken(ctx *fasthttp.RequestCtx) {
	var idToken string
	argMap := (*httpArgs)(ctx.QueryArgs())
	argMap.getStringArg("id_token", &idToken)
	token, err := third_sign.CheckIdToken(idToken)
	// token.UID -> 做存储和用户的唯一id
}
```

### qq

```golang
package main

import (
	"third_login"
)

func main() {
	third := third_login.NewAuthQQ(&third_login.Auth_conf{Appid: "xxx", Appkey: "xxx", Rurl: "http://123123.cn"})
	token, err := third.Get_Token("code")
	me, err := third.Get_Me(token)
	userinfo, _ := third.Get_User_Info(token, me.OpenID)
}
```

### weibo

```golang
package main

import (
	"third_login"
)

func main() {
	third := third_login.NewAuthWb(&third_login.Auth_conf{Appid: "xxx", Appkey: "xxx", Rurl: "http://123123.cn"})
	token, err := third.Get_Token("code")
	me, err := third.Get_Me(token)
	userinfo, _ := third.Get_User_Info(token, me.OpenID)
}
```

### wechat

```golang
package main

import (
	"third_login"
)

func main() {
	third := third_login.NewAuthWx(&third_login.Auth_conf{Appid: "xxx", Appkey: "xxx", Rurl: "http://123123.cn"})
	token, err := third.Get_Token("code")
	me, err := third.Get_Me(token)
	userinfo, _ := third.Get_User_Info(token, me.OpenID)
}
```

### github

```golang
package main

import (
	"third_login"
)

func main() {
	third := third_login.NewAuthGithub(&third_login.Auth_conf{Appid: "xxx", Appkey: "xxx", Rurl: "http://123123.cn"})
	token, err := third.Get_Token("code")
	me, err := third.Get_Me(token)
	userinfo, _ := third.Get_User_Info(token, me.OpenID)
}
```
