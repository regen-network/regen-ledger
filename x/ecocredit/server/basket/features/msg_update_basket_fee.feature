Feature: Msg/MsgUpdateBasketFees

  The basket fees can be updated:
  - when the authority is a governance account address
  - the basket fees are updated

  Rule: The authority is a governance account

    Scenario: the authority is a governance account
      When alice attempts to update basket fees with properties
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
      When alice attempts to update basket fees with properties
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

  Rule: The basket fees are updated

    Scenario: the basket fees are updated (single fee)
      When alice attempts to update basket fees with properties
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

    Scenario: the basket fees are updated (no fees)
      When alice attempts to update basket fees with properties
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

