# features/banking.features
Feature: CapitalOne Banking 
    A banking service for CapitalOne 
    That gets the credit card information
    From customers

    Scenario: Check the FICO Score of a Customer's SSN
        Given the SSN 1234567890 from the ClientDatabase
        Then the FICOScore should be 123
        And the Customer's name attached to 1234567890 is "John"