package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "godemo",
	Long:  "Add long description here. This is godemo for cli",
	Short: "Add short description here",

	Run: func(cmd *cobra.Command, args []string) { fmt.Println("Hello CLI") },
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func init() {
	fmt.Println("inside init")
	cobra.OnInitialize(initConfig)
}
func initConfig() {
	fmt.Println("inside initConfig")
}
