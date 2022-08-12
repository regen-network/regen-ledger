```trig
@prefix rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#> .
@prefix : <xrn://data/>
@prefix xrn: <http://regen.network/schema>

# default graph
{
  :poly1 geo:wkt "Polygon(....)"
  
  :verification1 rdf:type xrn:VerificationResult .
  :verification1 xrn:verificationProtocol :regen_index .
  :verification1 geo:polygon :poly1.
  :verification1 xrn:verificationStartDate "2018-01-01"
  :verification1 xrn:verificationEndDate "2018-12-31"
  :verification1 xrn:verificationResult :ab823nsdg83t.
  
  :verification1 rdf:type xrn:VerificationResult .
  :verification1 xrn:verificationProtocol :soil_health_idx .
  :verification1 geo:polygon :poly1.
  :verification1 xrn:verificationStartDate "2018-01-01"
  :verification1 xrn:verificationEndDate "2018-12-31"
  :verification1 xrn:verificationResult :xtyys3thv8.
}


# named graph
:ab823nsdg83t
{
  :poly1 :biodiversityScore 9.5 .
  :poly1 :carbonTonsInLastYear 115.0 .
  :poly1 :waterHealth 4.1 .
}

:xtyys3thv8
{
  :poly1 :soilHealthIndex 310 .
}
```

```sparql
CONSTRUCT
{
  $poly1 :soilHealthIndex ?score .
}
WHERE
{
  ?v xrn:verificationProtocol :regen_index .
  ?v geo?polygon $poly1 .
  ?v xrn:verificationResult ?data .
  GRAPH ?data
  {
    $poly1 :carbonTonsInLastYear ?carbon .
    $poly1 :waterHealth ?water .
    BINDING (3 * ?carbon + ?water AS ?score)
  }
}
```
