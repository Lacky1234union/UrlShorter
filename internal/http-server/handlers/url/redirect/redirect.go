package redirect

import (
	"log/slog"
	"net/http"

	"github.com/Lacky1234union/UrlShorter/internal/lib/api/response"
	"github.com/Lacky1234union/UrlShorter/internal/lib/logger/sl"
	"github.com/Lacky1234union/UrlShorter/internal/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func New(log *slog.Logger, urlService service.URLService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.redirect.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		alias := chi.URLParam(r, "alias")
		if alias == "" {
			log.Info("alias is empty")
			render.JSON(w, r, response.Error("alias is required"))
			return
		}

		url, err := urlService.GetURL(alias)
		if err != nil {
			log.Error("failed to get URL", sl.Err(err))
			render.JSON(w, r, response.Error("URL not found"))
			return
		}

		log.Info("redirecting to URL", slog.String("url", url))
		http.Redirect(w, r, url, http.StatusFound)
	}
}
