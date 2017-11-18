package main

import (
	"net/http"
  "github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"fmt"
)

func initRoute(m *martini.ClassicMartini) {
	m.Use(render.Renderer())
  m.Get("/", func(r render.Render) {
    r.HTML(200, "index", nil)
	})
	m.Get("/api/test", func(r render.Render) {
		r.JSON(200, map[string]interface{}{
			"test":"That's good",
		})
	})
	m.Post("/update/userinfo", func(r render.Render, req *http.Request) {
		fmt.Println(req.FormValue("firstname"))
		r.JSON(200, map[string]interface{}{
			"firstName": req.FormValue("firstname"),
			"lastName" : req.FormValue("lastname"),
			"userName" : req.FormValue("username"),
		})
	})
	m.NotFound(func(r render.Render, req *http.Request)  {
		fmt.Println("[martini] Page Not Found[500]")
		r.HTML(500, "notfound", req.URL.Path)
	})
}

func main() {
  m := martini.Classic()
	// render html templates from templates directory
	initRoute(m)
	m.Use(martini.Static("assets"))
	m.Run()
}