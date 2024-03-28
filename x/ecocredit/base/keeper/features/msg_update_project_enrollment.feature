Feature: Msg/UpdateProjectEnrollment

  Background:
    Given class "C01" with issuer "I01"
    And an application for project "P001" to class "C01" with status "UNSPECIFIED" and metadata "abc"

  Rule: an issuer can accept a project application from status unspecified or changes requested
    Scenario Outline:
      Given enrollment for "P001" to "C01" is "<cur_status>" with metadata "<cur_metadata>"
      When <issuer> updates enrollment for "P001" to "C01" with status "<status>" and metadata "<metadata>"
      Then expect error contains <err>
      And expect enrollment for "P001" to "C01" to be "<new_status>" with metadata "<new_metadata>"
      Examples:
        | cur_status        | cur_metadata | issuer | new_status | new_metadata | err          | expected_status   | expected_metadata |
        | UNSPECIFIED       | abc          | I01    | ACCEPTED   | foo123       |              | ACCEPTED          | foo123            |
        | UNSPECIFIED       | abc          | Bob    | ACCEPTED   | foo123       | unauthorized | UNSPECIFIED       | abc               |
        | CHANGES_REQUESTED | bar456       | I01    | ACCEPTED   | foo123       |              | ACCEPTED          | foo123            |
        | CHANGES_REQUESTED | bar456       | Bob    | ACCEPTED   | foo123       | unauthorized | CHANGES_REQUESTED | bar456            |
        | ACCEPTED          | foo123       | I01    | ACCEPTED   | bar357       |              | ACCEPTED          | bar357            |
        | ACCEPTED          | foo123       | Bob    | ACCEPTED   | bar357       | unauthorized | ACCEPTED          | foo123            |

    Scenario Outline:
      Given enrollment for "P001" to "C01" is "<cur_status>" with metadata "<cur_metadata>"
      When <issuer> updates enrollment for "P001" to "C01" with status "<status>" and metadata "<metadata>"
      Then expect error contains <err>
      And expect enrollment for "P001" to "C01" to be <exists>

      Examples:
        | cur_status  | cur_metadata | issuer | new_status | new_metadata | err          | exists |
        | UNSPECIFIED | abc          | I01    | REJECTED   | baz789       |              | false
        | UNSPECIFIED | abc          | Bob    | REJECTED   | baz789       | unauthorized | true   |
        | CHANGES_REQUESTED | bar456 | I01    | REJECTED   | baz789       |              | false  |
        | CHANGES_REQUESTED | bar456 | Bob    | REJECTED   | baz789       | unauthorized | true   |
        | ACCEPTED    | foo123       | I01    | REJECTED   | baz789       | invalid      | true   |
        | ACCEPTED    | foo123       | Bob    | REJECTED   | baz789       | unauthorized | true   |
        | ACCEPTED    | foo123       | I01    | TERMINATED | baz789       |              | false  |
        | ACCEPTED    | foo123       | Bob    | TERMINATED | baz789       | unauthorized | true   |

    Scenario: an issuer accepts a project application
      When "I01" updates enrollment for "P001" to "C01" with status "ACCEPTED" and metadata "foo123"
      Then expect no error
      And expect enrollment for "P001" to "C01" to be "ACCEPTED" with metadata "foo123"

    Scenario: non issuer attempts to accept a project application
      When "Bob" updates enrollment for "P001" to "C01" with status "ACCEPTED" and metadata "foo123"
      Then expect error contains "unauthorized"
      And expect enrollment for "P001" to "C01" to be "UNSPECIFIED" with metadata "abc"

  Rule: an issuer can request changes to a project application
    Scenario: an issuer requests changes to a project application
      When "I01" updates enrollment for "P001" to "C01" with status "CHANGES_REQUESTED" and metadata "bar456"
      Then expect no error
      And expect enrollment for "P001" to "C01" to be "CHANGES_REQUESTED" with metadata "bar456"

    Scenario: non issuer attempts to request changes to a project application
      When "Bob" updates enrollment for "P001" to "C01" with status "CHANGES_REQUESTED" and metadata "bar456"
      Then expect error contains "unauthorized"
      And expect enrollment for "P001" to "C01" to be "UNSPECIFIED" with metadata "abc"

  Rule: an issuer can reject a project application and it is removed from state
    Scenario: an issuer rejects a project application
      When "I01" updates enrollment for "P001" to "C01" with status "REJECTED" and metadata "baz789"
      Then expect no error

    Scenario: non issuer attempts to reject a project application
      When "Bob" updates enrollment for "P001" to "C01" with status "REJECTED" and metadata "baz789"
      Then expect error contains "unauthorized"
      And expect enrollment for "P001" to "C01" to be "UNSPECIFIED" with metadata "abc"

  Rule: an issuer can terminate a project enrollment and it is removed from state

    Rule: any issuer can change the status even if a different issuer changed it before

  Rule: events get emitted