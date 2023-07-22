package bankaccount

import (
	"net/http"
	"strconv"
	"synergize/api/bankaccount/usecase"
	usecasebankaccount "synergize/api/bankaccount/usecase"
	"synergize/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type BankAccountHttpRouterRegistry struct {
	bankAccount usecasebankaccount.BankAccountService
}

func NewBankAccountHttpRouterRegistry(bankAccount usecase.BankAccountService) *BankAccountHttpRouterRegistry {
	return &BankAccountHttpRouterRegistry{
		bankAccount: bankAccount,
	}
}

func (b *BankAccountHttpRouterRegistry) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	cmd := usecase.BankAccountRequest{}

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
	if err := b.bankAccount.Create(cmd); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}

	render.JSON(w, r, "created bank account")

}

func (b *BankAccountHttpRouterRegistry) FindbyId(w http.ResponseWriter, r *http.Request) {
	iduser := chi.URLParam(r, "id")

	userIdInt, _ := strconv.Atoi(iduser)
	bankAccount, bankAccountErr := b.bankAccount.FindbyId(userIdInt)
	if bankAccountErr != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, bankAccountErr.Error())
		return
	}
	render.JSON(w, r, bankAccount)
}

func (b *BankAccountHttpRouterRegistry) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	cmd := usecase.BankAccountRequestUpdate{}
	iduser := chi.URLParam(r, "id")
	userIdInt, _ := strconv.Atoi(iduser)

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

	cmd.ID = uint(userIdInt)
	cmd.UserId = userId

	if err := b.bankAccount.Update(cmd); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}

	render.JSON(w, r, "update bank account")
}
