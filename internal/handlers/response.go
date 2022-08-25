package handlers

type Response struct {
	Rows  any   `json:"rows"`
	Total int64 `json:"total"`
}

func NewResponse(rows any, total int64) *Response {
	return &Response{
		Rows:  rows,
		Total: total,
	}
}
