/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"

	"github.com/mmessmore/updatemgr/agent"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// agentCmd represents the agent command
var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "Host agent for update management",
	Long:  `Agent that manages updates on host`,
	Run: func(cmd *cobra.Command, args []string) {
		natsUrl := viper.GetString("nats-url")
		if natsUrl == "" {
			natsUrl, _ = rootCmd.Flags().GetString("nats-url")
		}
		nc := agent.NatsConnect(natsUrl)
		agent.Subscribe(nc)
		log.Println("ERROR: Exiting unnaturally")
	},
}

func init() {
	rootCmd.AddCommand(agentCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// agentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// agentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
