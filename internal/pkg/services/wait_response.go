package services

// WaitExitError container waiting error, if any
type WaitExitError struct {
	Message string `json:"Message,omitempty"`
}

// WaitResponse ContainerWaitResponse
// OK response to ContainerWait operation
type WaitResponse struct {
	// error
	Error *WaitExitError `json:"Error,omitempty"`

	// Exit code of the container
	// Required: true
	StatusCode int64 `json:"StatusCode"`
}
