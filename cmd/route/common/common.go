package common

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/wpf1118/api/cmd/errcode"
	"github.com/wpf1118/toolbox/tools/db"
	"github.com/wpf1118/toolbox/tools/help"
	"github.com/wpf1118/toolbox/tools/response"
	"gorm.io/gorm"
	"net/http"
)

func Route() func(chi.Router) {
	return func(r chi.Router) {
		r.Post("/set", set())
		r.Post("/get", get())
	}
}

func set() http.HandlerFunc {
	type Req struct {
		Key   string      `json:"key"`
		Value interface{} `json:"value"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &Req{}
		err := json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			response.Error(w, errcode.ParseParamError)
			return
		}

		key := req.Key
		if key == "" {
			response.Error(w, errcode.ParseParamRequired.AddF("key"))
			return
		}
		value := help.InterfaceToJson(req.Value)
		if value == "" {
			response.Error(w, errcode.ParseParamRequired.AddF("value"))
			return
		}

		ctx := context.Background()
		redis := db.NewRedis()
		_, err = redis.Set(ctx, key, value)
		if err != nil {
			response.Error(w, errcode.CommonSetError.Log())
			return
		}

		type Kv struct {
			ID        uint           `json:"id" gorm:"primarykey"`
			Key       string         `json:"key"`
			Value     string         `json:"value"`
			CreatedAt int64          `json:"created_at" gorm:"autoCreateTime"`
			UpdatedAt int64          `json:"updated_at" gorm:"autoUpdateTime"`
			DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
		}

		entity := &Kv{
			Key:   key,
			Value: value,
		}
		mysql := db.NewMysql()
		err = mysql.DB.Create(entity).Error
		if err != nil {
			response.Error(w, errcode.CommonSetError.Log())
			return
		}

		response.Ok(w, "success")
	}
}

func get() http.HandlerFunc {
	type Req struct {
		Key string `json:"key"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &Req{}
		err := json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			response.Error(w, errcode.ParseParamError)
			return
		}

		key := req.Key
		if key == "" {
			response.Error(w, errcode.ParseParamRequired.AddF("key"))
			return
		}

		ctx := context.Background()
		redis := db.NewRedis()
		v, err := redis.Get(ctx, key)
		if err != nil {
			response.Error(w, errcode.DataNotExists)
			return
		}

		var res interface{}
		json.Unmarshal([]byte(v), &res)

		response.Ok(w, res)
	}
}
