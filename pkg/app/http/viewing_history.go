package http

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/app/http/request"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/app/http/response"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/entity"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/usecase"
)

func (s *Service) routeViewingHistory(r chi.Router) {
	r.Post("/", s.createViewingHistory)
	r.Get("/", s.listViewingHistories)
}

func (s *Service) createViewingHistory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx, span := tracer.Start(ctx, "http.Service#createViewinghistory")
	defer span.End()

	body := &entity.ViewingHistory{}
	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		response.BadRequest(err, w, r)
		return
	}

	req := &usecase.CreateViewingHistoryRequest{ViewingHistory: body}
	resp, err := s.usecase.CreateViewingHistory(ctx, req)
	if err != nil {
		response.Error(err, w, r)
		return
	}
	response.OK(resp.ViewingHistory, w, r)
}

func (s *Service) listViewingHistories(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx, span := tracer.Start(ctx, "http.Service#createViewinghistory")
	defer span.End()

	episodeIDs := request.QueryStrings(r, "episodeIds")
	if len(episodeIDs) == 0 {
		response.OK(nil, w, r)
		return
	}
	req := &usecase.BatchGetViewingHistoriesRequest{
		UserID:     r.Header.Get("userId"),
		EpisodeIDs: request.QueryStrings(r, "episodeIds"),
	}
	resp, err := s.usecase.BatchGetViewingHistories(ctx, req)
	if err != nil {
		response.Error(err, w, r)
		return
	}

	response.OK(resp.ViewingHistories, w, r)
}
