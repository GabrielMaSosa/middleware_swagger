package web

type ResponsError struct {
	Status  int    `json:"status,omitempty"`
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type Response struct {
	Data interface{} `json:"data,omitempty"`
}
