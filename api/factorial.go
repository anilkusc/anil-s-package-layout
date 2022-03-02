package api

import (
	"net/http"
)

func (api *Api) CalculateHandler(w http.ResponseWriter, r *http.Request) {
	err := api.Domain.Calculate()
	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
	http.Error(w, "OK", http.StatusOK)
	return
}
