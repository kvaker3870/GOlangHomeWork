package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	Name    string
	Age     int
	Gender  string
	Address string
}

func (p *Person) Birthday() {
	p.Age++
}

func (p Person) Birthday2() {
	p.Age++
	fmt.Println("Not linked birthday increment " + strconv.Itoa(p.Age))
}

func (p *Person) SayHelloAndAge() {
	fmt.Println("Hi I'm "+p.Name+" and I just turned", p.Age, "years old")
}

func main() {

	var personNew = Person{}

	personNew.Birthday()
	fmt.Println(personNew)

	var personNew2 = new(Person)

	personNew2.Birthday()
	fmt.Println(personNew2)

	person1 := Person{
		Name:    "Dmitriy",
		Age:     34,
		Gender:  "Male",
		Address: "Kaskelen",
	}

	person1.Birthday()
	person1.SayHelloAndAge()

	person2 := Person{
		Name:    "Dima",
		Age:     33,
		Gender:  "Male",
		Address: "Kaldayakov",
	}
	person2.Birthday()
	person2.SayHelloAndAge()

	person3 := &Person{
		Name: "Dima3",
		Age:  31,
	}

	person3.Birthday()
	person3.SayHelloAndAge()

	person4 := person3

	person4.Name = "Dima4(copy of person3)"
	person4.Birthday2()
	person4.SayHelloAndAge()

	//полка
	bookShelf := &BookShelf{}

	book1 := &Book{}
	book1.Title = "Pohui Book"
	book1.Author = "Author Unknown"

	book2 := Book{}
	book2.Title = "Pirate bay"
	book2.Author = "Dmitriy Kauts"

	//copy
	book3 := *book1
	book3.Title = "Sasai Masai"

	book4 := new(Book)

	book4.Title = "Golden Goose Part 4"

	setBookOnShelf(book3, bookShelf)
	setBookOnShelf(*book1, bookShelf)
	setBookOnShelf(book2, bookShelf)

	fmt.Print("Books with letter S:")
	bookShelf.getBooksFromBookshelfByAlphabet('S')
	fmt.Print("Books with letter P:")
	bookShelf.getBooksFromBookshelfByAlphabet('P')

	fmt.Printf("%s\n%s\n%s\n%s\n%s\n", "ALLBOOOKS:", book1.Title, book2.Title, book3.Title, book4.Title)

}
