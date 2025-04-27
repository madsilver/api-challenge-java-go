package main

import (
	"github.com/madsilver/api-challenge/internal/adapter/handler"
	"github.com/madsilver/api-challenge/internal/infra/http"
)

func main() {
	manager := handler.NewManager()
	server := http.NewServer(manager)
	server.Start()
}
