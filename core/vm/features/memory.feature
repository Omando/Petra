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


  Scenario Outline: Populate memory with data of a given size starting at a given offset
    Given a new memory store is created
    And store is initialized with "<data>"
    When Setting offset "<offset>" and size "<size>" to "<value>"
    Then data should be "<newdata>"
    And error is "<error>"
    Examples:
      |data         |offset|size|value |newdata     |error                |
      |A0A1A2A3A4A5 |2     |  0 |      |A0A1A2A3A4A5|*vm.MemorySizeError  |
      |A0A1A2A3A4A5 |20    |  10|      |A0A1A2A3A4A5|*vm.MemoryOffsetError|
      |A0A1A2A3A4A5 |0     |  6 |QWERTY|QWERTYA3A4A5|                     |
      |A0A1A2A3A4A5 |2     |  4 |QWERTY|A0QWERA3A4A5|                     |
      |A0A1A2A3A4A5 |5     |  4 |QWERTY|A0A1AQWER4A5|                     |

    Scenario Outline: Get copy
      Given a new memory store is created
      And store is initialized with "<data>"
      When getting a copy at offset "<offset>" and size "<size>"
      Then copied data should be "<copieddata>"
      And error is "<error>"
      And data is a copy
      Examples:
      |data                 |offset|size|copieddata  |error                 |
      |A0A1A2A3A4A5A6A7A8A9 |2     |0   |            |*vm.MemorySizeError   |
      |A0A1A2A3A4A5A6A7A8A9 |20    |3   |            |*vm.MemoryOffsetError |
      |A0A1A2A3A4A5A6A7A8A9 |0     |10  |A0A1A2A3A4  |                      |
      |A0A1A2A3A4A5A6A7A8A9 |2     |6   |A1A2A3      |                      |

  @run
  Scenario Outline: Get ptr to data of a given size starting from a given offset
    Given a new memory store is created
    And store is initialized with "<data>"
    When getting a ptr at offset "<offset>" and size "<size>"
    Then getptr data should be "<copieddata>"
    And error is "<error>"
    And data is not a copy
    Examples:
      |data                 |offset|size|copieddata  |error                |
      |A0A1A2A3A4A5A6A7A8A9 |2     |0   |            |*vm.MemorySizeError  |
      |A0A1A2A3A4A5A6A7A8A9 |20    |3   |            |*vm.MemoryOffsetError|
      |A0A1A2A3A4A5A6A7A8A9 |0     |10  |A0A1A2A3A4  |                     |
      |A0A1A2A3A4A5A6A7A8A9 |2     |6   |A1A2A3      |                     |
