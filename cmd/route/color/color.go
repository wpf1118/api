package color

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/wpf1118/api/cmd/errcode"
	"github.com/wpf1118/toolbox/tools/help"
	"github.com/wpf1118/toolbox/tools/request"
	"github.com/wpf1118/toolbox/tools/response"
	"net/http"
)

func Route() func(chi.Router) {
	return func(r chi.Router) {
		r.Post("/list", colorList())
	}
}

func colorList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		listReq := request.NewListReq()
		err := listReq.Parse(r)
		if err != nil {
			response.Error(w, errcode.ParseParamError)
			return
		}

		page := listReq.Page
		size := listReq.Size

		if page > 5 {
			response.Ok(w, response.List{
				Total: 50,
				Page:  page,
				Size:  size,
				List:  nil,
			})

			return
		}

		var list []Color
		var i int
		max := 255
		for i = 0; i < size; i++ {
			id := (page-1)*size + i + 1
			list = append(list, Color{
				ID:  int64(id),
				Rgb: fmt.Sprintf("%d,%d,%d", help.RandLt(max), help.RandLt(max), help.RandLt(max)),
			})
		}

		response.Ok(w, response.List{
			Total: 50,
			Page:  page,
			Size:  size,
			List:  list,
		})

		return
	}
}
