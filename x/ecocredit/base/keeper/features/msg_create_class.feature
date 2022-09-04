Feature: Msg/CreateClass

  A credit class can be created:
  - when the credit type exists
  - when the allowlist is enabled and the admin is an approved credit class creator
  - when the required class fee is not set and no fee is provided
  - when the credit class fee denom matches the required class fee denom
  - when the credit class fee amount is greater than or equal to the required class fee amount
  - when the admin balance is greater than or equal to an allowed basket fee amount
  - the admin balance is updated and only the required fee is taken
  - the class sequence is updated
  - the class issuers are added
  - the class properties are added
  - the response includes the class id

  Rule: The credit type must exist

    Scenario: credit type exists
      Given a credit type with abbreviation "A"
      When alice attempts to create a credit class with credit type "A"
      Then expect no error

    Scenario: credit type exists
      Given a credit type with abbreviation "A"
      When alice attempts to create a credit class with credit type "B"
      Then expect the error "could not get credit type with abbreviation B: not found: invalid request"

  Rule: The admin must be an approved credit class creator if the allowlist is enabled

    Background:
      Given a credit type

    Scenario: allowlist disabled and admin is an approved creator
      Given allowlist enabled "false"
      And alice is an approved credit class creator
      When alice attempts to create a credit class
      Then expect no error

    Scenario: allowlist disabled and admin is not an approved creator
      Given allowlist enabled "false"
      When alice attempts to create a credit class
      Then expect no error

    Scenario: allowlist enabled and admin is an approved creator
      Given allowlist enabled "true"
      And alice is an approved credit class creator
      When alice attempts to create a credit class
      Then expect no error

    Scenario: allowlist enabled and admin is not an approved creator
      Given allowlist enabled "true"
      When alice attempts to create a credit class
      Then expect error contains "is not allowed to create credit classes: unauthorized"

  Rule: The credit class fee is not required if the required class fee is not set

    Background:
      Given a credit type

    Scenario: credit class fee provided and required class fee not set
      Given alice has a token balance "20regen"
      When alice attempts to create a credit class with fee "20regen"
      Then expect no error

    Scenario: credit class fee not provided and required class fee not set
      When alice attempts to create a credit class
      Then expect no error

    # no failing scenario - credit class fee is not required if required class fee is not set

  Rule: The credit class fee must match a credit class fee denom

    Background:
      Given a credit type
      And alice has a token balance "20regen"

    Scenario: credit class fee matches required class fee denom
      Given required class fee "20regen"
      When alice attempts to create a credit class with fee "20regen"
      Then expect no error

    Scenario: credit class fee does not match required class fee denom
      Given required class fee "20regen"
      When alice attempts to create a credit class with fee "20atom"
      Then expect the error "fee must be 20regen, got 20atom: insufficient fee"

    Scenario: credit class fee not provided and required class fee set
      Given required class fee "20regen"
      When alice attempts to create a credit class
      Then expect the error "fee cannot be empty: must be 20regen: insufficient fee"

  Rule: The credit class fee must be greater than or equal to a credit class fee

    Background:
      Given a credit type
      And alice has a token balance "20regen"

    Scenario Outline: credit class fee is greater than or equal to required class fee
      Given required class fee "20regen"
      When alice attempts to create a credit class with fee "<class-fee>"
      Then expect no error

      Examples:
        | description  | class-fee |
        | greater than | 30regen          |
        | equal to     | 20regen          |

    Scenario: credit class fee is less than required class fee
      Given required class fee "20regen"
      When alice attempts to create a credit class with fee "10regen"
      Then expect the error "fee must be 20regen, got 10regen: insufficient fee"

  Rule: The user must have a balance greater than or equal to the credit class fee amount

    Background:
      Given a credit type
      And required class fee "20regen"

    Scenario Outline: admin balance is greater than or equal to credit class fee amount
      Given alice has a token balance "<token-balance>"
      When alice attempts to create a credit class with fee "20regen"
      Then expect no error

      Examples:
        | description  | token-balance |
        | greater than | 30regen       |
        | equal to     | 20regen       |

    Scenario: admin balance is less than credit class fee amount
      Given alice has a token balance "10regen"
      When alice attempts to create a credit class with fee "20regen"
      Then expect the error "insufficient balance 10 for bank denom regen: insufficient funds"

  Rule: The class sequence is updated

    Background:
      Given a credit type with abbreviation "A"

    Scenario: the class sequence is updated
      Given a class sequence with credit type "A" and next sequence "1"
      When alice attempts to create a credit class with credit type "A"
      Then expect class sequence with credit type "A" and next sequence "2"

    Scenario: the class sequence is not updated
      Given a class sequence with credit type "A" and next sequence "1"
      When alice attempts to create a credit class with credit type "B"
      Then expect class sequence with credit type "A" and next sequence "1"

    # no failing scenario - state transitions only occur upon successful message execution

  Rule: The class issuers are added

    Background:
      Given a credit type

    Scenario: the class issuers are added
      When alice attempts to create a credit class with issuers
      """
      [
        "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4"
      ]
      """
      Then expect class issuers
      """
      [
        "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4"
      ]
      """

    # no failing scenario - state transitions only occur upon successful message execution

  Rule: The class properties are added

    Background:
      Given a credit type with abbreviation "A"

    Scenario: the class properties are added
      When alice attempts to create a credit class with properties
      """
      {
        "credit_type_abbrev": "A",
        "metadata": "regen:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """
      Then expect class properties
      """
      {
        "id": "A01",
        "credit_type_abbrev": "A",
        "metadata": "regen:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    # no failing scenario - state transitions only occur upon successful message execution

  Rule: The response includes the class id

    Background:
      Given a credit type with abbreviation "A"

    Scenario: the response includes the class id
      When alice attempts to create a credit class with credit type "A"
      Then expect the response
      """
      {
        "class_id": "A01"
      }
      """

    # no failing scenario - response should always be empty when message execution fails

  Rule: Event is emitted

  Background:
    Given a credit type with abbreviation "A"

    Scenario: EventCreateClass is emitted
      When  alice attempts to create a credit class with credit type "A"
      Then expect event with properties
      """
      {
        "class_id": "A01"
      }
      """
