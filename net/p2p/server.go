package p2p

import (
	"github.com/libp2p/go-libp2p-host"
	"io"
	"crypto/rand"
	mrand "math/rand"
	"github.com/libp2p/go-libp2p-crypto"
	"github.com/libp2p/go-libp2p"
	"fmt"
	"context"
	"github.com/multiformats/go-multiaddr"
	"github.com/libp2p/go-libp2p-net"
	"bufio"
)

var sendBuf = make(chan string,100)

func MakeBasicHost(port int,secio bool,randseed int64) (host.Host,error){
	var r io.Reader
	if randseed == 0{
		r = rand.Reader
	}else{
		r = mrand.New(mrand.NewSource(randseed))
	}

	priv,_,err := crypto.GenerateKeyPairWithReader(crypto.Ed25519,1024,r)
	if err != nil{
		return nil,err
	}

	opts := []libp2p.Option{
		libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/127.0.0.1/tcp/%d",port)),
		libp2p.Identity(priv),
	}

	if !secio{
		opts = append(opts,libp2p.NoSecurity)
	}
	basicHost,err := libp2p.New(context.Background(),opts...)
	if err != nil{
		return nil,err
	}

	hostAddr,_ := multiaddr.NewMultiaddr(fmt.Sprintf("/ipfs/%s",basicHost.ID().Pretty()))
	addr := basicHost.Addrs()[0]
	fullAddr := addr.Encapsulate(hostAddr)
	fmt.Println("my full addr:"+fullAddr.String())

	fmt.Printf("now run at port :%d",port)
	return basicHost,nil
}


func HandleStream(s net.Stream){
	fmt.Println("parse a steam")
	rw := bufio.NewReadWriter(bufio.NewReader(s),bufio.NewWriter(s))
	go handleReadData(rw)
	go writeData(rw)
}

func handleReadData(writer *bufio.ReadWriter){
	for {
		str,err := writer.ReadString('\n')
		if err != nil{
			fmt.Println(err)
		}
		if str == ""{
			continue
		}
		fmt.Printf("get msg :%s\n",str)
	}
}

func writeData(writer *bufio.ReadWriter){
	for{
		str := <- sendBuf
		writer.WriteString(str)
		writer.Flush()
		fmt.Println("send ...  "+str)
	}
}


func Send(data string){
	sendBuf <- data
}
