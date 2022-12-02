package cmd

import (
	"crypto/tls"
	"os"
	"os/signal"
	"syscall"

	"okp4/nemeton-leaderboard/app/system"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	FlagGraphQLAddress = "address"
	FlagMongoURI       = "mongo-uri"
	FlagDatabaseName   = "db"
	FlagGrpcAddress    = "grpc-address"
	FlagNoTLS          = "no-tls"
	FlagTLSSkipVerify  = "tls-skip-verify"
	FlagTwitterToken   = "twitter-token" // nolint:gosec
	FlagTwitterAccount = "twitter-account"
)

var (
	graphqlAddr    string
	mongoURI       string
	dbName         string
	grpcAddress    string
	noTLS          bool
	tlsSkipVerify  bool
	twitterToken   string
	twitterAccount string
	accessToken    string

	startCmd = &cobra.Command{
		Use:   "start",
		Short: "Start the leaderboard service",
		Run: func(cmd *cobra.Command, args []string) {
			var accessTokenOpt *string
			if len(accessToken) > 0 {
				accessTokenOpt = &accessToken
			}

			app := system.Bootstrap(
				graphqlAddr,
				mongoURI,
				dbName,
				grpcAddress,
				twitterToken,
				twitterAccount,
				getTransportCredentials(),
				accessTokenOpt,
			)

			kill := make(chan os.Signal, 1)
			signal.Notify(kill, syscall.SIGINT, syscall.SIGTERM)
			<-kill

			log.Info().Msg("🛑 Stopping app...")
			if err := app.Stop(); err != nil {
				log.Fatal().Err(err).Msg("❌ Unexpected error when stopping app")
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.PersistentFlags().StringVar(&graphqlAddr, FlagGraphQLAddress, ":8080", "GraphQL listen address")
	startCmd.PersistentFlags().StringVar(&mongoURI, FlagMongoURI, "mongodb://localhost:27017", "MongoDB connection string")
	startCmd.PersistentFlags().StringVar(&dbName, FlagDatabaseName, "nemeton", "Database name")
	startCmd.PersistentFlags().StringVar(&grpcAddress, FlagGrpcAddress, "127.0.0.1:9090", "The grpc okp4 server url")
	startCmd.PersistentFlags().BoolVar(&noTLS, FlagNoTLS, false, "No encryption with the GRPC endpoint")
	startCmd.PersistentFlags().BoolVar(&tlsSkipVerify,
		FlagTLSSkipVerify,
		false,
		"Encryption with the GRPC endpoint but skip certificates verification")
	startCmd.PersistentFlags().StringVar(&twitterToken, FlagTwitterToken, "", "Set the twitter bearer token")
	startCmd.PersistentFlags().StringVar(&twitterAccount,
		FlagTwitterAccount,
		"@OKP4_Protocol",
		"Set the twitter account that should be mentioned on tweet to be accepted for twitter task")
	startCmd.PersistentFlags().StringVar(
		&accessToken,
		"access-token",
		"",
		"The required access token for authenticated operations, an empty value = no auth")
}

func getTransportCredentials() credentials.TransportCredentials {
	switch {
	case noTLS:
		return insecure.NewCredentials()
	case tlsSkipVerify:
		return credentials.NewTLS(&tls.Config{InsecureSkipVerify: true}) // #nosec G402 : skip lint since it's an optional flag
	default:
		return credentials.NewTLS(&tls.Config{MinVersion: tls.VersionTLS12})
	}
}
