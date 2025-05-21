package save

// internal/http-server/handlers/url/save/save.go

import (
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/Lacky1234union/UrlShorter/internal/lib/api/response"
	"github.com/Lacky1234union/UrlShorter/internal/lib/logger/sl"
	"github.com/Lacky1234union/UrlShorter/internal/service"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

type Request struct {
	URL   string `json:"url" validate:"required,url"`
	Alias string `json:"alias,omitempty"`
}

type Response struct {
	response.Response
	Alias string `json:"alias,omitempty"`
}

//go:generate go run github.com/vektra/mockery/v2@v2.28.2 --name=URLSaver
type URLSaver interface {
	SaveURL(URL, alias string) (int64, error)
}

func New(log *slog.Logger, urlService service.URLService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.save.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req Request

		err := render.DecodeJSON(r.Body, &req)
		if errors.Is(err, io.EOF) {
			log.Error("request body is empty")
			render.JSON(w, r, response.Error("empty request"))
			return
		}
		if err != nil {
			log.Error("failed to decode request body", sl.Err(err))
			render.JSON(w, r, response.Error("failed to decode request"))
			return
		}

		log.Info("request body decoded", slog.Any("req", req))

		alias, err := urlService.SaveURL(req.URL, req.Alias)
		if err != nil {
			log.Error("failed to save URL", sl.Err(err))
			render.JSON(w, r, response.Error("failed to save URL"))
			return
		}

		log.Info("URL saved", slog.String("alias", alias))

		render.JSON(w, r, Response{
			Response: response.OK(),
			Alias:    alias,
		})
	}
}

func responseOK(w http.ResponseWriter, r *http.Request, alias string) {
	render.JSON(w, r, Response{
		Response: response.OK(),
		Alias:    alias,
	})
}
