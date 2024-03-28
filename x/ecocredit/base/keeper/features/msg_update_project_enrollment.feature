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
    Scenario Outline:
      Given enrollment for "P001" to "C01" is "<cur_status>" with metadata "<cur_metadata>"
      When "<issuer>" updates enrollment for "P001" to "C01" with status "<new_status>" and metadata "<new_metadata>"
      Then expect error contains "<err>"
      And expect enrollment for "P001" to "C01" to be "<expected_status>" with metadata "<expected_metadata>"
      Examples:
        | cur_status        | cur_metadata | issuer | new_status | new_metadata | err          | expected_status   | expected_metadata |
        | UNSPECIFIED       | abc          | I01    | ACCEPTED   | foo123       |              | ACCEPTED          | foo123            |
        | UNSPECIFIED       | abc          | Bob    | ACCEPTED   | foo123       | unauthorized | UNSPECIFIED       | abc               |
        | CHANGES_REQUESTED | bar456       | I01    | ACCEPTED   | foo123       |              | ACCEPTED          | foo123            |
        | CHANGES_REQUESTED | bar456       | Bob    | ACCEPTED   | foo123       | unauthorized | CHANGES_REQUESTED | bar456            |
        | ACCEPTED          | foo123       | I01    | ACCEPTED   | bar357       |              | ACCEPTED          | bar357            |
        | ACCEPTED          | foo123       | Bob    | ACCEPTED   | bar357       | unauthorized | ACCEPTED          | foo123            |

  Rule: valid state transitions performed by issuers which remove enrollment entries are:
  UNSPECIFIED -> REJECTED
  CHANGES_REQUESTED -> REJECTED
  ACCEPTED -> TERMINATED
    Scenario Outline:
      Given enrollment for "P001" to "C01" is "<cur_status>" with metadata "<cur_metadata>"
      When "<issuer>" updates enrollment for "P001" to "C01" with status "<new_status>" and metadata "<new_metadata>"
      Then expect error contains "<err>"
      And expect enrollment exists for "P001" to "C01" to be "<exists>"
      Examples:
        | cur_status        | cur_metadata | issuer | new_status | new_metadata | err          | exists |
        | UNSPECIFIED       | abc          | I01    | REJECTED   | baz789       |              | false  |
        | UNSPECIFIED       | abc          | Bob    | REJECTED   | baz789       | unauthorized | true   |
        | CHANGES_REQUESTED | bar456       | I01    | REJECTED   | baz789       |              | false  |
        | CHANGES_REQUESTED | bar456       | Bob    | REJECTED   | baz789       | unauthorized | true   |
        | ACCEPTED          | foo123       | I01    | REJECTED   | baz789       | invalid      | true   |
        | ACCEPTED          | foo123       | Bob    | REJECTED   | baz789       | unauthorized | true   |
        | UNSPECIFIED       | abc          | I01    | TERMINATED | baz789       | invalid      | true   |
        | UNSPECIFIED       | abc          | Bob    | TERMINATED | baz789       | unauthorized | true   |
        | CHANGES_REQUESTED | bar456       | I01    | TERMINATED | baz789       | invalid      | true   |
        | CHANGES_REQUESTED | bar456       | Bob    | TERMINATED | baz789       | unauthorized | true   |
        | ACCEPTED          | foo123       | I01    | TERMINATED | baz789       |              | false  |
        | ACCEPTED          | foo123       | Bob    | TERMINATED | baz789       | unauthorized | true   |

  Rule: events get emitted

    Rule: any issuer can transition states