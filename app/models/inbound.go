package models

type Inbound struct {
	ID       int    `json:"id"`
	Protocol string `json:"protocol"`
	Port     int    `json:"port"`
	Tag      string `json:"tag"`
}
