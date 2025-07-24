package controllers

import (
	"encoding/json"
	"net/http"
	"xray-panel/models"
)

type TunnelController struct {
}

func NewTunnelController() *TunnelController {
	return &TunnelController{}
}

func (c *TunnelController) GetTunnels(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement this method
}

func (c *TunnelController) AddTunnel(w http.ResponseWriter, r *http.Request) {
	var tunnel models.Tunnel
	err := json.NewDecoder(r.Body).Decode(&tunnel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: Implement this method
}

func (c *TunnelController) UpdateTunnel(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement this method
}

func (c *TunnelController) DeleteTunnel(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement this method
}
