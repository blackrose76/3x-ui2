package controllers

import (
	"encoding/json"
	"net/http"
	"xray-panel/app/models"
)

type ServerController struct {
}

func NewServerController() *ServerController {
	return &ServerController{}
}

func (c *ServerController) GetServers(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement this method
}

func (c *ServerController) AddServer(w http.ResponseWriter, r *http.Request) {
	var server models.Server
	err := json.NewDecoder(r.Body).Decode(&server)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: Implement this method
}

func (c *ServerController) UpdateServer(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement this method
}

func (c *ServerController) DeleteServer(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement this method
}
