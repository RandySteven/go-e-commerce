package responses

type Response struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
	Data    any    `json:"data"`
}
