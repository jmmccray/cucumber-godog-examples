package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/cucumber/godog"
	"fmt"
)

type asserter struct {
	err error
}

func (a *asserter) Errorf(format string, args ...interface{}) {
	a.err = fmt.Errorf(format, args...)
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
	  ScenarioInitializer: InitializeScenario,
	  Options: &godog.Options{
		Format:   "pretty",
		Paths:    []string{"features"},
		TestingT: t,
	  },
	}
  
	if suite.Run() != 0 {
	  t.Fatal("non-zero status returned, failed to run feature tests")
	}
  }

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^the SSN (\d+) from the ClientDatabase$`, testSSNExistsInClientDatabase)
	ctx.Step(`^the FICOScore should be (\d+)$`, theFICOScoreShouldBe)
	ctx.Step(`^the Customer\'s name attached to (\d+) is "([^"]*)"$`, testNameOfSSN)
}

func testSSNExistsInClientDatabase(ssn int) error {
	err := assertExpectedAndActual(
		assert.Equal, true, CheckSSNValid(ssn),
		"The SSN should be %d", ssn,
	)

	if err != nil {
		return err
	}
	return nil
}

func testNameOfSSN(ssn int, name string) error {

	err := assertExpectedAndActual(
		assert.Equal, name, CustomerDB[ssn].Name ,
		"The Customer's name should be %s", name,
	)
	if err != nil {
		return err
	}
	return nil
}

func theFICOScoreShouldBe(fico int) error {
	return assertExpectedAndActual(
		assert.Equal, fico, 123,
		"Then the FICOScore should be %d", 501,
	)
}

func assertExpectedAndActual(a expectedAndActualAssertion, expected, actual interface{}, msgAndArgs ...interface{}) error {
	var t asserter
	a(&t, expected, actual, msgAndArgs...)
	return t.err
}

type expectedAndActualAssertion func(t assert.TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool

func TestFICOScore(t *testing.T) {
	// Test if FICO Score is positive, test if its certain value, test if no FICO score.
	tests := []Customer{{
		FICOScore: 830,
	},
		{
			FICOScore: -1,
		},
		{},
	}

	expected := []int{
		830,
		-1,
		0,
	}

	for i, test := range tests {
		assert.Equal(t, expected[i], test.GetFICOScore())
	}
}
