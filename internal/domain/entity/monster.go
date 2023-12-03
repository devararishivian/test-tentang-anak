package entity

type MonsterStat struct {
	HP     int
	Attack int
	Def    int
	Speed  int
}

type Monster struct {
	ID           int
	Name         string
	CategoryName string
	TypeIDs      []int
	Description  string
	Height       float32
	Weight       float32
	Stats        MonsterStat
	ImagePath    string
	Options      string
}

type Monsters []Monster
