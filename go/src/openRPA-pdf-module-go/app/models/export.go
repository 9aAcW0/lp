package models

type Export struct {
	ExecDate          string      `json:"exec_date"`
	FromEndPoint      string      `json:"from_end_point"`
	Data              interface{} `json:"data"`
	Files             interface{} `json:"files"`
	HTTPRequestStatus int         `json:"http_request_status"`
	ErrorCode         int         `json:"error_code"`
	Message           string      `json:"message"`
}
