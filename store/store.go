package storage

type Storage struct{}

// 将一个key和对应的值保存到本地KV存储中，若此key存在则改变此key；若key不存在则新建一个key并保存相应的值
func (s *Storage) Store(key []byte, val []byte) {

}

// 查询某个key对应的值，若不存在则返回nil
func (s *Storage) Query(key []byte) []byte {
	return nil
}
