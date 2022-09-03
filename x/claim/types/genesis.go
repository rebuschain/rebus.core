package types

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
)

type Actions []Action

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params: Params{
			AirdropStartTime: time.Time{},
			AirdropDuration:  DefaultAirdropDuration,
			ClaimDenom:       DefaultClaimDenom,
			ClaimEnabled:     false,
		},
		ClaimRecords: []ClaimRecord{},
	}
}

// GetGenesisStateFromAppState returns x/claims GenesisState given raw application
// genesis state.
func GetGenesisStateFromAppState(cdc codec.Codec, appState map[string]json.RawMessage) *GenesisState {
	var genesisState GenesisState

	if appState[ModuleName] != nil {
		cdc.MustUnmarshalJSON(appState[ModuleName], &genesisState)
	}

	return &genesisState
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (genState GenesisState) Validate() error {
	seenClaims := make(map[string]bool)

	for _, claimsRecord := range genState.ClaimRecords {
		if seenClaims[claimsRecord.Address] {
			return fmt.Errorf("duplicated claims record entry %s", claimsRecord.Address)
		}
		if err := claimsRecord.Validate(); err != nil {
			return err
		}
		seenClaims[claimsRecord.Address] = true
	}

	return genState.Params.Validate()
}
