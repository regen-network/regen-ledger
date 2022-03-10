Feature: Fixed Start Date

  Scenario Outline: a basket with a fixed start date that automatically updates on a yearly basis
    Given: a basket with a fixed start date of <x>
    When: the block time matches the month and day of the fixed start date
    Then: the fixed start date is updated to <y>

    Examples:
      | x | y |
      | 2012-01-01 | 2013-01-01 |
      | 2012-06-06 | 2013-06-06 |
