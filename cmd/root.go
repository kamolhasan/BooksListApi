package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "BookListApi",
	Short: "This is the root command",
	Long: "This is the root command for BookListApi server",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello...!")
	},
}

func Execute() {


	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}