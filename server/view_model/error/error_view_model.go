package error

type ErrorViewModel struct {
	IsError bool   `json:"isError"`
	Message string `json:"message"`
}
