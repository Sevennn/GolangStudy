### xorm体验
基本操作不做赘述，代码里均可体现。具体谈谈xorm的体验。
可以说xorm确实极大简化了代码工作，就算不熟悉sql语言的程序员也可以轻松利用这套框架写出不错的网站数据库管理应用。
不过，真正的程序员还是应该主动从底层学起，框架是工具，但是我们对施工对象的理解则决定了我们能使用工具的效率能够有多高
### 实验测试部分
配置环境后，输入(最重要的是在数据库中建好test数据库)
`go run main.go`
服务跑起来后，进行测试
```
Seven@SevenBig:/mnt/c/Users/Administrator$ curl -d "username=ooo&departname=1" http://localhost:8080/service/userinfo
{
  "UID": 1,
  "UserName": "ooo",
  "DepartName": "1",
  "CreateAt": "2017-11-27T20:31:42.9758369+08:00"
}
Seven@SevenBig:/mnt/c/Users/Administrator$ curl http://localhost:8080/service/userinfo?userid=
[
  {
    "UID": 1,
    "UserName": "ooo",
    "DepartName": "1",
    "CreateAt": "2017-11-28T04:31:42+08:00"
  }
]
Seven@SevenBig:/mnt/c/Users/Administrator$ curl http://localhost:8080/service/userinfo?userid=1
{
  "UID": 1,
  "UserName": "ooo",
  "DepartName": "1",
  "CreateAt": "2017-11-28T04:31:42+08:00"
}
```