package controller

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
)

type spaHandler struct {
	staticPath string
	indexPath  string
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	path = filepath.Join(h.staticPath, path)

	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

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
