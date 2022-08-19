package route

import (
	"github.com/wpf1118/api/cmd/route/user"
	"github.com/wpf1118/toolbox/tools/logging"
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
	if !httpLoggerDisabled && verbose {
		r.Use(middleware.Logger)
	}

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/user", user.Route())
		})
	})

	return r
}

type CorsLogger struct {
}

func (l *CorsLogger) Printf(s string, v ...interface{}) {
	logging.DebugF(s, v...)
}
