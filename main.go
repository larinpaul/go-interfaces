//// 2023/10/15 // 21:02 //

//// Interfaces in Go

// Interfaces are collections of method signatures. A type "implements" an
// interface if it has all of the methods of the given interface defined on it.

// In the followign example, a "shape" must be able to return its area and
// perimeter. Both rect and circle fulfill the interface.

// type shape interface {
// 	area() float64
// 	perimeter() float64
// }

// type rect struct {
// 	width, height float64
// }

// func (r rect) area() float64 {
// 	return r.width * r.height
// }
// func (r rect) perimeter() float64 {
// 	return 2*r.width + 2*r.height
// }

// type circle struct {
// 	radius float64
// }

// func (c circle) area() float64 {
// 	return math.Pi * c.radius * c.radius
// }
// func (c circle) perimeter() float64 {
// 	return 2 * math.Pi * c.radius
// }

//// Assignment

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func sendMessage(msg message) {
// 	fmt.Println(msg.getMessage())
// }

// type message interface {
// 	getMessage() string
// }

// type birthdayMessage struct {
// 	birthdayTime  time.Time
// 	recipientName string
// }

// func (bm birthdayMessage) getMessage() string {
// 	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
// }

// type sendingReport struct {
// 	reportName    string
// 	numberOfSends int
// }

// func (sr sendingReport) getMessage() string {
// 	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
// }

// func test(m message) {
// 	sendMessage(m)
// 	fmt.Println("======================================")
// }

// func main() {
// 	test(sendingReport{
// 		reportName:    "First Report",
// 		numberOfSends: 10,
// 	})
// 	test(birthdayMessage{
// 		recipientName: "John Doe",
// 		birthdayTime:  time.Date(1994, 03, 21, 0, 0, 0, 0, time.UTC),
// 	})
// 	test(sendingReport{
// 		reportName:    "First Report",
// 		numberOfSends: 10,
// 	})
// 	test(birthdayMessage{
// 		recipientName: "Bill Deer",
// 		birthdayTime:  time.Date(1934, 05, 01, 0, 0, 0, 0, time.UTC),
// 	})
// }

//// 22:09

// Interface Implementation
// Interfaces are implemented implicitly.

// A type never declares that it implements a given interface. If an interface
// exists and a type has the proper methods defined, then the type
// automatically fulfills that interface.

// // Assignment

// package main

// import (
// 	"fmt"
// )

// type employee interface {
// 	getName() string
// 	getSalary() int
// }

// type contractor struct {
// 	name         string
// 	hourlyPay    int
// 	hoursPerYear int
// }

// func (c contractor) getName() string {
// 	return c.name
// }

// func (c contractor) getSalary() int {
// 	return c.hourlyPay * c.hoursPerYear
// }

// type fullTime struct {
// 	name   string
// 	salary int
// }

// func (ft fullTime) getSalary() int {
// 	return ft.salary
// }

// func (ft fullTime) getName() string {
// 	return ft.name
// }

// func test(e employee) {
// 	fmt.Println(e.getName(), e.getSalary())
// 	fmt.Println("======================================")
// }

// func main() {
// 	test(fullTime{
// 		name:   "Jack",
// 		salary: 50000,
// 	})
// 	test(contractor{
// 		name:         "Bob",
// 		hourlyPay:    100,
// 		hoursPerYear: 73,
// 	})
// 	test(contractor{
// 		name:         "Jill",
// 		hourlyPay:    872,
// 		hoursPerYear: 982,
// 	})
// }

//// 22:17

//// Multiple Interfaces

// A type can implement any number of interfaces in Go. For example, the
// empty interface, interface{}, is always implemented by every type
// because it has no requirements.

// Assignment
// Add the required methods so that the email type implements bot ht
// expense and printer interfaces.

// cost()
// If the email is not "subscribed", then the cost is 0.04 for each character in
// the body. If it is, then the cost if 0.01 per character.

// print()
// The print method should print to standard out the email's body text.

// package main

// import (
// 	"fmt"
// )

// // func (e email) cost() float64 {
// // 	if !e.isSubsribed {
// // 		return 0.05 * float64(len(e.body))
// // 	}
// // 	return 0.01 * float64(len(e.body))
// // }

// func (e email) cost() float64 {
// 	if !e.isSubsribed {
// 		return float64(len(e.body)) * .05
// 	}
// 	return float64(len(e.body)) * .01
// }

// func (e email) print() {
// 	fmt.Println(e.body)
// }

// // don't touch below this line

// type expense interface {
// 	cost() float64
// }

// type printer interface {
// 	print()
// }

// type email struct {
// 	isSubsribed bool
// 	body        string
// }

// func print(p printer) {
// 	p.print()
// }

// func test(e expense, p printer) {
// 	fmt.Printf("Printing with cost: $%.2f ...\n", e.cost())
// 	p.print()
// 	fmt.Println("====================================")
// }

// func main() {
// 	e := email{
// 		isSubsribed: true,
// 		body:        "hello there",
// 	}
// 	test(e, e)
// 	e = email{
// 		isSubsribed: false,
// 		body:        "I want my money back",
// 	}
// 	test(e, e)
// 	e = email{
// 		isSubsribed: true,
// 		body:        "Are you free for a chat?",
// 	}
// 	test(e, e)
// 	e = email{
// 		isSubsribed: false,
// 		body:        "This meeting could have been an email",
// 	}
// 	test(e, e)
// }

//// 23:05

// Name your Interface arguments

// type Copier interface {
// 	Copy(string, string) int
// }

// type Copier interface {
// 	Copy(sourceFile string, destinationFile string) (bytesCopied int)
// }

//// 23:07

//// Type Assertions in Go

// When workign with interfaces in Go, every once-in-awhile you'll need
// access to the underlying type of an interface value. You can cast an
// interface to its underlying type using a type assertion.

// type shape interface {
// 	area() float64
// }

// type circle struct {
// 	radius float64
// }

// "c" is a new circle cast from "s"
// which is an instance of a shape.
// "ok" is a bool that is true if s was a cricle
// or false if s isn't a circle

// Assignment

// Implement the getExpenseReport function.
// * If the expense is an email then it should return the email's
// toAddress and the cost of the email.
// * If the expense is an sms then it should return the sms's
// toPhoneNumber and its cost.
// * If the expense has any other underlying type, jsut return an empty
// string and 0.0 for the cost.

// package main

// import (
// 	"fmt"
// )

// func getExpanseReport(e expense) (string, float64) {
// 	em, ok := e.(email)
// 	if ok {
// 		return em.toAddress, em.cost()
// 	}
// 	s, ok := e.(sms)
// 	if ok {
// 		return s.toPhoneNumber, s.cost()
// 	}
// 	return "", 0.0
// }

// func (e email) cost() float64 {
// 	if !e.isSubscribed {
// 		return float64(len(e.body)) * .05
// 	}
// 	return float64(len(e.body)) * .01
// }

// func (s sms) cost() float64 {
// 	if !s.isSubscribed {
// 		return float64(len(s.body)) * .1
// 	}
// 	return float64(len(s.body)) * .1
// }

// func (i invalid) cost() float64 {
// 	return 0.0
// }

// type expense interface {
// 	cost() float64
// }

// type email struct {
// 	isSubscribed bool
// 	body         string
// 	toAddress    string
// }

// type sms struct {
// 	isSubscribed  bool
// 	body          string
// 	toPhoneNumber string
// }

// type invalid struct{}

// func estimateYearlyCost(e expense, averageMessagesPerYear int) float64 {
// 	return e.cost() * float64(averageMessagesPerYear)
// }

// func test(e expense) {
// 	address, cost := getExpanseReport(e)
// 	switch e.(type) {
// 	case email:
// 		fmt.Printf("Report: The email foing to %s will cost: %.2f\n", address, cost)
// 		fmt.Println("=====================================")
// 	case sms:
// 		fmt.Printf("Report: The sms going to %s wil lcost: %.2f\n", address, cost)
// 		fmt.Println("=====================================")
// 	default:
// 		fmt.Println("Report: Invalid expense")
// 		fmt.Println("=====================================")
// 	}
// }

// func main() {
// 	test(email{
// 		isSubscribed: true,
// 		body:         "hello there",
// 		toAddress:    "john@does.com",
// 	})
// 	test(email{
// 		isSubscribed: false,
// 		body:         "This meeting could have been an email",
// 		toAddress:    "jave@doe.com",
// 	})
// 	test(email{
// 		isSubscribed: false,
// 		body:         "This meeting could have been an email",
// 		toAddress:    "elon@doe.com",
// 	})
// 	test(sms{
// 		isSubscribed:  false,
// 		body:          "This meeting could have been an email",
// 		toPhoneNumber: "+155555509832",
// 	})
// 	test(sms{
// 		isSubscribed:  false,
// 		body:          "This meeting could have been an email",
// 		toPhoneNumber: "+155555504444",
// 	})
// 	test(invalid{})
// }

//// 23:21

//// Type Switches

// A type switch makes it easy to do several type assertions in a series.

// A type switch is similar to a regular switch statement, but the cases specify
// // types instead of values.

// func printNumericValue(num interface{}) {
// 	switch v := num.(type) {
// 	case int:
// 		fmt.Printf("%T\n", v)
// 	case string:
// 		fmt.Printf("%T\n", v)
// 	default:
// 		fmt.Printf("%T\n", v)
// 	}
// }

// func main() {
// 	printNumericValue(1)
// 	// prints "int"

// 	printNumericValue("1")
// 	// prints "string"

// 	printNumericValue(struct{}{})
// 	// prints "struct {}"

// }]

// Assignment

// Aften submitting our last snippet of code for review, a more experienced gopher told us to use a type
// switch instead of successive assertions. Let's make that improvement!

// Impelement the getExpenseReport function using a type switch.
// * If the expense is an email then it should return the email's toAddress and the cost of the email.
// * If the expense is an sms then it should return the sms's toPhoneNumber and its cost.
// * If the expense has any other underlying type, just return an empoty string and 0.0 for the cost.

package main

import (
	"fmt"
)

func getExpenseReport(e expense) (string, float64) {
	switch v := e.(type) {
	case email:
		return v.toAddress, v.cost()

	case sms:
		return v.toPhoneNumber, v.cost()
	default:
		return "", 0.0
	}
}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func estimateYearlyCost(e expense, averageMessagePerYear int) float64 {
	return e.cost() * float64(averageMessagePerYear)
}

func test(e expense) {
	address, cost := getExpenseReport(e)
	switch e.(type) {
	case email:
		fmt.Printf("Report: The email going to %s will cost: %.2f\n", address, cost)
		fmt.Println("====================================")

	case sms:
		fmt.Printf("Report: The sms going to %s will cost: %.2f\n", address, cost)
		fmt.Println("==================================")
	default:
		fmt.Println("Report: Invalid expense")
		fmt.Println("==================================")
	}
}

func main() {
	test(email{isSubscribed: true,
		body:      "hello there",
		toAddress: "john@does.com",
	})
	test(email{isSubscribed: false,
		body:      "This could have been an email",
		toAddress: "jave@doe.com",
	})
	test(email{
		isSubscribed: false,
		body:         "Wanna catch up later?",
		toAddress:    "elon@doe.com",
	})
	test(sms{
		isSubscribed:  false,
		body:          "I'm a Nigerian prince, please send me your bank info so I can deposit $1000 dollars",
		toPhoneNumber: "+155555509832",
	})
	test(sms{
		isSubscribed:  false,
		body:          "i don't need this",
		toPhoneNumber: "+15555504444",
	})
	test(invalid{})
}

//// 23:54

//// Clean Interfaces

// Writing clean interfaces is hard. Frankly, anytime you're dealing with abstractions in code, the simple can
// become complex very quickly if you're not careful. Let's go over some rules of thumb for keeping interfaces
// clean.

// 1. Keep Interfaces Small

// If there is only one piece of advice that you take away from this article, make it this: keep interfaces small!
// Interfaces are meant to define the minimal behavior necessary to accurately represent an idea or concept.

// Here is an example fro mthe stanard HTTP package of a larger interface that's a good example of defining
// minimal behavior:

// type File interface {
// 	io.Closer
// 	io.Reader
// 	io.Seeker
// 	Readdir(count int) ([]os.FileInfo, error)
// 	Stat() (os.FileInfo, error)
// }

// Any type that satisfies the interface's behaviors can be considered by the HTTP pacakge as a File. This is
// convenient because the HTTP package doesn't need to knwo if it's dealing with a file on disk, a network
// buffer, or a simple []byte.

// 2. Interfaces Should Have No Knowledge of Satisfying Types

// An interface should define what is necesary for other types to classify as a member of that interfafe. They
// shouldn't be away of any types that happen to satisfy the interface at design time.

// // Bad example:
// type car interface {
// 	Color() string
// 	Speed() int
// 	IsFiretruct() bool
// }
// // Good example:
// type firetruck interface {
// 	car
// 	HoseLength() int
// }

// 3. Interfaces Are Not Classes
// * Interfaces are not classes, they are slimmer.
// * Interfaces don't have constructors or deconstructors that require that data is created or destroyed.
// * Interfaces aren't hierarchical by nature, though there is syntactic sugar to create interfaces that happen
// to be supersets of other interfaces.
// * Interfaces define function signatures, but not underlying behavior. Making an interface often won't
// DRY up your code in regards to struct methods. For example, if five types satisfy the fmt.Stringer
// interface, they all need their own version of the String() function.
