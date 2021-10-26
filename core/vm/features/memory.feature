Feature: memory

  Scenario: new memory is empty
    When a new memory store is created
    Then store is empty

  Scenario Outline: resize
    Given a new memory store is created
    And size is "<existingSize>"
    When resized to "<newSize>"
    Then updated size is "<updatedSize>"
    Examples:
    |existingSize|newSize|updatedSize|
    | 10         | 20    | 20    |
    | 10         | 10    | 10    |
    | 10         | 5     | 10    |

  @run
  Scenario Outline: Populate memory with data of a given size starting at a given offset
    Given a new memory store is created
    And store is initialized with "<data>"
    When Setting offset "<offset>" and size "<size>" to "<value>"
    Then data should be "<newdata>"
    And error is "<error>"
    Examples:
      |data         |offset|size|value   |newdata                     |error                |
      |A0A1A2A3A4A5 |2     |  0 |        |A0A1A2A3A4A5                |vm.MemorySizeError   |
      |A0A1A2A3A4A5 |20    |  10|        |A0A1A2A3A4A5                |vm.MemoryOffsetError |
      |A0A1A2A3A4A5 |2     |  2 |B1B2    |A0A1B1B2A2A3A4A5            |                     |
      |A0A1A2A3A4A5 |0     |  3 |B1B2B3B4|B1B2B3A0A1A2A3A4A5A6A7A8A9  |                     |
      |A0A1A2A3A4A5 |5     |  4 |B1B2B3B4|A0A1A2A3A4A5A6A7A8A9B1B2B3B4|                     |

    Scenario Outline: Get copy
      Given a new memory store is created
      And store is initialized with "<data>"
      When getting a copy at offset "<offset>" and size "<size>"
      Then data should be "<copieddata>"
      And error is "<error>"
      And data is a copy
      Examples:
      |data                 |offset|size|copieddata          |error             |
      |A0A1A2A3A4A5A6A7A8A9 |2     |0   |                    |MemorySizeError   |
      |A0A1A2A3A4A5A6A7A8A9 |20    |3   |A2A3A4              |MemoryOffsetError |
      |A0A1A2A3A4A5A6A7A8A9 |2     |3   |A2A3A4              |                  |
      |A0A1A2A3A4A5A6A7A8A9 |0     |10  |A0A1A2A3A4A5A6A7A8A9|                  |

  Scenario Outline: Get ptr to data of a given size starting from a given offset
    Given a new memory store is created
    And store is initialized with "<data>"
    When getting a ptr at offset "<offset>" and size "<size>"
    Then data should be "<copieddata>"
    And error is "<error>"
    And data is not a copy
    Examples:
      |data                 |offset|size|copieddata          |error             |
      |A0A1A2A3A4A5A6A7A8A9 |2     |0   |                    |MemorySizeError   |
      |A0A1A2A3A4A5A6A7A8A9 |20    |3   |A2A3A4              |MemoryOffsetError |
      |A0A1A2A3A4A5A6A7A8A9 |2     |3   |A2A3A4              |                  |
      |A0A1A2A3A4A5A6A7A8A9 |0     |10  |A0A1A2A3A4A5A6A7A8A9|                  |
