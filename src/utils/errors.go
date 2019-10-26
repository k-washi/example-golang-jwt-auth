package utils

//APIError error interface
type APIError interface {
	GetStatus() int
	GetMessage() string
	GetError() string
}

type apiError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

func (e *apiError) GetStatus() int {
	return e.Status
}

func (e *apiError) GetMessage() string {
	return e.Message
}

func (e *apiError) GetError() string {
	return e.Error
}

//NewAPIError internal api error
func NewAPIError(statusCode int, message string) APIError {
	return &apiError{
		Status:  statusCode,
		Message: message,
	}
}
