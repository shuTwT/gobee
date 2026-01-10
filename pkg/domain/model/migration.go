package model

type MigrationResult struct {
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Total   int      `json:"total"`
	Success int      `json:"success"`
	Failed  int      `json:"failed"`
	Errors  []string `json:"errors,omitempty"`
}

type MigrationCheckResult struct {
	HasDuplicates bool       `json:"has_duplicates"`
	Duplicates    []DuplicateFile `json:"duplicates"`
}

type DuplicateFile struct {
	Filename string `json:"filename"`
	Title    string `json:"title"`
	PostID   *int    `json:"post_id,omitempty"`
}
