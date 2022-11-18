package cosmos

import (
	"context"

	"okp4/nemeton-leaderboard/app/messages"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type GrpcClient struct {
	grpcConn *grpc.ClientConn
}

func NewGrpcClient(address string, transportCreds credentials.TransportCredentials) (*GrpcClient, error) {
	grpcConn, err := grpc.Dial(
		address,
		grpc.WithTransportCredentials(transportCreds),
	)
	if err != nil {
		return nil, err
	}

	return &GrpcClient{grpcConn: grpcConn}, nil
}

func (client *GrpcClient) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *actor.Started:
		log.Info().Msg("GRPC connection started.")
		break
	case *actor.Stopping:
		if err := client.grpcConn.Close(); err != nil {
			log.Warn().Err(err).Msg("😥 Could not close grpc connection.")
		}
		break
	case *messages.GetBlock:
		goCTX, cancelFunc := context.WithCancel(context.Background())
		defer cancelFunc()

		block, err := client.GetBlock(goCTX, msg.Height)
		if err != nil {
			panic(err)
		}
		ctx.Respond(&messages.GetBlockResponse{
			Block: block,
		})

		log.Info().Fields(block).Msg("Request GetBlock")
	default:
		log.Info().Fields(msg).Msg("No message")
	}
}

func (client *GrpcClient) GetBlock(context context.Context, height int64) (*tmservice.Block, error) {
	serviceClient := tmservice.NewServiceClient(client.grpcConn)

	query, err := serviceClient.GetBlockByHeight(context, &tmservice.GetBlockByHeightRequest{Height: height})
	if err != nil {
		return nil, err
	}

	return query.GetSdkBlock(), nil
}
