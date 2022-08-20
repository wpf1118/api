package user

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/wpf1118/api/cmd/errcode"
	"github.com/wpf1118/toolbox/tools/response"
	"net/http"
)

func Route() func(chi.Router) {
	return func(r chi.Router) {
		r.Post("/detail", userDetail())
	}
}

type User struct {
	ID        int64  `json:"id"` //全局ID
	Name      string `json:"name"`
	NickName  string `json:"nickName"`
	Sex       int    `json:"sex"`
	Age       int    `json:"age"`
	UpdatedAt int64  `json:"updatedAt"` //更新时间
	CreatedAt int64  `json:"createdAt"` //创建时间
}

func userDetail() http.HandlerFunc {
	type Req struct {
		ID int64 `json:"id"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &Req{}
		err := json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			response.Error(w, errcode.ParseParamError)
			return
		}

		if req.ID <= 0 {
			response.Error(w, errcode.ParseParamInvalid.AddF("ID"))
			return
		}

		userMap := mockUserList()
		if userInfo, ok := userMap[req.ID]; ok {
			response.Ok(w, userInfo)
			return
		}

		response.Error(w, errcode.UserNotExists)
	}
}

func mockUserList() map[int64]*User {
	userMap := make(map[int64]*User)

	userMap[1] = &User{
		ID:        1,
		Name:      "王鹏飞",
		NickName:  "wpf",
		Sex:       1,
		Age:       29,
		UpdatedAt: 1660819712,
		CreatedAt: 0,
	}

	userMap[2] = &User{
		ID:        1,
		Name:      "吴老师",
		NickName:  "wmf",
		Sex:       2,
		Age:       19,
		UpdatedAt: 1660820108,
		CreatedAt: 0,
	}

	return userMap
}
