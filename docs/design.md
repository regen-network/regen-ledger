# xrn

## Schema Design Principles

* Attribute, class and function API definition never change or at least never break backwards compatibility
* If different functionality is needed a new versioned attribute, class or function is created, or we create some convenient facility for ensuring forward compatbility 
* The problem with trying to design for forward compatibility is that we need to think about both sender and receiver forward compatibility
* It should generally be safe to add more optional attributes to an existing class
* Attributes can neither be more specific (breaks sender forward compatibility) nor general (breaks receiver forward compatibility)
* It should generally be safe to add more optional arguments to existing functions

So maybe the way to do this is have schema versions be content (hash) addressable and then in later schema versions classes or function definitions that extend earlier versions specify explicitly that they extend an earlier content-addressed schema.

The issue with this is that then there is no persistent namespacing logic. A better approach is o address schemas first by authority and then have only-forward compatible natural language definitions within the scope of that authority. So for example, the `name` attribute in regen.network could be set to a string and that would be its persistent definition for lifetime. There could of course be namespacing within an authority, so regen.network could have an attribute `common/name`.

Since what we are talking about is some rules surrounding the evolution of names within the scope of what's defined by an authority, a blockchain itself might be a good place to define and track schemas.

 

