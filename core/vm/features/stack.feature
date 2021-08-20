Feature: stack

  Scenario: new stack is empty
    When a stack is created
    Then stack should be empty

