package controller

import (
	"encoding/json"
	"net/http"
)

func (h handler) getMortgagePayment(w http.ResponseWriter, r *http.Request) {
	params := &MortgageRequest{}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	if err := decoder.Decode(params); err != nil {
		setErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	payment, err := h.uc.GetMortgagePayment(params.ToDomain())
	if err != nil {
		setErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	setResponse(w, payment)
}

func setResponse(w http.ResponseWriter, payment float64) {
	response := MortgageResponse{Payment: payment}
	encoder := json.NewEncoder(w)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoder.Encode(response)
}

func setErrorResponse(w http.ResponseWriter, status int, err error) {
	errorResponse := ErrorResponse{
		Error: err.Error(),
	}
	encoder := json.NewEncoder(w)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	encoder.Encode(errorResponse)
}
