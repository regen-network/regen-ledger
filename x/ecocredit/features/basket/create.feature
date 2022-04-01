Feature: Create

  Scenario: basket fee exceeds minimum basket fee
    When user tries to create a basket
    And user provides a basket fee that is more than the minimum basket fee
    Then basket is successfully created
    And minimum basket fee is deducted from user account

  Scenario: basket fee does not meet minimum basket fee
    When user tries to create a basket
    And user provides a basket fee that is less than the minimum basket fee
    Then basket is NOT created
    And minimum basket fee is NOT deducted from user account

  Scenario: user has enough funds to pay for the basket fee
    Given user account balance is equal to or more than the minimum basket fee
    When user tries to create a basket
    And user provides a basket fee that is equal to or more than the minimum basket fee
    Then basket is created
    And minimum basket fee is deducted from user account

  Scenario: user does not have enough funds to pay for the basket fee
    Given user account balance is less than the minimum basket fee
    When user tries to create a basket
    And user provides a basket fee that is less than the minimum basket fee
    Then basket is NOT created
    And minimum basket fee is NOT deducted from user account

  Scenario: basket criteria includes a valid credit class
    Given credit class C01 exists
    When user tries to create a basket
    And user provides credit_classes=C01
    Then basket is created

  Scenario: basket criteria includes an invalid credit class
    Given credit class X01 does not exist
    When user tries to create a basket
    And user provides credit_classes=X01
    Then basket is NOT created

  Scenario: basket denom generated from name, exponent, and credit type is valid
    Given a credit type with abbreviation C exists
    When user tries to create a basket
    And user provides name=Foo, exponent=6, credit_type_abbrev=C
    Then basket is created
    And basket denom is eco.dC.Foo

  Scenario: basket denom generated from name, exponent, and credit type is not unique
    Given a credit type with abbreviation C exists
    And a basket with denom eco.dC.Foo exists
    When user tries to create a basket
    And user provides name=Foo, exponent=6, credit_type_abbrev=C
    Then basket is NOT created

  Scenario: basket name must start with alphabetic character
    Given a credit type with abbreviation C exists
    When user tries to create a basket
    And user provides name=1Foo, exponent=6, credit_type_abbrev=C
    Then basket is NOT created

  Scenario: basket name does not meet required length
    Given a credit type with abbreviation C exists
    When user tries to create a basket
    And user provides name=fo, exponent=6, credit_type_abbrev=C
    Then basket is NOT created

  Scenario: basket name exceeds required length
    Given a credit type with abbreviation C exists
    When user tries to create a basket
    And user provides name=foobarbaz, exponent=6, credit_type_abbrev=C
    Then basket is NOT created

  Scenario: basket exponent must be greater than or equal to the precision of the credit type
    Given a credit type with abbreviation C exists
    And the precision of credit type C is 6
    When user tries to create a basket
    And user provides name=foobarbaz, exponent=3, credit_type_abbrev=C
    Then basket is NOT created

  Scenario: basket exponent must be a value with an official SI prefix
    Given a credit type with abbreviation C exists
    When user tries to create a basket
    And user provides name=foobarbaz, exponent=7, credit_type_abbrev=C
    Then basket is NOT created

  Scenario: basket credit type abbreviation must reference a valid credit type
    Given a credit type with abbreviation X does not exist
    When user tries to create a basket
    And user provides name=foobarbaz, exponent=6, credit_type_abbrev=X
    Then basket is NOT created

  Scenario: basket description must not exceed 256 characters
    When user tries to create a basket
    And user provides a description that exceeds 256 characters
    Then basket is NOT created

  Scenario: basket can only specify one date criteria option
    When user tries to create a basket
    And user provides values for both a minimum start date and a start date window
    Then basket is NOT created

  Scenario: basket minimum start date must be a valid timestamp
    When user tries to create a basket
    And user provides a minimum start date that is before 1900-01-01
    Then basket is NOT created

  Scenario: basket start date window must be a valid duration
    When user tries to create a basket
    And user provides a minimum start date that is less than 1 day
    Then basket is NOT created
