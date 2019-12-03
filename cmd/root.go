package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string

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
	rootCmd.MarkPersistentFlagRequired("config")
}

// initConfig reads in config file
func initConfig() {
	// Проверка файла конфигурации
	if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
		fmt.Println("Файл конфигурации не найден. Будут выставлены значения по умолчанию.")
	}

	viper.SetConfigFile(cfgFile)

	// Значения по умолчанию.
	viper.SetDefault("log_http", map[string]string{
		"level": "info",
		"path":  "http.log",
	})
	viper.SetDefault("log_event", map[string]string{
		"level": "info",
		"path":  "event.log",
	})
	viper.SetDefault("http", map[string]string{
		"host":     "127.0.0.1",
		"port":     "8080",
		"shutdown": "5",
	})

	// Чтения настроек
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
