package remove

import (
	"errors"
	"log/slog"
	"net/http"
	"url-shortner/internal/storage"

	resp "url-shortner/internal/lib/api/response"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type URLRemover interface {
	DeleteURL(alias string) error
}

func New(log *slog.Logger, urlRemover URLRemover) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.remove.New"

		log = log.With(slog.String("op", op), slog.String(
			"request_id",
			middleware.GetReqID(r.Context()),
		))

		alias := chi.URLParam(r, "alias")
		if alias == "" {
			log.Info(
				"alias not found",
				slog.String("alias", alias),
			)

			render.JSON(w, r, resp.Response{
				Status: resp.StatusError,
				Error:  "not found",
			})

			return
		}

		err := urlRemover.DeleteURL(alias)
		if errors.Is(err, storage.ErrURLNotFound) {
			log.Info(
				"url not found",
				slog.String("alias", alias),
			)

			render.JSON(w, r, resp.Response{
				Status: resp.StatusError,
				Error:  "not found",
			})

			return
		}

		render.JSON(w, r, resp.Response{
			Status: resp.StatusOk,
			Error:  "removed successfully",
		})

		log.Info("alias removed successfully", slog.String(
			"alias",
			alias,
		))

		return
	}
}
