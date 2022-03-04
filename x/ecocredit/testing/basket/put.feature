Feature: Put In Basket

  Scenario: user provides a valid basket denom
    Given basket with denom eco.dC.Foo exists
    When user tries to put credits into a basket
    And user provides basket denom eco.dC.Foo
    Then credits are put into the basket

  Scenario: user provides an invalid basket denom
    Given basket with denom eco.dX.Foo does not exist
    When user tries to put credits into a basket
    And user provides basket denom eco.dX.Foo
    Then credits are NOT put into the basket

  Scenario: user has enough credits
    Given user has a positive credit balance
    When user tries to put credits into a basket
    And user provides an amount of credits less than or equal to user credit balance
    Then credits are put into the basket
    And credits are deducted from user credit balance

  Scenario: user does not have enough credits
    Given user has a positive credit balance
    When user tries to put credits into a basket
    And user provides an amount of credits more than user credit balance
    Then credits are NOT put into the basket
    And credits are NOT deducted from user credit balance

  Scenario: user provides an invalid credit denom
    Given credit denom X01-20200101-20210101-001 does not exist
    When user tries to put credits into a basket
    And user provides credits with denom X01-20200101-20210101-001
    Then credits are NOT put into the basket

  Scenario: user provides an invalid amount of credits
    When user tries to put credits into a basket
    And user provides credits with an amount equal to or less than zero
    Then credits are NOT put into the basket

  Scenario: credits do not satisfy minimum start date
    When user tries to put credits into a basket
    And user provides credits with a start date before the minimum start date
    Then credits are NOT put into the basket

  Scenario: credits do not satisfy start date window
    When user tries to put credits into a basket
    And user provides credits with a start date before the start date window
    Then credits are NOT put into the basket

  Scenario: credits must be in allowed credit classes
    When user tries to put credits into a basket
    And user provides credits from a credit class that is not in the list of allowed credit classes
    Then credits are NOT put into the basket
