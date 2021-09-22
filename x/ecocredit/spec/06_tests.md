# Tests

This document includes acceptance tests for the ecocredit module.

### Create Credit Class

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

If a user tries to create a credit class and the user provides metadata that is base64 encoded, then the credit class is created.

GIVEN - user provides metadata that is base64 encoded
WHEN - user tries to create a credit class
THEN - credit class is created

If a user tries to create a credit class and the user provides metadata that is NOT base54 encoded, then the credit class is NOT created.

GIVEN - user provides metadata that is NOT base54 encoded
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

If a user tries to create a credit class and the user provides a valid credit type (credit type is included in the list of approved credit types), then the credit class is created.

GIVEN - user provides a valid credit type
WHEN - user tries to create a credit class
THEN - credit class is created

If a user tries to create a credit class and the user provide an invalid credit type (credit type is NOT included in the list of approved credit types), then the credit class is NOT created.

GIVEN - user provides an invalid credit type
WHEN - user tries to create a credit class
THEN - credit class is NOT created

If a user tries to create a credit class and the user account balance is less than the credit class fee, then the credit class is NOT created and the credit class fee is NOT deducted from the user account balance.

GIVEN - user account balance is less than the credit class fee
WHEN - user tries to create a credit class
THEN - credit class is NOT created and the credit class fee is NOT deducted from the user account balance

If a user tries to create a credit class and the user account balance is equal to or more than the credit class fee, then the credit class is created and the credit class fee is deducted from the user account balance.

GIVEN - user account balance is equal to or more than the credit class fee
WHEN - user tries to create a credit class
THEN - credit class is NOT created and the credit class fee is deducted from the user account balance

### Create Credit Batch

If a user tries to create a credit batch and their account address is on the list of approved credit issuers for the given credit class, then the credit batch is created.

GIVEN - user is on list of approved credit issuers for the given credit class
WHEN - user tries to create a credit batch
THEN - credit batch is created

If a user tries to create a credit batch and their account address is NOT on the list of approved credit issuers for the given credit class, then the credit batch is NOT created.

GIVEN - user is NOT on list of approved credit issuers for the given credit class
WHEN - user tries to create a credit batch
THEN - credit batch is NOT created

If a user tries to create a credit batch and the user provides a valid recipient address, then the credit batch is created.

GIVEN - user provides a valid recipient address
WHEN - user tries to create a credit batch
THEN - credit batch is created

If a user tries to create a credit batch and the user provides an invalid recipient address, then the credit batch is NOT created.

GIVEN - user provides an invalid recipient address
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

If a user tries to create a credit batch and the user provides a valid project location, then the credit batch is created.

GIVEN - user provides a valid project location
WHEN - user tries to create a credit batch
THEN - credit batch is created

If a user tries to create a credit batch and the user provides an invalid project location, then the credit batch is NOT created.

GIVEN - user provides an invalid project location
WHEN - user tries to create a credit batch
THEN - credit batch is NOT created

If a user tries to create a credit batch and the user provides a valid start and end date, then the credit batch is created.

GIVEN - user provides a valid start and end date
WHEN - user tries to create a credit batch
THEN - credit batch is created

If a user tries to create a credit batch and the user provides an invalid start and end date, then the credit batch is NOT created.

GIVEN - user provides an invalid start and end date
WHEN - user tries to create a credit batch
THEN - credit batch is NOT created

If a user tries to create a credit batch and the user includes retired credits with a retirement location, then the credit batch is created.

GIVEN - user includes retired credits with a retirement location
WHEN - user tries to create a credit batch
THEN - credit batch is created

If a user tries to create a credit batch and the user includes retired credits without a retirement location, then the credit batch is NOT created.

GIVEN - user includes retired credits without a retirement location
WHEN - user tries to create a credit batch
THEN - credit batch is NOT created

### Transfer Credits

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

### Retire Credits

If a user tries to retire 20 credits and their tradable balance is less than 20 credits, then the credits are NOT retired.

GIVEN - tradable balance is less than 20 credits
WHEN - user tries to transfer 20 credits
THEN - credits are NOT retired

If a user tries to retire 20 credits and their tradable balance is more than 20 credits, then the credits are retired.

GIVEN - tradable balance is more than 20 credits
WHEN - user tries to transfer 20 credits
THEN - credits are retired

...

### Cancel Credits

If a user tries to cancel 20 credits and their tradable balance is less than 20 credits, then the credits are NOT cancelled.

GIVEN - tradable balance is less than 20 credits
WHEN - user tries to cancel 20 credits
THEN - credits are NOT cancelled

If a user tries to cancel 20 credits and their tradable balance is more than 20 credits, then the credits are cancelled.

GIVEN - tradable balance is more than 20 credits
WHEN - user tries to cancel 20 credits
THEN - credits are cancelled

...