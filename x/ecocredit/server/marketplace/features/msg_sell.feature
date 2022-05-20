Feature: MsgSell

  A sell order can be created:
    - when the credit batch exists
    - when the seller owns credits from the credit batch
    - when the seller owns greater than or equal to the quantity of credits
    - when the ask denom is an allowed denom

  Rule: The credit batch must exist

    Background:
      Given an allowed denom

    Scenario: credit batch exists
      Given alice owns credits with batch denom "C01-001-20200101-20210101-001"
      When alice attempts to create a sell order with batch denom "C01-001-20200101-20210101-001"
      Then expect no error

    Scenario: credit batch does not exist
      When alice attempts to create a sell order with batch denom "C01-001-20200101-20210101-001"
      Then expect the error "batch denom C01-001-20200101-20210101-001: not found: invalid request"

  Rule: The seller must own credits from the credit batch

    Background:
      Given an allowed denom

    Scenario: sell owns credits
      Given alice owns credits with batch denom "C01-001-20200101-20210101-001"
      When alice attempts to create a sell order with batch denom "C01-001-20200101-20210101-001"
      Then expect no error

    Scenario: seller does not own credits
      Given a credit batch with batch denom "C01-001-20200101-20210101-001"
      When alice attempts to create a sell order with batch denom "C01-001-20200101-20210101-001"
      Then expect the error "not found"

  Rule: The seller must own greater than or equal to the quantity of credits

    Background:
      Given an allowed denom

    Scenario Outline: seller owns greater than or equal to credit quantity
      Given alice owns credit quantity "<balance>"
      When alice attempts to create a sell order with credit quantity "<quantity>"
      Then expect no error

      Examples:
        | description  | balance | quantity |
        | greater than | 100     | 50       |
        | equal to     | 100     | 100      |

    Scenario: seller owns less than credit quantity
      Given alice owns credit quantity "50"
      When alice attempts to create a sell order with credit quantity "100"
      Then expect the error "tradable balance: 50, sell order request: 100 - insufficient funds: invalid request"

  Rule: The ask denom must be an allowed denom

    Scenario: ask denom is allowed
      Given an allowed denom with denom "regen"
      And alice owns credits
      When alice attempts to create a sell order with ask price "100regen"
      Then expect no error

    Scenario: ask denom is not allowed
      Given alice owns credits
      When alice attempts to create a sell order with ask price "100regen"
      Then expect the error "regen is not allowed to be used in sell orders: invalid request"
