package main

import (
	"fmt"
	queueRW "github.com/saeidraei/go-crawler-clean/implem/redis.queueRW"
	urlValidator "github.com/saeidraei/go-crawler-clean/implem/url.validator"

	"github.com/saeidraei/go-crawler-clean/implem/gin.server"
	"github.com/saeidraei/go-crawler-clean/implem/logrus.logger"
	"github.com/saeidraei/go-crawler-clean/infra"
	"github.com/saeidraei/go-crawler-clean/uc"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Build number and versions injected at compile time, set yours
var (
	Version = "unknown"
	Build   = "unknown"
)

// the command to run the server
var rootCmd = &cobra.Command{
	Use:   "go-crawler-clean",
	Short: "Runs the server",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show build and version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Build: %s\nVersion: %s\n", Build, Version)
	},
}

func main() {
	rootCmd.AddCommand(versionCmd)
	cobra.OnInitialize(infra.CobraInitialization)

	infra.LoggerConfig(rootCmd)
	infra.ServerConfig(rootCmd)
	infra.DatabaseConfig(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		logrus.WithError(err).Fatal()
	}
}

func run() {
	ginServer := infra.NewServer(
		viper.GetInt("server.port"),
		infra.DebugMode,
	)

	routerLogger := logger.NewLogger("TEST",
		viper.GetString("log.level"),
		viper.GetString("log.format"),
	)

	server.NewRouterWithLogger(
		uc.HandlerConstructor{
			Logger:       routerLogger,
			QueueRW:      queueRW.New(),
			UrlValidator: urlValidator.New(),
		}.New(),
		routerLogger,
	).SetRoutes(ginServer.Router)

	ginServer.Start()
}
