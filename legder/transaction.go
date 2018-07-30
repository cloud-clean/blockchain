package legder

import (
	"blockchain/entity"
	"github.com/golang/protobuf/proto"
	"crypto/sha256"
)

func GetTxHash(transaction *entity.Transaction) []byte{
	data,_ := proto.Marshal(transaction);
	sha := sha256.New()
	sha.Write(data)
	hash := sha.Sum(nil)
	transaction.Hash = hash
	return transaction.Hash
}