package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MinterKey is the key to use for the keeper store.
var MinterKey = []byte{0x00}

const (
	// module name
	ModuleName = "mint"

	// StoreKey is the default store key for mint
	StoreKey = ModuleName

	// QuerierRoute is the querier route for the minting store.
	QuerierRoute = StoreKey

	// Query endpoints supported by the minting querier
	QueryParameters       = "parameters"
	QueryInflation        = "inflation"
	QueryAnnualProvisions = "annual_provisions"

	// Ratio for block reward distribution
	PosRatio = "0.75"
	TrRatio  = "0.11"
	LiqRatio = "0.04"
	EtRatio  = "0.08"
)

var PosRewardProportion sdk.Dec
var TreasuryProportion sdk.Dec
var LiquidityProportion sdk.Dec
var EthicalProportion sdk.Dec

func init() {
	PosRewardProportion, _ = sdk.NewDecFromStr(PosRatio)
	TreasuryProportion, _ = sdk.NewDecFromStr(TrRatio)
	LiquidityProportion, _ = sdk.NewDecFromStr(LiqRatio)
	EthicalProportion, _ = sdk.NewDecFromStr(EtRatio)
}
