package console

import (
	"context"
	"github.com/wpf1118/api/cmd/route"
	"github.com/wpf1118/toolbox/tools"
	"github.com/wpf1118/toolbox/tools/flag"
	"github.com/wpf1118/toolbox/tools/logging"
	"net/http"
	"sync"
	"time"
)

// Console represents the Vegeta Console server.
type Console struct {
	tools.Service
	server *http.Server
}

// NewConsole is to create a new Console struct.
func NewConsole(httpOpts *flag.HTTPOpts, verbose bool) (*Console, error) {
	return &Console{
		server: &http.Server{
			Addr:    httpOpts.HTTPListen,
			Handler: route.SetupChiRouter(httpOpts.HTTPLoggerDisabled, verbose),
		},
	}, nil
}

// Run is to run the service.
func (c *Console) Run() func() {
	logging.InfoF("service started")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := c.server.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				logging.ErrorF("error in http.Server.ListenAndServe")
			}
		}
	}()

	return func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := c.server.Shutdown(ctx); err != nil {
			logging.ErrorF("error in shutting down HTTP server")
		}
		wg.Wait()

		logging.InfoF("service stopped")
	}
}
