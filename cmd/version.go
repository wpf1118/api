package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wpf1118/toolbox/tools/help"
	"time"
)

var (
	BuildTime string
	PWD       string
	BRANCH    string
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print out the software version",
	Long:  `Print out the software version`,
	Run: func(cmd *cobra.Command, args []string) {
		printVersion()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func printVersion() {
	fmt.Printf("应用名称: api\n")
	timeUnix := help.StrToInt64(BuildTime)
	timestamp := time.Unix(timeUnix, 0)
	fmt.Printf("构建时间: %s\n", timestamp.Format("2006-01-02 15:04:05"))
	fmt.Printf("构建位置: %s\n", PWD)
	fmt.Printf("构建分支: %s\n", BRANCH)
}
