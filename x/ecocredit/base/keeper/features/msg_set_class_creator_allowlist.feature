Feature: Msg/SetClassCreatorAllowlist

  A class allow list can be enabled/disabled:
  - when the authority is a governance account
  - the class allow list setting is updated

  Rule: The authority is a governance account

    Scenario: the authority is a governance account
      When alice attempts to set class creator allowlist with properties
      """
      {
        "authority": "regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "enabled": true
      }
      """
      Then expect no error
      And expect class allowlist flag to be "true" 

    Scenario: the authority is not a governance account
      When alice attempts to set class creator allowlist with properties
      """
      {
        "authority": "regen1fua8speyxgempgy06gpfs0p4z32zznkqakm57s",
        "enabled": true
      }
      """
      Then expect error contains "expected gov account as only signer for proposal message"

  Rule: The class allow list setting is updated
    
    Scenario: The class allow list is enabled
      When alice attempts to set class creator allowlist with properties
      """
      {
        "authority": "regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "enabled": true
      }
      """
      Then expect class allowlist flag to be "true" 
    
    Scenario: The class allow list is disabled
      When alice attempts to set class creator allowlist with properties
      """
      {
        "authority": "regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "enabled": false
      }
      """
      Then expect class allowlist flag to be "false" 