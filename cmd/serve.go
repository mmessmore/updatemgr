/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/mmessmore/updatemgr/srv"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run web interface",
	Long:  `Run web-based interface for updatemgr`,
	//TODO: make this serve stuff
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetInt("port")
		ttl, _ := cmd.Flags().GetInt("ttl")
		purge, _ := cmd.Flags().GetInt("purge")
		natsUrl, _ := cmd.Flags().GetString("nats-url")
		srv.RunServer(port, ttl, purge, natsUrl)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	serveCmd.Flags().IntP("port", "p", 8080, "Listen port for Web Server")
	serveCmd.Flags().IntP("purge", "P", 3,
		"Minutes between host purge intervals")
	serveCmd.Flags().IntP("ttl", "t", 300,
		"Seconds before host is considered offline")
}
