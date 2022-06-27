Feature: Define Resolver

  A resolver can be defined:
  - when the url and manager pair is unique

  Rule: The url and manager pair must be unique

    Scenario: the url and manager pair is unique
      Given alice has defined a resolver with url "https://foo.bar"
      When bob attempts to define a resolver with url "https://foo.bar"
      Then expect the resolver with id "1" and url "https://foo.bar" and manager alice

    Scenario: the url and manager pair is not unique
      Given alice has defined a resolver with url "https://foo.bar"
      When alice attempts to define a resolver with url "https://foo.bar"
      Then expect the error "a resolver with the same URL and manager already exists: unique key violation"
