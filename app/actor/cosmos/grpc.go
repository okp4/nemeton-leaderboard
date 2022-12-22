package cosmos

import (
	"context"

	"okp4/nemeton-leaderboard/app/message"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
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
	case *actor.Stopping:
		if err := client.grpcConn.Close(); err != nil {
			log.Warn().Err(err).Msg("😥 Could not close grpc connection.")
		}
	case *message.GetBlock:
		block, err := client.GetBlock(context.Background(), msg.Height)
		if err != nil {
			log.Panic().Err(err).Msg("❌ Failed request get block.")
		}
		ctx.Respond(&message.GetBlockResponse{
			Block: block,
		})

	case *message.GetLatestBlock:
		block, err := client.GetLatestBlock(context.Background())
		if err != nil {
			log.Panic().Err(err).Msg("❌ Failed request get latest block.")
		}
		ctx.Respond(&message.GetBlockResponse{
			Block: block,
		})

	case *message.GetValidator:
		validator, err := client.GetValidator(context.Background(), msg.Valoper)
		if err != nil {
			log.Panic().Err(err).Msg("❌ Failed request validator.")
		}
		ctx.Respond(&message.GetValidatorResponse{
			Validator: validator,
		})
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

func (client *GrpcClient) GetLatestBlock(context context.Context) (*tmservice.Block, error) {
	serviceClient := tmservice.NewServiceClient(client.grpcConn)

	query, err := serviceClient.GetLatestBlock(context, &tmservice.GetLatestBlockRequest{})
	if err != nil {
		return nil, err
	}

	return query.GetSdkBlock(), nil
}

func (client *GrpcClient) GetValidator(ctx context.Context, addr types.ValAddress) (*stakingtypes.Validator, error) {
	res, err := stakingtypes.NewQueryClient(client.grpcConn).
		Validator(
			ctx,
			&stakingtypes.QueryValidatorRequest{
				ValidatorAddr: addr.String(),
			},
		)
	if err != nil {
		return nil, err
	}

	val := res.GetValidator()
	return &val, nil
}
