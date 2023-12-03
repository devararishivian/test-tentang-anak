package model

type StoreMonsterRequest struct {
	Name         string                  `json:"name" validate:"required,max=255"`
	CategoryName string                  `json:"categoryName" validate:"required,max=255"`
	TypeIDs      []int                   `json:"typeIDs" validate:"required"`
	Description  string                  `json:"description" validate:"required"`
	Height       float32                 `json:"height" validate:"required"`
	Weight       float32                 `json:"weight" validate:"required"`
	Stats        StoreMonsterStatRequest `json:"stats" validate:"required"`
	ImagePath    string                  `json:"imagePath"`
}

type StoreMonsterStatRequest struct {
	HP     int `json:"hp" validate:"required"`
	Attack int `json:"attack" validate:"required"`
	Def    int `json:"def" validate:"required"`
	Speed  int `json:"speed" validate:"required"`
}
