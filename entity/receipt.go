package entity

type Result struct {
	Hash    []byte
	State   string //"SUCCESS" or "FAIL"
	Message string //失败的原因
}
type Receipt struct {
	Results []Result
}
