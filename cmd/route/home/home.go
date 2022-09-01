package home

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/wpf1118/toolbox/tools/response"
	"net/http"
)

func Route() func(chi.Router) {
	return func(r chi.Router) {
		r.Post("/swiper-list", swiperList())
	}
}

func swiperList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var list []Swiper
		for i := 1; i < 5; i++ {
			list = append(list, Swiper{
				Image:   fmt.Sprintf("https://www.zzrs.xyz/images/swiper/%d.jpeg", i),
				GoodsID: i,
			})
		}

		response.Ok(w, response.List{
			Total: int64(len(list)),
			Page:  1,
			Size:  10,
			List:  list,
		})

		return
	}
}
