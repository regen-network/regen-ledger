Feature: Define Resolver

  Rule: the resolver is defined with no unique constraints on the URL

    Scenario: the url is unique
      When alice attempts to define a resolver with url "https://foo.bar"
      Then expect the resolver with id "1" and url "https://foo.bar" and manager alice

    Scenario: the url is not unique
      Given alice has defined a resolver with url "https://foo.bar"
      When alice attempts to define a resolver with url "https://foo.bar"
      Then expect the resolver with id "2" and url "https://foo.bar" and manager alice
