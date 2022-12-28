package database

import (
	"context"

	"github.com/CyberAgentHack/server-performance-tuning-2023/ent"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/repository"
)

type Season struct {
}

func NewSeason() *Season {
	return &Season{}
}

func (e *Season) List(ctx context.Context, params *repository.ListSeasonsParams) (ent.Seasons, error) {
	// TODO
	return ent.Seasons{{ID: 1}}, nil
}
