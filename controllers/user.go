package controllers

import "github.com/gin-gonic/gin"

type User struct {
}

func NewUserController() *User {
	return &User{}
}

func (userController *User) Login(c *gin.Context) {

}
