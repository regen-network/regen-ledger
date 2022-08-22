Feature: Msg/RemoveClassCreators

  A class creators can be remove:
  - when the class creator is exist
  - when the authority is a governance account address
  - the class creator is removed

  Rule: The class creator is exist

    Scenario: The class creator is exist
      Given class creators with properties
      """
      {
        "creators":[
            "regen156d26rl52y3wl865pr5x9q2vqetuw9kf0642sa"
        ]
      }
      """
      When alice attempts to remove class creators with properties
      """
      {
        "authority":"regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "creators":[
            "regen156d26rl52y3wl865pr5x9q2vqetuw9kf0642sa"
        ]
      }
      """
      Then expect no error

    Scenario: More than one class creator
      Given class creators with properties
      """
      {
        "creators":[
          "regen156d26rl52y3wl865pr5x9q2vqetuw9kf0642sa",
          "regen19tk63w56szvyvz5msdqtu8y8fzk2qc70vnczmd"
        ]
      }
      """
      When alice attempts to remove class creators with properties
      """
      {
        "authority":"regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "creators":[
            "regen156d26rl52y3wl865pr5x9q2vqetuw9kf0642sa",
            "regen19tk63w56szvyvz5msdqtu8y8fzk2qc70vnczmd"
        ]
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
      When alice attempts to remove class creators with properties
      """
      {
        "authority":"regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "creators":[
            "regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68"
        ]
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
      When alice attempts to remove class creators with properties
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
      When alice attempts to remove class creators with properties
      """
      {
        "authority":"regen1fua8speyxgempgy06gpfs0p4z32zznkqakm57s",
        "creators":[
            "regen156d26rl52y3wl865pr5x9q2vqetuw9kf0642sa"
        ]
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
      When alice attempts to remove class creators with properties
      """
      {
        "authority":"regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "creators":[
            "regen156d26rl52y3wl865pr5x9q2vqetuw9kf0642sa"
        ]
      }
      """
      Then expect no error
      And expect class creators list to be empty

    Scenario: The class creators are removed
    Given class creators with properties
      """
      {
        "creators":[
            "regen156d26rl52y3wl865pr5x9q2vqetuw9kf0642sa",
            "regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
            "regen1awp3csw2f6dw36f5mdfventta3g3pqk64fprsr"

        ]
      }
      """
      When alice attempts to remove class creators with properties
      """
      {
        "authority":"regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68",
        "creators":[
            "regen156d26rl52y3wl865pr5x9q2vqetuw9kf0642sa",
            "regen1awp3csw2f6dw36f5mdfventta3g3pqk64fprsr"
        ]
      }
      """
      Then expect no error
      And expect class creators with properties
      """
      {
        "creators": [
          "regen1nzh226hxrsvf4k69sa8v0nfuzx5vgwkczk8j68"
        ]
      }
      """

