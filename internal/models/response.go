package models

// GenericResponse provides a consistent JSON structure for API responses
type GenericResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
