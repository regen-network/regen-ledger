Feature: Msg/AddCreditType

  A credit type can be added:
  - when the credit type does not exist
  - when the authority is a governance account address
  - the credit type is added

  Rule: The credit type does not exist

    Scenario: The credit type does not exist
      When alice attempts to add a credit type with properties
      """
      {
        "authority":"regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "credit_type":{
          "abbreviation":"C",
          "name":"carbon",
          "unit":"kg",
          "precision": 6
        }
      }
      """
      Then expect no error
  
    Scenario: The credit type already exists
      Given a credit type with properties
      """
      {
        "authority":"regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "credit_type":{
          "abbreviation":"C",
          "name":"carbon",
          "unit":"kg",
          "precision": 6
        }
      }
      """
      When alice attempts to add a credit type with name "carbon"
      Then expect the error "credit type with carbon name already exists: conflict"

    Scenario: The credit type abbreviation already exists
      Given a credit type with properties
      """
      {
        "authority":"regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "credit_type":{
          "abbreviation":"C",
          "name":"carbon",
          "unit":"kg",
          "precision": 6
        }
      }
      """
      When alice attempts to add a credit type with properties
      """
      {
        "authority":"regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "credit_type":{
          "abbreviation":"C",
          "name":"bio",
          "unit":"kg",
          "precision": 6
        }
      }
      """
      Then expect the error "credit type abbreviation C already exists: conflict"

  Rule: The authority must be governance account address

    Scenario: The authority is a governance account address
      When alice attempts to add a credit type with properties
      """
      {
        "authority":"regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "credit_type":{
          "abbreviation":"C",
          "name":"carbon",
          "unit":"kg",
          "precision": 6
        }
      }
      """
      Then expect no error

    Scenario: The authority is not a governance account address
      When alice attempts to add a credit type with properties
      """
      {
        "authority":"regen1fua8speyxgempgy06gpfs0p4z32zznkqakm57s",
        "credit_type":{
          "abbreviation":"C",
          "name":"carbon",
          "unit":"kg",
          "precision": 6
        }
      }
      """
      Then expect error contains "expected gov account as only signer for proposal message"
