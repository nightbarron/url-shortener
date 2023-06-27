package models

type UrlResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Url     string `json:"url"`
}
