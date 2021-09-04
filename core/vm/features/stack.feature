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
    Given a stack is created
    And "<Data>" is pushed
    When pop is called "<NumberOfPops>" times
    Then popped data is "<PopData>"
    And stack size is "<Size>"
    And error should be "<StackError>"
    Examples:
      | Data       | NumberOfPops | PopData    | Size | StackError     |
      | A0,B0,C0   | 2            | C0,B0      | 1    |                |
      | A1,B1,C1   | 3            | C1,B1,A1   | 0    |                |
      | A2,B2,C2   | 4            | C2,B2,A2   | 0    | stack is empty |