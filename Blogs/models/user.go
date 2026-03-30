package models

type User struct {
	Uid       int    `json:"uid"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Avatar    string `json:"avatar"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// TableName 指定表名
func (User) TableName() string {
	return "user"
}

type UserInfo struct {
	Uid      int    `json:"uid"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}
