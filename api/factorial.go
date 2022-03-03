package api

import (
	"net/http"
	"strconv"
)

func (api *Api) CalculateHandler(w http.ResponseWriter, r *http.Request) {
	number, err := strconv.Atoi(r.URL.Query().Get("number"))
	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
	result, err := api.Domain.Calculate(number)
	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
	http.Error(w, strconv.Itoa(result), http.StatusOK)
	return
}
