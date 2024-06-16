package storage

type Storage struct{}

func (s *Storage) Store(key []byte, val []byte) {

}

func (s *Storage) Query(key []byte) []byte {
	return nil
}
