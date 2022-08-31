package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	DefaultClaimDenom      = "arebus"
	DefaultAirdropDuration = time.Hour * 5840
)

func (params Params) Validate() error {
	if params.AirdropDuration <= 0 {
		return fmt.Errorf("duration of airdrop must be positive: %d", params.AirdropDuration)
	}

	if err := sdk.ValidateDenom(params.ClaimDenom); err != nil {
		return err
	}
	return nil
}
