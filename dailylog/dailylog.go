package dailylog

import (
	"net/http"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"

	"appengine"
	"appengine/user"
)

func init() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))

	m.Get("/", func(r render.Render) {
		r.HTML(200, "index", "This is The Daily Log")
	})

	m.Get("/log", func(r render.Render, w http.ResponseWriter, req *http.Request) {
		c := appengine.NewContext(req)
		u := user.Current(c)
		if u == nil {
			url, err := user.LoginURL(c, req.URL.String())
			if err != nil {
				w.WriteHeader(500)
			} else {
				w.Header().Set("Location", url)
				w.WriteHeader(http.StatusFound)
			}
		} else {
			r.HTML(200, "log", "foo")
		}
	})

	//m.Get("/login", func(r render.Render) {
	//	r.HTML(200, "login", "")
	//})


	http.Handle("/", m)
}

