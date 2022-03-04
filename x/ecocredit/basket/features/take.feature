Feature: takes credits from a basket

  Background:
    Given a basket foo, with auto-retire enabled

  Scenario: must retire when taking from foo
    When I try to take credits from foo
    Then expect must retire error
