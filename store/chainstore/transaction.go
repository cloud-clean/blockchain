package chainstore


import (
	"blockchain/entity"
	"time"
	"github.com/golang/protobuf/proto"
	"blockchain/common"
	"blockchain/encrypt/ed25519"
	"bytes"
)

func NewTransaction(data,sender []byte,typ entity.TransactionType,key *ed25519.EdKeys)(*entity.Transaction,error){
	var tx entity.Transaction
	tx.Timestamp = time.Now().UnixNano()
	tx.Data = data
	tx.Sender = sender
	tx.Type = typ
	b,err :=proto.Marshal(&tx)
	if err != nil{
		return nil,err
	}
	hash := common.GetHash(b)
	tx.Hash = hash
	sig := key.Sign(tx.Hash)
	tx.Sign = sig
	return &tx,nil
}

func TxVerify(tx entity.Transaction)bool{
	if tx.Hash == nil {
		return false
	}
	if tx.Sign == nil {
		return false
	}
	hash := tx.Hash
	sig := tx.Sign
	tx.Hash = nil
	tx.Sign = nil
	b,err := proto.Marshal(&tx)
	if err != nil{
		return false
	}
	if bytes.Compare(hash,common.GetHash(b)) != 1{
		return false
	}
	if ed25519.Verify(hash,sig,tx.Sender){
		tx.Hash = hash
		tx.Sign = sig
		return true
	}else{
		return false
	}
}
