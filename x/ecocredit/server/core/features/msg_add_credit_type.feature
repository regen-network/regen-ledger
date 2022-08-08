Feature: Msg/AddCreditType

  A credit type can be added:
  - when the credit type does not exists
  - when the credit type exists
  - when the authority is a governance account address
  - the credit type is added

  Rule: The credit type does not exist

    Scenario: The credit type does not exist
      When alice attempts to add a credit type with properties
      """
      {
        "authority":"",
        "credit_type":{
          "name":"carbon",
          "abbreviation":"C",
          "precision"6,
          "unit":"kg"
        }
      }
      """
      Then expect no error
  
  Rule: The credit type exists

    Scenario: The credit type already exists
      Given a credit type with abbreviation properties
      """
      {
        "authority":"",
        "credit_type":{
          "name":"carbon",
          "abbreviation":"C",
          "precision"6,
          "unit":"kg"
        }
      }
      """
      When alice attempts to add a credit type with name "carbon"
      Then expect error "credit type already exists"

  Rule: The authority must be governance account address

    Scenario: The authority is a governance account address
      When alice attempts to add a credit type with properties
      """
      {
        "authority":"",
        "credit_type":{
          "name":"carbon",
          "abbreviation":"C",
          "precision"6,
          "unit":"kg"
        }
      }
      """
      Then expect no error

    Scenario: The authority is not a governance account address
      When alice attempts to add a credit type with properties
      """
      {
        "authority":"",
        "credit_type":{
          "name":"carbon",
          "abbreviation":"C",
          "precision"6,
          "unit":"kg"
        }
      }
      """
      Then expect no error
