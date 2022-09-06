Feature: AllowedClassCreator

  Scenario: a valid allowed class creator
    Given the allowed class creator
    """
    {
      "address": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y="
    }
    """
    When the allowed class creator is validated
    Then expect no error

  Scenario: an error is returned if address is empty
    Given the allowed class creator
    """
    {}
    """
    When the allowed class creator is validated
    Then expect the error "address: empty address string is not allowed: parse error"
