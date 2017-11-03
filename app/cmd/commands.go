package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"github.com/spf13/viper"
	"os/exec"
	"riesling-cms-shop/app/routes"
	"riesling-cms-shop/app/utils"
)

/**
 * := Coded with love by Sakib Sami on 3/11/17.
 * := root@sakib.ninja
 * := www.sakib.ninja
 * := Coffee : Dream : Code
 */

var RootCmd = cobra.Command{}
var serverCmd = cobra.Command{
	Use:   "server",
	Short: "Handle application server",
}

func Init() {
	RootCmd.AddCommand(&cobra.Command{
		Use:   "info",
		Short: "Check info of application",
		Run:   info,
	})
	serverCmd.AddCommand(&cobra.Command{
		Use:   "start",
		Short: "Start application server",
		Run:   serverStart,
	})
	serverCmd.AddCommand(&cobra.Command{
		Use:   "stop",
		Short: "Stop application server",
		Run:   serverStop,
	})
	RootCmd.AddCommand(&serverCmd)
}

func info(cmd *cobra.Command, args []string) {
	fmt.Println(viper.GetString("app.name"), ":", viper.GetString("app.version"))
	fmt.Println("Author \t\t\t  : ", viper.GetString("credit.author"))
	fmt.Println("License \t\t  : ", viper.GetString("credit.license"))
}

func serverStart(cmd *cobra.Command, args []string) {
	routes.InitRoutes()
	fmt.Println("Application Started")
	utils.PutAppPID()
	routes.WaitGroup.Wait()
}

func serverStop(cmd *cobra.Command, args []string) {
	osCmd := exec.Command("kill", utils.ReadPID())
	osCmd.Run()
	osCmd.Wait()
	fmt.Println("Application Stoped")
}
