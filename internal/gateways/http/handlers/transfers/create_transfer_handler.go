package transfers

import (
	"context"
	"net/http"

	"bank-api/internal/domain/transfers"
	"bank-api/internal/gateways/http/middlewares"
	"bank-api/internal/gateways/http/utils"
)

type createTransfertUC interface {
	CreateTransfer(ctx context.Context, transfer transfers.CreateTransferInput) (transfers.CreateTransferOutput, error)
}

type createTransferHandler struct {
	createTransfertUC createTransfertUC
	validator         *utils.StructValidator
}

func (h createTransferHandler) CreateTransfer(w http.ResponseWriter, r *http.Request) {
	var reqBody PerformRequest

	authAccountID, ok := middlewares.GetAccountID(r.Context())
	if !ok || authAccountID == "" {
		_ = utils.SendError(w, utils.ErrUnauthorized, http.StatusUnauthorized)
		return
	}

	err := utils.Decode(r, &reqBody)
	if err != nil {
		_ = utils.SendError(w, utils.ErrDecode, http.StatusBadRequest)
		return
	}

	var valErr ValidationErrResponse
	err = h.validator.Validate(reqBody, &valErr)
	if err != nil {
		_ = utils.Send(w, valErr, http.StatusBadRequest)
		return
	}

	transfer, err := h.createTransfertUC.CreateTransfer(r.Context(), transfers.CreateTransferInput{
		AccountOriginID:      authAccountID,
		AccountDestinationID: reqBody.AccountDestinationID,
		Amount:               reqBody.Amount,
	})

	if err != nil {
		_ = utils.SendError(w, utils.ErrIntervalServer, http.StatusInternalServerError)
		return
	}

	_ = utils.Send(w, ResponseBody{
		ID: transfer.TransferID,
	}, http.StatusCreated)
}
