package controllers

import (
	"net/http"
)

type SubscriptionController struct {
}

func NewSubscriptionController() *SubscriptionController {
	return &SubscriptionController{}
}

func (c *SubscriptionController) GetSubscription(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement this method
}
