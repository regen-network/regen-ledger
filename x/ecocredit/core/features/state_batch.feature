Feature: Batch

  Scenario: a valid batch
    Given the batch
    """
    {
      "denom": "C01-001-20200101-20210101-001",
      "project_key": 1,
      "start_date": "2020-01-01T00:00:00Z",
      "end_date": "2021-01-01T00:00:00Z",
      "issuer": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y="
    }
    """
    When the batch is validated
    Then expect no error

  Scenario: an error is returned if batch denom is empty
    Given the batch
    """
    {}
    """
    When the batch is validated
    Then expect the error "batch denom cannot be empty: parse error"

  Scenario: an error is returned if batch denom is not formatted
    Given the batch
    """
    {
      "denom": "foo"
    }
    """
    When the batch is validated
    Then expect the error "invalid batch denom: expected format A00-000-00000000-00000000-000: parse error"

  Scenario: an error is returned if project key is empty
    Given the batch
    """
    {
      "denom": "C01-001-20200101-20210101-001"
    }
    """
    When the batch is validated
    Then expect the error "project key cannot be zero: parse error"

  Scenario: an error is returned if start date is empty
    Given the batch
    """
    {
      "denom": "C01-001-20200101-20210101-001",
      "project_key": 1
    }
    """
    When the batch is validated
    Then expect the error "must provide a start date for the credit batch: parse error"

  Scenario: an error is returned if end date is empty
    Given the batch
    """
    {
      "denom": "C01-001-20200101-20210101-001",
      "project_key": 1,
      "start_date": "2020-01-01T00:00:00Z"
    }
    """
    When the batch is validated
    Then expect the error "must provide an end date for the credit batch: parse error"

  Scenario: an error is returned if start date is equal to end date
    Given the batch
    """
    {
      "denom": "C01-001-20200101-20210101-001",
      "project_key": 1,
      "start_date": "2020-01-01T00:00:00Z",
      "end_date": "2020-01-01T00:00:00Z"
    }
    """
    When the batch is validated
    Then expect the error "the batch end date (2020-01-01T00:00:00Z) must be the same as or after the batch start date (2020-01-01T00:00:00Z): parse error"

  Scenario: an error is returned if start date after end date
    Given the batch
    """
    {
      "denom": "C01-001-20200101-20210101-001",
      "project_key": 1,
      "start_date": "2021-01-01T00:00:00Z",
      "end_date": "2020-01-01T00:00:00Z"
    }
    """
    When the batch is validated
    Then expect the error "the batch end date (2020-01-01T00:00:00Z) must be the same as or after the batch start date (2021-01-01T00:00:00Z): parse error"

  Scenario: an error is returned if issuer is empty
    Given the batch
    """
    {
      "denom": "C01-001-20200101-20210101-001",
      "project_key": 1,
      "start_date": "2020-01-01T00:00:00Z",
      "end_date": "2021-01-01T00:00:00Z"
    }
    """
    When the batch is validated
    Then expect the error "issuer: empty address string is not allowed: parse error"
