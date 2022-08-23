Feature: Basket

  Scenario: a valid basket
    Given the basket
    """
    {
      "id": 1,
      "basket_denom": "eco.uC.NCT",
      "name": "NCT",
      "credit_type_abbrev": "C",
      "curator": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y="
    }
    """
    When the basket is validated
    Then expect no error

  Scenario: a valid basket with date criteria
    Given the basket
    """
    {
      "id": 1,
      "basket_denom": "eco.uC.NCT",
      "name": "NCT",
      "credit_type_abbrev": "C",
      "date_criteria": {
        "years_in_the_past": 10
      },
      "curator": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y="
    }
    """
    When the basket is validated
    Then expect no error

  Scenario: an error is returned if id is empty
    Given the basket
    """
    {}
    """
    When the basket is validated
    Then expect the error "id cannot be zero: parse error"
