package goods

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/wpf1118/api/cmd/errcode"
	"github.com/wpf1118/api/pkg/service/goods"
	"github.com/wpf1118/toolbox/tools/response"
	"net/http"
)

// err = e.Orm.Debug().Model(&data).
//		Scopes(
//			cDto.MakeCondition(c.GetNeedSearch()),
//			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
//			actions.Permission(data.TableName(), p),
//		).
//		Find(list).Limit(-1).Offset(-1).
//		Count(count).Error

func Route() func(chi.Router) {
	return func(r chi.Router) {
		r.Post("/list", goodsList())
	}
}

func goodsList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s := goods.NewGoodsServ()
		var req goods.Req
		err := json.NewDecoder(r.Body).Decode(&req)
		total, list, err := s.Paginate(req)
		if err != nil {
			response.Error(w, errcode.GetDataError)
			return
		}

		response.Ok(w, response.List{
			Total: total,
			Page:  req.GetPage(),
			Size:  req.GetSize(),
			List:  list,
		})

		return
	}
}
