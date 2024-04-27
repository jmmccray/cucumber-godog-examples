package main

// Date: 4.19.24
// Interviewer: Amrit Khadka
// Description: Create a service that returns private banking information from CapitalOne Customers
// Sources: https://github.com/cucumber/godog/tree/main
import (
	"errors"
	"fmt"
)

type Customer struct {
	Name          string
	SSN           int
	FICOScore     int
	NumberOfCards int
}

var CustomerDB = map[int]Customer{1234567890:{Name: "John", SSN: 1234567890, FICOScore: 123, NumberOfCards: 1},
	987654321:{Name: "Jane", SSN: 987654321, FICOScore: 777, NumberOfCards: 1},
	331902445:{Name: "Bill", SSN: 331902445, FICOScore: 501, NumberOfCards: 3},
	2247071331:{Name: "Sarah", SSN: 2247071331, FICOScore: 850, NumberOfCards: 2},
	9917071331:{Name: "Matthew", SSN: 9917071331, FICOScore: 620, NumberOfCards: 1},
}

func (p Customer) GetFICOScore() int {
	if p.FICOScore < 0 {
		return -1
	}
	return p.FICOScore
}

func (p Customer) GetCustomerSSN() int {
	if p.SSN <= 0 {
		return -1
	}
	return p.SSN
}

func CheckSSNValid(ssn int) bool {
	if _, ok := CustomerDB[ssn]; ok{
		return true
	}
	return false
}

func GetFICOScoreFromSSN(ssn int) int {
	if _, ok := CustomerDB[ssn]; ok {
		return CustomerDB[ssn].FICOScore
	}
	return -1
}

func AddCustomer(customer Customer) error {
	if customer.SSN == 0 {
		return errors.New("ssn is 0 so it's not valid")
	}

	if _, ok := CustomerDB[customer.SSN]; !ok {
		CustomerDB[customer.SSN] = customer
		return nil
	}
	return errors.New("ssn of a customer already exists in database.")
}

func GetCustomers() map[int]Customer {
	return CustomerDB
}

func main() {

	John := Customer{SSN: 123456789, FICOScore: 830, NumberOfCards: 2}
	Jane := Customer{}

	// Print John's FICO Score
	fmt.Printf("John's FICO Score: %d\n", John.GetFICOScore())

	// Print Jane FICO score
	fmt.Printf("Jane FICO Score: %d\n", Jane.GetFICOScore())

	// Print Jon's FICO score from SSN
	fmt.Printf("John's FICO Score: %d\n", GetFICOScoreFromSSN(123456789))

}
