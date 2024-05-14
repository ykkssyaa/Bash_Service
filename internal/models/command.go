package models

const (
	StatusStarted = "started"
	StatusStopped = "stopped"
	StatusSuccess = "success"
	StatusError   = "error"
)

type Command struct {
	Id     int    `json:"id,omitempty"`
	Script string `json:"script,omitempty"`
	Status string `json:"status,omitempty"`
	Output string `json:"output,omitempty"`
}
