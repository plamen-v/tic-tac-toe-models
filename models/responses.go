package models

type Response struct {
	StatusCode int           `json:"statusCode"`
	Error      *ErrorMessage `json:"error"`
	Payload    interface{}   `json:"d"`
}
