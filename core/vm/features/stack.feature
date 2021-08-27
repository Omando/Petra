Feature: stack

  Scenario: new stack is empty
    When a stack is created
    Then stack should be empty

  Scenario: Peek on empty stack returns error
    Given an empty stack
    When peek is called
    Then  error should be:
    """
    stack is empty
    """

  Scenario: Pop on empty stack returns error
    Given an empty stack
    When pop is called
    Then  error should be:
      """
      stack is empty
      """

  Scenario Outline: pushes and pops
    Given "<Data>" is pushed
    When pop is called "<NumberOfPops>" times
    Then popped data is "<PopData>"
    And stack size is "<Size>"
    And error should be "<Error>"
    Examples:
      | Data      | NumberOfPops | PopData | Size | Error|
      | 1,2,3     | 2            | 3,2     | 1    | nil  |
