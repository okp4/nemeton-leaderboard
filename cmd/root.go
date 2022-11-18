package cmd

import (
	"crypto/tls"
	"os"

	"github.com/spf13/cobra"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	grpcAddress   string
	noTLS         bool
	tlsSkipVerify bool
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "nemeton-leaderboard",
	Short: "🧙 Nemeton Leaderboard 🏆",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.PersistentFlags().StringVar(&grpcAddress, FlagGrpcAddress, "127.0.0.1:9090", "The grpc okp4 server url")
	rootCmd.PersistentFlags().BoolVar(&noTLS, FlagNoTLS, false, "No encryption with the GRPC endpoint")
	rootCmd.PersistentFlags().BoolVar(&tlsSkipVerify,
		FlagTLSSkipVerify,
		false,
		"Encryption with the GRPC endpoint but skip certificates verification")

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
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
