package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/vmorais34/rocketseat-go-pratica/internal/store/pgstore"
)

// Implementação concreta x implementação de interface
type apiHandler struct {
	q *pgstore.Queries
	r *chi.Mux
}

func (h apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.r.ServeHTTP(w, r)
}

func newHandler(q *pgstore.Queries) http.Handler{
		a := apiHandler{
			q: q,
		}

		r := chi.NewRouter()

		a.r = r
		return a
}