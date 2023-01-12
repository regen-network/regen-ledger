Feature: MsgUpdateCurator

  Scenario: a valid message
    Given the message
    """
    {
        "curator": "regen1ua97smk5yv26wpqmftgdg0sx0q0d38vky7998g",
        "new_curator": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
        "denom": "eco.uC.NCT"
    }
    """
    When the message is validated
    Then expect no error

  Scenario: an error is returned if curator address is empty
    Given the message
    """
    {}
    """
    When the message is validated
    Then expect the error "curator: empty address string is not allowed: invalid address"

  Scenario: an error is returned if curator address is not a valid bech32 address 
    Given the message
    """
    {
        "curator": "invalid-address"
    }
    """
    When the message is validated
    Then expect the error "curator: decoding bech32 failed: invalid separator index -1: invalid address"

  Scenario: an error is returned if new curator address is empty 
    Given the message
    """
    {
        "curator": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw"
    }
    """
    When the message is validated
    Then expect the error "new curator: empty address string is not allowed: invalid address"

  Scenario: an error is returned if new curator address is not a valid bech32 address 
    Given the message
    """
    {
        "curator": "regen1ua97smk5yv26wpqmftgdg0sx0q0d38vky7998g",
        "new_curator": "invalid-address"
    }
    """
    When the message is validated
    Then expect the error "new curator: decoding bech32 failed: invalid separator index -1: invalid address"
	

  Scenario: an error is returned if curator and new curator are same
    Given the message
    """
    {
        "curator": "regen1ua97smk5yv26wpqmftgdg0sx0q0d38vky7998g",
        "new_curator": "regen1ua97smk5yv26wpqmftgdg0sx0q0d38vky7998g"
    }
    """
    When the message is validated
    Then expect the error "curator and new curator cannot be the same: invalid address"

  Scenario: an error is returned if basket denom is empty 
    Given the message
    """
    {
        "curator": "regen1ua97smk5yv26wpqmftgdg0sx0q0d38vky7998g",
        "new_curator": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw"
    }
    """
    When the message is validated
    Then expect the error "basket denom: empty string is not allowed: parse error: invalid request"

  Scenario: an error is returned if basket denom is invalid 
    Given the message
    """
    {
        "curator": "regen1ua97smk5yv26wpqmftgdg0sx0q0d38vky7998g",
        "new_curator": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
        "denom":"basket1+"
    }
    """
    When the message is validated
    Then expect the error "basket denom: expected format eco.<exponent-prefix><credit-type-abbrev>.<name>: parse error: invalid request"

  Scenario: a valid amino message
    Given the message
    """
    {
        "curator": "regen1ua97smk5yv26wpqmftgdg0sx0q0d38vky7998g",
        "new_curator": "regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw",
        "denom": "eco.uC.NCT"
    }
    """
    When message sign bytes queried
    Then expect the sign bytes
    """
    {
      "type":"regen.basket/MsgUpdateCurator",
      "value":{
        "curator":"regen1ua97smk5yv26wpqmftgdg0sx0q0d38vky7998g",
        "denom":"eco.uC.NCT",
        "new_curator":"regen1elq7ys34gpkj3jyvqee0h6yk4h9wsfxmgqelsw"
      }
    }
    """
