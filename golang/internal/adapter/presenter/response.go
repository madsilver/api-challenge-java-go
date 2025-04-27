package presenter

import "github.com/madsilver/api-challenge/internal/entity"

type Response struct {
	Status int `json:"status"`
	Body   any `json:"body"`
}

type FilterUser struct {
	Timestamp       string        `json:"timestamp"`
	ExecutionTimeMs int           `json:"execution_time_ms"`
	Data            []entity.User `json:"data"`
}

type CreateUser struct {
	Message   string `json:"message"`
	UserCount int    `json:"user_count"`
}

type TopCountry struct {
	Timestamp       string           `json:"timestamp"`
	ExecutionTimeMs int              `json:"execution_time_ms"`
	Countries       []entity.Country `json:"countries"`
}

type TeamInsights struct {
	Timestamp       string            `json:"timestamp"`
	ExecutionTimeMs int               `json:"execution_time_ms"`
	Teams           []entity.TeamInfo `json:"teams"`
}

type Logins struct {
	Timestamp       string         `json:"timestamp"`
	ExecutionTimeMs int            `json:"execution_time_ms"`
	Logins          []entity.Login `json:"logins"`
}

type Evaluation struct {
	TestedEndpoints any `json:"tested_endpoints"`
}
