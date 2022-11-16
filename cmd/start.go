package cmd

import (
	"okp4/nemeton-leaderboard/app/system"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

const (
	FlagGraphQLAddress = "address"
)

var graphqlAddr string
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the leaderboard service",
	Run: func(cmd *cobra.Command, args []string) {
		app := system.Bootstrap(graphqlAddr)

		kill := make(chan os.Signal, 0)
		signal.Notify(kill, syscall.SIGINT, syscall.SIGTERM)
		<-kill

		log.Info().Msg("Stopping app")
		if err := app.Stop(); err != nil {
			log.Fatal().Err(err).Msg("Unexpected error when stopping app")
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.PersistentFlags().StringVar(&graphqlAddr, FlagGraphQLAddress, ":8080", "graphql listen address")
}
