Feature: memory

  Scenario: new memory is empty
    When a new memory store is created
    Then store is empty

