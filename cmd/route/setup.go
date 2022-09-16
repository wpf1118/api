package route

import (
	"github.com/go-chi/chi"
	"github.com/wpf1118/api/cmd/route/category"
	"github.com/wpf1118/api/cmd/route/color"
	"github.com/wpf1118/api/cmd/route/common"
	"github.com/wpf1118/api/cmd/route/goods"
	"github.com/wpf1118/api/cmd/route/home"
	"github.com/wpf1118/api/cmd/route/image"
	"github.com/wpf1118/api/cmd/route/user"
	"github.com/wpf1118/toolbox/tools/response"
	"net/http"
)

func SetupRouter(r *chi.Mux) {
	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Post("/ping", ping)
			r.Get("/ping", ping)
			r.Route("/user", user.Route())
			r.Route("/common", common.Route())
			r.Route("/image", image.Route())
			r.Route("/color", color.Route())
			r.Route("/home", home.Route())
			r.Route("/category", category.Route())
			r.Route("/goods", goods.Route())
		})
	})
}

func ping(w http.ResponseWriter, r *http.Request) {
	response.Ok(w, "pong")
}
