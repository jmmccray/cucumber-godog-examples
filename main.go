package main

// Date: 4.19.24
// Interviewer: Amrit Khadka
// Description: Create a service that returns private banking information from CapitalOne Customers

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

var CustomerDB = []Customer{{Name: "John", SSN: 1234567890, FICOScore: 123, NumberOfCards: 1},
	{Name: "Jane", SSN: 987654321, FICOScore: 777, NumberOfCards: 1},
	{Name: "Bill", SSN: 331902445, FICOScore: 501, NumberOfCards: 3},
	{Name: "Sarah", SSN: 2247071331, FICOScore: 850, NumberOfCards: 2},
	{Name: "Matthew", SSN: 9917071331, FICOScore: 620, NumberOfCards: 1}}

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

// Checks if SSN exists in CustomerDB
func CheckSSNValid(ssn int) bool {
	for _, customer := range CustomerDB {
		if customer.SSN == ssn {
			return true
		}
	}
	return false
}


func GetFICOScoreFromSSN(ssn int) int {
	for _, person := range CustomerDB {
		if person.SSN == ssn {
			return person.FICOScore
		}
	}
	return -1
}

func CheckFICOScoreFromSSN(ssn int) error {
	for _, person := range CustomerDB {
		if person.SSN == ssn {
			return nil
		}
	}
	return errors.New("False")
}

func AddCustomers() {
	// TODO: add "Customer" to CustomerDB
}

func GetCustomers() []Customer {
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
