package recordoutput

type CreateRecordOutput struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}
