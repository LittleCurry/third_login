# third_login


### firebase(已和客户端调通)

> 类似国内友盟, 可以在firebase 里直接集成 google(已测),facebook(已测),匿名(已测),twitter,github...登录
 
[申请流程参考](https://www.freesion.com/article/66471214400)

```golang
package main

import (
	"third_login"
	"github.com/valyala/fasthttp"
)

func main() {
	third_login.InitFirebase("./configs/firebase.json") // 参考申请流程获取 firebase.json
	r := NewRouter() // 路由自行实现
	r.Use(Middleware)
	r.Add("/firebase/login", FirebaseCheckIdToken)
	err := fasthttp.ListenAndServe("127.0.0.1:9999", r.Handler)
}

func FirebaseCheckIdToken(ctx *fasthttp.RequestCtx) {
	var idToken string
	argMap := (*httpArgs)(ctx.QueryArgs())
	argMap.getStringArg("id_token", &idToken)
	token, err := third_sign.CheckIdToken(idToken)
	// token.UID -> 做存储和用户的唯一id
}
```

### qq(未测试)

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

### weibo(未测试)

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

### wechat(未测试)

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


> 欢迎加入golang微信群：
> 

<img width="151" height="211" src="https://bhhz.oss-cn-shanghai.aliyuncs.com/wechat08-02.jpg"/>
