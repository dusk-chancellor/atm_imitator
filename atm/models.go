package atm

// здесь реализованы структура, интерфейс
// и основные методы, описанные в тз(1),
// а также некоторые модели для флоу сервиса

import "sync"

const (
	logChanSize 	= 100 // размер буфера канала логов
	opCreateAccount = "create account"
	opDeposit 		= "deposit"
	opWithdraw 		= "withdraw"
	opGetBalance 	= "get balance"
)

type Account struct {
	ID 		int // айди юзера
	Balance float64 // баланс юзера
	mu 		sync.RWMutex // потокобезопасность
}

type BankAccount interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	GetBalance() float64
}

func (a *Account) Deposit(amount float64) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if amount <= 0 {
		return ErrInvalidAmount
	}
	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if amount <= 0 {
		return ErrInvalidAmount
	}
	if a.Balance < amount {
		return ErrInsufficientFunds
	}
	a.Balance -= amount
	return nil
}

func (a *Account) GetBalance() float64 {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.Balance
}
