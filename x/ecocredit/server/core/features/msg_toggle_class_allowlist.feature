Feature: Msg/MsgToggleClassAllowlist

  A class allow list can be enabled/disabled:
  - when the authority is a governance account
  - the class allow list is enabled

  Rule: The authority is a governance account

    Scenario: the authority is a governance account
      When alice attempts to toggle class allowlist with properties
      """
      {
        "authority": "regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "setting": true
      }
      """
      Then expect no error
      And expect class allowlist flag to be true 

    Scenario: the authority is not a governance account
      When alice attempts to toggle class allowlist with properties
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

  Rule: The class allow list is setting is updated
    
    Scenario: The class allow list is enabled
      When alice attempts to toggle class allowlist with properties
      """
      {
        "authority": "regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "setting": true
      }
      """
      Then expect no error
      And expect class allowlist flag to be true 
    
    Scenario: The class allow list is disabled
      When alice attempts to toggle class allowlist with properties
      """
      {
        "authority": "regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "setting": false
      }
      """
      Then expect no error
      And expect class allowlist flag to be false 