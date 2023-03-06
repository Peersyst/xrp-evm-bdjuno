package bank

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/forbole/bdjuno/v4/database"
	"github.com/forbole/bdjuno/v4/modules/bank/source"

	junomessages "github.com/forbole/juno/v4/modules/messages"

	"github.com/forbole/juno/v4/modules"
)

var (
	_ modules.Module                   = &Module{}
	_ modules.PeriodicOperationsModule = &Module{}
)

// Module represents the x/bank module
type Module struct {
	cdc codec.Codec
	db  *database.Db

	messageParser junomessages.MessageAddressesParser
	keeper        source.Source
	stakingModule StakingModule
}

// NewModule returns a new Module instance
func NewModule(
	messageParser junomessages.MessageAddressesParser, keeper source.Source, stakingModule StakingModule, cdc codec.Codec, db *database.Db,
) *Module {
	return &Module{
		cdc:           cdc,
		db:            db,
		messageParser: messageParser,
		keeper:        keeper,
		stakingModule: stakingModule,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "bank"
}
