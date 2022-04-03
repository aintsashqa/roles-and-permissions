package rest

import (
	"net/http"

	"github.com/aintsashqa/roles-and-permissions/internal/delivery"
	v1 "github.com/aintsashqa/roles-and-permissions/internal/delivery/rest/v1"
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(container *delivery.Container) http.Handler {
	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			v1.RegisterRoleRoutes(r, container)
		})
	})

	return r
}
