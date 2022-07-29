package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	ethermint "github.com/evmos/ethermint/types"
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
	PosRatio = "0.72"
	TrRatio  = "0.11"
	IncRatio = "0.04"
	EtRatio  = "0.08"
	ComRatio = "0.05"
)

var TreasuryAddressList = [1]string{"rebus1kr9etd0k4se82regqrxkhjvpstl2u9yd2cptp4"}
var EthicalAddressList = [1]string{"rebus1l4l243c8mee7kh6spawlz6zzad745j2vlrdmdq"}
var IncentiveAddressList = [1]string{"rebus1hzm429073g3mumg2dt3v67sfsnnmshq6nrufy2"}

var PosRewardProportion sdk.Dec
var TreasuryProportion sdk.Dec
var IncentiveProportion sdk.Dec
var EthicalProportion sdk.Dec
var CommunityProportion sdk.Dec

var PosAllocation = sdk.NewInt(585_000_000).Mul(ethermint.PowerReduction)

func init() {
	PosRewardProportion, _ = sdk.NewDecFromStr(PosRatio)
	TreasuryProportion, _ = sdk.NewDecFromStr(TrRatio)
	IncentiveProportion, _ = sdk.NewDecFromStr(IncRatio)
	EthicalProportion, _ = sdk.NewDecFromStr(EtRatio)
	CommunityProportion, _ = sdk.NewDecFromStr(ComRatio)
}
