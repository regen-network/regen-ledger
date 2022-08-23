Feature: Class

  Scenario: a valid class
    Given the class
    """
    {
      "key": 1,
      "id": "C01",
      "admin": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
      "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
      "credit_type_abbrev": "C"
    }
    """
    When the class is validated
    Then expect no error

  Scenario: an error is returned if key is empty
    Given the class
    """
    {}
    """
    When the class is validated
    Then expect the error "key cannot be zero: parse error"

  Scenario: an error is returned if id is empty
    Given the class
    """
    {
      "key": 1
    }
    """
    When the class is validated
    Then expect the error "class id cannot be empty: parse error"

  Scenario: an error is returned if id is not formatted
    Given the class
    """
    {
      "key": 1,
      "id": "foo"
    }
    """
    When the class is validated
    Then expect the error "class ID didn't match the format: expected A00, got foo: parse error"

  Scenario: an error is returned if admin is empty
    Given the class
    """
    {
      "key": 1,
      "id": "C01"
    }
    """
    When the class is validated
    Then expect the error "admin: empty address string is not allowed: parse error"

  Scenario: an error is returned if metadata exceeds 256 characters
    Given the class
    """
    {
      "key": 1,
      "id": "C01",
      "admin": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y="
    }
    """
    And metadata with length "257"
    When the class is validated
    Then expect the error "credit class metadata cannot be more than 256 characters: parse error"

  Scenario: an error is returned if credit type is empty
    Given the class
    """
    {
      "key": 1,
      "id": "C01",
      "admin": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
      "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf"
    }
    """
    When the class is validated
    Then expect the error "credit type abbreviation cannot be empty: parse error"

  Scenario: an error is returned if credit type is not formatted
    Given the class
    """
    {
      "key": 1,
      "id": "C01",
      "admin": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
      "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
      "credit_type_abbrev": "1"
    }
    """
    When the class is validated
    Then expect the error "credit type abbreviation must be 1-3 uppercase latin letters: got 1: parse error"
