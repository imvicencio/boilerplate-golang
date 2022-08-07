package factory

import (
	"github.com/mcholismalik/boilerplate-golang/internal/abstraction"
	"github.com/mcholismalik/boilerplate-golang/internal/factory/repository"
	"github.com/mcholismalik/boilerplate-golang/internal/factory/usecase"
)

type Factory struct {
	Repository repository.Factory
	Usecase    usecase.Factory
	WsHub      *abstraction.WsHub
}

func Init() Factory {
	f := Factory{}
	// f.Repository = repository.Init()
	// f.Usecase = usecase.Init(f.Repository)
	f.WsHub = abstraction.NewWsHub()

	return f
}
