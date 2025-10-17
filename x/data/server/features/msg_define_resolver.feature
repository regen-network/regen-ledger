Feature: Msg/DefineResolver

  A resolver can be defined:
  - when the url and manager combination is unique

  Rule: Message validation
  
    Scenario: a valid message
      Given the message
      """
      {
        "definer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "resolver_url": "https://foo.bar"
      }
      """
      When the message is validated
      Then expect no error

    Scenario: an error is returned if resolver url is empty
      Given the message
      """
      {
        "definer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
      }
      """
      When the message is validated
      Then expect the error "invalid resolver url: invalid request"

    Scenario: an error is returned if resolver url is missing a protocol
      Given the message
      """
      {
        "definer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "resolver_url": "foo.bar"
      }
      """
      When the message is validated
      Then expect the error "invalid resolver url: invalid request"

  Rule: The url and manager combination must be unique

    Scenario: the url and manager combination is unique
      Given alice has defined a resolver with url "https://foo.bar"
      When bob attempts to define a resolver with url "https://foo.bar"
      Then expect the resolver with id "2" and url "https://foo.bar" and manager bob

    Scenario: the url and manager combination is not unique
      Given alice has defined a resolver with url "https://foo.bar"
      When alice attempts to define a resolver with url "https://foo.bar"
      Then expect the error "a resolver with the same URL and manager already exists: unique key violation"

    Scenario: public resolvers can only be defined once per URL
      Given a public resolver is defined for the url "ipfs:"
      When alice attempts to define a public resolver with url "ipfs:"
      Then expect the error "a resolver with the same URL and manager already exists: unique key violation"

  Rule: Event is emitted

    Scenario: EventDefineResolver is emitted
      Given alice has defined a resolver with url "https://foo.bar"
      When bob attempts to define a resolver with url "https://foo.bar"
      Then expect the resolver with id "2" and url "https://foo.bar" and manager bob
      And expect event with properties
      """
      {
        "id": 2
      }
      """