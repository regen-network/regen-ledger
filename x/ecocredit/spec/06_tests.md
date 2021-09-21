# Tests

This document includes acceptance tests for the ecocredit module.

### when user tries to create a credit class

If a user tries to create a credit class and their address is on the list of approved credit class creators, then the credit class is created.

GIVEN - user is on list of approved credit class creators
WHEN - user tries to create a credit class
THEN - credit class is created

If a user tries to create a credit class and their address is NOT on the list of approved credit class creators, then the credit class is NOT created.

GIVEN - user is NOT on list of approved credit class creators
WHEN - user tries to create a credit class
THEN - credit class is NOT created

If a user tries to create a credit class and the credit class includes a credit type on the list of approved credit types, then the credit class is created.

GIVEN - credit type is on list of approved credit types
WHEN - user tries to create a credit class
THEN - credit class is created

If a user tries to create a credit class and the credit class includes a credit type NOT on the list of approved credit types, then the credit class is created.

GIVEN - credit type is NOT on list of approved credit types
WHEN - user tries to create a credit class
THEN - credit class is NOT created

If a user tries to create a credit class and the user provides metadata that is less than 256 bytes, then the credit class is created.

GIVEN - user provides metadata that is less than 256 bytes
WHEN - user tries to create a credit class
THEN - credit class is created

If a user tries to create a credit class and the user provides metadata that is more than 256 bytes, then the credit class is NOT created.

GIVEN - user provides metadata that is more than 256 bytes
WHEN - user tries to create a credit class
THEN - credit class is NOT created

If a user tries to create a credit class and the user provides a valid issuer address, then the credit class is created.

GIVEN - user provides a valid issuer address
WHEN - user tries to create a credit class
THEN - credit class is created

If a user tries to create a credit class and the user provides an invalid issuer address, then the credit class is NOT created.

GIVEN - user provides an invalid issuer address
WHEN - user tries to create a credit class
THEN - credit class is NOT created

If a user tries to create a credit class and the user does not provide an issuer address, then the credit class is NOT created.

GIVEN - user does not provide an issuer address
WHEN - user tries to create a credit class
THEN - credit class is NOT created

...

### when user tries to create a credit batch

If a user tries to create a credit batch and their address is on the list of approved credit issuers, then the credit batch is created.

GIVEN - user is on list of approved credit issuers
WHEN - user tries to create a credit batch
THEN - credit batch is created

If a user tries to create a credit batch and their address is NOT on the list of approved credit issuers, then the credit batch is NOT created.

GIVEN - user is NOT on list of approved credit issuers
WHEN - user tries to create a credit batch
THEN - credit batch is NOT created

If a user tries to create a credit batch and the user provides metadata that is less than 256 bytes, then the credit batch is created.

GIVEN - user provides metadata that is less than 256 bytes
WHEN - user tries to create a credit batch
THEN - credit batch is created

If a user tries to create a credit batch and the user provides metadata that is more than 256 bytes, then the credit batch is NOT created.

GIVEN - user provides metadata that is more than 256 bytes
WHEN - user tries to create a credit batch
THEN - credit batch is NOT created

...

### when user tries to transfer credits

If a user tries to transfer 20 credits and their tradable balance is less than 20 credits, then the credits are NOT transferred.

GIVEN - tradable balance is less than 20 credits
WHEN - user tries to transfer 20 credits
THEN - credits are NOT transferred

If a user tries to transfer 20 credits and their tradable balance is less than 20 credits, then the credits are transferred.

GIVEN - tradable balance is more than 20 credits
WHEN - user tries to transfer 20 credits
THEN - credits are transferred

If a user tries to transfer credits and the user provides a valid recipient address, then the credits are transferred.

GIVEN - user provides a valid recipient address
WHEN - user tries to transfer credits
THEN - credits are transferred

If a user tries to transfer credits and the user provides an invalid recipient address, then the credits are NOT transferred.

GIVEN - user provides an invalid recipient address
WHEN - user tries to transfer credits
THEN - credits are NOT transferred

...

### when user tries to retire credits

If a user tries to retire 20 credits and their tradable balance is less than 20 credits, then the credits are NOT retired.

GIVEN - tradable balance is less than 20 credits
WHEN - user tries to transfer 20 credits
THEN - credits are NOT retired

If a user tries to retire 20 credits and their tradable balance is more than 20 credits, then the credits are retired.

GIVEN - tradable balance is more than 20 credits
WHEN - user tries to transfer 20 credits
THEN - credits are retired

...

### when user tries to cancel credits

If a user tries to cancel 20 credits and their tradable balance is less than 20 credits, then the credits are NOT cancelled.

GIVEN - tradable balance is less than 20 credits
WHEN - user tries to cancel 20 credits
THEN - credits are NOT cancelled

If a user tries to cancel 20 credits and their tradable balance is more than 20 credits, then the credits are cancelled.

GIVEN - tradable balance is more than 20 credits
WHEN - user tries to cancel 20 credits
THEN - credits are cancelled

...