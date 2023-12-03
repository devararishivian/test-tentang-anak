package entity

type MonsterStat struct {
	HP     int `json:"hp"`
	Attack int `json:"attack"`
	Def    int `json:"def"`
	Speed  int `json:"speed"`
}

type MonsterOptions struct {
	Catch string
}

type Monster struct {
	ID           int         `json:"id"`
	Name         string      `json:"name"`
	CategoryName string      `json:"categoryName"`
	TypeIDs      []int       `json:"typeIDs"`
	Description  string      `json:"description"`
	Height       float32     `json:"height"`
	Weight       float32     `json:"weight"`
	Stats        MonsterStat `json:"stats"`
	ImagePath    string      `json:"imagePath"`
	Options      string      `json:"options"`
}

type Monsters []Monster

type FetchMonstersRequest struct {
	Options MonsterOptions
	Name    string
	TypeIDs *[]int
	Order   string
	Sort    string
}
