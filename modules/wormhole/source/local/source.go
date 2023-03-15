package local

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/forbole/juno/v4/node/local"

	wormholesource "github.com/forbole/bdjuno/v3/modules/wormhole/source"
	wormholetypes "github.com/wormhole-foundation/wormchain/x/wormhole/types"
)

var (
	_ wormholesource.Source = &Source{}
)

// Source implements wormholesource.Source using a local node
type Source struct {
	*local.Source
	querier wormholetypes.QueryServer
}

// NewSource implements a new Source instance
func NewSource(source *local.Source, querier wormholetypes.QueryServer) *Source {
	return &Source{
		Source:  source,
		querier: querier,
	}
}

// GetGuardianSetAll implements wormholesource.Source
func (s Source) GetGuardianSetAll(height int64) ([]wormholetypes.GuardianSet, error) {
	ctx, err := s.LoadHeight(height)
	if err != nil {
		return nil, fmt.Errorf("error while loading height: %s", err)
	}

	var guardianSet []wormholetypes.GuardianSet
	var nextKey []byte
	var stop = false
	for !stop {
		res, err := s.querier.GuardianSetAll(
			sdk.WrapSDKContext(ctx),
			&wormholetypes.QueryAllGuardianSetRequest{
				Pagination: &query.PageRequest{
					Key:   nextKey,
					Limit: 100, // Query 100 guardians set at once
				},
			},
		)
		if err != nil {
			return nil, err
		}

		nextKey = res.Pagination.NextKey
		stop = len(res.Pagination.NextKey) == 0
		guardianSet = append(guardianSet, res.GuardianSet...)
	}

	return guardianSet, nil
}

// GetGuardianValidatorAll implements wormholesource.Source
func (s Source) GetGuardianValidatorAll(height int64) ([]wormholetypes.GuardianValidator, error) {
	ctx, err := s.LoadHeight(height)
	if err != nil {
		return nil, fmt.Errorf("error while loading height: %s", err)
	}

	var guardianValidatorList []wormholetypes.GuardianValidator
	var nextKey []byte
	var stop = false
	for !stop {
		res, err := s.querier.GuardianValidatorAll(
			sdk.WrapSDKContext(ctx),
			&wormholetypes.QueryAllGuardianValidatorRequest{
				Pagination: &query.PageRequest{
					Key:   nextKey,
					Limit: 100, // Query 100 guardians set at once
				},
			},
		)
		if err != nil {
			return nil, err
		}

		nextKey = res.Pagination.NextKey
		stop = len(res.Pagination.NextKey) == 0
		guardianValidatorList = append(guardianValidatorList, res.GuardianValidator...)
	}

	return guardianValidatorList, nil
}
