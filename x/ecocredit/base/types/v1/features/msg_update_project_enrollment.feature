Feature: MsgUpdateProjectEnrollment

  Rule: issuer is address and signer
    Scenario Outline: validate issuer
      Given issuer "<issuer>"
      * project ID "P001"
      * class ID "C01"
      * new status "ACCEPTED"
      When the message is validated
      Then expect error contains "<error>"

      Examples:
        | issuer                                       | error          |
        |                                              | issuer: empty  |
        | 0x0                                          | invalid bech32 |
        | regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6 |                |

    Scenario: issuer is signer
      Given issuer "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
      Then expect GetSigners returns "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"

  Rule: project_id and class_id must be non-empty
    Scenario Outline: validate project and class IDs
      Given issuer "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
      * project ID "<project_id>"
      * class ID "<class_id>"
      * new status "ACCEPTED"
      When the message is validated
      Then expect error contains "<error>"

      Examples:
        | project_id | class_id | error             |
        |            | C01      | project id: empty |
        | P001       | C01      |                   |
        | P001       |          | class id: empty   |

  Rule: new_status is valid
    Scenario Outline: validate status
      Given issuer "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
      * project ID "P001"
      * class ID "C01"
      * new status "<status>"
      When the message is validated
      Then expect error contains "<error>"

      Examples:
        | status            | error               |
        | UNSPECIFIED       |                     |
        | ACCEPTED          |                     |
        | REJECTED          |                     |
        | CHANGES_REQUESTED |                     |
        | TERMINATED        |                     |
        | 6                 | new status: invalid |

  Rule: metadata is optional and at most 256 characters
    Scenario Outline: validate metadata
      Given issuer "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
      * project ID "P001"
      * class ID "C01"
      * new status "ACCEPTED"
      * metadata "<metadata>"
      When the message is validated
      Then expect error contains "<error>"

      Examples:
        | metadata                                                                                                                                                                                                                                                          | error    |
        |                                                                                                                                                                                                                                                                   |          |
        | a                                                                                                                                                                                                                                                                 |          |
        | This is a string with 256 characters. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed ac facilisis arcu. Nullam nec dui ac nunc dapibus cursus. Sed sit amet massa rutrum, auctor sapien ut, euismod dolor. Nullam vehicula tellus laoreet tincidun  |          |
        | This is a string with 257 characters. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed ac facilisis arcu. Nullam nec dui ac nunc dapibus cursus. Sed sit amet massa rutrum, auctor sapien ut, euismod dolor. Nullam vehicula tellus laoreet tincidunt | metadata |
