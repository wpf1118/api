package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wpf1118/api/cmd/console"
	"github.com/wpf1118/toolbox/tools"
	"github.com/wpf1118/toolbox/tools/db"
	"github.com/wpf1118/toolbox/tools/flag"
	"github.com/wpf1118/toolbox/tools/logging"
	"os"
	"strings"
)

var (
	verbose bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "console",
	Short: "The centralized server",
	Long:  `The centralized server`,
	RunE: func(cmd *cobra.Command, args []string) error {
		_init()

		logging.DebugF("test")

		httpOpts := flag.GetHTTPOpts(cmd)
		logging.DebugF("listen: %s", httpOpts.HTTPListen)

		console, err := console.NewConsole(httpOpts, verbose)
		if err != nil {
			return err
		}

		tools.NewApplication(
			console,
		).Run()

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func _init() {
	// 开启debug日志
	if verbose {
		logging.SetVerbose()
	}

	mysqlOpts := flag.NewDefaultMysqlOpts()
	db.MysqlInit(mysqlOpts)

	redisOpts := flag.NewDefaultRedisOpts()
	db.RedisInit(redisOpts)
}

func init() {
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "", false, "日志DEBUG模式")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()

	flag.AddHTTPFlags(rootCmd)

	//rootCmd.PersistentFlags().StringVarP(&mode, "mode", "m", "k8s", "run mode[k8s|file]")
}
