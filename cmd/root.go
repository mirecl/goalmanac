package cmd

import (
	"os"

	"github.com/mirecl/goalmanac/internal/adapters"
	"github.com/spf13/cobra"
)

var cfgFile string
var cfg adapters.Config

var rootCmd = &cobra.Command{
	Use:  "goalmanac",
	Long: "CLI для запуска Сервиса - Альманах",
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file")
}

// initConfig reads in config file
func initConfig() {
	err := adapters.CreateConfig(cfgFile, &cfg)
	if err != nil {
		os.Exit(0)
	}
}
