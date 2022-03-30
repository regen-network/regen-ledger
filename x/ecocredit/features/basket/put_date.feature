Feature: Put Date

  Scenario Outline: batch start date is more than basket start date calculated from years into the past
    Given a current block timestamp of <x>
    And a basket with date criteria years into the past of <y>
    And a user owns credits from a batch with start date of <z>
    When the user attempts to put the credits into the basket
    Then the credits are NOT put into the basket

    Examples:
      | x            | y    | z            |
      | "2022-04-01" | "10" | "2011-01-01" |
      | "2022-04-01" | "10" | "2011-04-01" |
      | "2022-04-01" | "10" | "2011-07-01" |

  Scenario Outline: batch start date is equal to or less than basket start date calculated from years into the past
    Given a current block timestamp of <x>
    And a basket with date criteria years into the past of <y>
    And a user owns credits from a batch with start date of <z>
    When the user attempts to put the credits into the basket
    Then the credits are put into the basket

    Examples:
      | x            | y    | z            |
      | "2022-04-01" | "10" | "2012-01-01" |
      | "2022-04-01" | "10" | "2012-04-01" |
      | "2022-04-01" | "10" | "2012-07-01" |
      | "2022-04-01" | "10" | "2013-01-01" |
      | "2022-04-01" | "10" | "2013-04-01" |
      | "2022-04-01" | "10" | "2013-07-01" |
