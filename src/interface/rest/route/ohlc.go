package route

import (
	"net/http"

	handlers "solution_ch_part1/src/interface/rest/handlers/ohlc"

	"github.com/go-chi/chi/v5"
)

// HealthRouter a completely separate router for health check routes
func OHLCRouter(h handlers.OHLCHandlerInterface) http.Handler {
	r := chi.NewRouter()

	r.Post("/", h.Calculation)

	return r
}
