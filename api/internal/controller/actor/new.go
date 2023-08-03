package actor

import (
	"context"
	"github.com/hxcuber/hollywood-api/api/internal/repository"
)

type Controller interface {
	GetAllActors(ctx context.Context)
}

func New(repo repository.Registry) Controller {
	return impl{repo: repo}
}

type impl struct {
	repo repository.Registry
}
