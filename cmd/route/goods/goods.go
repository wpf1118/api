package goods

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/wpf1118/api/cmd/errcode"
	"github.com/wpf1118/api/pkg/service/goods"
	"github.com/wpf1118/api/pkg/service/image"
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
		r.Post("/add", add)
		r.Post("/detail", detail())
	}
}

func goodsList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s := goods.NewGoodsServ()
		var req goods.Req
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			response.Error(w, errcode.ParseParamError)
			return
		}
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

func add(w http.ResponseWriter, r *http.Request) {
	s := goods.NewGoodsServ()
	var reqAdd goods.ReqAdd
	err := json.NewDecoder(r.Body).Decode(&reqAdd)
	if err != nil {
		response.Error(w, errcode.ParseParamError)
		return
	}

	name := reqAdd.Name
	cid := reqAdd.Cid
	pic := reqAdd.Pic
	oPrice := reqAdd.OriginalPrice
	sPrice := reqAdd.SellPrice
	goodsPics := reqAdd.GoodsPics

	if name == "" {
		response.Error(w, errcode.ParseParamRequired.AddF("name"))
		return
	}

	if pic == "" {
		response.Error(w, errcode.ParseParamRequired.AddF("pic"))
		return
	}

	if cid == 0 {
		response.Error(w, errcode.ParseParamRequired.AddF("cid"))
		return
	}

	if oPrice == 0 {
		response.Error(w, errcode.ParseParamRequired.AddF("original_price"))
		return
	}

	if len(goodsPics) == 0 {
		response.Error(w, errcode.ParseParamRequired.AddF("goods_pics"))
		return
	}

	if sPrice == 0 {
		reqAdd.SellPrice = oPrice
	}

	id, err := s.Add(reqAdd)
	if err != nil {
		response.Error(w, errcode.ValidationErrors.AddError(err))
		return
	}

	type Resp struct {
		ID uint `json:"id"`
	}
	response.Ok(w, &Resp{id})

	return
}

func detail() http.HandlerFunc {
	type Req struct {
		ID uint `json:"id"`
	}

	type Resp struct {
		ID            uint     `json:"id" gorm:"primarykey"`
		Name          string   `json:"name" gorm:"size:20;comment:商品名称"`
		Cid           uint     `json:"cid" gorm:"comment:商品分类"`
		Pic           string   `json:"pic" gorm:"size:100;comment:商品主图"`
		Status        int      `json:"status" gorm:"default:0;comment:0 未上架 1 正常 2 下架"`
		OriginalPrice uint     `json:"original_price" gorm:"default:9999999;comment:原价 单位分"`
		SellPrice     uint     `json:"sell_price" gorm:"default:9999999;comment:售卖价 单位分"`
		Stock         uint     `json:"stock" gorm:"default:0;comment:库存"`
		ImagePics     []string `json:"image_pics"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var req Req
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			response.Error(w, errcode.ParseParamError)
			return
		}

		id := req.ID

		if id == 0 {
			response.Error(w, errcode.ParseParamRequired.AddF("id"))
			return
		}

		s := goods.NewGoodsServ()
		goodsDetail, err := s.Detail(id)
		if err != nil {
			response.Error(w, errcode.GoodsDetailError.AddError(err))
			return
		}

		imageS := image.NewImageServ()
		pics, _ := imageS.Pics(id)

		resp := &Resp{
			ID:            goodsDetail.ID,
			Name:          goodsDetail.Name,
			Cid:           goodsDetail.Cid,
			Pic:           goodsDetail.Pic,
			Status:        goodsDetail.Status,
			OriginalPrice: goodsDetail.OriginalPrice,
			SellPrice:     goodsDetail.SellPrice,
			Stock:         goodsDetail.Stock,
			ImagePics:     pics,
		}
		response.Ok(w, resp)
		return
	}
}
