package model

type Live struct {
	ID          string `gorm:"id" json:"id"`
	Title       string `gorm:"title" json:"title"`
	Description string `gorm:"description" json:"description"`
	Img         string `gorm:"img" json:"img"`
	Sharer      string `gorm:"sharer" json:"sharer"`
}

func (l *Live) TableName() string {
	return "lives"
}

func GetAllLive() ([]Live, error) {
	var l []Live
	err := DB().Model(Live{}).Find(&l)
	if err != nil {
		return l, nil
	}
	return l, nil
}
