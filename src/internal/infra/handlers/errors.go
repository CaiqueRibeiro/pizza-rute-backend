package handlers

type Error struct {
	Message  string   `json:"error,omitempty"`
	Messages []string `json:"errors,omitempty"`
}
