# Messages

## Credit Issuance, Transfer and Consumption

```go
// MsgCreateCreditClass creates a class of credits and returns a new CreditClassID
type MsgCreateCreditClass struct {
  // Designer is the entity which designs a credit class at the top-level and
  // certifies issuers
  Designer sdk.AccAddress
  // Name is the name the Designer gives to the credit, internally credits
  // are identified by their CreditClassID
  Name string
  // Issuers are those entities authorized to issue credits via MsgIssueCredit
  Issuers []sdk.AccAddress
}

type []byte CreditClassID

// MsgIssueCredit issues a credit to the Holder with the number of Units provided
// for the provided credit class, polygon, and start and end dates. A new CreditID
// is returned. It is illegal to issue a credit where the provided polygon and dates
// overlaps with those of an existing credit of the same class 
type MsgIssueCredit struct {
  CreditClass CreditClassID 
  Polygon geo.Polygon
  StartDate time.Time
  EndDate time.Time
  // Units specifies how many total units of this credit are issued for this polygon
  Units sdk.Dec
  Issuer sdk.AccAddress
  // Holder receives the credit from the issuer and can send it to other holders
  // or consume it
  Holder sdk.AccAddress
}

type []byte CreditID

// MsgSendCredit sends the provided number of units of the credit from the from
// address to the to address
type MsgSendCredit struct {
  Credit CreditID
  From sdk.AccAddress
  To sdk.AccAddress
  Units sdk.Dec
}

// MsgConsumeCredit consumes the provided number of units of the credit, essentially
// burning or retiring those units. This operation is used to actually use
// the credit as an offset. Otherwise, the holder of the credit is simply
// holding the credit as an asset but has not claimed the offset. Once a
// credit has been consumed, it can no longer be transferred
type MsgConsumeCredit struct {
  Credit CreditID
  Holder sdk.AccAddress
  Units sdk.Dec
}
```


## Credit Exchange

### Buying and Selling Credits with Coins

```go
type []byte OfferID

type ManageCreditOffer struct {
  Credits []CreditID
  Account sdk.AccAddress
  // Units should be set to 0 to delete an offer
  Units sdk.Dec
  CoinsPerUnit sdk.Coins
  // Offer should be set to nil to create a new offer
  Offer OfferID
}

type MsgManageCreditSellOffer struct {
  ManageCreditOffer
}

type MsgManageCreditBuyOffer struct {
  ManageCreditOffer
}

// MsgManageCreditClassBuyOffer can be used to generically buy credits of a
// given class irregardless of the specific credit being purchased
type MsgManageCreditClassBuyOffer struct {
  CreditClass CreditClassID
  Account sdk.AccAddress
  // Units should be set to 0 to delete an offer
  Units sdk.Dec
  CoinsPerUnit sdk.Coins
  // Offer should be set to nil to create a new offer
  Offer OfferID
}
```

### Exchanging Credits of Different Classes

This would effectively allow credits of a single class to be treated as an 
effectively fungible asset and allow trading pairs between two credit classes.

```go

type MsgManageCreditClassExchangeOffer struct {
  SellCredits []CreditID
  BuyCreditClass CreditClassID
  Account sdk.AccAddress
  // Units should be set to 0 to delete an offer
  SellUnits sdk.Dec
  BuyUnits sdk.Dec
  // Offer should be set to nil to create a new offer
  Offer OfferID
}
```
