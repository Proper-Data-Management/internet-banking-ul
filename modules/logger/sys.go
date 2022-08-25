package logger

type Metadata struct {
	Ch    string `json:"ch"`
	ReqID string `json:"req_id"`
	SsnID int64  `json:"ssn_id"`
	CstID int64  `json:"cst_id"`
	UsrID int64  `json:"usr_id"`
	PrsID int64  `json:"prs_id"`
	UsrAg string `json:"usr_ag"`
	AppVr string `json:"app_vr,omitempty"`
}
