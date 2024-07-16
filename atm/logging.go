package atm

// здесь реализован логгер(канал логов)

import (
	"fmt"
	"log"
	"time"
)
// общая структура логов
type Logger struct {
	AccountID int
	Operation string
	Amount    float64
	Time	  time.Time
	Error     error
}

var LogChan = make(chan Logger, logChanSize)

func InitLogger() { // постоянная бэкграунд горутина записи логов
	go func()  {
		log.Println("Starting logger")
		for entry := range LogChan {
			log.Printf("Account: %d | Operation: %s | Amount: %.2f | Time: %s | Error: %v\n",
			entry.AccountID,
			entry.Operation,
			entry.Amount,
			entry.Time.Format("2006-01-02 15:04:05"),
			entry.Error,
			)
		}
	}()
}

func LogOperation(accountID int, operation string, amount float64, err error) {
	entry := Logger{
		AccountID: accountID,
		Operation: operation,
		Amount: amount,
		Time: time.Now(),
		Error: err,
	}

	select {
	case LogChan <- entry:
		// успешно записан лог
	default:
		// канал полный, лог не записался
		fmt.Println("LogChan is full, log entry discarded")
	}
}
