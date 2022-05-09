Feature: Take

  Scenario: user provides a valid basket denom
    Given basket with denom eco.dC.Foo exists
    When user tries to take credits from a basket
    And user provides eco.dC.Foo as the basket denom
    Then credits are taken from the basket

  Scenario: user provides an invalid basket denom
    Given basket with denom eco.dX.Foo does not exist
    When user tries to take credits from a basket
    And user provides eco.dX.Foo as the basket denom
    Then credits are NOT taken from the basket

  Scenario: user has enough basket tokens
    Given user has a positive basket token balance
    When user tries to take credits from a basket
    And user provides an amount less than or equal to user basket token balance
    Then credits are taken from the basket
    And basket tokens are deducted from user account

  Scenario: user does not have enough basket tokens
    Given user has a positive basket token balance
    When user tries to take credits from a basket
    And user provides an amount more than user basket token balance
    Then credits are NOT taken from the basket
    And basket tokens are NOT deducted from user account

  Scenario: user provides an invalid amount of basket tokens
    When user tries to take credits from a basket
    And user provides an amount of basket tokens that is less than or equal to zero
    Then credits are NOT taken from the basket

  Scenario: user sets retire_on_take to true and basket retire_on_take is true
    Given basket retire_on_take is true
    When user tries to take credits from a basket
    And user sets retire_on_take to true
    And user provides a valid retirement jurisdiction
    Then credits are taken from the basket
    And credits are received in a retired state

  Scenario: user sets retire_on_take to false and basket retire_on_take is true
    Given basket retire_on_take is true
    When user tries to take credits from a basket
    Given user sets retire_on_take to false
    Then credits are NOT taken from the basket
    And credits are NOT received in a retired or tradable state

  Scenario: user does not provide retirement jurisdiction when retire_on_take is false
    Given basket retire_on_take is false
    When user tries to take credits from a basket
    And user sets retire_on_take to false
    And user does not specify a retirement jurisdiction
    Then credits are taken from the basket
    And credits are received in a tradable state

  Scenario: user does not provide retirement jurisdiction when retire_on_take is true
    Given basket retire_on_take is true
    When user tries to take credits from a basket
    And user sets retire_on_take to true
    And user does not provide a retirement jurisdiction
    Then credits are NOT taken from the basket
    And credits are NOT received in a retired or tradable state

  Scenario: user provides an invalid retirement jurisdiction
    Given basket retire_on_take is true
    When user tries to take credits from a basket
    And user sets retire_on_take to true
    And user does not provide a valid retirement jurisdiction
    Then credits are NOT taken from the basket
    And credits are NOT received in a retired or tradable state
