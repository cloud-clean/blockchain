package gkvstore

import (
	"gitee.com/johng/gkvdb/gkvdb"
	"sync"
	"log"
)

var once sync.Once
var gkv *Gkv
type Gkv struct{
 	db *gkvdb.DB
}

func NewGkv() *Gkv{
	once.Do(func() {
		db,err := gkvdb.New("db")
		if err != nil{
			log.Fatal(err)
		}
		gkv = &Gkv{db:db}
	});
	return gkv
}

func (self *Gkv)Close(){
	self.db.Close()
}



func(self *Gkv)Get(key []byte)([]byte){
		value := self.db.Get(key)
		return value
}

func(self *Gkv)Set(key,value []byte)error{
	err := self.db.Set(key,value)
	return err
}

func (self *Gkv)Del(key []byte)error{
	err := self.db.Remove(key)
	return err
}
