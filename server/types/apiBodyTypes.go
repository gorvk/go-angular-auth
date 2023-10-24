package types

type Response struct {
	Success    bool
	StatusCode int
	Data       interface{}
}
