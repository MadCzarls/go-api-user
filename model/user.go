package model

type User struct {
	Id       string `json:"id"`
	Username string `json:"username" binding:"required,min=3,max=20"`
	Age      uint8  `json:"age" binding:"required,min=0,max=150"`
}

func (u *User) SetId(id string) {
	u.Id = id
}

