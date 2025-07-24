package models

type User struct {
	ID           int   `json:"id"`
	TotalDataUsage int64 `json:"total_data_usage"`
}
