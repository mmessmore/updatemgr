/*
Copyright © 2022 Mike Messmore <mike@messmore.org>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "updatemgr",
		Short: "Service for Cluster Upgrade Manager",
		Long: `System upgrade management for clusters.  Made for Raspberry Pi OS,
but would work on any Debian/apt based system`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "",
		"config file (default is $HOME/.updatemgr.yaml)")
	rootCmd.PersistentFlags().StringP("nats-url", "n",
		"nats://localhost:4222", "URL for NATS")
	rootCmd.PersistentFlags().StringP("log-level", "l",
		"INFO", "Miniumum Logging Level")
	rootCmd.PersistentFlags().StringP("log-file", "f",
		"-", "Log file, '-' for STDERR")
	rootCmd.PersistentFlags().IntP("log-backups", "b",
		5, "Number of log files to keep if using file")
	rootCmd.PersistentFlags().IntP("log-max-size", "m",
		10, "Maximum size (in MB) of log files if using file")
	rootCmd.PersistentFlags().IntP("log-max-age", "a",
		5, "Maximum age of log files if using file")
	rootCmd.PersistentFlags().BoolP("log-fancy", "F",
		false, "Do pretty console logging instead of JSON. Disables file logging")
	viper.BindPFlags(rootCmd.PersistentFlags())
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".updatemgr")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
