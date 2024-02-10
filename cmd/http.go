/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/franthescomarchesi/ports_and_adapters/adapters/web/server"
	"github.com/spf13/cobra"
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		server := server.MakeNewWebserver()
		server.Service = productService
		fmt.Println("Webserver has been started")
		server.Serve()
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
}
