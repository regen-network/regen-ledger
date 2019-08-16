# Messages

```go
type MsgCreateCreditClass struct {
  Curator sdk.AccAddress
  Name string
  Issuers []sdk.AccAddress
}

type []byte CreditClassID

type MsgIssueCredit struct {
  CreditClass CreditClassID 
  Polygon geo.Polygon
  Issuer sdk.AccAddress
  Vendor sdk.AccAddress
}

type []byte CreditID

type MsgAssignCredit struct {
  Credit CreditID
  Vendor sdk.AccAddress
  Offsetter sdk.AccAddress
  SquareMeters sdk.Dec
}

type MsgOfferCredit struct {
  Credit CreditID
  Vendor sdk.AccAddress
  MaxSquareMeters sdk.Dec
  CoinsPerSquareMeter sdk.Coins
}

type MsgBuyCredit struct {
  Credit CreditID
  SquareMeters sdk.Dec
  Buyer sdk.AccAddress
  Coins sdk.Coins
}
```
