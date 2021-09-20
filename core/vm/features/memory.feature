Feature: memory

  Scenario: new memory is empty
    When a new memory store is created
    Then store is empty

  Scenario Outline: resize
    Given a new memory store is created
    And size is "<oldSize>"
    When resized to "<newSize>"
    Then updated size is "<updatedSize>"
    Examples:
    |oldSize|newSize|updatedSize|
    | 10    | 20    | 20    |
