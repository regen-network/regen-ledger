Feature: MsgCreateUnregisteredProject

  Rule: admin must be signer address
    Scenario Outline: validate admin
      Given admin "<admin>"
      And jurisdiction "US"
      When the message is validated
      Then expect error contains "<error>"

      Examples:
        | admin                                        | error                          |
        |                                              | "admin is required"            |
        | 0x0                                          | "admin is not a valid address" |
        | regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6 |                                |

    Scenario: admin is signer
      Given admin "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
      Then expect GetSigners returns "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"

  Rule: jurisdiction is required and must be valid
    Scenario Outline: validate jurisdiction
      Given admin "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
      And jurisdiction "<jurisdiction>"
      When the message is validated
      Then expect error contains "<error>"

      Examples:
        | jurisdiction | error                      |
        |              | "jurisdiction is required" |
        | "US"         |                            |
        | "US123"      | "jurisdiction is invalid"  |
        | "US-NY"      |                            |
        | "US-NY123"   | "jurisdiction is invalid"  |
        | "US-NY-123"  |                            |

  Rule: metadata is optional and must be at most 256 characters
    Scenario Outline: validate metadata
      Given admin "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
      And jurisdiction "US"
      And metadata "<metadata>"
      When the message is validated
      Then expect error contains "<error>"

      Examples:
        | metadata                                                                                                                                                                                                                                                                               | error                  |
        |                                                                                                                                                                                                                                                                                        |                        |
        | a                                                                                                                                                                                                                                                                                      |                        |
        | This is a string with 256 characters. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed ac facilisis arcu. Nullam nec dui ac nunc dapibus cursus. Sed sit amet massa rutrum, auctor sapien ut, euismod dolor. Nullam vehicula tellus laoreet tincidunt efficitur. Vivamus.  |                        |
        | This is a string with 257 characters. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed ac facilisis arcu. Nullam nec dui ac nunc dapibus cursus. Sed sit amet massa rutrum, auctor sapien ut, euismod dolor. Nullam vehicula tellus laoreet tincidunt efficitur. Vivamus.! | "metadata is too long" |

  Rule: reference is optional and at most 32 characters
    Scenario Outline: validate reference
      Given admin "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
      And jurisdiction "US"
      And reference "<reference>"
      When the message is validated
      Then expect error contains "<error>"

      Examples:
        | reference                        | error                   |
        |                                  |                         |
        | a                                |                         |
        | This is a string with 32 chars.  |                         |
        | This is a string with 33 chars.! | "reference is too long" |