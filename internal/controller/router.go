package controller

import (
	"net/http"

	"github.com/edgarSucre/moca/internal/domain"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type (
	Usecase interface {
		GetMortgagePayment(domain.Mortgage) (float64, error)
	}

	handler struct {
		uc Usecase
	}
)

func New(uc Usecase, logger *log.Entry) http.Handler {
	h := handler{
		uc: uc,
	}

	router := mux.NewRouter()
	router.HandleFunc("/api/payment", loggerMiddleware(logger)(h.getMortgagePayment))

	return router
}
