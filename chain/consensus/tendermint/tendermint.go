package tendermint

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"golang.org/x/net/context"
)

type TMC struct {
}

// Implement functions below first.
//func NewTMC(sm *stmgr.StateManager, beacon beacon.Schedule, verifier storiface.Verifier, genesis chain.Genesis) consensus.Consensus {
//
//}

func (tmc *TMC) ValidateBlock(ctx context.Context, b *types.FullBlock) (err error) {
	return nil
}

func (tmc *TMC) ValidateBlockPubsub(ctx context.Context, self bool, msg *pubsub.Message) (pubsub.ValidationResult, string) {
	return pubsub.ValidationAccept, ""
}

func (tmc *TMC) IsEpochBeyondCurrMax(epoch abi.ChainEpoch) bool {
	return false
}

func (tmc *TMC) CreateBlock(ctx context.Context, w api.Wallet, bt *api.BlockTemplate) (*types.FullBlock, error) {
	return &types.FullBlock{}, nil
}
