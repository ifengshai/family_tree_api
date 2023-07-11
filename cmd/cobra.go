package cmd

import (
	"errors"
	"fmt"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"os"

	"github.com/spf13/cobra"

	"go-admin/cmd/api"
)

var rootCmd = &cobra.Command{
	// 命令的名称
	Use: "go-admin",
	//短介绍
	Short:        "go-admin",
	SilenceUsage: true,
	//长介绍
	Long: `go-admin`,
	Args: func(cmd *cobra.Command, args []string) error {
		fmt.Println(pkg.Green("cobra校验参数逻辑"))
		if len(args) < 1 {
			tip()
			return errors.New(pkg.Red("需要至少一个参数"))
		}
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(pkg.Green("执行cobra.run"))
		tip()
	},
}

func tip() {
	fmt.Println("执行tip")
}

func init() {
	fmt.Println(pkg.Green("执行cobra.init"))
	rootCmd.AddCommand(api.StartCmd)
	//rootCmd.AddCommand(migrate.StartCmd)
	//rootCmd.AddCommand(version.StartCmd)
	//rootCmd.AddCommand(config.StartCmd)
	//rootCmd.AddCommand(app.StartCmd)
}

// Execute : apply commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
