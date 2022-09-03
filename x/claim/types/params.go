package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	DefaultClaimDenom      = "arebus"
	DefaultAirdropDuration = time.Hour * 5840
)

// Parameter store key
var (
	AirdropStartTime = []byte("AirdropStartTime")
	AirdropDuration  = []byte("AirdropDuration")
	ClaimDenom       = []byte("ClaimDenom")
	ClaimEnabled     = []byte("ClaimEnabled")
)

// ParamTable for minting module.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(AirdropStartTime, &p.AirdropStartTime, validateStartTime),
		paramtypes.NewParamSetPair(AirdropDuration, &p.AirdropDuration, validateDuration),
		paramtypes.NewParamSetPair(ClaimDenom, &p.ClaimDenom, validateDenom),
		paramtypes.NewParamSetPair(ClaimEnabled, &p.ClaimEnabled, validateBool),
	}
}

func (params Params) Validate() error {
	if params.AirdropDuration <= 0 {
		return fmt.Errorf("duration of airdrop must be positive: %d", params.AirdropDuration)
	}

	if err := sdk.ValidateDenom(params.ClaimDenom); err != nil {
		return err
	}
	return validateBool(params.ClaimEnabled)
}

func validateStartTime(i interface{}) error {
	return nil
}

func validateDuration(i interface{}) error {
	v, ok := i.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if v <= 0 {
		return fmt.Errorf("duration must be positive: %s", v)
	}
	return nil
}

func validateDenom(i interface{}) error {
	denom, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return sdk.ValidateDenom(denom)
}

func validateBool(i interface{}) error {
	_, ok := i.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}
