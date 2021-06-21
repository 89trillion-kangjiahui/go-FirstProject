package entity

type Gift struct {
	Gid  uint32 `json:"gid"`
	Name string `json:"name"`
	Num  uint64 `json:"num"`
}
