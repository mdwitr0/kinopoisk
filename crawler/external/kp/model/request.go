package model

type Request[T any] struct {
	OperationName string `json:"operationName"`
	Variables     T      `json:"variables"`
	Query         string `json:"query"`
}
