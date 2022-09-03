package v1

const (
	// UpgradeName is the shared upgrade plan name for mainnet and testnet
	UpgradeName = "v0.0.4"
	// MainnetUpgradeHeight defines the Rebus mainnet block height on which the upgrade will take place
	TestNetUpgradeHeight = (24*60*60)/3 + 1
	// UpgradeInfo defines the binaries that will be used for the upgrade
	UpgradeInfo = `{"binaries": {"linux/x86_64": "https://github.com/rebuschain/rebus.core/releases/download/v0.0.4/rebus_0.0.4_Linux_x86_64.tar.gz"}}`
)
