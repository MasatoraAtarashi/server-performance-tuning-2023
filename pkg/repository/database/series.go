package database

import (
	"context"
	"database/sql"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/entity"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/errcode"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/repository"
)

type Series struct {
	db *sql.DB
}

func NewSeries(db *sql.DB) *Series {
	return &Series{db: db}
}

func (e *Series) List(ctx context.Context, params *repository.ListSeriesParams) (entity.SeriesMulti, error) {
	ctx, span := tracer.Start(ctx, "database.Series#List")
	defer span.End()

	query := "SELECT seriesID, displayName, description, imageURL, genreID FROM series"
	args := make([]any, 0, 3)
	if params.SeriesID != "" {
		query += " WHERE seriesID = ?"
		args = append(args, params.SeriesID)
	}
	query += " LIMIT ? OFFSET ?"
	args = append(args, params.Limit, params.Offset)

	rows, err := e.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, errcode.New(err)
	}

	var seriesMulti entity.SeriesMulti
	for rows.Next() {
		series := &entity.Series{}
		err = rows.Scan(
			&series.ID,
			&series.DisplayName,
			&series.Description,
			&series.ImageURL,
			&series.GenreID,
		)
		if err != nil {
			break
		}
		seriesMulti = append(seriesMulti, series)
	}

	if closeErr := rows.Close(); closeErr != nil {
		return nil, errcode.New(closeErr)
	}
	if err != nil {
		return nil, errcode.New(err)
	}

	return seriesMulti, nil
}
