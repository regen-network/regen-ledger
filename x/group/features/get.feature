Feature: Get group details
  Scenario:
    Given a group ID
    When a user gets the group details on the command line
    Then they should get back the group details in JSON format