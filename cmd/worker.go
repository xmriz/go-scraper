package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/xmriz/go-scraper/internal/handler"
)

var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Scrape data concurrently",
	RunE: func(cmd *cobra.Command, args []string) error {
		workers, _ := cmd.Flags().GetInt("workers")
		if workers <= 0 {
			return fmt.Errorf("number of workers must be greater than 0")
		}
		appID := os.Getenv("APP_ID")
		if appID == "" {
			return fmt.Errorf("APP_ID not set in environment")
		}
		return handler.ScrapeData(appID, workers)
	},
}

func init() {
	workerCmd.Flags().IntP("workers", "w", 1, "Number of concurrent workers")
}
