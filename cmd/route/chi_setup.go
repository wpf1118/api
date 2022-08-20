package route

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func SetupChiRouter(
	httpLoggerDisabled bool,
	verbose bool,
) http.Handler {
	r := chi.NewRouter()
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		Debug:            verbose,
	})

	cors.Log = &CorsLogger{}

	r.Use(cors.Handler)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.StripSlashes)
	//r.Use(middleware.Compress(5))
	r.Use(middleware.Timeout(60 * time.Second))
	if !httpLoggerDisabled {
		r.Use(middleware.Logger)
	}

	// 路由
	SetupRouter(r)

	return r
}

type CorsLogger struct {
}

func (l *CorsLogger) Printf(s string, v ...interface{}) {
	//logging.DebugF(s, v...)
}
