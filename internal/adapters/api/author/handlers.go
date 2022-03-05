package author

import (
	"ca-library-app/internal/adapters/api"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	authorURL  = "/authors/:authors_id"
	authorsURL = "/authors"
)

type handler struct {
	authorService Service
}

func NewHandler(service Service) api.Handler {
	return &handler{authorService: service}
}
func (h *handler) Register(router *httprouter.Router) {
	router.GET(authorsURL, h.GetAllAuthor)
}
func (h *handler) GetAllAuthor(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//author := h.authorService.GetAllAuthor(context.Background(),0,0)
	w.Write([]byte("author"))
	w.WriteHeader(http.StatusOK)

}
