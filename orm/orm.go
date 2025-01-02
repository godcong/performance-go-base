package orm

import (
	"time"
)

// User 测试数据结构
type User struct {
	ID int `xorm:"pk" gorm:"primary_key"`
	// Username holds the value of the "username" field.
	Username string
	// Email holds the value of the "email" field.
	Email string
	// Password holds the value of the "password" field.
	Password string
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time
	//Posts     []Post `gorm:"foreignKey:UserID"`
	//Tags      []Tag  `gorm:"many2many:user_tags"` // 多对多关系
}

type Post struct {
	// ID of the ent.
	ID int
	// Title holds the value of the "title" field.
	Title string
	// Content holds the value of the "content" field.
	Content string
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time
	AuthorID  int
	// AuthorID holds the value of the "author_id" field.
	Author *User `gorm:"foreignKey:AuthorID"`
}

type Tag struct {
	// ID of the ent.
	ID int
	// Name holds the value of the "name" field.
	Name string
	// Description holds the value of the "description" field.
	Description string
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time
	Users     []User `gorm:"many2many:user_tags"`
}
