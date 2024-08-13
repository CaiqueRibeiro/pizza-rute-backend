package errors

type HandlerError struct {
	Message  string   `json:"error,omitempty"`
	Messages []string `json:"errors,omitempty"`
}
