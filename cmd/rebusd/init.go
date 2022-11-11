package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	cfg "github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/libs/cli"
	tmos "github.com/tendermint/tendermint/libs/os"
	tmrand "github.com/tendermint/tendermint/libs/rand"
	"github.com/tendermint/tendermint/types"

	"github.com/cosmos/go-bip39"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/input"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	"github.com/rebuschain/rebus.core/v1/app"
	rebuscfg "github.com/rebuschain/rebus.core/v1/cmd/config"
)

const (
	// FlagTimeoutCommit defines a flag to set timeout commit of node.
	FlagTimeoutCommit = "timeout-commit"
)

type printInfo struct {
	Moniker    string          `json:"moniker" yaml:"moniker"`
	ChainID    string          `json:"chain_id" yaml:"chain_id"`
	NodeID     string          `json:"node_id" yaml:"node_id"`
	GenTxsDir  string          `json:"gentxs_dir" yaml:"gentxs_dir"`
	AppMessage json.RawMessage `json:"app_message" yaml:"app_message"`
}

func newPrintInfo(moniker, chainID, nodeID, genTxsDir string, appMessage json.RawMessage) printInfo {
	return printInfo{
		Moniker:    moniker,
		ChainID:    chainID,
		NodeID:     nodeID,
		GenTxsDir:  genTxsDir,
		AppMessage: appMessage,
	}
}

func displayInfo(info printInfo) error {
	out, err := json.MarshalIndent(info, "", " ")
	if err != nil {
		return err
	}

	if _, err := fmt.Fprintf(os.Stderr, "%s\n", string(sdk.MustSortJSON(out))); err != nil {
		return err
	}

	return nil
}

// InitCmd returns a command that initializes all files needed for Tendermint
// and the respective application.
func InitCmd(mbm module.BasicManager, defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init [moniker]",
		Short: "Initialize private validator, p2p, genesis, and application configuration files",
		Long:  `Initialize validators's and node's configuration files.`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			cdc := clientCtx.Codec

			serverCtx := server.GetServerContextFromCmd(cmd)
			config := serverCtx.Config
			config.SetRoot(clientCtx.HomeDir)

			timeoutCommit, err := cmd.Flags().GetInt(FlagTimeoutCommit)
			if err != nil {
				return err
			}
			config.Consensus.TimeoutCommit = time.Duration(timeoutCommit) * time.Millisecond

			// Set peers in and out to an 8:1 ratio to prevent choking
			config.P2P.MaxNumInboundPeers = 240
			config.P2P.MaxNumOutboundPeers = 30

			config.Mempool.Size = 10000
			config.StateSync.TrustPeriod = 112 * time.Hour

			config.SetRoot(clientCtx.HomeDir)

			chainID, _ := cmd.Flags().GetString(flags.FlagChainID)
			if chainID == "" {
				chainID = fmt.Sprintf("reb_3333-%v", tmrand.Str(6))
			}

			if strings.Contains(chainID, "reb_3333") {
				seeds := []string{
					"a6d710cd9baac9e95a55525d548850c91f140cd9@3.211.101.169:26656",
					"c296ee829f137cfe020ff293b6fc7d7c3f5eeead@54.157.52.47:26656",
				}
				config.P2P.Seeds = strings.Join(seeds, ",")
			}

			// Get bip39 mnemonic
			var mnemonic string
			recover, _ := cmd.Flags().GetBool(genutilcli.FlagRecover)
			if recover {
				inBuf := bufio.NewReader(cmd.InOrStdin())
				value, err := input.GetString("Enter your bip39 mnemonic", inBuf)
				if err != nil {
					return err
				}

				mnemonic = value
				if !bip39.IsMnemonicValid(mnemonic) {
					return errors.New("invalid mnemonic")
				}
			}

			nodeID, _, err := genutil.InitializeNodeValidatorFilesFromMnemonic(config, mnemonic)
			if err != nil {
				return err
			}

			config.Moniker = args[0]

			genFile := config.GenesisFile()
			overwrite, _ := cmd.Flags().GetBool(genutilcli.FlagOverwrite)

			if !overwrite && tmos.FileExists(genFile) {
				return fmt.Errorf("genesis.json file already exists: %v", genFile)
			}

			appState, err := json.MarshalIndent(
				app.GenesisState(mbm.DefaultGenesis(cdc)).
					ConfigureBondDenom(cdc, rebuscfg.BaseDenom), "", " ",
				// ConfigureICA(cdc), "", " ",
			)
			if err != nil {
				return errors.Wrap(err, "Failed to marshall default genesis state")
			}

			genDoc := &types.GenesisDoc{}
			if _, err := os.Stat(genFile); err != nil {
				if !os.IsNotExist(err) {
					return err
				}
			} else {
				genDoc, err = types.GenesisDocFromFile(genFile)
				if err != nil {
					return errors.Wrap(err, "Failed to read genesis doc from file")
				}
			}

			genDoc.ChainID = chainID
			genDoc.Validators = nil
			genDoc.AppState = appState

			if err := genutil.ExportGenesisFile(genDoc, genFile); err != nil {
				return errors.Wrap(err, "Failed to export gensis file")
			}

			toPrint := newPrintInfo(config.Moniker, chainID, nodeID, "", appState)

			cfg.WriteConfigFile(filepath.Join(config.RootDir, "config", "config.toml"), config)
			return displayInfo(toPrint)
		},
	}

	cmd.Flags().String(cli.HomeFlag, defaultNodeHome, "node's home directory")
	cmd.Flags().BoolP(genutilcli.FlagOverwrite, "o", false, "overwrite the genesis.json file")
	cmd.Flags().Bool(genutilcli.FlagRecover, false, "provide seed phrase to recover existing key instead of creating")
	cmd.Flags().String(flags.FlagChainID, "", "genesis file chain-id, if left blank will be randomly created")
	cmd.Flags().Int(FlagTimeoutCommit, 2000, "timeout commit of the node in ms")

	return cmd
}
