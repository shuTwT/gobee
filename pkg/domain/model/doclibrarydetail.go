package model

type DocLibraryDetailCreateReq struct {
	LibraryID int    `json:"library_id" validate:"required"`
	Title     string `json:"title" validate:"required"`
	Version   string `json:"version"`
	Content   string `json:"content"`
	ParentID  int    `json:"parent_id"`
	Path      string `json:"path"`
	URL       string `json:"url"`
	Language  string `json:"language"`
}

type DocLibraryDetailUpdateReq struct {
	ID        int    `json:"id" validate:"required"`
	Title     string `json:"title" validate:"required"`
	Version   string `json:"version"`
	Content   string `json:"content"`
	ParentID  int    `json:"parent_id"`
	Path      string `json:"path"`
	URL       string `json:"url"`
	Language  string `json:"language"`
}

type DocLibraryDetailResp struct {
	ID        int        `json:"id"`
	LibraryID int        `json:"library_id"`
	Title     string     `json:"title"`
	Version   string     `json:"version"`
	Content   string     `json:"content"`
	ParentID  *int       `json:"parent_id"`
	Path      string     `json:"path"`
	URL       string     `json:"url"`
	Language  string     `json:"language"`
	CreatedAt *LocalTime `json:"created_at"`
	UpdatedAt *LocalTime `json:"updated_at"`
}

type DocLibraryDetailTreeResp struct {
	ID        int                         `json:"id"`
	LibraryID int                         `json:"library_id"`
	Title     string                      `json:"title"`
	Version   string                      `json:"version"`
	Content   string                      `json:"content"`
	ParentID  *int                        `json:"parent_id"`
	Path      string                      `json:"path"`
	URL       string                      `json:"url"`
	Language  string                      `json:"language"`
	CreatedAt *LocalTime                  `json:"created_at"`
	UpdatedAt *LocalTime                  `json:"updated_at"`
	Children  []DocLibraryDetailTreeResp  `json:"children,omitempty"`
}
