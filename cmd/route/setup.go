package route

import (
	"github.com/go-chi/chi"
	"github.com/wpf1118/api/cmd/route/color"
	"github.com/wpf1118/api/cmd/route/common"
	"github.com/wpf1118/api/cmd/route/image"
	"github.com/wpf1118/api/cmd/route/user"
)

func SetupRouter(r *chi.Mux) {
	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/user", user.Route())
			r.Route("/common", common.Route())
			r.Route("/image", image.Route())
			r.Route("/color", color.Route())
		})
	})
}
