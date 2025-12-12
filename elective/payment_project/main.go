package paymentService

import (
	"errors"
	"fmt"
	"log"
	"time"
)

// Money - специальный тип деньги - в центах
type Money int

// PaymentStatus статус платежа
type PaymentStatus int

// PaymentType - тип платежа
type PaymentType int

const (
	STATUS_NEW     PaymentStatus = 1
	STATUS_PROCESS PaymentStatus = 2
	STATUS_SUCCESS PaymentStatus = 3
	STATUS_FAILED  PaymentStatus = 4

	PAYMENT_TYPE_IN  PaymentType = 1
	PAYMENT_TYPE_OUT PaymentType = 2
)

var (
	ErrPermissionDenied        = errors.New("permission denied")         // ошибка если недостаточно прав
	ErrInsufficientFunds       = errors.New("insufficient amount")       // ошибка если недостаточно средств
	ErrAccountBlocked          = errors.New("account is blocked")        // ошибка если аккаунт блокирован
	ErrInvalidStatusTransition = errors.New("invalid status transition") // ошибка статуса платежа
)

// Account - entity аккаунта с балансаом
type Account struct {
	Id          string    // id uuid
	Title       string    // название аккаунта
	Balance     Money     // баланс
	Blocked     bool      // статус блокировки
	BlockReason string    // причина блокировки
	CreatedAt   time.Time // дата создания
	UpdatedAt   time.Time // дата обновления
}

type Payment struct {
	Id        string        // id uuid (можно просто любой)
	Amount    Money         // сумма платежа
	Account   Account       // аккаунт на который или с которого зачиляются средства
	Type      PaymentType   // тип операции PAYMENT_TYPE_IN PAYMENT_TYPE_OUT
	Status    PaymentStatus // статус платежа STATUS_NEW | STATUS_PROCESS | STATUS_SUCCESS | STATUS_FAILED
	CreatedAt time.Time     // дата создания
	UpdatedAt time.Time     // дата обновления платежа
}

type User struct {
	Id   string
	Name string
}

type Operator struct {
	User
	CanCreateAccounts bool
	CanTransferMoney  bool
	CanBlockAccounts  bool
}

type AccountService interface {
	// CreateNewAccount - созаем новый аккаунт с указанным балансом
	CreateNewAccount(amount Money) (*Account, error)
	UpdateAccountBalance(account *Account, amount Money) error
	TransferMoney(from *Account, to *Account, amount Money) error
	BlockAccount(account *Account, reason string) error
	SetPaymentService(paymentService IPaymentService)
	SetOperator(operator Operator)
	SetLogger(logger ILogger)
}

type IPaymentService interface {
	// CreateNewPayment - создание платежа,
	// возвращает новый платеж с указанием аккаунта и суммой в статусе new
	CreateNewPayment(accountOut Account, amount Money, paymentType PaymentType, comment string) (*Payment, error)
	// UpdatePaymentStatus - автомат статусов, переводит платеж по жизненному циклу
	UpdatePaymentStatus(payment *Payment, status PaymentStatus) error
	SetLogger(logger ILogger)
	SetOperator(operator Operator)
}

type ILogger interface {
	Info(msg string)
	Fatal(msg string)
}

