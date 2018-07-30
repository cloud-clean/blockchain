package main

import (
	"blockchain/entity"
	"time"
	"blockchain/legder"
	"fmt"
	"encoding/hex"
	"runtime"
	"sync"
	"blockchain/net/p2p"
	"log"
)

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func test_gomaxprocs() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i:", i)
			defer wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("j: ", i)
			defer wg.Done()
		}(i)
	}
	//sleep
	wg.Wait()
}


func main(){
	tran := &entity.Transaction{Type:entity.TransactionType_TRANF,Timestamp:time.Now().UnixNano(),Data:[]byte("adfafdasdfasdfasdfas")}
	legder.GetTxHash(tran)
	fmt.Println(hex.EncodeToString(tran.Hash))
	basicHost,err := p2p.MakeBasicHost(3012,false,45)
	if err != nil{
		log.Fatal(err)
	}
	basicHost.SetStreamHandler("/p2p/1.0.0",p2p.HandleStream)


	for{
		time.Sleep(time.Second*4)
		p2p.Send("hahaha")
	}

}
