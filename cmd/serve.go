/*
Copyright © 2022 Mike Messmore <mike@messmore.org>

*/
package cmd

import (
	"github.com/mmessmore/updatemgr/srv"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run Server",
	Long:  `Run server with web-based interface for updatemgr`,
	//TODO: Fix viper + cobra to not be backwards
	Run: func(cmd *cobra.Command, args []string) {
		ConfigureLogger()
		sc := ConfigServer()
		sc.RunServer()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().IntP("port", "p", 1138, "Listen port for Web Server")
	serveCmd.Flags().IntP("purge", "P", 5,
		"Minutes between host purge intervals")
	serveCmd.Flags().IntP("ttl", "t", 300,
		"Seconds before host is considered offline")
	serveCmd.Flags().IntP("refresh", "r", 30,
		"Seconds between refreshing host data")
	viper.BindPFlags(serveCmd.Flags())
}

func ConfigServer() *srv.ServerConfig {
	return &srv.ServerConfig{
		NatsUrl: viper.GetString("nats-url"),
		Port:    viper.GetInt("port"),
		Purge:   viper.GetInt("purge"),
		Refresh: viper.GetInt("refresh"),
		Ttl:     viper.GetInt("ttl"),
	}
}
