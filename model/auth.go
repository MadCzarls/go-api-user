package model

type Auth struct {
	Id string `json:"id" binding:"required"`
	//Password string `json:"password" binding:"required"` //@TODO implement password handling in the future
}
