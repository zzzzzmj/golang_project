package controllers

import "github.com/astaxie/beego/context"

func MyFilter(ctx *context.Context) {
	ua := ctx.Input.Session("userAccount")
	if ua == nil {
		ctx.WriteString("<html lang=\"en\">\n<head>\n    <meta charset=\"UTF-8\" http-equiv=\"refresh\" content=\"2 url=http://localhost:8080/login\">\n    <title>jumpPage</title>\n</head>\n<body>\n未登录\n</body>\n</html>")
	}
}

/*
页面跳转已修改并测试
*/
