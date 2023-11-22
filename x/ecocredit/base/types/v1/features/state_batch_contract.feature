Feature: BatchContract

  Scenario: a valid batch contract
    Given the batch contract
    """
    {
      "batch_key": 1,
      "class_key": 1,
      "contract": "0x0e65079a29d7793ab5ca500c2d88e60ee99ba606"
    }
    """
    When the batch contract is validated
    Then expect no error

  Scenario: an error is returned if batch key is empty
    Given the batch contract
    """
    {}
    """
    When the batch contract is validated
    Then expect the error "batch key cannot be zero: parse error"

  Scenario: an error is returned if class key is empty
    Given the batch contract
    """
    {
      "batch_key": 1
    }
    """
    When the batch contract is validated
    Then expect the error "class key cannot be zero: parse error"

  Scenario: an error is returned if contract is empty
    Given the batch contract
    """
    {
      "batch_key": 1,
      "class_key": 1
    }
    """
    When the batch contract is validated
    Then expect the error "contract must be a valid ethereum address: parse error"

  Scenario: an error is returned if contract is not an ethereum address
    Given the batch contract
    """
    {
      "batch_key": 1,
      "class_key": 1,
      "contract": "foo"
    }
    """
    When the batch contract is validated
    Then expect the error "contract must be a valid ethereum address: parse error"
