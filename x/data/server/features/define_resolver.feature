Feature: Define Resolver

  Background: the message has been validated
    Given the resolver url "https://foo.bar"

  Rule: a resolver is defined if the resolver url is unique

    Scenario: a resolver with the same url has not been defined
      When alice attempts to define a resolver
      Then the resolver info entry exists

    Scenario: a resolver with the same url has already been defined
      Given alice has defined a resolver
      When alice attempts to define a resolver
      Then expect the error "resolver URL already exists"
