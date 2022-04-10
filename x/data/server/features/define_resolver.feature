Feature: Define Resolver

  Scenario: a resolver is defined when the resolver url is unique
    When alice attempts to define a resolver with url "https://foo.bar"
    Then an error of ""
    And the resolver info entry exists with url "https://foo.bar"

  Scenario: an error is returned when a resolver with the same url has been defined
    Given alice has defined a resolver with url "https://foo.bar"
    When alice attempts to define a resolver with url "https://foo.bar"
    Then an error of "resolver with url https://foo.bar already exists"

  # Note: see ../features/msg_define_resolver.feature for resolver url validation
