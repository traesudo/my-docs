package model

import "time"

type Direction struct {
	ID         int        `gorm:"id" json:"id"`
	Direction  string     `gorm:"direction" json:"direction"`
	PCreatedAt time.Time  `gorm:"p_created_at" json:"p_created_at"`
	PUpdatedAt time.Time  `gorm:"p_updated_at" json:"p_updated_at"`
	PDeletedAt *time.Time `gorm:"p_deleted_at" json:"p_deleted_at"`
}

func (l *Direction) TableName() string {
	return "directions"
}

func GetAllDirection() ([]Direction, error) {
	var l []Direction
	err := DB().Model(Direction{}).Find(&l)
	if err != nil {
		return l, nil
	}
	return l, nil
}
