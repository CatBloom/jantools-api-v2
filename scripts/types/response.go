package types

type Response struct {
	Count   int         `json:"count"`
	Results interface{} `json:"results"`
}
