package transfers

import (
	"context"
	"net/http"

	"bank-api/internal/domain/transfers"
	"bank-api/internal/gateways/http/middlewares"
	"bank-api/internal/gateways/http/utils"
)

type getTransferstUC interface {
	GetTransfers(ctx context.Context, accountID string) (transfers.GetTransfersOutput, error)
}

type getTransfersHandler struct {
	getTransferstUC getTransferstUC
}

func (h getTransfersHandler) GetTransfers(w http.ResponseWriter, r *http.Request) {
	accountID, ok := middlewares.GetAccountID(r.Context())
	if !ok || accountID == "" {
		_ = utils.SendError(w, utils.ErrUnauthorized, http.StatusUnauthorized)
		return
	}

	transfers, err := h.getTransferstUC.GetTransfers(r.Context(), accountID)
	if err != nil {
		_ = utils.SendError(w, utils.ErrIntervalServer, http.StatusInternalServerError)
		return
	}

	transfersResponse := formatSliceResponse(transfers.Transfers)
	_ = utils.Send(w, transfersResponse, http.StatusOK)
}
