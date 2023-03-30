package store

import (
	"context"
	"encoding/json"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/consensus/tendermint"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	dstore "github.com/ipfs/go-datastore"
	"golang.org/x/xerrors"
)

func (cs *ChainStore) GetMessage(ctx context.Context, c cid.Cid) (*types.Message, error) {
	var msg *types.Message
	// TODO: should I change modify this function to get message from tendermint?
	err := cs.chainBlockstore.View(ctx, c, func(b []byte) (err error) {
		msg, err = types.DecodeMessage(b)
		return err
	})
	return msg, err
}

func (cs *ChainStore) GetSignedMessage(ctx context.Context, c cid.Cid) (*types.SignedMessage, error) {
	var msg *types.SignedMessage
	err := cs.chainBlockstore.View(ctx, c, func(b []byte) (err error) {
		msg, err = types.DecodeSignedMessage(b)
		return err
	})
	return msg, err
}

func (cs *ChainStore) Load(ctx context.Context) error {
	if err := cs.loadHead(ctx); err != nil {
		return err
	}
	if err := cs.loadCheckpoint(ctx); err != nil {
		return err
	}
	return nil
}

func (cs *ChainStore) loadHead(ctx context.Context) error {
	head, err := cs.metadataDs.Get(ctx, chainHeadKey)
	if err == dstore.ErrNotFound {
		log.Warn("no previous chain state found")
		return nil
	}
	if err != nil {
		return xerrors.Errorf("failed to load chain state from datastore: %w", err)
	}

	var tscids []cid.Cid
	if err := json.Unmarshal(head, &tscids); err != nil {
		return xerrors.Errorf("failed to unmarshal stored chain head: %w", err)
	}

	ts, err := cs.LoadTipSet(ctx, types.NewTipSetKey(tscids...))
	if err != nil {
		return xerrors.Errorf("loading tipset: %w", err)
	}

	cs.heaviest = ts

	return nil
}

// GetTipsetByHeight returns the tipset on the chain behind 'ts' at the given
// height. In the case that the given height is a null round, the 'prev' flag
// selects the tipset before the null round if true, and the tipset following
// the null round if false.
func (cs *ChainStore) GetTipsetByHeight(ctx context.Context, h abi.ChainEpoch, ts *types.TipSet, prev bool) (*types.TipSet, error) {
	if ts == nil {
		ts = cs.GetHeaviestTipSet()
	}

	if h > ts.Height() {
		return nil, xerrors.Errorf("looking for tipset with height greater than start point")
	}

	if h == ts.Height() {
		return ts, nil
	}

	blk := cs.tmNode.BlockStore().LoadBlock(int64(h))
	return tendermint.TipsetByBlock(blk)
}
