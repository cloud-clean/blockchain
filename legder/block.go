package legder

import (
	"blockchain/entity"
	"crypto"
	"errors"
)

func BuildRootHash(block entity.Block)[]byte{
	trans := block.GetTrans()
	var bs []byte
	for _,tran := range trans{
		bs = append(bs,tran.GetHash()...)
	}
	sha := crypto.SHA256.New()
	sha.Write(bs)
	return sha.Sum(nil)
}

func MakeMerkleRoot(block entity.Block)([]byte,error){
	trans := block.GetTrans()
	if len(trans) == 0{
		return nil,errors.New("trans is null")
	}
	var hashes [][]byte
	for _,tran := range trans{
		hashes = append(hashes,tran.Hash)
	}
	return makeLeafHash(hashes)[0],nil
}

func makeLeafHash(hashes [][]byte)[][]byte{
	var res [][]byte
	if len(hashes)%2 == 1{
		hashes = append(hashes,hashes[len(hashes)-1])
	}
	if len(hashes) == 2 {
		sha := crypto.SHA256.New()
		sha.Write(append(hashes[0],hashes[1]...))
		res = append(res,sha.Sum(nil))
		return res
	}else{
		mid := len(hashes)/2
		return makeLeafHash(append(makeLeafHash(hashes[:mid]),makeLeafHash(hashes[mid:])...))
	}
}



