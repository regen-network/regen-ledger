# Data Module (gRPC)

For the purposes of this guide, we'll use [gRPCurl](https://github.com/fullstorydev/grpcurl), a command-line tool
which acts as a `curl` replacement for gRPC services. Follow the instructions on their [github](https://github.com/fullstorydev/grpcurl)
to download the binary.

Now let's get to querying!

Assuming you have a local node running (either a localnet, or connected to our devnet), you should be able
to run the following to list the query services available:

```sh
grpcurl -plaintext localhost:9090 list
```

You should see a list of gRPC services like `cosmos.bank.v1beta1.Query`, `regen.data.v1alpha1.Query`. Each of these
represents a different API endpoint which you can query for some relevant state from the blockchain.

Let's see if we can query the same CID that we anchored and signed from the previous section. To find out
which methods are exposed via the `regen.data.v1alpha1.Query` service, we need to provide `grpcurl` with
the .proto files corresponding to this service:

```sh
$ grpcurl -proto ../proto/regen/data/v1alpha1/query.proto describe regen.data.v1alpha1.Query
# regen.data.v1alpha1.Query is a service:
# // Query is the regen.data.v1alpha1 Query service
# service Query {
#   // ByCid queries data based on its CID.
#   rpc ByCid ( .regen.data.v1alpha1.QueryByCidRequest ) returns ( .regen.data.v1alpha1.QueryByCidResponse );
# }
```

Here we see that there is one method, `ByCid`, which takes a `QueryByCidRequest`, and returns a `QueryByCidResponse`.
The structure of these request & response messages can be similarly queried with `grpcurl`, or you can find
details on them in our [protobuf documentation](./modules/data/protobuf.md#regen-data-v1alpha1-query-proto).

Now that we know the structure of our `QueryByCidRequest`, we can query the `ByCid` method directly using a JSON
encoding of the `QueryByCidRequest` message.

_Note: Since gRPCurl requires bytes to be encoded as base64 strings, we have to do some gymnastics to decode our CID
and re-encode the raw bytes using base64. The base64 CID string below is the correct one for the `$EXAMPLE_CID` in
the CLI tutorial above. For more details on how CIDs work see the [CID spec](https://github.com/multiformats/cid)._

```sh
grpcurl -proto ../proto/regen/data/v1alpha1/query.proto \
    -d '{"cid": "AVUSIG5v95UKNhh6gBYTQm6Fjc5obNfX48D8Qu4DMActJFyV"}' \
    -plaintext localhost:9090 regen.data.v1alpha1.Query/ByCid
```

The result should look something like this:

```json
{
  "timestamp": "2020-11-14T03:45:50.924123Z",
  "signers": [
    "regen1qx9tsl46vkgf9tf8pl3chm7mkmqa5s30ydjdzx"
  ]
}
```
