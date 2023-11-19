package types

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/forbole/juno/v4/node/remote"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypesv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	govtypesv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/forbole/juno/v4/node/local"

	nodeconfig "github.com/forbole/juno/v4/node/config"

	banksource "github.com/forbole/bdjuno/v4/modules/bank/source"
	remotebanksource "github.com/forbole/bdjuno/v4/modules/bank/source/remote"
	distrsource "github.com/forbole/bdjuno/v4/modules/distribution/source"
	remotedistrsource "github.com/forbole/bdjuno/v4/modules/distribution/source/remote"
	govsource "github.com/forbole/bdjuno/v4/modules/gov/source"
	remotegovsource "github.com/forbole/bdjuno/v4/modules/gov/source/remote"
	mintsource "github.com/forbole/bdjuno/v4/modules/mint/source"
	remotemintsource "github.com/forbole/bdjuno/v4/modules/mint/source/remote"
	slashingsource "github.com/forbole/bdjuno/v4/modules/slashing/source"
	remoteslashingsource "github.com/forbole/bdjuno/v4/modules/slashing/source/remote"
	stakingsource "github.com/forbole/bdjuno/v4/modules/staking/source"
	remotestakingsource "github.com/forbole/bdjuno/v4/modules/staking/source/remote"
)

type Sources struct {
	BankSource     banksource.Source
	DistrSource    distrsource.Source
	GovSource      govsource.Source
	MintSource     mintsource.Source
	SlashingSource slashingsource.Source
	StakingSource  stakingsource.Source
}

func BuildSources(nodeCfg nodeconfig.Config, encodingConfig *params.EncodingConfig) (*Sources, error) {
	switch cfg := nodeCfg.Details.(type) {
	case *remote.Details:
		return buildRemoteSources(cfg)
	case *local.Details:
		return nil, fmt.Errorf("local mode is currently not supported")

	default:
		return nil, fmt.Errorf("invalid configuration type: %T", cfg)
	}
}

func buildRemoteSources(cfg *remote.Details) (*Sources, error) {
	source, err := remote.NewSource(cfg.GRPC)
	if err != nil {
		return nil, fmt.Errorf("error while creating remote source: %s", err)
	}

	return &Sources{
		BankSource:     remotebanksource.NewSource(source, banktypes.NewQueryClient(source.GrpcConn)),
		DistrSource:    remotedistrsource.NewSource(source, distrtypes.NewQueryClient(source.GrpcConn)),
		GovSource:      remotegovsource.NewSource(source, govtypesv1.NewQueryClient(source.GrpcConn), govtypesv1beta1.NewQueryClient(source.GrpcConn)),
		MintSource:     remotemintsource.NewSource(source, minttypes.NewQueryClient(source.GrpcConn)),
		SlashingSource: remoteslashingsource.NewSource(source, slashingtypes.NewQueryClient(source.GrpcConn)),
		StakingSource:  remotestakingsource.NewSource(source, stakingtypes.NewQueryClient(source.GrpcConn)),
	}, nil
}
