package main

import (
	"fmt"
	"time"
)

// --- Часть 1: Greeter ---

type Greeter interface {
	Greet() string
}

type User struct {
	Name     string
	Birthday time.Time
}

type Robot struct {
	Model        string
	CreationDate time.Time
}

func (u User) Greet() string {
	return "Hello " + u.Name
}

func (r Robot) Greet() string {
	return "Hello " + r.Model
}

// --- Часть 2: Worker ---

type Worker interface {
	Work() // 1. Убрали 'string', так как реализация ничего не возвращает
}

type WorkImpl struct {
	WorkDone bool
}

// Метод меняет состояние структуры, поэтому обязательно указатель (w *WorkImpl)
func (w *WorkImpl) Work() {
	fmt.Println("Doing work")
	w.WorkDone = true
}

// Функция-конструктор
func NewWorker() Worker {
	// 2. Исправили опечатку: WorkerImpl -> WorkImpl
	// Возвращаем указатель, так как метод Work определен для *WorkImpl
	return &WorkImpl{}
}

func PrintGreeter(g Greeter) {
	u, ok := g.(User)

	if ok {
		fmt.Println(u.Greet())
	}
}

func PrintType(x interface{}) {

	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case time.Time:
		fmt.Println("time")
	case int8, int16, int32, int64:
		fmt.Println("int8")
	case User:
		fmt.Println("User")
	}
}

type MyType struct {
	Name string
}

func (MyType) Val() {

}

func (*MyType) Ptr() {

}

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Saved    bool
}

type Saver interface {
	Save() error
}

func (c Config) IsSaved() bool {
	return c.Saved
}

func (c Config) Save() error {
	fmt.Println(fmt.Sprintf("Saving config file to %s:%s:%s:%s", c.Host, c.Port, c.User, c.Password))
	c.Saved = true
	return nil
}

func main() {

	var dbConfig = Config{
		Host:     "localhost",
		Port:     "3306",
		User:     "dkzee",
		Password: "dkzee",
	}

	err := dbConfig.Save()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(dbConfig.IsSaved())
	}

	// Проверка Greeter
	var someGreeter Greeter

	someGreeter = Robot{Model: "Robot1", CreationDate: time.Now()}
	fmt.Println(someGreeter.Greet())

	someGreeter = User{Name: "Dima"}
	fmt.Println(someGreeter.Greet())

	// Проверка Worker
	var w Worker
	w = NewWorker() // w теперь хранит указатель на WorkImpl внутри интерфейса
	w.Work()        // Выведет "Doing work"
	var u User = User{Name: "Test"}
	fmt.Println(u)

	PrintType(u)

}
