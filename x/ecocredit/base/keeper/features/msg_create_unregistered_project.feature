Feature: CreateUnregisteredProject
  Background:
    Given project creation fee "100regen"
    Given I have balance "150regen"

  Rule: a project can be created if the user pays the fee and has enough balance
    Scenario: fee amount is sufficient
      When I create a project with jurisdiction "US" and fee "100regen"
      Then expect the project is created successfully
      And expect balance "50regen"

    Scenario: fee amount is insufficient
      When I create a project with jurisdiction "US" and fee "50regen"
      Then expect error contains "insufficient fee"
      And expect balance "150regen"

    Scenario: fee amount is more than required
      When I create a project with jurisdiction "US" and fee "150regen"
      Then expect the project is created successfully
      And expect balance "50regen"

  Rule: a project cannot be created if has insufficient balance
    Given I have balance "50regen"
    When I create a project with jurisdiction "US" and fee "100regen"
    Then expect error contains "insufficient balance"
    And expect balance "50regen"

  Rule: creator is admin
    When I create a project with jurisdiction "US" and fee "100regen"
    Then expect the project is created successfully
    And expect I am the admin

  Rule: jurisdiction and metadata are saved
    When I create a project with jurisdiction "US", metadata "foobar" and fee "100regen"
    Then expect the project is created successfully
    And expect jurisdiction "US"
    And expect metadata "foobar"
