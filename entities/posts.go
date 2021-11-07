package entities

type Post struct {
	ID      string `json:"id" gorm:"column:post_id"`
	Title   string `json:"title" gorm:"column:title"`
	Content string `json:"content" gorm:"column:content"`
}

type Posts []Post

func (Post) TableName() string {
	return "posts"
}
