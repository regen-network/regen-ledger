Feature: MsgUpdateProjectEnrollment

  Rule: issuer is address and signer
    Scenario Outline: validate issuer
      Given issuer "<issuer>"
      * project ID "P1"
      * class ID "C1"
      * new status "PROJECT_ENROLLMENT_STATUS_ACCEPTED"
      When the message is validated
      Then expect error contains "<error>"

      Examples:
        | issuer                                       | error                         |
        |                                              | issuer is required            |
        | 0x0                                          | issuer is not a valid address |
        | regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6 |                               |

    Scenario: issuer is signer
      Given issuer "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
      Then expect GetSigners returns "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"

  Rule: project_id is non-empty

  Rule: class_id is non-empty

  Rule: new_status is specified and valid

  Rule: metadata is optional and at most 256 characters