package store

type Store interface{
	Get([]byte) []byte
	Set([]byte,[]byte)error
	Del([]byte)
	Close()

}
