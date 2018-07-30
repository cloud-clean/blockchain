package chainstore

import (
	"blockchain/store"
	"blockchain/common/constant"
	"encoding/binary"
	"bytes"
	"errors"
	"blockchain/entity"
	"github.com/golang/protobuf/proto"
)

type ChainDb struct{
	Db store.Store
}

func (self *ChainDb) GetCurHeight()int{
	b := self.Db.Get([]byte(constant.BLOCK_CUR_HEIGHT))
	bbuf := bytes.NewBuffer(b)
	var height int
	binary.Read(bbuf,binary.LittleEndian,&height)
	if height >0 {
		return height
	}
	return 0

}

func (self *ChainDb) SetCurHeight(h int) error{
	bbuf := bytes.NewBuffer([]byte{})
	err := binary.Write(bbuf,binary.LittleEndian,h)
	if err != nil{
		return err
	}
	self.Db.Set([]byte(constant.BLOCK_CUR_HEIGHT),bbuf.Bytes())
	return nil
}

func (self *ChainDb) GetHashByHeight(h int)([]byte,error){
	bbuf := bytes.NewBuffer([]byte{})
	err := binary.Write(bbuf,binary.LittleEndian,h)
	if err != nil{
		return nil,err
	}
	buf := self.Db.Get(append([]byte(constant.PREFIX_BLOCK_HASH),bbuf.Bytes()...))
	if buf != nil{
		return buf,nil
	}else{
		return nil,errors.New("don't find hash")
	}
}

func (self *ChainDb) SetHashOfHeight(h int,hash []byte){
	bbuf := bytes.NewBuffer([]byte{})
	binary.Write(bbuf,binary.LittleEndian,h)
	self.Db.Set(append([]byte(constant.PREFIX_BLOCK_HASH),bbuf.Bytes()...),hash)
}

func (self *ChainDb) GetBlocByHash(hash []byte) *entity.Block{
	b := self.Db.Get(append([]byte(constant.PREFIX_BLOCK),hash...))
	if b == nil {
		return nil
	}else{
		var block = new(entity.Block)
		err := proto.Unmarshal(b,block)
		if err != nil{
			return nil
		}else{
			return block
		}
	}
}

func (self *ChainDb) SetBlock(block *entity.Block){
	if block == nil{
		return
	}
	key := append([]byte(constant.PREFIX_BLOCK),block.Header.Hash...)
	data,err := proto.Marshal(block)
	if err != nil{
		return
	}
	self.Db.Set(key,data)
}

func (self *ChainDb) GetTxByHash(hash []byte) *entity.Transaction{
	tx := new(entity.Transaction)
	data := self.Db.Get(append([]byte(constant.PREFIX_TRANSATION),hash...))
	err := proto.Unmarshal(data,tx)
	if err != nil{
		return nil
	}
	return tx
}

func (self *ChainDb) SetTx(tx *entity.Transaction){
	if tx == nil{
		return
	}
	data,err := proto.Marshal(tx)
	if err != nil{
		return
	}
	self.Db.Set(append([]byte(constant.PREFIX_TRANSATION),tx.Hash...),data)
}

