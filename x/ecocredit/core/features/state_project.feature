Feature: Project

  Scenario: a valid project
    Given the project
    """
    {
      "key": 1,
      "id": "C01-001",
      "admin": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
      "class_key": 1,
      "jurisdiction": "US-WA",
      "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf"
    }
    """
    When the project is validated
    Then expect no error

  Scenario: a valid project without metadata
    Given the project
    """
    {
      "key": 1,
      "id": "C01-001",
      "admin": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
      "class_key": 1,
      "jurisdiction": "US-WA"
    }
    """
    When the project is validated
    Then expect no error

  Scenario: a valid project with reference id
    Given the project
    """
    {
      "key": 1,
      "id": "C01-001",
      "admin": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
      "class_key": 1,
      "jurisdiction": "US-WA",
      "metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
      "reference_id": "VCS-001"
    }
    """
    When the project is validated
    Then expect no error

  Scenario: an error is returned if key is empty
    Given the project
    """
    {}
    """
    When the project is validated
    Then expect the error "key cannot be zero: parse error"

  Scenario: an error is returned if id is empty
    Given the project
    """
    {
      "key": 1
    }
    """
    When the project is validated
    Then expect the error "project id cannot be empty: parse error"

  Scenario: an error is returned if id is not formatted
    Given the project
    """
    {
      "key": 1,
      "id": "foo"
    }
    """
    When the project is validated
    Then expect the error "invalid project id: foo: parse error"

  Scenario: an error is returned if admin is empty
    Given the project
    """
    {
      "key": 1,
      "id": "C01-001"
    }
    """
    When the project is validated
    Then expect the error "admin: empty address string is not allowed: parse error"

  Scenario: an error is returned if class key is empty
    Given the project
    """
    {
      "key": 1,
      "id": "C01-001",
      "admin": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y="
    }
    """
    When the project is validated
    Then expect the error "class key cannot be zero: parse error"

  Scenario: an error is returned if jurisdiction is empty
    Given the project
    """
    {
      "key": 1,
      "id": "C01-001",
      "admin": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
      "class_key": 1
    }
    """
    When the project is validated
    Then expect the error "jurisdiction cannot be empty, expected format <country-code>[-<region-code>[ <postal-code>]]: parse error"

  Scenario: an error is returned if jurisdiction is not formatted
    Given the project
    """
    {
      "key": 1,
      "id": "C01-001",
      "admin": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
      "class_key": 1,
      "jurisdiction": "foo"
    }
    """
    When the project is validated
    Then expect the error "invalid jurisdiction: foo, expected format <country-code>[-<region-code>[ <postal-code>]]: parse error"

  Scenario: an error is returned if metadata exceeds 256 characters
    Given the project
    """
    {
      "key": 1,
      "id": "C01-001",
      "admin": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
      "class_key": 1,
      "jurisdiction": "US-WA"
    }
    """
    And metadata with length "257"
    When the project is validated
    Then expect the error "metadata exceeds 256 characters: parse error"
