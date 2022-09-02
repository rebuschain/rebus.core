package types

import (
	"errors"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/rebuschain/rebus.core/v1/cmd/config"
)

// Parameter store keys
var (
	KeyMintDenom     = []byte("MintDenom")
	KeyBlocksPerYear = []byte("BlocksPerYear")
	KeyMintEnabled   = []byte("MintEnabled")
)

// ParamTable for minting module.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(
	mintDenom string, blocksPerYear uint64, mintEnabled bool,
) Params {
	return Params{
		MintDenom:     mintDenom,
		BlocksPerYear: blocksPerYear,
		MintEnabled:   mintEnabled,
	}
}

// default minting module parameters
func DefaultParams() Params {
	return Params{
		MintDenom:     config.BaseDenom,
		BlocksPerYear: uint64(60 * 60 * 8766 / 5), // assuming 5 second block times
		MintEnabled:   false,
	}
}

// validate params
func (p Params) Validate() error {
	if err := validateMintDenom(p.MintDenom); err != nil {
		return err
	}
	if err := validateBlocksPerYear(p.BlocksPerYear); err != nil {
		return err
	}

	return validateBool(p.MintEnabled)
}

// String implements the Stringer interface.
/*
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}*/

// Implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMintDenom, &p.MintDenom, validateMintDenom),
		paramtypes.NewParamSetPair(KeyBlocksPerYear, &p.BlocksPerYear, validateBlocksPerYear),
		paramtypes.NewParamSetPair(KeyMintEnabled, &p.MintEnabled, validateBool),
	}
}

func validateMintDenom(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if strings.TrimSpace(v) == "" {
		return errors.New("mint denom cannot be blank")
	}
	if err := sdk.ValidateDenom(v); err != nil {
		return err
	}

	return nil
}

func validateBlocksPerYear(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("blocks per year must be positive: %d", v)
	}

	return nil
}

func validateBool(i interface{}) error {
	_, ok := i.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}
