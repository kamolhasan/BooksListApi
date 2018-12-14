package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)


var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of the product",
	Long: `Print the version of the product
like in the format of V[0-9,.]+ `,


	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("BookListApi v0.3")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

}
