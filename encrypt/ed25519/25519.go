package ed25519

import (
	"golang.org/x/crypto/ed25519"
	"bytes"
	"fmt"
	"encoding/hex"
	"golang.org/x/crypto/ripemd160"
	"github.com/tyler-smith/go-bip39"
	"os"
	"path/filepath"
	"errors"
)

type WalletJson struct{
	Version string 	`json:"version"`
	Prikey string 	`json:"prikey"`
	address string 	`json:"address"`
}

type EdKeys struct{
	Prikey ed25519.PrivateKey
	Pubkey ed25519.PublicKey
	Mnemonic string
	Password string
	Address string
}

func NewKeys(passwd string) *EdKeys{
	encrypt,_ :=bip39.NewEntropy(128);
	mnemonic,_ :=bip39.NewMnemonic(encrypt)
	seed := bip39.NewSeed(mnemonic,passwd)
	pub,pri,_ := ed25519.GenerateKey(bytes.NewReader(seed))
	return &EdKeys{Prikey:pri,Pubkey:pub,Password:passwd,Mnemonic:mnemonic}
}


func NewKeysBymnenonic(mnemonic,passwd string) *EdKeys{
	seed := bip39.NewSeed(mnemonic,passwd)
	fmt.Println(mnemonic)
	pub,pri,_ := ed25519.GenerateKey(bytes.NewReader(seed))
	fmt.Println(hex.EncodeToString(pub))
	return &EdKeys{Prikey:pri,Pubkey:pub,Password:passwd,Mnemonic:mnemonic}
}

func LoadKeys(path string) *EdKeys{
	_,err := os.Stat(path)
	if os.IsNotExist(err){
		return nil
	}

}

func (self *EdKeys)Save(path string) error{
	_,err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err){
			filepath,_ := filepath.Split(path)
			os.Mkdir(filepath,os.ModeDir)
			file,err := os.Create(path)
			if err != nil{
				return err
			}



		}else{
			return err
		}
	}
	return errors.New("file "+path+" is exists")





}

func (self *EdKeys)GetAddress() string{
	if self.Address == ""{
		haser := ripemd160.New()
		haser.Write(self.Pubkey)
		bs := haser.Sum(nil)
		self.Address = hex.EncodeToString(bs)
	}
	return self.Address
}

func (self *EdKeys)Sign(data []byte)[]byte{
	return ed25519.Sign(self.Prikey,data)
}

func Verify(data,sig,pub []byte)bool{
	return ed25519.Verify(pub,data,sig)
}

