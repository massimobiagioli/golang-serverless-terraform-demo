package models

type HealthResponse struct {
	Status string `json:"status"`
}

func HealthResponseOk() HealthResponse {
	return HealthResponse{Status: "ok"}
}
