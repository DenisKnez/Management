package http

type Reponse struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
	Total int         `json:"total,omitempty"`
}
