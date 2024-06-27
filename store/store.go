package storage

import "github.com/syndtr/goleveldb/leveldb"

var db *leveldb.DB

func check() {
	if db != nil {
		return
	}
	db, _ = leveldb.OpenFile("minichain.db", nil)
}

// 将一个key和对应的值保存到本地KV存储中，若此key存在则改变此key；若key不存在则新建一个key并保存相应的值
func Store(key []byte, val []byte) {
	check()
	db.Put(key, val, nil)

}

// 查询某个key对应的值，若不存在则返回nil
func Query(key []byte) []byte {
	check()
	ret, err := db.Get(key, nil)
	if err == leveldb.ErrNotFound {
		return nil
	}
	return ret
}
