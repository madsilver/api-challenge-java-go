package presenter

import (
	"time"
)

func MapperUsers(status, count int, message string) Response {
	return Response{
		Status: status,
		Body: CreateUser{
			Message:   message,
			UserCount: count,
		},
	}
}

func MapperResponse[T any](status int, took time.Time, data T) Response {
	return Response{
		Status: status,
		Body: ResponseBody[T]{
			Timestamp:       time.Now().Format("2006-01-02 15:04:05:00Z"),
			ExecutionTimeMs: int(time.Since(took).Milliseconds()),
			Data:            data,
		},
	}
}

func MapperEvaluation(status int, data any) Response {
	return Response{
		Status: status,
		Body: Evaluation{
			TestedEndpoints: data,
		},
	}
}
