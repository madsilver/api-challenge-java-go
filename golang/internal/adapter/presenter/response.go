package presenter

type Response struct {
	Status int `json:"status"`
	Body   any `json:"body"`
}

type ResponseBody[T any] struct {
	Timestamp       string `json:"timestamp"`
	ExecutionTimeMs int    `json:"execution_time_ms"`
	Data            T      `json:"data"`
}

type CreateUser struct {
	Message   string `json:"message"`
	UserCount int    `json:"user_count"`
}

type Evaluation struct {
	TestedEndpoints any `json:"tested_endpoints"`
}
