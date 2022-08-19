Feature: Msg/MsgUpdateBasketFees

  The basket fee can be updated:
  - when the authority is a governance account address
  - when the provided basket fee is valid
  - the basket fee is updated

  Rule: The authority is a governance account

    Scenario: the authority is a governance account
      When alice attempts to update basket fee with properties
      """
      {
        "authority": "regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "basket_fees": [
            {
                "denom": "uregen",
                "amount": "1000"
            }
        ]
      }
      """
      Then expect no error

    Scenario: the authority is not a governance account
      When alice attempts to update basket fee with properties
      """
      {
        "authority": "regen1fua8speyxgempgy06gpfs0p4z32zznkqakm57s",
        "basket_fees": [
            {
                "denom": "uregen",
                "amount": "1000"
            }
        ]
      }
      """
      Then expect error contains "expected gov account as only signer for proposal message"

  Rule: The basket fee is valid

    Scenario: the basket fee is valid
      When alice attempts to update basket fee with properties
      """
      {
        "authority":"regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "basket_fees":[
            {
                "denom":"uregen",
                "amount":"1000"
            }
        ]
      }
      """
      Then expect no error
      And expect basket fees with properties 
      """
      {
        "fees":[
          {
                "denom":"uregen",
                "amount":"1000"
          }
        ]
      }
      """

    Scenario: valid basket fee empty 
      When alice attempts to update basket fee with properties
      """
      {
        "authority":"regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "basket_fees":[]
      }
      """
      Then expect no error
      And expect basket fees with properties 
      """
      {
        "fees":[]
      }
      """
    
    Scenario: the basket fees are updated (multiple fees)
      When alice attempts to update basket fees with properties
      """
      {
        "authority":"regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "basket_fees":[
          {
            "denom": "uregen",
            "amount": "1000"
          },
          {
            "denom": "uatom",
            "amount": "1000"
          }
        ]
      }
      """
      Then expect no error
      And expect basket fees with properties 
      """
      {
        "fees":[
          {
            "denom": "uregen",
            "amount": "1000"
          },
          {
            "denom": "uatom",
            "amount": "1000"
          }
        ]
      }
      """

