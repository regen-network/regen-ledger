# Definitions and Basic Concepts

Regardless of how the term is used elsewhere, we use the word "oracle" to 
refer to a computer or computer system that reliably executes "compute functions"
which have access to "well-known" on and off-chain data. A "compute function"
in this context refers to a piece of code that:
1) is deterministic - that is given the same set of inputs, ir will always
produce the same output, and
2) conforms to a well-defined specification that spells out how these functions
can access external data

Once we have an agreed upon specification of how compute functions are supposed
to behave, in order to trust that oracles are faithfully executing we these
functions, we need a sufficiently robust "consensus protocol" around oracles
computations. The consensus needed from an oracle can be context dependent -
there may be scenarios where the downstream users of oracle computation results
fully trust the oracle in all cases and other scenarios where for whatever reason
we need to have a more "byzantine fault tolerant" protocol.

Later in this document we will describe both the specification for compute
functions and outline a few consensus mechanisms which can give us varying
degrees of certainty that a computation was faithfully executed. 

# Motivation and Rationale

The motivation for creating such a framework is to create a scalable, cost
effective and decentralized system for the verification of ecological state
and change of state. In a world where climate change and other environmental
catastrophes put the very future of human civilization into question, it is
essential that we have mechanisms for ensuring that the vast amounts of money
that are being and will need to be spent on mitigation efforts are having the
most beneficial effects possible. Since this is a global issue, we feel that it
requires global scale infrastructure which we take to mean a number of things:
- the system should scale to an arbitrarily large number of computations
- it should be possible for third parties to independently audit and contest
results (given the appropriate permissions to the underlying data)
- the underlying compute infrastructure shouldn't be in the hands of a single
profit seeking entity

Many of the computations that we envision for this system involve complex
analysis of remote sensing imagery. If preprocessing of a time series of images
is involved, the load for running the computation even once can be quite high
(several hours on a single modern machine). In order to scale a system like this
we need to be able to have reasonable trust in results even if the computation
was only done by one or two computers - unlike a full blockchain where the
desired level of consensus is generally much higher.

# Compute Function Specification

# Consensus Mechanisms

