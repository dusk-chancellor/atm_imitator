package atm

// services выполняет основную логики сервиса
// оно связывает модельные методы с хендлерами

import (
	"log"
	"sync"
)

type AccountService struct {
	accounts  map[int]*Account // мап хранения аккаунтов в кеше
	idCounter int // автоинкрементирование айди
	mu 		  sync.Mutex // потокобезопасность
	operations chan func() // канал операций
}

func NewAccountService() *AccountService {
	s := &AccountService{
		accounts: make(map[int]*Account),
		idCounter: 1,
		operations: make(chan func()),
	}
	log.Println("Starting account service")
	go s.processOperations() // запускаем все операции в горутине
	return s
}

func (s *AccountService) processOperations() { // процессирование операций с канала
	for op := range s.operations {
		op()
	}
}

// CreateAccount
func (s *AccountService) CreateAccount() *Account {
	s.mu.Lock()
	defer s.mu.Unlock()
	account := &Account{
		ID: s.idCounter,
	}
	s.idCounter++
	s.accounts[account.ID] = account
	LogOperation(account.ID, opCreateAccount, 0, nil)
	return account
}

// Deposit
func (s *AccountService) Deposit(id int, amount float64) error {
	res := make(chan error) // канал для ретурна
	s.operations <- func() { // передается операция в канал для процессирования
		s.mu.Lock()
		account, exists := s.accounts[id]
		s.mu.Unlock()
		if !exists {
			res <- ErrAccountNotFound
			LogOperation(id, opDeposit, amount, ErrAccountNotFound)
			return
		}
		res <- account.Deposit(amount)
	}
	LogOperation(id, opDeposit, amount, <-res)
	return <-res
}

func (s *AccountService) Withdraw(id int, amount float64) error {
	res := make(chan error)
	s.operations <- func() {
		s.mu.Lock()
		account, exists := s.accounts[id]
		s.mu.Unlock()
		if !exists {
			res <- ErrAccountNotFound
			LogOperation(id, opWithdraw, amount, ErrAccountNotFound)
			return
		}
		res <- account.Withdraw(amount)
	}
	LogOperation(id, opWithdraw, amount, <-res)
	return <-res
}

func (s *AccountService) GetBalance(id int) (float64, error) {
	resStruct := make(chan struct{
		balance float64
		err 	error
	})
	s.operations <- func() {
		s.mu.Lock()
		account, exists := s.accounts[id]
		s.mu.Unlock()
		if !exists {
			resStruct <- struct{
				balance float64
				err 	error
			}{
				balance: 0,
				err: ErrAccountNotFound,
			}
			LogOperation(id, opGetBalance, 0, ErrAccountNotFound)
			return
		}
		resStruct <- struct{
			balance float64
			err 	error
		}{
			balance: account.GetBalance(),
			err: nil,
		}
	}
	res := <-resStruct
	LogOperation(id, opGetBalance, res.balance, res.err)
	return res.balance, res.err
}
