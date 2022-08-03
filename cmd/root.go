package cmd

import (
	"errors"
	"github.com/spf13/cobra"
)

var vers bool

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "demo-api",
	Short: "demo-api 后端API",
	Long:  "demo-api 后端API",
	RunE: func(cmd *cobra.Command, args []string) error {
		if vers {
			//fmt.Println(version.FullVersion())
			return errors.New("no flag find")
		}
		return nil
	},
}

func init() {
	//RootCmd.PersistentFlags().BoolVarP(&vers, "version", "v", false, "print demo-api version")
	//RootCmd.AddCommand(StartCmd)
}
