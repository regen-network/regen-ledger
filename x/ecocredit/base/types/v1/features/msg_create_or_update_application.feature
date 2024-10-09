Feature: MsgCreateOrUpdateApplication

  Rule: project_admin is address and signer
    Scenario Outline: validate admin
      Given project admin "<admin>"
      * project ID "P1"
      * class ID "C01"
      When the message is validated
      Then expect error contains "<error>"

      Examples:
        | admin                                        | error         |
        |                                              | project admin |
        | 0x0                                          | project admin |
        | regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6 |               |

    Scenario: project admin is signer
      Given project admin "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
      Then expect GetSigners returns "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"

  Rule: project_id should be non-empty
    Scenario Outline: validate project ID
      Given project admin "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
      * project ID "<project_id>"
      * class ID "C01"
      When the message is validated
      Then expect error contains "<error>"

      Examples:
        | project_id | error      |
        |            | project id |
        | P1         |            |

  Rule: class_id should be non-empty
    Scenario Outline: validate class ID
      Given project admin "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
      * project ID "P1"
      * class ID "<class_id>"
      When the message is validated
      Then expect error contains "<error>"

      Examples:
        | class_id | error    |
        |          | class id |
        | C01       |          |

  Rule: metadata is optional and at most 256 characters
    Scenario Outline: validate metadata
      Given project admin "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
      * project ID "P1"
      * class ID "C01"
      * metadata "<metadata>"
      When the message is validated
      Then expect error contains "<error>"

      Examples:
        | metadata                                                                                                                                                                                                                                                          | error    |
        |                                                                                                                                                                                                                                                                   |          |
        | a                                                                                                                                                                                                                                                                 |          |
        | This is a string with 256 characters. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed ac facilisis arcu. Nullam nec dui ac nunc dapibus cursus. Sed sit amet massa rutrum, auctor sapien ut, euismod dolor. Nullam vehicula tellus laoreet tincidun  |          |
        | This is a string with 257 characters. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed ac facilisis arcu. Nullam nec dui ac nunc dapibus cursus. Sed sit amet massa rutrum, auctor sapien ut, euismod dolor. Nullam vehicula tellus laoreet tincidunt | metadata |
