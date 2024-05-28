Feature: Msg/CreateOrUpdateApplication

  Background:
    Given project "P001" with admin "Alice"
    And credit class "C01"

  Rule: when there is no application for the project, the project admin can create an application
    Background:
      Given an application for "P001" to "C01" does not exist

    Scenario: admin creates an application
      When "Alice" creates or updates an application for "P001" to "C01" with metadata "xyz123"
      Then expect no error
      * expect the application for "P001" to "C01" exists with metadata "xyz123"
      * expect EventUpdateApplication with properties
      """
      { "project_id": "P001", "class_id": "C01", "action": "ACTION_CREATE",
        "new_application_metadata": "xyz123" }
      """

    Scenario: non-admin attempts to create an application
      When "Bob" creates or updates an application for "P001" to "C01" with metadata "xyz123"
      Then expect error contains "unauthorized"
      And an application for "P001" to "C01" does not exist

  Rule: when there is an existing application, the project admin can update the metadata
    Background:
      Given an application for "P001" to "C01" with metadata "abc123"

    Scenario: admin updates the application metadata
      When "Alice" creates or updates an application for "P001" to "C01" with metadata "foobar379"
      Then expect no error
      * expect the application for "P001" to "C01" exists with metadata "foobar379"
      * expect EventUpdateApplication with properties
      """
      { "project_id": "P001", "class_id": "C01", "action": "ACTION_UPDATE",
        "new_application_metadata": "foobar379" }
      """

    Scenario: non-admin attempts to update the application metadata
      When "Bob" creates or updates an application for "P001" to "C01" with metadata "foobar379"
      Then expect error contains "unauthorized"
      And expect the application for "P001" to "C01" exists with metadata "abc123"

  Rule: when there is an existing application, the project admin can withdraw the application and it is removed from state
    Background:
      Given an application for "P001" to "C01" with metadata "abc123"

    Scenario: admin withdraws the application
      When "Alice" attempts to withdraw the application for "P001" to "C01" with metadata "foobar379"
      Then expect no error
      * an application for "P001" to "C01" does not exist
      * expect EventUpdateApplication with properties
      """
      { "project_id": "P001", "class_id": "C01", "action": "ACTION_WITHDRAW",
        "new_application_metadata": "foobar379" }
      """

    Scenario: non-admin attempts to withdraw the application
      When "Bob" attempts to withdraw the application for "P001" to "C01" with metadata "bob123"
      Then expect error contains "unauthorized"
      And expect the application for "P001" to "C01" exists with metadata "abc123"

  Rule: project admins cannot withdraw a project with an accepted enrollment (a credit class issuer must do that)
    Scenario: project is accepted and admin attempts to withdraw the application
      Given an application for "P001" to "C01" with metadata "abc123"
      And the application for "P001" to "C01" is accepted
      When "Alice" attempts to withdraw the application for "P001" to "C01" with metadata "foobar379"
      Then expect error contains "cannot withdraw accepted enrollment"
      And expect the application for "P001" to "C01" exists with metadata "abc123"

