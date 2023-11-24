package transfers

import (
	"bank-api/internal/domain/entities"
)

type (
	PerformRequest struct {
		AccountDestinationID string `json:"account_destination_id" validate:"required"`
		Amount               int64  `json:"amount" validate:"required"`
	}
	ValidationErrResponse struct {
		AccountDestinationID string `json:"account_destination_id,omitempty"`
		Amount               string `json:"amount,omitempty"`
	}
	ResponseBody struct {
		ID string `json:"id"`
	}
)

func formatSliceResponse(transfers []entities.Transfer) []ResponseBody {
	transfersResponse := make([]ResponseBody, len(transfers))

	for i, trans := range transfers {
		transfersResponse[i] = ResponseBody{
			ID: trans.ID,
		}
	}

	return transfersResponse
}
