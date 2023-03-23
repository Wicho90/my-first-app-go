package model

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  uint32 `json:"age"`
}

func NewUser() *User {
	return &User{}
}

// func (u *User) Name() string {
// 	return u.name
// }
