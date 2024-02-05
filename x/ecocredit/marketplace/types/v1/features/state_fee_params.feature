Feature: FeeParams

  Scenario Outline: validate fee params
    Given buyer percentage fee "<buyer-percentage-fee>"
    And seller percentage fee "<seller-percentage-fee>"
    When I validate the fee params
    Then expect error to be <error>

    Examples:
      | buyer-percentage-fee | seller-percentage-fee | error |
      | 0.0                  | 0.0                   | false |
      | 0.0                  | 0.1                   | false |
      | 0.1                  | 0.0                   | false |
      | 0.1                  | 0.1                   | false |
      | -0.1                 | 0.0                   | true  |
      | 0.0                  | -0.1                  | true  |
      |                      | 0.1                   | false  |
      | 0.1                  |                       | false  |
      |                      |                       | false  |




