package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	//	RootCmd.AddCommand(RootCmd)
}

var RootCmd = &cobra.Command{
	Use:   "sar",
	Short: "SAR is a demonstrator for split data across multiple stores",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
