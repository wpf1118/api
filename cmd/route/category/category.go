package category

import (
	"github.com/go-chi/chi"
	"github.com/wpf1118/api/cmd/errcode"
	"github.com/wpf1118/api/pkg/service/category"
	"github.com/wpf1118/toolbox/tools/response"
	"net/http"
)

func Route() func(chi.Router) {
	return func(r chi.Router) {
		r.Post("/tree-list", treeList())
	}
}

func treeList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s := category.NewCategoryServ()
		allList, err := s.AllList()
		if err != nil {
			response.Error(w, errcode.GetDataError)
			return
		}

		treeList := s.TreeList(allList, 0)

		response.Ok(w, response.List{
			Total: int64(len(treeList)),
			List:  treeList,
		})

		return
	}
}
