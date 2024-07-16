package atm

// здесь идет обработка хендлеров

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)
// Request общая структура запросов для amount-а
type Request struct {
	Amount float64 `json:"amount"`
}

// CreateAccountHandler godoc
// @Summary Создать новый аккаунт
// @Description Создает новый аккаунт
// @Tags accounts
// @Produce plain
// @Success 201 {string} string "Account created with ID: {id}"
// @Router /accounts [post]
func CreateAccountHandler(accountService *AccountService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		account := accountService.CreateAccount()
		w.WriteHeader(http.StatusCreated)
		resp := fmt.Sprintf("Account successfully created, id: %d", account.ID)
		w.Write([]byte(resp))
	}
}

// DepositHandler godoc
// @Summary Депозит денег в аккаунт
// @Description Вкладывает определенную сумму в аккаунт
// @Tags accounts
// @Accept  json
// @Produce plain
// @Param id path int true "Account ID"
// @Param amount body float64 true "Deposit amount"
// @Success 200 {string} string "Deposited {amount} to account ID: {id}"
// @Failure 400 {string} string "Invalid account ID or amount"
// @Router /accounts/{id}/deposit [post]
func DepositHandler(accountService *AccountService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"]) // вытягиваем id из параметров
		if err != nil {
			http.Error(w, ErrInvalidAccountID.Error(), http.StatusBadRequest)
			return
		}
		// декод реквеста
		var req Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, ErrInvalidAmount.Error(), http.StatusBadRequest)
			return
		}
		if err := accountService.Deposit(id, req.Amount); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		resp := fmt.Sprintf("Amount of %.2f successfully deposited", req.Amount)
		w.Write([]byte(resp))
	}
}

// WithdrawHandler godoc
// @Summary Забрать деньги с аккаунта
// @Description Забирает определенную сумму с аккаунта
// @Tags accounts
// @Accept  json
// @Produce plain
// @Param id path int true "Account ID"
// @Param withdraw body float64 true "Withdraw amount"
// @Success 200 {string} string "Withdrew {amount} from account ID: {id}"
// @Failure 400 {string} string "Invalid account ID or amount"
// @Router /accounts/{id}/withdraw [post]
func WithdrawHandler(accountService *AccountService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			http.Error(w, ErrInvalidAccountID.Error(), http.StatusBadRequest)
			return
		}
		var req Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, ErrInvalidAmount.Error(), http.StatusBadRequest)
			return
		}
		if err := accountService.Withdraw(id, req.Amount); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		resp := fmt.Sprintf("Amount of %.2f successfully withdrawn", req.Amount)
		w.Write([]byte(resp))
	}
}


// BalanceHandler godoc
// @Summary Получить баланса аккаунта
// @Description Возвращает баланс аккаунта
// @Tags accounts
// @Produce plain
// @Param id path int true "Account ID"
// @Success 200 {string} string "Account ID: {id} | Balance: {balance}"
// @Failure 400 {string} string "Invalid account ID"
// @Router /accounts/{id}/balance [get]
func GetBalance(accountService *AccountService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			http.Error(w, ErrInvalidAccountID.Error(), http.StatusBadRequest)
			return
		}

		balance, err := accountService.GetBalance(id)
		if err != nil {
			http.Error(w, ErrAccountNotFound.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		resp := fmt.Sprintf("Balance: %.2f", balance)
		w.Write([]byte(resp))
	}
}
