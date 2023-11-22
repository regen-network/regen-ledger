Feature: Msg/UpdateCurator

  The curator of a basket can be updated:
  - when the basket exists
  - when the curator is the admin of the basket
  - the basket curator is updated

  Rule: The basket must exist

    Background:
      Given a basket with properties and curator alice
      """
      {
        "name":"basket1",
        "basket_denom":"eco.uC.NCT",
        "credit_type_abbrev": "C"
      }
      """

    Scenario: the basket exists
      When alice attempts to update basket curator with denom "eco.uC.NCT"
      Then expect no error

    Scenario: the basket does not exist
      When alice attempts to update basket curator with denom "eco.uC.rNCT"
      Then expect the error "basket with denom eco.uC.rNCT: not found"

  Rule: The curator must be the curator of the basket

    Background:
      Given a basket with properties and curator alice
      """
      {
        "name":"basket1",
        "basket_denom":"eco.uC.NCT",
        "credit_type_abbrev": "C"
      }
      """

    Scenario: the curator is the curator of the basket
      When alice attempts to update basket curator with denom "eco.uC.NCT"
      Then expect no error

    Scenario: the curator is not the curator of the basket
      When bob attempts to update basket curator with denom "eco.uC.NCT"
      Then expect error contains "unauthorized"

  Rule: The basket curator is updated

    Background:
      Given a basket with properties and curator alice
      """
      {
        "name":"basket1",
        "basket_denom":"eco.uC.NCT",
        "credit_type_abbrev": "C"
      }
      """

    Scenario: the basket curator is updated
      When alice attempts to update basket curator with denom "eco.uC.NCT" and new curator bob
      Then expect basket with denom "eco.uC.NCT" and curator bob

    # no failing scenario - state transitions only occur upon successful message execution

  Rule: Event is emitted

    Background:
      Given a basket with properties and curator alice
      """
      {
        "name":"basket1",
        "basket_denom":"eco.uC.NCT",
        "credit_type_abbrev": "C"
      }
      """

    Scenario: the basket curator is updated
      When alice attempts to update basket curator with denom "eco.uC.NCT"
      Then expect basket with denom "eco.uC.NCT" and curator bob
      And expect event with properties
      """
      {
        "denom": "eco.uC.NCT"
      }
      """
