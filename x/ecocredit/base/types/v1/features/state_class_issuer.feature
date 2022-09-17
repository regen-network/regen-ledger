Feature: ClassIssuer

  Scenario: a valid class issuer
    Given the class issuer
    """
    {
      "class_key": 1,
      "issuer": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y="
    }
    """
    When the class issuer is validated
    Then expect no error

  Scenario: an error is returned if class key is empty
    Given the class issuer
    """
    {}
    """
    When the class issuer is validated
    Then expect the error "class key cannot be zero: parse error"

  Scenario: an error is returned if issuer is empty
    Given the class issuer
    """
    {
      "class_key": 1
    }
    """
    When the class issuer is validated
    Then expect the error "issuer: empty address string is not allowed: parse error"
