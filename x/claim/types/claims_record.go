package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Validate performs a stateless validation of the fields
func (claimRecord ClaimRecord) Validate() error {
	if _, err := sdk.AccAddressFromBech32(claimRecord.Address); err != nil {
		return err
	}

	if !claimRecord.InitialClaimableAmount.IsValid() {
		return fmt.Errorf("initial claimable amount is not valid, %s", claimRecord.InitialClaimableAmount)
	}

	if len(Action_value)-1 != len(claimRecord.ActionCompleted) {
		return fmt.Errorf("action length mismatch, expected %d, got %d", len(Action_value)-1, len(claimRecord.ActionCompleted))
	}

	return nil
}
