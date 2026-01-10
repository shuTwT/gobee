package model

type MigrationResult struct {
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Total   int      `json:"total"`
	Success int      `json:"success"`
	Failed  int      `json:"failed"`
	Errors  []string `json:"errors,omitempty"`
}
