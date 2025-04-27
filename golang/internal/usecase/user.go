package usecase

import (
	"context"
	"github.com/madsilver/api-challenge/internal/entity"
	"sort"
)

const (
	targetScore       = 900
	limitTopCountries = 4
)

type UserUseCase interface {
	StoreUsers(ctx context.Context, users []entity.User) error
	GetSuperUsers(ctx context.Context) ([]entity.User, error)
	GetTopCountries(ctx context.Context) ([]entity.Country, error)
	GetTeamInsights(ctx context.Context) ([]entity.TeamInfo, error)
	GetActiveUsersPerDay(ctx context.Context, min int) ([]entity.Login, error)
}

type userUseCase struct {
	userRepo UserRepository
}

func NewUserUseCase(userRepo UserRepository) UserUseCase {
	return &userUseCase{userRepo}
}

func (u *userUseCase) StoreUsers(ctx context.Context, users []entity.User) error {
	return u.userRepo.StoreUsers(ctx, users)
}

func (u *userUseCase) GetSuperUsers(ctx context.Context) ([]entity.User, error) {
	users, _ := u.userRepo.GetUsers(ctx)
	var superUsers []entity.User
	for _, user := range users {
		if user.Score > targetScore && user.Active {
			superUsers = append(superUsers, user)
		}
	}
	return superUsers, nil
}

func (u *userUseCase) GetTopCountries(ctx context.Context) ([]entity.Country, error) {
	var topCountries []entity.Country
	var mapCountries = map[string]int{}
	users, _ := u.userRepo.GetUsers(ctx)
	for _, user := range users {
		mapCountries[user.Country]++
	}
	for i, name := range SortedNamesByTotal(mapCountries) {
		if i > limitTopCountries {
			break
		}
		topCountries = append(topCountries, entity.Country{
			Country: name,
			Total:   mapCountries[name],
		})
	}
	return topCountries, nil
}

func (u *userUseCase) GetTeamInsights(ctx context.Context) ([]entity.TeamInfo, error) {
	users, _ := u.userRepo.GetUsers(ctx)
	var mapTeam = map[string]*entity.TeamInfo{}
	for _, user := range users {
		_, ok := mapTeam[user.Team.Name]
		if !ok {
			mapTeam[user.Team.Name] = &entity.TeamInfo{
				Team: user.Team.Name,
			}
		}

		info := mapTeam[user.Team.Name]
		info.TotalMembers++
		if user.Team.Leader {
			info.Leaders++
		}
		if user.Active {
			info.ActivePercentage++
		}

		for _, proj := range user.Team.Projects {
			if proj.Completed {
				info.CompletedProjects++
			}
		}
	}
	var team []entity.TeamInfo
	for _, info := range mapTeam {
		info.ActivePercentage = (info.ActivePercentage / float64(info.TotalMembers)) * 100
		team = append(team, *info)
	}
	return team, nil
}

func (u *userUseCase) GetActiveUsersPerDay(ctx context.Context, min int) ([]entity.Login, error) {
	users, _ := u.userRepo.GetUsers(ctx)
	logs := map[string]int{}
	for _, user := range users {
		for _, log := range user.Logs {
			if log.Action == "login" {
				logs[log.Date]++
			}
		}
	}

	var logins []entity.Login
	for k, v := range logs {
		if min > 0 && v < min {
			continue
		}
		logins = append(logins, entity.Login{
			Date:  k,
			Total: v,
		})
	}

	return logins, nil
}

func SortedNamesByTotal(data map[string]int) []string {
	var names []string
	for name, _ := range data {
		names = append(names, name)
	}
	sort.Slice(names, func(i, j int) bool {
		return data[names[i]] > data[names[j]]
	})
	return names
}
