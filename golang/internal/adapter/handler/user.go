package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/madsilver/api-challenge/internal/adapter/presenter"
	"github.com/madsilver/api-challenge/internal/entity"
	"github.com/madsilver/api-challenge/internal/usecase"
	"strconv"
	"time"
)

type UserHandler interface {
	StoreUsers(ctx *gin.Context)
	GetSuperUsers(ctx *gin.Context)
	GetTopCountries(ctx *gin.Context)
	GetTeamInsights(ctx *gin.Context)
	GetActiveUsers(ctx *gin.Context)
}

type userHandler struct {
	useCase usecase.UserUseCase
}

func NewUserHandler(useCase usecase.UserUseCase) UserHandler {
	return &userHandler{useCase}
}

// StoreUsers godoc
// @Summary Store users
// @Description Load a users JSON file im memory
// @Tags Users
// @Accept multipart/form-data
// @Produce json
// @Success 200 {object} presenter.CreateUser
// @Param file formData file true "Upload file"
// @Router /users [post]
func (h *userHandler) StoreUsers(ctx *gin.Context) {
	fileHandler, _ := ctx.FormFile("file")
	file, _ := fileHandler.Open()
	defer file.Close()
	var users []entity.User
	_ = json.NewDecoder(file).Decode(&users)
	_ = h.useCase.StoreUsers(ctx, users)
	ctx.JSON(200, presenter.MapperUsers(200, len(users), "Arquivo recebido com sucesso"))
}

// GetSuperUsers godoc
// @Summary Get the super users
// @Description Return the users with a given score and active
// @Tags Users
// @Produce json
// @Success 200 {object} presenter.FilterUser
// @Router /superusers [get]
func (h *userHandler) GetSuperUsers(ctx *gin.Context) {
	start := time.Now()
	users, _ := h.useCase.GetSuperUsers(ctx)
	ctx.JSON(200, presenter.MapperResponse[[]entity.User](200, start, users))
}

// GetTopCountries godoc
// @Summary Get the top countries
// @Description Group users by country and return the top 5 countries by user
// @Tags Users
// @Produce json
// @Success 200 {object} presenter.TopCountry
// @Router /top-countries [get]
func (h *userHandler) GetTopCountries(ctx *gin.Context) {
	start := time.Now()
	countries, _ := h.useCase.GetTopCountries(ctx)
	ctx.JSON(200, presenter.MapperResponse[[]entity.Country](200, start, countries))
}

// GetTeamInsights godoc
// @Summary Get the team insights
// @Description Return statistics by team
// @Tags Users
// @Produce json
// @Success 200 {object} presenter.TeamInsights
// @Router /team-insights [get]
func (h *userHandler) GetTeamInsights(ctx *gin.Context) {
	start := time.Now()
	insights, _ := h.useCase.GetTeamInsights(ctx)
	ctx.JSON(200, presenter.MapperResponse[[]entity.TeamInfo](200, start, insights))
}

// GetActiveUsers godoc
// @Summary Get the active users
// @Description Return the number of logins by day
// @Tags Users
// @Produce json
// @Success 200 {object} presenter.TeamInsights
// @Param        min    query     int  false  "filter by days with at least MIN value"
// @Router /active-users-per-day [get]
func (h *userHandler) GetActiveUsers(ctx *gin.Context) {
	start := time.Now()
	var minQuery int
	param := ctx.Query("min")
	if param != "" {
		minQuery, _ = strconv.Atoi(param)
	}
	logins, _ := h.useCase.GetActiveUsersPerDay(ctx, minQuery)
	ctx.JSON(200, presenter.MapperResponse[[]entity.Login](200, start, logins))
}
