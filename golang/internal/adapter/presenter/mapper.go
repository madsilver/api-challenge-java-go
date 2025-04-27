package presenter

import (
	"github.com/madsilver/api-challenge/internal/entity"
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

func MapperSuperUsers(status int, took time.Time, data []entity.User) Response {
	return Response{
		Status: status,
		Body: FilterUser{
			Timestamp:       time.Now().Format("2006-01-02 15:04:05:00Z"),
			ExecutionTimeMs: int(time.Since(took).Milliseconds()),
			Data:            data,
		},
	}
}

func MapperTopCountries(status int, took time.Time, data []entity.Country) Response {
	return Response{
		Status: status,
		Body: TopCountry{
			Timestamp:       time.Now().Format("2006-01-02 15:04:05:00Z"),
			ExecutionTimeMs: int(time.Since(took).Milliseconds()),
			Countries:       data,
		},
	}
}

func MapperTeamInsights(status int, took time.Time, data []entity.TeamInfo) Response {
	return Response{
		Status: status,
		Body: TeamInsights{
			Timestamp:       time.Now().Format("2006-01-02 15:04:05:00Z"),
			ExecutionTimeMs: int(time.Since(took).Milliseconds()),
			Teams:           data,
		},
	}
}

func MapperLogins(status int, took time.Time, data []entity.Login) Response {
	return Response{
		Status: status,
		Body: Logins{
			Timestamp:       time.Now().Format("2006-01-02 15:04:05:00Z"),
			ExecutionTimeMs: int(time.Since(took).Milliseconds()),
			Logins:          data,
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
