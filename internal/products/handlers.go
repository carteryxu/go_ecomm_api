package products

import (
	"log"
	"net/http"

	"github.com/carteryxu/go_ecomm_api/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	// call service -> listproducts
	err := h.service.ListProducts(r.Context())	
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return json in http response
	products := []string{"Hello", "World"}
	json.Write(w, http.StatusOK, products)
}
