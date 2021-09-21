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
    | 10    | 10    | 10    |
    | 10    | 5     | 10    |

    Scenario Outline: Get copy
      Given a memory is created and initialized with "<data>"
      When getting a copy at offset "<>" and size "<>"
      Then data should be "<copieddata>"
      And error is "<error>"
      And data is a copy
      Examples:
      |data                 |offset|size|copieddata          |error                                 |
      |A0A1A2A3A4A5A6A7A8A9 |2     |0   |                    |size is zero                          |
      |A0A1A2A3A4A5A6A7A8A9 |20    |3   |A2A3A4              |offset 20 + size 3 is > data length 10|
      |A0A1A2A3A4A5A6A7A8A9 |2     |3   |A2A3A4              |                                      |
