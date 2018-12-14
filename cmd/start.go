package cmd

import (
	"fmt"
	"github.com/kamolhasan/BookListApi/api"
	"github.com/spf13/cobra"
)

var port string
var bypass bool

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "This will start the sever",
	Long: `This will start the server 
at a given port.
-p,--port :insert port number(8000 default)`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start is called")
		api.CreateSever()
		api.SetValue(port,bypass)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.PersistentFlags().StringVarP(&port,"port","p","8000","This flag will set the port")
	startCmd.PersistentFlags().BoolVarP(&bypass,"bypass","b",false,"This flag allows to bypass the authentication")

}
