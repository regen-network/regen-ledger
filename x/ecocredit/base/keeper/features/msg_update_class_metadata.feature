Feature: Msg/UpdateClassMetadata

  The metadata of a credit class can be updated:
  - message validations
  - when the credit class exists
  - when the admin is the admin of the credit class
  - the credit class metadata is updated

  Rule: Message validations

    Scenario: a valid message
      Given the message
      """
      {
        "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "class_id": "C01",
        "new_metadata": "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf"
      }
      """
      When the message is validated
      Then expect no error

    Scenario: a valid message removing metadata
      Given the message
      """
      {
        "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "class_id": "C01"
      }
      """
      When the message is validated
      Then expect no error

    
    Scenario: an error is returned if class id is empty
      Given the message
      """
      {
        "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
      }
      """
      When the message is validated
      Then expect the error "class id: empty string is not allowed: parse error: invalid request"

    Scenario: an error is returned if class id is not formatted
      Given the message
      """
      {
        "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "class_id": "foo"
      }
      """
      When the message is validated
      Then expect the error "class id: expected format <credit-type-abbrev><class-sequence>: parse error: invalid request"

    Scenario: an error is returned if new metadata exceeds 256 characters
      Given the message
      """
      {
        "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "class_id": "C01"
      }
      """
      And new metadata with length "257"
      When the message is validated
      Then expect the error "metadata: max length 256: limit exceeded"


  Rule: The credit class must exist

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with class id "C01" and admin alice

    Scenario: the credit class exists
      When alice attempts to update class metadata with class id "C01"
      Then expect no error

    Scenario: the credit class does not exist
      When alice attempts to update class metadata with class id "C02"
      Then expect the error "could not get credit class with id C02: not found: invalid request"

  Rule: The admin must be the admin of the credit class

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with class id "C01" and admin alice

    Scenario: the admin is the admin of the credit class
      When alice attempts to update class metadata with class id "C01"
      Then expect no error

    Scenario: the admin is not the admin of the credit class
      When bob attempts to update class metadata with class id "C01"
      Then expect error contains "is not the admin of credit class C01: unauthorized"

  Rule: The credit class metadata is updated

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with class id "C01" and admin alice

    Scenario: the credit class metadata is updated
      When alice attempts to update class metadata with class id "C01" and new metadata
      """
      regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf
      """
      Then expect credit class with class id "C01" and metadata
      """
      regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf
      """

      # no failing scenario - state transitions only occur upon successful message execution

  Rule: Event is emitted

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with class id "C01" and admin alice

    Scenario: EventUpdateClassMetadata is emitted
      When alice updates the class metadata
      Then expect event with properties
      """
      {
        "class_id": "C01"
      }
      """
