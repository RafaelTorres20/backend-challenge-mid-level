package cmd

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/RafaelTorres20/backend-challenge-mid-level/pkg/api"
	"github.com/RafaelTorres20/backend-challenge-mid-level/pkg/domain/assets"
	"github.com/RafaelTorres20/backend-challenge-mid-level/pkg/domain/users"
	"github.com/RafaelTorres20/backend-challenge-mid-level/pkg/gateways"
	"github.com/RafaelTorres20/backend-challenge-mid-level/pkg/gateways/postgres"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Serves user services",
	RunE: func(cmd *cobra.Command, args []string) error {

		db := newPostgresqlConection("postgres://postgres:postgres@localhost:5432/assets?sslmode=disable")
		defer db.Close()
		err := db.Ping()
		if err != nil {
			log.Fatal(err)
		}
		apiKey := os.Getenv("KEY")
		assetsRepository := postgres.NewAssetsRepository(db)
		usersRepository := postgres.NewUsersRepository(db)
		yahooClient := gateways.NewYahooClient(http.DefaultClient, "https://yfapi.net/v6/finance/quote", apiKey)
		assetsService := assets.NewAssetService(assetsRepository, yahooClient)
		usersService := users.NewUsersService(usersRepository)
		svc := api.NewServer(assetsService, usersService)
		go svc.Serve(8080)
		log.Println("Running port 8080")

		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

		<-stop

		log.Println("Gracefully shutting down...")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
