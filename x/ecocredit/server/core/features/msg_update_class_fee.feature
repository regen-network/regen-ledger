Feature: Msg/MsgUpdateClassFees

  The class fees can be updated:
  - when the authority is a governance account address
  - the class fees are updated

  Rule: The authority is a governance account

    Scenario: the authority is a governance account
      When alice attempts to update class fees with properties
      """
      {
        "authority": "regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "fees": [
            {
                "denom": "uregen",
                "amount": "1000"
            }
        ]
      }
      """
      Then expect no error

    Scenario: the authority is not a governance account
      When alice attempts to update class fees with properties
      """
      {
        "authority": "regen1fua8speyxgempgy06gpfs0p4z32zznkqakm57s",
        "fees": [
            {
                "denom": "uregen",
                "amount": "1000"
            }
        ]
      }
      """
      Then expect error contains "expected gov account as only signer for proposal message"

  Rule: The class fees are updated
    
    Scenario: valid empty class fees
      When alice attempts to update class fees with properties
      """
      {
        "authority":"regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "fees":[]
      }
      """
      Then expect class fees with properties 
      """
      {
        "fees":[]
      }
      """
    
  Scenario: valid class fees (multiple tokens)
      When alice attempts to update class fees with properties
      """
      {
        "authority":"regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
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
      Then expect class fees with properties
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