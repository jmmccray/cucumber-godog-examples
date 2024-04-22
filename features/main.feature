# features/banking.features
Feature: CapitalOne Banking 
    A banking service for CapitalOne 
    That gets the credit card information
    From customers

    Scenario: GetFICOScore
        Given the SSN of 1234567890 from ClientDatabase
        Then the FICOScore of John should be 123
