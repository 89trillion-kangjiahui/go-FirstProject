package entity

type SoldierDTO struct {
	Id          string `json:"id"`
	Name        string `json:"Name"`
	UnlockArena string `json:"UnlockArena"`
	Rarity      string `json:"Rarity"`
	Atk         string `json:"Atk"`
	Cvc         string `json:"Cvc"`
}
