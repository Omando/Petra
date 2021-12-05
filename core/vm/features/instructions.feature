Feature: Instructions

  @run
  Scenario Outline: Add
    Given "<x>" and "<y>" operands
    When Add is called
    Then result is "<sum>"
    Examples:
    |x    |y    |sum  |
    |01   |02   |3    |
    |aaaa |bbbb |16665|
