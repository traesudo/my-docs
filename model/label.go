package model

import "time"

type Label struct {
	ID         int        `gorm:"id" json:"id"`
	Label      string     `gorm:"label" json:"label"`
	Direction  int        `gorm:"direction" json:"direction"`
	PCreatedAt time.Time  `gorm:"p_created_at" json:"p_created_at"`
	PUpdatedAt time.Time  `gorm:"p_updated_at" json:"p_updated_at"`
	PDeletedAt *time.Time `gorm:"p_deleted_at" json:"p_deleted_at"`
}

func (l *Label) TableName() string {
	return "labels"
}

func GetAllLabel() ([]Label, error) {
	var l []Label
	err := DB().Model(Label{}).Find(&l)
	if err != nil {
		return l, nil
	}
	return l, nil
}
