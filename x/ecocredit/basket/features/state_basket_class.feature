Feature: BasketClass

  Scenario: a valid basket class
    Given the basket class
    """
    {
      "basket_id": 1,
      "class_id": "C01"
    }
    """
    When the basket class is validated
    Then expect no error

  Scenario: an error is returned if basket id is empty
    Given the basket class
    """
    {}
    """
    When the basket class is validated
    Then expect the error "basket id cannot be zero: parse error"
