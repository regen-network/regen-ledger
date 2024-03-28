Feature: Msg/UpdateProjectEnrollment

  Background:
    Given project "P001"
    * class "C01"
    * class issuer "I01" for "C01"
    * class issuer "I02" for "C01"

  Rule: valid state transitions performed by issuers which don't remove enrollment entries are:
  UNSPECIFIED -> CHANGES_REQUESTED or ACCEPTED
  CHANGES_REQUESTED -> ACCEPTED
  ACCEPTED -> ACCEPTED with new metadata
    Scenario Outline: <scenario>
      Given enrollment for "P001" to "C01" is "<cur_status>" with metadata "<cur_metadata>"
      When "<issuer>" updates enrollment for "P001" to "C01" with status "<new_status>" and metadata "<new_metadata>"
      Then expect error contains "<err>"
      And expect enrollment for "P001" to "C01" to be "<expected_status>" with metadata "<expected_metadata>"
      And if no error expect EventUpdateProjectEnrollment with properties
      """
        {
            "issuer": "<issuer>", "project_id": "P001", "class_id": "C01",
            "old_status": "PROJECT_ENROLLMENT_STATUS_<cur_status>",
            "new_status": "PROJECT_ENROLLMENT_STATUS_<expected_status>",
            "new_enrollment_metadata": "<expected_metadata>"
        }
      """
      Examples:
        | scenario                          | cur_status        | cur_metadata | issuer | new_status | new_metadata | err          | expected_status   | expected_metadata |
        | I01 unspecified to accepted       | UNSPECIFIED       | abc          | I01    | ACCEPTED   | foo123       |              | ACCEPTED          | foo123            |
        | Bob unspecified to accepted       | UNSPECIFIED       | abc          | Bob    | ACCEPTED   | foo123       | unauthorized | UNSPECIFIED       | abc               |
        | I01 changes requested to accepted | CHANGES_REQUESTED | bar456       | I01    | ACCEPTED   | foo123       |              | ACCEPTED          | foo123            |
        | Bob changes requested to accepted | CHANGES_REQUESTED | bar456       | Bob    | ACCEPTED   | foo123       | unauthorized | CHANGES_REQUESTED | bar456            |
        | I01 update accepted metadata      | ACCEPTED          | foo123       | I01    | ACCEPTED   | bar357       |              | ACCEPTED          | bar357            |
        | Bob updated accepted metadata     | ACCEPTED          | foo123       | Bob    | ACCEPTED   | bar357       | unauthorized | ACCEPTED          | foo123            |

  Rule: valid state transitions performed by issuers which remove enrollment entries are:
  UNSPECIFIED -> REJECTED
  CHANGES_REQUESTED -> REJECTED
  ACCEPTED -> TERMINATED
    Scenario Outline: <scenario>
      Given enrollment for "P001" to "C01" is "<cur_status>" with metadata "<cur_metadata>"
      When "<issuer>" updates enrollment for "P001" to "C01" with status "<new_status>" and metadata "<new_metadata>"
      Then expect error contains "<err>"
      And expect enrollment exists for "P001" to "C01" to be "<exists>"
      And if no error expect EventUpdateProjectEnrollment with properties
      """
        {
            "issuer": "<issuer>", "project_id": "P001", "class_id": "C01",
            "old_status": "PROJECT_ENROLLMENT_STATUS_<cur_status>",
            "new_status": "PROJECT_ENROLLMENT_STATUS_<new_status>",
            "new_enrollment_metadata": "<new_metadata>"
        }
      """
      Examples:
        | scenario                            | cur_status        | cur_metadata | issuer | new_status | new_metadata | err          | exists |
        | IO1 unspecified to rejected         | UNSPECIFIED       | abc          | I01    | REJECTED   | baz789       |              | false  |
        | Bob unspecified to rejected         | UNSPECIFIED       | abc          | Bob    | REJECTED   | baz789       | unauthorized | true   |
        | I01 changes requested to rejected   | CHANGES_REQUESTED | bar456       | I01    | REJECTED   | baz789       |              | false  |
        | Bob changes requested to rejected   | CHANGES_REQUESTED | bar456       | Bob    | REJECTED   | baz789       | unauthorized | true   |
        | I01 accepted to rejected            | ACCEPTED          | foo123       | I01    | REJECTED   | baz789       | invalid      | true   |
        | Bob accepted to rejected            | ACCEPTED          | foo123       | Bob    | REJECTED   | baz789       | unauthorized | true   |
        | I01 unspecified to terminated       | UNSPECIFIED       | abc          | I01    | TERMINATED | baz789       | invalid      | true   |
        | Bob unspecified to terminated       | UNSPECIFIED       | abc          | Bob    | TERMINATED | baz789       | unauthorized | true   |
        | I01 changes requested to terminated | CHANGES_REQUESTED | bar456       | I01    | TERMINATED | baz789       | invalid      | true   |
        | Bob changes requested to terminated | CHANGES_REQUESTED | bar456       | Bob    | TERMINATED | baz789       | unauthorized | true   |
        | I01 accepted to terminated          | ACCEPTED          | foo123       | I01    | TERMINATED | baz789       |              | false  |
        | Bob accepted to terminated          | ACCEPTED          | foo123       | Bob    | TERMINATED | baz789       | unauthorized | true   |

  Rule: any issuer can transition states
    Scenario: Issuer 1 requests changes, issuer 2 accepts
      Given enrollment for "P001" to "C01" is "UNSPECIFIED" with metadata "abc"
      When "I01" updates enrollment for "P001" to "C01" with status "CHANGES_REQUESTED" and metadata "def"
      Then expect no error
      And expect enrollment for "P001" to "C01" to be "CHANGES_REQUESTED" with metadata "def"
      And expect EventUpdateProjectEnrollment with properties
      """
        {
            "issuer": "I01", "project_id": "P001", "class_id": "C01",
            "old_status": "PROJECT_ENROLLMENT_STATUS_UNSPECIFIED",
            "new_status": "PROJECT_ENROLLMENT_STATUS_CHANGES_REQUESTED",
            "new_enrollment_metadata": "def"
        }
      """
      When "I02" updates enrollment for "P001" to "C01" with status "ACCEPTED" and metadata "ghi"
      Then expect no error
      And expect enrollment for "P001" to "C01" to be "ACCEPTED" with metadata "ghi"
      And expect EventUpdateProjectEnrollment with properties
      """
        {
            "issuer": "I02", "project_id": "P001", "class_id": "C01",
            "old_status": "PROJECT_ENROLLMENT_STATUS_CHANGES_REQUESTED",
            "new_status": "PROJECT_ENROLLMENT_STATUS_ACCEPTED",
            "new_enrollment_metadata": "ghi"
        }
      """
