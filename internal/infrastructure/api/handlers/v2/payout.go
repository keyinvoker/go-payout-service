package v2

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/keyinvoker/go-payout-service/internal/application/services"
)

type PayoutHandler struct {
	payoutService *services.PayoutService
}

func NewPayoutHandler(payoutService *services.PayoutService) *PayoutHandler {
	return &PayoutHandler{payoutService: payoutService}
}

func (h *PayoutHandler) GetPayoutByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid payout ID", http.StatusBadRequest)
		return
	}

	payout, err := h.payoutService.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Payout not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payout)
}
