package controllers

import (
	"net/http"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (c *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement this method
}

func (c *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement this method
}
