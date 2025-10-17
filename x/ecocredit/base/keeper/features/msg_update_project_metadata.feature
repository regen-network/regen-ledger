Feature: Msg/UpdateProjectMetadata

  The metadata of a project can be updated:
  - message validations
  - when the project exists
  - when the admin is the admin of the project
  - the project metadata is updated

  Rule: Message Validations

  Scenario: a valid message
  Given the message
  """
  {
  "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
  "project_id": "C01-001",
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
  "project_id": "C01-001"
  }
  """
  When the message is validated
  Then expect no error

  Scenario: an error is returned if project id is empty
  Given the message
  """
  {
  "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
  }
  """
  When the message is validated
  Then expect the error "project id: empty string is not allowed: parse error: invalid request"

  Scenario: an error is returned if project id is not formatted
  Given the message
  """
  {
  "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
  "project_id": "foo"
  }
  """
  When the message is validated
  Then expect the error "project id: expected format <class-id>-<project-sequence>: parse error: invalid request"

  Scenario: an error is returned if new metadata exceeds 256 characters
  Given the message
  """
  {
  "admin": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
  "project_id": "C01-001"
  }
  """
  And new metadata with length "257"
  When the message is validated
  Then expect the error "metadata: max length is 256: limit exceeded"


  Rule: The project must exist

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with class id "C01"
      And a project with project id "C01-001" and admin alice

    Scenario: the project exists
      When alice attempts to update project metadata with project id "C01-001"
      Then expect no error

    Scenario: the project does not exist
      When alice attempts to update project metadata with project id "C01-002"
      Then expect the error "could not get project with id C01-002: not found: invalid request"

  Rule: The admin must be the admin of the project

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with class id "C01"
      And a project with project id "C01-001" and admin alice

    Scenario: the admin is the admin of the project
      When alice attempts to update project metadata with project id "C01-001"
      Then expect no error

    Scenario: the admin is not the admin of the project
      When bob attempts to update project metadata with project id "C01-001"
      Then expect error contains "is not the admin of project C01-001: unauthorized"

  Rule: The project metadata is updated

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with class id "C01"
      And a project with project id "C01-001" and admin alice

    Scenario: the project metadata is updated
      When alice attempts to update project metadata with project id "C01-001" and new metadata
      """
      regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf
      """
      Then expect project with project id "C01-001" and metadata
      """
      regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf
      """

    # no failing scenario - state transitions only occur upon successful message execution

  Rule: Event is emitted

    Background:
      Given a credit type with abbreviation "C"
      And a credit class with class id "C01"
      And a project with project id "C01-001" and admin alice

    Scenario: EventUpdateProjectMetadata is emitted
      When alice attempts to update project metadata with project id "C01-001"
      Then expect event with properties
      """
      {
        "project_id": "C01-001"
      }
      """