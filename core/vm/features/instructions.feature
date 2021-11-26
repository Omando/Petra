Feature: Instructions

  Scenario Outline: Add
    Given "<x>" and "<y>" operands
    When Add is called
    Then result is "<sum>"
    Examples:
    |x   |y   |sum  |
    |1   |2   |3    |
    |aaaa|bbbb|16665|
