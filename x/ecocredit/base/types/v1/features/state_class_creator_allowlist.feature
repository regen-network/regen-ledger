Feature: ClassCreatorAllowlist

  Scenario: a valid class creator allowlist with enabled true
    Given the class creator allowlist
    """
    {
      "enabled": true
    }
    """
    When the class creator allowlist is validated
    Then expect no error

  Scenario: a valid class creator allowlist with enabled false
    Given the class creator allowlist
    """
    {
      "enabled": false
    }
    """
    When the class creator allowlist is validated
    Then expect no error
