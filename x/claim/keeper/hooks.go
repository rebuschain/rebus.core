package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/rebuschain/rebus.core/v1/x/claim/types"

	"github.com/ethereum/go-ethereum/core"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

func (k Keeper) isClaimActive(ctx sdk.Context) bool {
	params, _ := k.GetParams(ctx)
	if !params.ClaimEnabled {
		return false
	}
	return true
}

func (k Keeper) AfterProposalVote(ctx sdk.Context, proposalID uint64, voterAddr sdk.AccAddress) {
	if !k.isClaimActive(ctx) {
		return
	}
	claimRecord, err := k.GetClaimRecord(ctx, voterAddr)
	if err != nil {
		panic(err.Error())
	}
	if claimRecord.Address == "" {
		return
	}

	_, err = k.ClaimCoinsForAction(ctx, claimRecord, voterAddr, types.ActionVote)
	if err != nil {
		panic(err.Error())
	}
}

func (k Keeper) AfterDelegationModified(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) {
	if !k.isClaimActive(ctx) {
		return
	}
	claimRecord, err := k.GetClaimRecord(ctx, delAddr)
	if err != nil {
		panic(err.Error())
	}
	if claimRecord.Address == "" {
		return
	}

	_, err = k.ClaimCoinsForAction(ctx, claimRecord, delAddr, types.ActionDelegate)
	if err != nil {
		panic(err.Error())
	}
}

// AfterEVMStateTransition implements the ethermint evm PostTxProcessing hook.
// After a EVM state transition is successfully processed, the claimable amount
// for the users's claims record evm action is claimed and transferred to the
// user address.
func (k Keeper) AfterEVMStateTransition(ctx sdk.Context, msg core.Message, receipt *ethtypes.Receipt) error {
	/*
		fromAddr := sdk.AccAddress(msg.From().Bytes())

		claimRecord, err := k.GetClaimRecord(ctx, fromAddr)
		if err != nil {
			panic(err.Error())
		}

		_, err = k.ClaimCoinsForAction(ctx, claimRecord, fromAddr, types.ActionNftID)
		if err != nil {
			k.Logger(ctx).Error(
				"failed to claim EVM action",
				"address", fromAddr.String(),
				"error", err.Error(),
			)
		}
	*/
	return nil
}

// ________________________________________________________________________________________

// Hooks wrapper struct for slashing keeper
type Hooks struct {
	k Keeper
}

var (
	_ govtypes.GovHooks         = Hooks{}
	_ stakingtypes.StakingHooks = Hooks{}
)

// Return the wrapper struct
func (k Keeper) Hooks() Hooks {
	return Hooks{k}
}

// evm hook
func (h Hooks) PostTxProcessing(ctx sdk.Context, msg core.Message, receipt *ethtypes.Receipt) error {
	return h.k.AfterEVMStateTransition(ctx, msg, receipt)
}

// governance hooks
func (h Hooks) AfterProposalFailedMinDeposit(ctx sdk.Context, proposalID uint64) {
}

func (h Hooks) AfterProposalVotingPeriodEnded(ctx sdk.Context, proposalID uint64) {
}

func (h Hooks) AfterProposalSubmission(ctx sdk.Context, proposalID uint64) {}
func (h Hooks) AfterProposalDeposit(ctx sdk.Context, proposalID uint64, depositorAddr sdk.AccAddress) {
}

func (h Hooks) AfterProposalVote(ctx sdk.Context, proposalID uint64, voterAddr sdk.AccAddress) {
	h.k.AfterProposalVote(ctx, proposalID, voterAddr)
}

func (h Hooks) AfterProposalInactive(ctx sdk.Context, proposalID uint64) {}
func (h Hooks) AfterProposalActive(ctx sdk.Context, proposalID uint64)   {}

// staking hooks
func (h Hooks) AfterValidatorCreated(ctx sdk.Context, valAddr sdk.ValAddress)   {}
func (h Hooks) BeforeValidatorModified(ctx sdk.Context, valAddr sdk.ValAddress) {}
func (h Hooks) AfterValidatorRemoved(ctx sdk.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) {
}

func (h Hooks) AfterValidatorBonded(ctx sdk.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) {
}

func (h Hooks) AfterValidatorBeginUnbonding(ctx sdk.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) {
}

func (h Hooks) BeforeDelegationCreated(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) {
}

func (h Hooks) BeforeDelegationSharesModified(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) {
}

func (h Hooks) BeforeDelegationRemoved(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) {
}

func (h Hooks) AfterDelegationModified(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) {
	h.k.AfterDelegationModified(ctx, delAddr, valAddr)
}
func (h Hooks) BeforeValidatorSlashed(ctx sdk.Context, valAddr sdk.ValAddress, fraction sdk.Dec) {}
