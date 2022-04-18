Feature: Define Resolver

  Background:
    Given the resolver url "https://foo.bar"

  Rule: the resolver is defined if the resolver url is unique

    Scenario: the resolver has not been defined
      When alice attempts to define the resolver
      Then the resolver exists and alice is the manager

    Scenario: the resolver has already been defined
      Given alice has defined the resolver
      When alice attempts to define the resolver
      Then expect the error "resolver URL already exists"
