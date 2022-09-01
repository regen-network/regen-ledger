Feature: Msg/RemoveClassCreator

  A class creator can be removed:
  - when the class creator exists
  - when the authority is a governance account address
  - the class creator is removed

  Rule: The class creator exists

    Scenario: The class creator exists
      Given class creators with properties
      """
      {
        "creators":[
            "regen156d26rl52y3wl865pr5x9q2vqetuw9kf0642sa"
        ]
      }
      """
      When alice attempts to remove a class creator with properties
      """
      {
        "authority":"regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "creator": "regen156d26rl52y3wl865pr5x9q2vqetuw9kf0642sa"
      }
      """
      Then expect no error
  
    Scenario: The class creator does not exist
      Given class creators with properties
      """
      {
        "creators":[
            "regen156d26rl52y3wl865pr5x9q2vqetuw9kf0642sa"
        ]
      }
      """
      When alice attempts to remove a class creator with properties
      """
      {
        "authority":"regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "creator": "regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68"
      }
      """
      Then expect the error "class creator regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68: not found"

  Rule: The authority must be governance account address

    Scenario: The authority is a governance account address
      Given class creators with properties
      """
      {
        "creators":[
            "regen156d26rl52y3wl865pr5x9q2vqetuw9kf0642sa"
        ]
      }
      """
      When alice attempts to remove a class creator with properties
      """
      {
        "authority":"regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "creator": "regen156d26rl52y3wl865pr5x9q2vqetuw9kf0642sa"
      }
      """
      Then expect no error

    Scenario: The authority is not a governance account address
      When alice attempts to remove a class creator with properties
      """
      {
        "authority":"regen1fua8speyxgempgy06gpfs0p4z32zznkqakm57s",
        "creator":"regen156d26rl52y3wl865pr5x9q2vqetuw9kf0642sa"
      }
      """
      Then expect error contains "expected gov account as only signer for proposal message"

  Rule: The class creator is removed

    Scenario: The class creator is removed
    Given class creators with properties
      """
      {
        "creators":[
            "regen156d26rl52y3wl865pr5x9q2vqetuw9kf0642sa"
        ]
      }
      """
      When alice attempts to remove a class creator with properties
      """
      {
        "authority":"regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "creator":"regen156d26rl52y3wl865pr5x9q2vqetuw9kf0642sa"
      }
      """
      Then expect class creators list to be empty

    Scenario: The class creator is removed
    Given class creators with properties
      """
      {
        "creators":[
            "regen156d26rl52y3wl865pr5x9q2vqetuw9kf0642sa",
            "regen1awp3csw2f6dw36f5mdfventta3g3pqk64fprsr"

        ]
      }
      """
      When alice attempts to remove a class creator with properties
      """
      {
        "authority":"regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "creator":"regen156d26rl52y3wl865pr5x9q2vqetuw9kf0642sa"
      }
      """
      Then expect class creators with properties
      """
      {
        "creators": [
          "regen1awp3csw2f6dw36f5mdfventta3g3pqk64fprsr"
        ]
      }
      """

