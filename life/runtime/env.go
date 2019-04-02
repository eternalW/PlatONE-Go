package runtime

import (
	"github.com/BCOSnetwork/BCOS-Go/common"
	"github.com/BCOSnetwork/BCOS-Go/core"
	"github.com/BCOSnetwork/BCOS-Go/core/vm"
)

// create env for wasmvm
func NewEnv(cfg *Config) *vm.EVM {
	context := vm.Context{
		CanTransfer: core.CanTransfer,
		Transfer:    core.Transfer,
		GetHash:     func(uint64) common.Hash { return common.Hash{} },

		Origin:      cfg.Origin,
		Coinbase:    cfg.Coinbase,
		BlockNumber: cfg.BlockNumber,
		Time:        cfg.Time,
		Difficulty:  cfg.Difficulty,
		GasLimit:    cfg.GasLimit,
		GasPrice:    cfg.GasPrice,
	}

	return vm.NewEVM(context, cfg.State, cfg.ChainConfig, cfg.EVMConfig)
}
