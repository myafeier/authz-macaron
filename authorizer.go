package authz_macaron


//授权人接口，只要实现这个接口就可以接入任意的授权模型
type Authorizer interface {

	//检查用户名密码,及用户所处客户端
	CheckPwd(username,password,client string)error
	CheckPermission(username,method,path string)error

}