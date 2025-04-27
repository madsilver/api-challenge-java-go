package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/madsilver/api-challenge/internal/adapter/presenter"
	"github.com/madsilver/api-challenge/internal/entity"
	"net/http"
	"time"
)

type MetricHandler interface {
	GetEvaluations(ctx *gin.Context)
}

type metricHandler struct {
}

func NewMetricHandler() MetricHandler {
	return &metricHandler{}
}

func (h *metricHandler) GetEvaluations(ctx *gin.Context) {
	evaluations := map[string]entity.Evaluation{}
	for _, uri := range []string{"/superusers", "/top-countries", "/team-insights", "/active-users-per-day"} {
		start := time.Now()
		res, err := http.Get("http://localhost:8080" + uri)
		evaluations[uri] = ProcessEvaluation(start, res, err)
	}
	ctx.JSON(200, presenter.MapperEvaluation(200, evaluations))
}

func ProcessEvaluation(start time.Time, res *http.Response, err error) entity.Evaluation {
	valid := false
	if err == nil {
		var data map[string]interface{}
		if err = json.NewDecoder(res.Body).Decode(&data); err == nil {
			valid = true
		}
		_ = res.Body.Close()
	}

	return entity.Evaluation{
		Status:        res.StatusCode,
		TimeMs:        time.Since(start).Milliseconds(),
		ValidResponse: valid,
	}
}
