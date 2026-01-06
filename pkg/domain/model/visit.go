package model

type VisitLogReq struct {
	IP        string `json:"ip"`
	UserAgent string `json:"user_agent"`
	Path      string `json:"path"`
	OS        string `json:"os"`
	Browser   string `json:"browser"`
	Device    string `json:"device"`
}
