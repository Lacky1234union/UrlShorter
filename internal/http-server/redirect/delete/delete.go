package delete

import (
	"log/slog"
	"net/http"

	"github.com/Lacky1234union/UrlShorter/internal/lib/api/response"
	"github.com/Lacky1234union/UrlShorter/internal/lib/logger/sl"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type URLDeleter interface {
	DeleteURL(alias string) error
}
type Request struct {
	Alias string `json:"alias,omitempty"`
}

type Response struct {
	response.Response
}

func New(log *slog.Logger, urlRedirect URLDeleter) http.HandlerFunc {
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
			log.Error("invalid alias", slog.String("alias", alias))
			render.JSON(w, r, response.Error("invalid alias"))
			return
		}

		err = urlRedirect.DeleteURL(alias)
		if err != nil {

			log.Error("failed to found url", sl.Err(err))

			render.JSON(w, r, response.Error("failed to found url"))
			return
		}
		log.Info("delete", slog.String("alias", alias))
		http.Redirect(w, r, alias, http.StatusFound)
		return
	}
}
