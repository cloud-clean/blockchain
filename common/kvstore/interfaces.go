package kvstore


type Store interface {
	Set() error
	Get([]byte)([]byte,error)
	Del([]byte)error
	Close()
	Open()
}


