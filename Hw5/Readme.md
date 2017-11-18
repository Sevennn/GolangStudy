# 本应用基于Martini进行服务搭建

### 静态服务支持
```
	m.Use(martini.Static("assets"))
```
martini提供了简单结果为我们所使用
### 简单的js访问与提交表单并且返回表单
此二者功能上较为类似（简单的视为ajax）
martini参考了express框架，写法上较为简单，此处仅给出POST实例,GET实例请在server.go中查看，也可以命令行执行 `curl http://localhost:3000/api/test`
```
	m.Post("/update/userinfo", func(r render.Render, req *http.Request) {
		fmt.Println(req.FormValue("firstname"))
		r.JSON(200, map[string]interface{}{
			"firstName": req.FormValue("firstname"),
			"lastName" : req.FormValue("lastname"),
			"userName" : req.FormValue("username"),
		})
	})
```
此处接受了表单数据，并把表单数据返回前端
### 对于未知路径给出开发中提示，返回码 5xx
martini的路由提供了相应的方法处理
```
	m.NotFound(func(r render.Render, req *http.Request)  {
		fmt.Println("[martini] Page Not Found[500]")
		r.HTML(500, "notfound", req.URL.Path)
	})
```

### Extra
使用了模板输出