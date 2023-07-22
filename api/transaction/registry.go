package transaction

import (
	"net/http"
	usecasetransaction "synergize/api/transaction/usecase"
	"synergize/utils"

	"github.com/go-chi/render"
)

type TransactionHttpRouterRegistry struct {
	transactionService usecasetransaction.TransactionService
}

func NewTransactionHttpRouterRegistry(transactionService usecasetransaction.TransactionService) *TransactionHttpRouterRegistry {
	return &TransactionHttpRouterRegistry{
		transactionService: transactionService,
	}
}

func (t *TransactionHttpRouterRegistry) TopUp(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	cmd := usecasetransaction.TransactionTopUpRequest{}

	if decodeErr := render.DecodeJSON(r.Body, &cmd); decodeErr != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, decodeErr.Error())
		return
	}

	userId, userIdErr := utils.GetUserIdByToken(ctx)
	if userIdErr != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, userIdErr.Error())
		return
	}
	cmd.UserId = userId

	if err := t.transactionService.CreateTransactionTopUp(cmd); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, "success top up")
}
