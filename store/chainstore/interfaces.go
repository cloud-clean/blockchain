package chainstore

import "blockchain/entity"

type ChainStore interface{
	GetCurHeight() int
	GetHashByHeight(int) []byte
	GetBlockByHash([]byte) *entity.Block
	GetTxByHash([]byte)*entity.Transaction
	SetCurHeight(int)
	SetHashOfHeight(int,[]byte)
	SetBlock(*entity.Block)
	SetTx(*entity.Transaction)
}
