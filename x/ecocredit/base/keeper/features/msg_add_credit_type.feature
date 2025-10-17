Feature: Msg/AddCreditType

  A credit type can be added:
  - when the credit type does not exist
  - when the authority is a governance account address
  - the credit type is added

  Rule: Message Validations

    Scenario: a valid message
    Given the message
    """
    {
    "authority":"regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
    "credit_type": {
    "abbreviation":"C",
    "name":"carbon",
    "unit":"ton",
    "precision":6
    }
    }
    """
    When the message is validated
    Then expect no error
    
    
    
    Scenario: an error is returned if credit type is empty
    Given the message
    """
    {
    "authority": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
    }
    """
    When the message is validated
    Then expect the error "credit type cannot be empty: invalid request"
    
    Scenario: an error is returned if credit type abbreviation is empty
    Given the message
    """
    {
    "authority": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
    "credit_type": {
    "abbreviation":"",
    "name":"carbon",
    "unit":"ton",
    "precision":6
    }
    }
    """
    When the message is validated
    Then expect the error "credit type: abbreviation: empty string is not allowed: parse error: invalid request"
    
    Scenario: an error is returned if credit type name is empty
    Given the message
    """
    {
    "authority": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
    "credit_type": {
    "abbreviation":"C",
    "name":"",
    "unit":"ton",
    "precision":6
    }
    }
    """
    When the message is validated
    Then expect the error "credit type: name cannot be empty: parse error: invalid request"
    
    Scenario: an error is returned if credit type unit is empty
    Given the message
    """
    {
    "authority": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
    "credit_type": {
    "abbreviation":"C",
    "name":"carbon",
    "unit":"",
    "precision":6
    }
    }
    """
    When the message is validated
    Then expect the error "credit type: unit cannot be empty: parse error: invalid request"
    
    Scenario: an error is returned if credit type precision is not 6
    Given the message
    """
    {
    "authority": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
    "credit_type": {
    "abbreviation":"C",
    "name":"carbon",
    "unit":"ton",
    "precision":60
    }
    }
    """
    When the message is validated
    Then expect the error "credit type: precision is currently locked to 6: parse error: invalid request"


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
      Then expect the error "credit type abbreviation C already exists: conflict"

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
