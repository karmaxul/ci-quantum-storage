package healchain

import (
    "fmt"
    "log"

    "github.com/ethereum/go-ethereum/core/vm"
)

func RegisterWithEVM(evm *vm.EVM) error {
    pre, err := NewPrecompile(10, 4)
    if err != nil {
        return err
    }

    evm.SetPrecompile(pre.Address(), pre)
    fmt.Printf("✅ HealChain Self-Healing Precompile registered at %s\n", HealChainPrecompileAddress)
    return nil
}
