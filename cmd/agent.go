/*
Copyright © 2022 Mike Messmore <mike@messmore.org>
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
		ConfigureLogger()
		natsUrl := viper.GetString("nats-url")
		nc := agent.NatsConnect(natsUrl)
		agent.Subscribe(nc)
		log.Println("ERROR: Exiting unnaturally")
	},
}

func init() {
	rootCmd.AddCommand(agentCmd)
}
