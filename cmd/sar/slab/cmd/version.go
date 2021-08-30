package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of SAR",
	Long:  `All software has versions. This is SAR's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("SAR v0.1 -- HEAD")
	},
}
