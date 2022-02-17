Feature: takes credits from a basket

  Background:
    Given a basket "foo"
    * a basket "bar", with auto-retire disables

  Scenario: must retire when taking from "foo"
    When something
    Then the other thing
