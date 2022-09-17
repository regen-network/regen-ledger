Feature: AllowedBridgeChain

  Scenario: a valid allowed bridge chain
    Given the allowed bridge chain
    """
    {
      "chain_name": "polygon"
    }
    """
    When the allowed bridge chain is validated
    Then expect no error

  Scenario: an error is returned if chain name is empty
    Given the allowed bridge chain
    """
    {}
    """
    When the allowed bridge chain is validated
    Then expect the error "name cannot be empty: parse error"
