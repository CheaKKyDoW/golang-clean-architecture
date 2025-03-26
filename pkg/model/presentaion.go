package model

type Response struct {
	Status      int    `json:"status"`
	Header      string `json:"header"`
	Description string `json:"description"`
	Data        any    `json:"data,omitempty"`
}
