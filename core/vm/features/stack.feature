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

