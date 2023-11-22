Feature: OriginTxIndex

  Scenario: a valid origin tx index
    Given the origin tx index
    """
    {
      "class_key": 1,
      "id": "0x7a70692a348e8688f54ab2bdfe87d925d8cc88932520492a11eaa02dc128243e",
      "source": "polygon"
    }
    """
    When the origin tx index is validated
    Then expect no error

  Scenario: an error is returned if class key is empty
    Given the origin tx index
    """
    {}
    """
    When the origin tx index is validated
    Then expect the error "class key cannot be zero: parse error"

  Scenario: an error is returned if id is empty
    Given the origin tx index
    """
    {
      "class_key": 1
    }
    """
    When the origin tx index is validated
    Then expect the error "id cannot be empty: parse error"

  Scenario: an error is returned if id exceeds 128 characters
    Given the origin tx index
    """
    {
      "class_key": 1
    }
    """
    And id with length "129"
    When the origin tx index is validated
    Then expect the error "id must be at most 128 characters long, valid characters: alpha-numberic, space, '-' or '_': parse error"

  Scenario: an error is returned if source is empty
    Given the origin tx index
    """
    {
      "class_key": 1,
      "id": "0x7a70692a348e8688f54ab2bdfe87d925d8cc88932520492a11eaa02dc128243e"
    }
    """
    When the origin tx index is validated
    Then expect the error "source cannot be empty: parse error"

  Scenario: an error is returned if source exceeds 32 characters
    Given the origin tx index
    """
    {
      "class_key": 1,
      "id": "0x7a70692a348e8688f54ab2bdfe87d925d8cc88932520492a11eaa02dc128243e"
    }
    """
    And source with length "33"
    When the origin tx index is validated
    Then expect the error "source must be at most 32 characters long, valid characters: alpha-numberic, space, '-' or '_': parse error"
