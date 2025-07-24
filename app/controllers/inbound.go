package controllers

import (
	"encoding/json"
	"net/http"
	"xray-panel/app/models"
)

type InboundController struct {
}

func NewInboundController() *InboundController {
	return &InboundController{}
}

func (c *InboundController) GetInbounds(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement this method
}

func (c *InboundController) AddInbound(w http.ResponseWriter, r *http.Request) {
	var inbound models.Inbound
	err := json.NewDecoder(r.Body).Decode(&inbound)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: Implement this method
}

func (c *InboundController) UpdateInbound(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement this method
}

func (c *InboundController) DeleteInbound(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement this method
}
