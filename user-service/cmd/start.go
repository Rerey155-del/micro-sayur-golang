package cmd

import (
	"user-service/config"
	"user-service/internal/adapter/repository"
	"user-service/internal/app"
	"user-service/internal/core/service"

	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start API server",
	Long:  `Start API server`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.NewConfig()
		db, err := config.ConnectionPostgres(cfg)
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		userRepo := repository.NewUserRepository(db.DB)
		userService := service.NewUserService(userRepo)

		app.RunServer(userService)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
