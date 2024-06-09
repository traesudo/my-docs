package model

import (
	"strconv"
	"time"
)

type Post struct {
	ID         string     `gorm:"id" json:"id"`
	Label      int        `gorm:"label" json:"label"`
	Title      string     `gorm:"title" json:"title"`
	Content    string     `gorm:"content" json:"content"`
	PCreatedAt time.Time  `gorm:"p_created_at" json:"p_created_at"`
	PUpdatedAt time.Time  `gorm:"p_updated_at" json:"p_updated_at"`
	PDeletedAt *time.Time `gorm:"p_deleted_at" json:"p_deleted_at"`
}

func GetPostById(id string) (*Post, error) {
	var post Post
	err := DB().Model(Post{}).Where("id = ?", id).Find(&post).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func GetLastPost(query string) ([]Post, error) {
	var posts []Post
	var direction int
	direction, err := strconv.Atoi(query)
	if err != nil {
		direction = 1
	}
	err = DB().Model(Post{}).Where("direction = ?", direction).Order("p_created_at desc").Find(&posts).Error
	if err != nil {
		return posts, nil
	}
	return posts, nil
}

func GetAnyPosts(query string) ([]Post, error) {
	var posts []Post
	err := DB().Model(Post{}).Where("content like ? OR title like ?", "%"+query+"%", "%"+query+"%").Order("p_created_at desc").Find(&posts).Error
	if err != nil {
		return posts, nil
	}
	return posts, nil
}
