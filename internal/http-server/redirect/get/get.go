package get

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/Lacky1234union/UrlShorter/internal/lib/api/response"
	"github.com/Lacky1234union/UrlShorter/internal/lib/logger/sl"
	"github.com/Lacky1234union/UrlShorter/internal/storage"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type URLGetter interface {
	GetURL(alias string) (string, error)
}

type Request struct {
	Alias string `json:"alias,omitempty"`
}

type Response struct {
	response.Response
	Url string `json:"url"`
}

// TODO: go:generate go run github.com/vektra/mock/ ...
func New(log *slog.Logger, urlGetter URLGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.get.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req Request
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("failed to decode request body", sl.Err(err))

			render.JSON(w, r, response.Error("failed to decode request"))
			return

		}

		log.Info("request body decoded", "Request", req)

		if err := validator.New().Struct(req); err != nil {
			log.Error("invalid request", sl.Err(err))
		}
		alias := chi.URLParam(r, "alias")
		if alias == "" {
			log.Error("invalid alias", slog.String("alias", req.Alias))
			render.JSON(w, r, response.Error("invalid alias"))
			return
		}

		url, err := urlGetter.GetURL(req.Alias)
		if errors.Is(err, storage.ErrURLNotFound) {
			log.Error("url not found", slog.String("alias", req.Alias))
			render.JSON(w, r, "url not found")
			return
		}
		log.Info("get", slog.String("url", url))
		http.Redirect(w, r, url, http.StatusFound)
	}
}
