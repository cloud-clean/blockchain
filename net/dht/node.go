package dht

import (
	"github.com/shiyanhui/dht"
	"fmt"
)

func NewNode(){
	downloader := dht.NewWire(65535,1024,256)
	go func(){
		for resp:= range downloader.Response(){
			fmt.Println(resp.InfoHash,resp.MetadataInfo)
		}
	}()
	downloader.Run()

	config := dht.NewCrawlConfig()
	config.OnAnnouncePeer = func(infoHash, ip string, port int) {
		// request to download the metadata info
		downloader.Request([]byte(infoHash), ip, port)
	}
	d := dht.New(config)
	d.Run()

}
