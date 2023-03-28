package tendermint

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/types"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	tmtypes "github.com/tendermint/tendermint/types"
)

// These functions transfer types from tendermint packages to filecoin packages

func BlockByBlock(blk *tmtypes.Block) (blocks.Block, error) {
	return blocks.NewBlock(blk.Data.Hash()), nil
}

func BlockHeaderByBlock(blk *tmtypes.Block) (*types.BlockHeader, error) {
	fakeMiner, err := address.NewFromString("lotus-tm-test-address")
	if err != nil {
		return nil, err
	}

	ret := &types.BlockHeader{
		Miner:                 fakeMiner,
		Ticket:                nil,
		ElectionProof:         nil,
		BeaconEntries:         nil,
		WinPoStProof:          nil,
		Parents:               nil,
		ParentWeight:          types.NewInt(0),
		Height:                abi.ChainEpoch(blk.Height),
		ParentStateRoot:       cid.Cid{},
		ParentMessageReceipts: cid.Cid{},
		Messages:              cid.Cid{},
		BLSAggregate:          nil,
		Timestamp:             0,
		BlockSig:              nil,
		ForkSignaling:         0,
		ParentBaseFee:         abi.TokenAmount{},
	}

	return ret, nil
}

func TipsetByBlock(blk *tmtypes.Block) (*types.TipSet, error) {
	header, err := BlockHeaderByBlock(blk)
	if err != nil {
		return nil, err
	}

	return types.NewTipSet([]*types.BlockHeader{header})
}
