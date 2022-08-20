module github.com/wpf1118/api

go 1.14

replace github.com/wpf1118/toolbox => ./modules/toolbox

require (
	github.com/clipperhouse/fsnotify v1.1.0 // indirect
	github.com/clipperhouse/gen v4.1.1+incompatible // indirect
	github.com/clipperhouse/slice v0.0.0-20200107170738-a74fc3888fd9 // indirect
	github.com/go-chi/chi v1.5.4
	github.com/go-chi/cors v1.2.1
	github.com/rs/zerolog v1.27.0
	github.com/spf13/cobra v1.5.0
	github.com/spf13/viper v1.12.0
	github.com/wpf1118/toolbox v1.0.0
	go.etcd.io/etcd/client/v3 v3.5.4
	golang.org/x/sys v0.0.0-20220818161305-2296e01440c6 // indirect
	golang.org/x/tools v0.1.12 // indirect
	golang.org/x/xerrors v0.0.0-20220609144429-65e65417b02f // indirect
	gorm.io/gorm v1.23.8
)
