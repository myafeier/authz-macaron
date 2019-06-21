package authz_macaron

import (
	"gopkg.in/macaron.v1"
	"net/http"
)

func Auth(authorizer Authorizer)macaron.Handler{

	return func(resp macaron.ResponseWriter,req *macaron.Request) {

		username,password,ok:=req.BasicAuth()
		if !ok{
			http.Error(resp,"",400)
			return
		}
		client:=req.Header.Get("client")
		if err:=authorizer.CheckPwd(username,password,client);err!=nil{
			http.Error(resp,err.Error(),403)
			return
		}

		method:=req.Method
		path:=req.URL.Path

		if err:=authorizer.CheckPermission(username,method,path);err!=nil{
			http.Error(resp,err.Error(),403)
			return
		}

	}


}
