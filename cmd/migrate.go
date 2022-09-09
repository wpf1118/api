package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wpf1118/api/pkg/entity"
	"github.com/wpf1118/toolbox/tools/db"
	"github.com/wpf1118/toolbox/tools/flag"
	"github.com/wpf1118/toolbox/tools/logging"
)

// migrateCmd
var migrateCmd = &cobra.Command{
	Use: "migrate",
	RunE: func(cmd *cobra.Command, args []string) error {
		migrateInit()

		// 数据表 初始化
		mysqlClient := db.NewMysql()
		mysqlClient.Set("gorm:table_options", "ENGINE=InnoDB")
		mysqlClient.AutoMigrate(&entity.Category{})
		mysqlClient.AutoMigrate(&entity.Kv{})

		logging.DebugF("migrate success")
		return nil
	},
}

func migrateInit() {
	// 开启debug日志
	if verbose {
		logging.SetVerbose()
	}

	mysqlOpts := flag.NewDefaultMysqlOpts()
	db.MysqlInit(mysqlOpts)
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
