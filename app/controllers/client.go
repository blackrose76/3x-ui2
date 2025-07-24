package controllers

import (
	"encoding/json"
	"net/http"
	"xray-panel/app/models"
)

type ClientController struct {
}

func NewClientController() *ClientController {
	return &ClientController{}
}

func (c *ClientController) GetClients(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement this method
}

func (c *ClientController) AddClient(w http.ResponseWriter, r *http.Request) {
	var client models.Client
	err := json.NewDecoder(r.Body).Decode(&client)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: Implement this method
}

func (c *ClientController) UpdateClient(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement this method
}

func (c *ClientController) DeleteClient(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement this method
}
