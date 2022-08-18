Feature: Msg/AllowedClassCreator

  A class creator can be added:
  - when the class creator does not exist
  - when the authority is a governance account address
  - the class creator is added

  Rule: The class creator does not exist

    Scenario: The class creator does not exist
      When alice attempts to add class creators with properties
      """
      {
        "authority":"regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "creators":[
            "regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68"
        ]
      }
      """
      Then expect no error
  
    Scenario: The class creator already exist
      Given class creators with properties
      """
      {
        "authority":"regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "creators":[
            "regen156d26rl52y3wl865pr5x9q2vqetuw9kf0642sa"
        ]
      }
      """
      When alice attempts to add class creators with properties
      """
      {
        "authority":"regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "creators":[
            "regen156d26rl52y3wl865pr5x9q2vqetuw9kf0642sa"
        ]
      }
      """
      Then expect the error "class creator regen156d26rl52y3wl865pr5x9q2vqetuw9kf0642sa already exists: invalid request"

    Scenario: duplicate class creator address
      Given class creators with properties
      """
      {
        "authority":"regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "creators":[
            "regen156d26rl52y3wl865pr5x9q2vqetuw9kf0642sa"
        ]
      }
      """
      When alice attempts to add class creators with properties
      """
      {
        "authority":"regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "creators":[
            "regen1k82wewrfkhdmegw6uxrgwwzrsd7593t8tej2d5",
            "regen1k82wewrfkhdmegw6uxrgwwzrsd7593t8tej2d5"
        ]
      }
      """
      Then expect the error "class creator regen1k82wewrfkhdmegw6uxrgwwzrsd7593t8tej2d5 already exists: invalid request"

  Rule: The authority must be governance account address

    Scenario: The authority is a governance account address
      When alice attempts to add class creators with properties
      """
      {
        "authority":"regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "creators":[
            "regen156d26rl52y3wl865pr5x9q2vqetuw9kf0642sa"
        ]
      }
      """
      Then expect no error

    Scenario: The authority is not a governance account address
      When alice attempts to add class creators with properties
      """
      {
        "authority":"regen1fua8speyxgempgy06gpfs0p4z32zznkqakm57s",
        "creators":[
            "regen156d26rl52y3wl865pr5x9q2vqetuw9kf0642sa"
        ]
      }
      """
      Then expect error contains "expected gov account as only signer for proposal message"
