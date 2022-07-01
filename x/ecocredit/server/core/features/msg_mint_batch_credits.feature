Feature: Msg/MintBatchCredits

  Credits can be dynamically minted to an existing credit batch:
  - when the credit batch is open
  - when the origin tx id and source is unique within the scope of a credit class
  - ...

  Rule: The credit batch must be open

    Background:
      Given a credit type
      And a credit class with issuer alice
      And a project

    Scenario: the credit batch is closed
      Given a credit batch with open "false" and issuer alice
      When alice attempts to mint credits
      Then expect the error "unable to mint credits: credits cannot be minted in a closed batch: invalid request"

    Scenario: the credit batch is open
      Given a credit batch with open "true" and issuer alice
      When alice attempts to mint credits
      Then expect no error

  Rule: The origin tx must be unique within the scope of the credit class

    Background:
      Given a credit type
      And a credit class with issuer alice
      And a project
      And a credit batch with open "true" and issuer alice

    Scenario: the origin tx is not unique within credit class
      Given an origin tx index
      """
      {
        "class_key": 1,
        "id": "0x0",
        "source": "polygon"
      }
      """
      When alice attempts to mint credits with origin tx
      """
      {
        "id": "0x0",
        "source": "polygon"
      }
      """
      Then expect the error "credits already issued with tx id: 0x0: invalid request"

    Scenario: the origin tx is unique within credit class
      Given an origin tx index
      """
      {
        "class_key": 2,
        "id": "0x0",
        "source": "polygon"
      }
      """
      When alice attempts to mint credits with origin tx
      """
      {
        "id": "0x0",
        "source": "polygon"
      }
      """
      Then expect no error
