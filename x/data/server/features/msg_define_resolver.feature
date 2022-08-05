Feature: Msg/DefineResolver

  A resolver can be defined:
  - when the url and manager combination is unique

  Rule: The url and manager combination must be unique

    Scenario: the url and manager combination is unique
      Given alice has defined a resolver with url "https://foo.bar"
      When bob attempts to define a resolver with url "https://foo.bar"
      Then expect the resolver with id "2" and url "https://foo.bar" and manager bob

    Scenario: the url and manager combination is not unique
      Given alice has defined a resolver with url "https://foo.bar"
      When alice attempts to define a resolver with url "https://foo.bar"
      Then expect the error "a resolver with the same URL and manager already exists: unique key violation"
