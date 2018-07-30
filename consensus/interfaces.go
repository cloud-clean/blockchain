package consensus

import "blockchain/entity"

type TxPool interface {
	Clear()
	add(transaction entity.Transaction)
	GetAll()[]entity.Transaction
}

