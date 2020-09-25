package models

// Error to send a error response from an API
type Error struct {
	Message string `json:"message"`
}
