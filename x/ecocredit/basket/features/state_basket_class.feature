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

  Scenario: an error is returned if class id is empty
    Given the basket class
    """
    {
      "basket_id": 1
    }
    """
    When the basket class is validated
    Then expect the error "class id cannot be empty: parse error"

  Scenario: an error is returned if class id is not formatted
    Given the basket class
    """
    {
      "basket_id": 1,
      "class_id": "foo"
    }
    """
    When the basket class is validated
    Then expect the error "class ID didn't match the format: expected A00, got foo: parse error"
