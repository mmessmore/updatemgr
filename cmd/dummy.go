/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/mmessmore/updatemgr/dummy"
	"github.com/spf13/cobra"
)

// dummyCmd represents the dummy command
var dummyCmd = &cobra.Command{
	Use:   "dummy",
	Short: "Host dummy for update management",
	Long:  `Agent that manages updates on host`,
	Run: func(cmd *cobra.Command, args []string) {
		natsUrl, _ := cmd.Flags().GetString("nats-url")
		dummy.NatsConnect(natsUrl)
	},
}

func init() {
	rootCmd.AddCommand(dummyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dummyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dummyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
