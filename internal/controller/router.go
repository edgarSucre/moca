package controller

import (
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

func New(uc Usecase, logger *log.Entry) *mux.Router {
	h := handler{
		uc: uc,
	}

	spa := spaHandler{staticPath: "./public", indexPath: "./public/index.html"}
	router := mux.NewRouter()
	router.HandleFunc("/api/payment", loggerMiddleware(logger)(h.getMortgagePayment))
	router.PathPrefix("/").Handler(spa)

	return router
}
