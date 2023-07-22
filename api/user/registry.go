package user

import (
	"net/http"
	"strconv"
	usecaseuser "synergize/api/user/usecase"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type UserHttpRouterRegistry struct {
	serviceUser usecaseuser.ServiceUser
}

func NewUserHttpRouterRegistry(serviceUser usecaseuser.ServiceUser) *UserHttpRouterRegistry {
	return &UserHttpRouterRegistry{serviceUser: serviceUser}
}

func (u *UserHttpRouterRegistry) UserList(w http.ResponseWriter, r *http.Request) {
	var balanceStart float64
	var balanceEnd float64

	var balanceStartParam string = r.URL.Query().Get("balanceStart")
	var balanceEndParam string = r.URL.Query().Get("balanceEnd")
	var registerAtParam string = r.URL.Query().Get("registerAt")

	if balanceStartParam != "" {
		balanceStartfloat, balanceStartfloatErr := strconv.ParseFloat(balanceStartParam, 64)
		if balanceStartfloatErr != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, balanceStartfloatErr.Error())
			return
		}
		balanceStart = balanceStartfloat
	}

	if balanceEndParam != "" {
		balanceEndFloat, balanceEndfloatErr := strconv.ParseFloat(balanceEndParam, 64)
		if balanceEndfloatErr != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, balanceEndfloatErr.Error())
			return
		}
		balanceEnd = balanceEndFloat
	}

	if registerAtParam != "" {
		_, formatRegisterAtErr := time.Parse("2006-01-02", registerAtParam)
		if formatRegisterAtErr != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, formatRegisterAtErr.Error())
			return
		}
	}

	param := usecaseuser.UserQueryParam{
		Username:        r.URL.Query().Get("username"),
		AccountName:     r.URL.Query().Get("accountName"),
		AccountNumber:   r.URL.Query().Get("accountNumber"),
		AccountBankName: r.URL.Query().Get("accountBankName"),
		BalanceStart:    balanceStart,
		BalanceEnd:      balanceEnd,
		RegisterAt:      registerAtParam,
	}

	userList, userListErr := u.serviceUser.UserList(param)
	if userListErr != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, userListErr.Error())
		return
	}

	result := map[string]interface{}{
		"result": userList,
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, result)
}

func (u *UserHttpRouterRegistry) UserDetail(w http.ResponseWriter, r *http.Request) {

	parseUintIdUser, parseUintIdUserErr := strconv.ParseUint(chi.URLParam(r, "userId"), 10, 32)
	if parseUintIdUserErr != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, parseUintIdUserErr.Error())
		return
	}

	userDetail, userDetailErr := u.serviceUser.UserDetail(uint(parseUintIdUser))
	if userDetailErr != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, userDetailErr.Error())
		return
	}

	result := map[string]interface{}{
		"result": userDetail,
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, result)
}
