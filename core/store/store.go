package store

import (
	"github.com/zhaohaijun/blockchain-crypto/keypair"
	"github.com/zhaohaijun/matrixchain/common"
	"github.com/zhaohaijun/matrixchain/core/payload"
	"github.com/zhaohaijun/matrixchain/core/states"
	"github.com/zhaohaijun/matrixchain/core/types"
	"github.com/zhaohaijun/matrixchain/smartcontract/event"
	cstates "github.com/zhaohaijun/matrixchain/smartcontract/states"
)

// LedgerStore provides func with store package.
type LedgerStore interface {
	InitLedgerStoreWithGenesisBlock(genesisblock *types.Block, defaultBookkeeper []keypair.PublicKey) error
	Close() error
	AddHeaders(headers []*types.Header) error
	AddBlock(block *types.Block) error
	GetCurrentBlockHash() common.Uint256
	GetCurrentBlockHeight() uint32
	GetCurrentHeaderHeight() uint32
	GetCurrentHeaderHash() common.Uint256
	GetBlockHash(height uint32) common.Uint256
	GetHeaderByHash(blockHash common.Uint256) (*types.Header, error)
	GetHeaderByHeight(height uint32) (*types.Header, error)
	GetBlockByHash(blockHash common.Uint256) (*types.Block, error)
	GetBlockByHeight(height uint32) (*types.Block, error)
	GetTransaction(txHash common.Uint256) (*types.Transaction, uint32, error)
	IsContainBlock(blockHash common.Uint256) (bool, error)
	IsContainTransaction(txHash common.Uint256) (bool, error)
	GetBlockRootWithNewTxRoot(txRoot common.Uint256) common.Uint256
	GetMerkleProof(m, n uint32) ([]common.Uint256, error)
	GetContractState(contractHash common.Address) (*payload.DeployCode, error)
	GetBookkeeperState() (*states.BookkeeperState, error)
	GetStorageItem(key *states.StorageKey) (*states.StorageItem, error)
	PreExecuteContract(tx *types.Transaction) (*cstates.PreExecResult, error)
	GetEventNotifyByTx(tx common.Uint256) (*event.ExecuteNotify, error)
	GetEventNotifyByBlock(height uint32) ([]*event.ExecuteNotify, error)
}
