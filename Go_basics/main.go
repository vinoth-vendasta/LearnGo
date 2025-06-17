package main

import (
	"errors"
	"fmt"
	"learngo/go_basics/math"
	"time"
)

type Book struct {
	Title       string
	Author      string
	Year        int
	Price       float64
	IsAvailable bool
}
type Library struct {
	Name     string
	Location string
	Books    []Book
	rating   float64
}

type Notifier interface {
	Notify() string
}

type Email struct {
	sender   string
	receiver string
	subject  string
	body     string
}
type SMS struct {
	sender   string
	receiver string
	message  string
}

func varTypesConst() {
	var str string = "Hello, Vinoth"
	var boolean bool = true
	//const pi float64 = 3.14
	// byte := byte(100)
	// fmt.Println("Byte: ", byte)
	fmt.Println("\n---Variables, Types amd Constants-------------")
	fmt.Println(str)
	fmt.Println(boolean)
}

func ForIfSwitchDefer() {
	fmt.Println("\n----For, If, Switch and Defer-------------")
	for i := range 5 {
		if i == 3 {
			fmt.Println("i is 3")
		}
	}
	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning")
	case t.Hour() < 18:
		fmt.Println("Good Afternoon")
	case t.Hour() < 22:
		fmt.Println("Good Evening")
	default:
		fmt.Println("Good Night")
	}
	defer fmt.Println("defer: This will be printed last")
	fmt.Println("Current time is:", t.Format("15:04:05"))
}

func checkError(msg string) (string, error) {
	fmt.Println("\n----Error Handling-------------")
	if msg == "" {
		return "", errors.New("message cannot be empty")
	}
	return "Hello, " + msg, nil
}
func mapPointerSlice() {
	fmt.Println("\n----Map, Pointer and Slice-------------")
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	m["one"] = 122
	fmt.Println("Map: ", m)

	fmt.Println("\n---- Pointer-------------")
	num := 42
	p := &num
	fmt.Println("Pointer value: ", *p)

	fmt.Println("\n---- Slice-------------")
	s := []int{4, 5, 6, 7, 8}
	tem := s[1:3]
	fmt.Println("Slice: ", tem)
}

func structureExample() {
	fmt.Println("\n----Structure Example-------------")
	bookArr := []Book{
		{Title: "Go Programming", Author: "John Doe", Year: 2021, Price: 29.99, IsAvailable: true},
		{Title: "Learning Go", Author: "Jane Smith", Year: 2020, Price: 39.99, IsAvailable: false},
		{Title: "Advanced Go", Author: "Alice Johnson", Year: 2019, Price: 49.99, IsAvailable: true},
	}
	myBook := Book{Title: "Atomic Habits", Author: "James Clear", Year: 2018, Price: 299.99, IsAvailable: true}
	bookArr = append(bookArr, myBook)

	lib := Library{
		Name:     "Vendasta Library",
		Location: "Chennai",
		Books:    bookArr,
		rating:   9.5,
	}

	fmt.Println("\nMy Library Dertails: *********")
	fmt.Println("Name: ", lib.Name)
	fmt.Println("Location: ", lib.Location)
	fmt.Println("Rating: ", lib.rating)
	fmt.Println("\nBooks: ")

	for _, book := range lib.Books {
		fmt.Println("Title: ", book.Title)
		fmt.Println("Author: ", book.Author)
		fmt.Println("Year: ", book.Year)
		fmt.Println("Price: ", book.Price)
		fmt.Println("Is Available: ", book.IsAvailable)
	}
}
func countCharacter(s string) map[rune]int {
	var counts = make(map[rune]int)
	for _, ch := range s {
		if counts[ch] == 0 {
			counts[ch] = 1
		} else {
			counts[ch]++
		}
	}
	return counts
}

// methods for specific use
func (b Book) PrintBookTitle() string {
	return b.Title
}

func (b *Book) changeBookAuthor(newAuthor string) {
	b.Author = newAuthor
}

func methods() {
	b := Book{Title: "The man called Ove", Author: "Fredrik Backman", Year: 2012, Price: 19.99, IsAvailable: true}
	fmt.Println("\n----Methods-------------")
	fmt.Println("Book Title: ", b.PrintBookTitle())

	b.changeBookAuthor("Vinoth Kumar N")
	fmt.Println("Updated Author: ", b.Author)
}

func (e Email) Notify() string {
	return fmt.Sprintf("Email sent from %s to %s with subject '%s'", e.sender, e.receiver, e.subject)
}
func (s SMS) Notify() string {
	return fmt.Sprintf("SMS sent from %s to %s with message '%s'", s.sender, s.receiver, s.message)
}

func interfaceExample() {
	fmt.Println("\n----Interface Example-------------")
	email := Email{
		sender:   "vinoth",
		receiver: "vendasta.com",
		subject:  "request for Leave",
		body:     "I would like to request a leave for 2 days",
	}
	sms := SMS{
		sender:   "9385562690",
		receiver: "1234567890",
		message:  "Hello, this is a test SMS",
	}
	var notifier1, notifier2 Notifier
	notifier1 = email
	fmt.Println(notifier1.Notify())

	notifier2 = sms
	fmt.Println(notifier2.Notify())
}

func main() {
	var str string = "Hello, Vinoth"
	fmt.Println(str)
	fmt.Println("Addition: :", math.Add(100, 200))
	varTypesConst()
	ForIfSwitchDefer()
	mapPointerSlice()
	structureExample()
	methods()
	interfaceExample()

	fmt.Println("\n----Count Character Example -------------")
	counts := countCharacter("Hello, Vendasta Vinoth")
	for ch, count := range counts {
		fmt.Println("Character: ", string(ch), "Count : ", count)
	}
	//fmt.Println("Character Counts: ", counts)

	// Error Handling
	if msg, err := checkError(""); err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(msg)
	}
}