func BusinessLogic(operator Operator, accountService AccountService, paymentService IPaymentService, log ILogger) {
	// TODO сделай так, чтобы это заработало с твоими реализацими
	// хранение аккаунтов логов и платежей сделать in memory

	// даем сервису платежей свои зависимости
	paymentService.SetLogger(log)
	paymentService.SetOperator(operator)

	// даем сервису аккаунтов свои зависимости
	accountService.SetPaymentService(paymentService)
	accountService.SetOperator(operator)
	accountService.SetLogger(log)

	// созлаем 2 аккаунта ---------

	// создаем первый аккаунт, проверяем права на кадлое действие у оператора
	// внутри метода CreateNewAccount - проводим операцию пополнения через сервис платежей и смену статуса через автомат
	// статус платежа может устанавливаться только в определенной последовательности new -> process -> (success | failed)
	// все должно логироваться и сопровождаться соответсвующими ошибками
	accountA, err := accountService.CreateNewAccount(Money(100000))
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	log.Info("Account A created")

	// создаем второй аккаунт соответственно, сопровождая логами и ошибками
	accountB, err := accountService.CreateNewAccount(Money(0))
	if err != nil {
		e := accountService.BlockAccount(accountA, "второй аккаунт не смог создаться, по этому первый заблокировали")
		log.Fatal(err.Error())
		if e != nil {
			log.Fatal(e.Error())
		}
		return
	}
	log.Info("account B created")

	// делаем перевод между аккаунтами
	log.Info("creating transfer from account A to account B")
	err = accountService.TransferMoney(accountA, accountB, Money(1000))
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	// ! примечание ! все сервисы должны поддерживать ошибки
	//если у оператора нет CanCreateAccounts — CreateNewAccount возвращает ошибку;
	//если нет CanTransferMoney — ошибка в TransferMoney;
	//если нет CanBlockAccounts — ошибка в BlockAccount;
	//если баланс меньше суммы перевода — ошибка ErrInsufficientFunds;
	//если аккаунт заблокирован — ошибки в операциях по нему;
	//если нарушен автомат статусов (например, NEW -> SUCCESS напрямую) — ErrInvalidStatusTransition.

	// меняйте код с разными вводными чтобы проверить случаи ошибок

}
func NewOperator(name string) Operator {
	user := User{Id: "1234", Name: name}
	return Operator{
		User:              user,
		CanCreateAccounts: true,
		CanTransferMoney:  true,
		CanBlockAccounts:  true,
	}
}

func NewLogger(driver string) ILogger {
	switch driver {
	case LOG_DRIVER_STDOUT:
		return &LoggerStdout{}
	case LOG_DRIVER_MOCK:
		return &LoggerMock{}
	case LOG_DRIVER_COMBINED:
		return &LoggerCombined{&LoggerStdout{}, &LoggerMock{}}
	default:
		return nil
	}
}

type LoggerMock struct {
	log []string
}

func (l LoggerMock) Info(msg string) {
	log.Println(fmt.Sprintf("[INFO] " + msg))
}

func (l LoggerMock) Fatal(msg string) {
	log.Println(fmt.Sprintf("[FATAL] " + msg))
}

type LoggerStdout struct {
}

func (l LoggerStdout) Info(msg string) {
	log.Println(fmt.Sprintf("[INFO] " + msg))
}

func (l LoggerStdout) Fatal(msg string) {
	log.Println(fmt.Sprintf("[FATAL] " + msg))
}

const LOG_DRIVER_STDOUT = "stdout"
const LOG_DRIVER_MOCK = "mock"
const LOG_DRIVER_COMBINED = "combined"

type Config struct {
	LogDriver string
}

func NewConfig(env string) Config {
	switch env {
	case "dev":
		return Config{
			LogDriver: LOG_DRIVER_STDOUT,
		}

	case "test":
		return Config{
			LogDriver: LOG_DRIVER_MOCK,
		}

	case "qa":
		return Config{
			LogDriver: LOG_DRIVER_COMBINED,
		}
	}
	return Config{
		LogDriver: LOG_DRIVER_STDOUT,
	}
}

type LoggerCombined struct {
	stdout *LoggerStdout
	mock   *LoggerMock
}

func (l *LoggerCombined) Info(msg string) {
	l.stdout.Info(msg)
	l.mock.Info(msg)
}

func (l *LoggerCombined) Fatal(msg string) {
	l.mock.Fatal(msg)
	l.stdout.Fatal(msg)
}

func main() {
	// TODO сделай все реализации и запусти с ними бизнес логику
	//operator := NewOperator(...аргументы),
	//accountService := NewAccountService(...аргументы),
	//paymentService := NewPaymentService(...аргументы),
	cfg := NewConfig("qa")
	operator := NewOperator("Вася")
	lg := NewLogger(cfg.LogDriver)

	BusinessLogic(operator, accountService, paymentService, lg)
}
