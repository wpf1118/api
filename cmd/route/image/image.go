package image

import (
	"github.com/go-chi/chi"
	"github.com/wpf1118/api/cmd/errcode"
	"github.com/wpf1118/toolbox/tools/logging"
	"github.com/wpf1118/toolbox/tools/response"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func Route() func(chi.Router) {
	return func(r chi.Router) {
		r.Post("/upload", upload())
		r.Get("/*", image)
	}
}

func upload() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// todo 判断文件类型为图片

		//设置内存大小
		r.ParseMultipartForm(32 << 20)

		//获取上传文件
		file, handler, err := r.FormFile("file")
		if err != nil {
			response.Error(w, errcode.FromFileError.AddError(err))
			return
		}
		defer file.Close()
		//fmt.Fprintf(w, "%v", handler.Header)

		//创建上传目录
		err = os.Mkdir("/data/images", os.ModePerm)
		if err != nil {
			response.Error(w, errcode.CreateDirError.AddError(err))
			return
		}

		filePath := "/data/images/" + handler.Filename
		if _, err = os.Stat(filePath); err == nil {
			response.Error(w, errcode.FileExists.AddError(err))
			return
		}

		//创建上传文件
		f, err := os.Create(filePath)
		if err != nil {
			response.Error(w, errcode.CreateFileError.AddError(err))
			return
		}
		defer f.Close()
		_, err = io.Copy(f, file)
		if err != nil {
			response.Error(w, errcode.CopyFileError.AddError(err))
			return
		}

		response.Ok(w, "data/images/"+handler.Filename)
	}
}

func image(w http.ResponseWriter, r *http.Request) {
	filesDir := "/data/images"
	filePath := strings.Replace(r.RequestURI, "/api/v1/image/", "/", 1)

	logging.DebugF(filepath.Join(filesDir, filePath))
	if _, err := os.Stat(filepath.Join(filesDir, filePath)); os.IsNotExist(err) {
		response.Error(w, errcode.FileNotExists.Log())
		return
	}

	http.StripPrefix("/api/v1/image/", http.FileServer(http.Dir(filesDir))).ServeHTTP(w, r)
}
