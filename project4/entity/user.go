package entity

type User struct {
	Uid string `json:"uid"`
	UserName string `json:"userName"`
	GemBalance Gift `json:"gemBalance"`
	GoldBalance Gift `json:"goldBalance"`
}
