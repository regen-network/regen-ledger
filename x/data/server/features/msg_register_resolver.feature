Feature: Msg/RegisterResolver

  Background:
    Given the content hash
    """
    {
      "raw": {
        "hash": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=",
        "digest_algorithm": 1,
        "file_extension": "bin"
      }
    }
    """

  Rule: Message Validation
    Scenario: a valid message
      Given the message
      """
      {
        "signer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "resolver_id": 1,
        "content_hashes": [
          {
            "raw": {
              "hash": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=",
              "digest_algorithm": 1,
              "file_extension": "txt"
            }
          }
        ]
      }
      """
      When the message is validated
      Then expect no error

    Scenario: an error is returned if resolver id is empty
      Given the message
      """
      {
        "signer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6"
      }
      """
      When the message is validated
      Then expect the error "resolver id cannot be empty: invalid request"

    Scenario: an error is returned if content hashes is empty
      Given the message
      """
      {
        "signer": "regen1depk54cuajgkzea6zpgkq36tnjwdzv4ak663u6",
        "resolver_id": 1
      }
      """
      When the message is validated
      Then expect the error "content hashes cannot be empty: invalid request"

  Rule: the data is registered if the resolver is defined

    Scenario: the resolver has not been defined
      When alice attempts to register the data to a resolver with id "1"
      Then expect the error "resolver with id 1 does not exist: not found"

    Scenario: the resolver has already been defined
      Given alice has defined the resolver with url "https://foo.bar"
      When alice attempts to register the data to the resolver
      Then the data resolver entry exists

  Rule: the data is anchored if not already anchored

    Background: the resolver has already been defined
      Given alice has defined the resolver with url "https://foo.bar"

    Scenario: the data has not been anchored
      When alice attempts to register the data to the resolver at block time "2020-01-01"
      Then the anchor entry exists with timestamp "2020-01-01"

    Scenario: the data has already been anchored
      Given alice has anchored the data at block time "2020-01-01"
      When alice attempts to register the data to the resolver at block time "2020-01-02"
      Then the anchor entry exists with timestamp "2020-01-01"

  Rule: the data is registered if not already registered

    Background: the resolver has already been defined
      Given alice has defined the resolver with url "https://foo.bar"

    Scenario: the data has not been registered
      When alice attempts to register the data to the resolver
      Then the data resolver entry exists

    Scenario: the data has already been registered
      Given alice has registered the data to the resolver
      When alice attempts to register the data to the resolver
      Then the data resolver entry exists

  Rule: the data is registered if the registrant is the manager

    Background: the resolver has already been defined
      Given alice has defined the resolver with url "https://foo.bar"

    Scenario: the registrant is the manager
      When alice attempts to register the data to the resolver
      Then the data resolver entry exists

    Scenario: the registrant is not the manager
      When bob attempts to register the data to the resolver
      Then expect the error "unauthorized resolver manager"

  Rule: event is emitted

    Scenario: EventRegisterResolver is emitted
      Given alice has defined the resolver with url "https://foo.bar"
      When alice attempts to register the data to the resolver
      Then the data resolver entry exists
      And expect event with properties
      """
      {
        "id": 1,
        "iri": "regen:112wkBET2rRgE8pahuaczxKbmv7ciehqsne57F9gtzf1PVhwuFTX.bin"
      }
      """

  Rule: public resolvers allow anyone to register data

    Scenario: register data to public resolver
      Given alice has defined a public resolver with url "ipfs:"
      When bob attempts to register the data to the resolver
      Then the data resolver entry exists
