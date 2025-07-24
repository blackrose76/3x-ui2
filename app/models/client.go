package models

import "time"

type Client struct {
	ID          int       `json:"id"`
	Email       string    `json:"email"`
	UUID        string    `json:"uuid"`
	DataUsage   int64     `json:"data_usage"`
	ExpirationDate time.Time `json:"expiration_date"`
	Online      bool      `json:"online"`
}
