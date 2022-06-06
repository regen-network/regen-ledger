# Tests

This document includes acceptance tests for the ecocredit module.

### Create Credit Class

If a user tries to create a credit class and their address is on the list of approved credit class creators and the allowlist is enabled, then the transaction is successful and the credit class is created.

- GIVEN - user is on list of approved credit class creators and the allowlist is enabled
- WHEN - user tries to create a credit class
- THEN - transaction is successful, credit class is created

If a user tries to create a credit class and their address is NOT on the list of approved credit class creators and the allowlist is enabled, then the transaction fails and the credit class is NOT created.

- GIVEN - user is NOT on list of approved credit class creators and the allowlist is enabled
- WHEN - user tries to create a credit class
- THEN - transaction fails, credit class is NOT created

If a user tries to create a credit class and their address is NOT on the list of approved credit class creators and the allowlist is disabled, then the transaction is successful and the credit class is created.

- GIVEN - user is NOT on list of approved credit class creators and the allowlist is disabled
- WHEN - user tries to create a credit class
- THEN - transaction is successful, credit class is created

If a user tries to create a credit class and the user provides a valid credit type (credit type is included in the list of approved credit types), then the transaction is successful and the credit class is created.

- GIVEN - user provides a valid credit type
- WHEN - user tries to create a credit class
- THEN - transaction is successful, credit class is created

If a user tries to create a credit class and the user provide an invalid credit type (credit type is NOT included in the list of approved credit types), then the transaction fails and the credit class is NOT created.

- GIVEN - user provides an invalid credit type
- WHEN - user tries to create a credit class
- THEN - transaction fails, credit class is NOT created

If a user tries to create a credit class and the user provides metadata that is base64 encoded, then the transaction is successful and the credit class is created.

- GIVEN - user provides metadata that is base64 encoded
- WHEN - user tries to create a credit class
- THEN - transaction is successful, credit class is created

If a user tries to create a credit class and the user provides metadata that is NOT base64 encoded, then the transaction fails and the credit class is NOT created.

- GIVEN - user provides metadata that is NOT base64 encoded
- WHEN - user tries to create a credit class
- THEN - transaction fails, credit class is NOT created

If a user tries to create a credit class and the user provides metadata that is less than 256 bytes, then the transaction is successful and the credit class is created.

- GIVEN - user provides metadata that is less than 256 bytes
- WHEN - user tries to create a credit class
- THEN - transaction is successful, credit class is created

If a user tries to create a credit class and the user provides metadata that is equal to 256 bytes, then the transaction is successful and the credit class is created.

- GIVEN - user provides metadata that is equal to 256 bytes
- WHEN - user tries to create a credit class
- THEN - transaction is successful, credit class is created

If a user tries to create a credit class and the user provides metadata that is more than 256 bytes, then the transaction fails and the credit class is NOT created.

- GIVEN - user provides metadata that is more than 256 bytes
- WHEN - user tries to create a credit class
- THEN - transaction fails, credit class is NOT created

If a user tries to create a credit class and the user provides a valid issuer address, then the transaction is successful and the credit class is created.

- GIVEN - user provides a valid issuer address
- WHEN - user tries to create a credit class
- THEN - transaction successful, credit class is created

If a user tries to create a credit class and the user provides an invalid issuer address, then the transaction fails and the credit class is NOT created.

- GIVEN - user provides an invalid issuer address
- WHEN - user tries to create a credit class
- THEN - transaction fails, credit class is NOT created

If a user tries to create a credit class and the user account balance is equal to or more than the credit class fee, then the transaction is successful and the credit class is created.

- GIVEN - user account balance is equal to or more than the credit class fee
- WHEN - user tries to create a credit class
- THEN - transaction is successful, credit class is created

If a user tries to create a credit class and the user account balance is less than the credit class fee, then the transaction fails and the credit class is NOT created.

- GIVEN - user account balance is less than the credit class fee
- WHEN - user tries to create a credit class
- THEN - transaction fails, credit class is NOT created

### Create Project

If a user tries to create a project and the issuer is on the list of approved credit issuers for the given credit class, then the transaction is successful and the project is created.

- GIVEN - issuer is on the list of approved credit issuers for the given credit class
- WHEN - user tries to create a project
- THEN - transaction is successful, project is created

If a user tries to create a project and the issuer is NOT on the list of approved credit issuers for the given credit class, then the transaction fails and the project is NOT created.

- GIVEN - issuer is NOT on the list of approved credit issuers for the given credit class
- WHEN - user tries to create a project
- THEN - transaction fails, project is NOT created

If a user tries to create a project and the credit class is a valid credit class, then the transaction is successful and the project is created.

- GIVEN - credit class is a valid credit class
- WHEN - user tries to create a project
- THEN - transaction is successful, project is created

If a user tries to create a project and the credit class is NOT a valid credit class, then the transaction fails and the project is NOT created.

- GIVEN - credit class is NOT a valid credit class
- WHEN - user tries to create a project
- THEN - transaction fails, project is NOT created

If a user tries to create a project and the project location is a valid location, then the transaction is successful and the project is created.

- GIVEN - project location is a valid location
- WHEN - user tries to create a project
- THEN - transaction is successful, project is created

If a user tries to create a project and the project location is NOT a valid location, then the transaction fails and the project is NOT created.

- GIVEN - project location is NOT a valid location
- WHEN - user tries to create a project
- THEN - transaction fails, project is NOT created

If a user tries to create a project and the project id is a valid project id, then the transaction is successful and the project is created.

- GIVEN - project id is a valid project id
- WHEN - user tries to create a project
- THEN - transaction is successful, project is created

If a user tries to create a project and the project id is NOT a valid project id, then the transaction fails and the project is NOT created.

- GIVEN - project id is NOT a valid project id
- WHEN - user tries to create a project
- THEN - transaction fails, project is NOT created

If a user tries to create a project and the project id already exists, then the transaction fails and the project is NOT created.

- GIVEN - project id already exists
- WHEN - user tries to create a project
- THEN - transaction fails, project is NOT created

If a user tries to create a project and the user provides metadata that is base64 encoded, then the transaction is successful and the project is created.

- GIVEN - user provides metadata that is base64 encoded
- WHEN - user tries to create a project
- THEN - transaction is successful, project is created

If a user tries to create a project and the user provides metadata that is NOT base64 encoded, then the transaction fails and the project is NOT created.

- GIVEN - user provides metadata that is NOT base64 encoded
- WHEN - user tries to create a project
- THEN - transaction fails, project is NOT created

If a user tries to create a project and the user provides metadata that is less than 256 bytes, then the transaction is successful and the project is created.

- GIVEN - user provides metadata that is less than 256 bytes
- WHEN - user tries to create a project
- THEN - transaction is successful, project is created

If a user tries to create a project and the user provides metadata that is equal to 256 bytes, then the transaction is successful and the project is created.

- GIVEN - user provides metadata that is equal to 256 bytes
- WHEN - user tries to create a project
- THEN - transaction is successful, project is created

If a user tries to create a project and the user provides metadata that is more than 256 bytes, then the transaction fails and the project is NOT created.

- GIVEN - user provides metadata that is more than 256 bytes
- WHEN - user tries to create a project
- THEN - transaction fails, project is NOT created

### Create Credit Batch

If a user tries to create a credit batch and their account address is on the list of approved credit issuers for a credit class, then the transaction is successful and the credit batch is created.

- GIVEN - user is on list of approved credit issuers for a credit class
- WHEN - user tries to create a credit batch
- THEN - transaction is successful, credit batch is created

If a user tries to create a credit batch and their account address is NOT on the list of approved credit issuers for a credit class, then the transaction fails and the credit batch is NOT created.

- GIVEN - user is NOT on list of approved credit issuers for a credit class
- WHEN - user tries to create a credit batch
- THEN - transaction fails, credit batch is NOT created

If a user tries to create a credit batch and the user provides a valid recipient address, then the transaction is successful and the credit batch is created.

- GIVEN - user provides a valid recipient address
- WHEN - user tries to create a credit batch
- THEN - transaction is successful, credit batch is created

If a user tries to create a credit batch and the user provides an invalid recipient address, then the transaction fails and the credit batch is NOT created.

- GIVEN - user provides an invalid recipient address
- WHEN - user tries to create a credit batch
- THEN - transaction fails, credit batch is NOT created

If a user tries to create a credit batch and the user provides metadata that is less than 256 bytes, then the transaction is successful and the credit batch is created.

- GIVEN - user provides metadata that is less than 256 bytes
- WHEN - user tries to create a credit batch
- THEN - transaction is successful, credit batch is created

If a user tries to create a credit batch and the user provides metadata that is equal to 256 bytes, then the transaction is successful and the credit batch is created.

- GIVEN - user provides metadata that is equal to 256 bytes
- WHEN - user tries to create a credit batch
- THEN - transaction is successful, credit batch is created

If a user tries to create a credit batch and the user provides metadata that is more than 256 bytes, then the transaction fails and the credit batch is NOT created.

- GIVEN - user provides metadata that is more than 256 bytes
- WHEN - user tries to create a credit batch
- THEN - transaction fails, credit batch is NOT created

If a user tries to create a credit batch and the user provides a valid project, then the transaction is successful and the credit batch is created.

- GIVEN - user provides a valid project
- WHEN - user tries to create a credit batch
- THEN - transaction successful, credit batch is created

If a user tries to create a credit batch and the user provides an invalid project, then the transaction fails and the credit batch is NOT created.

- GIVEN - user provides an invalid project
- WHEN - user tries to create a credit batch
- THEN - transaction fails, credit batch is NOT created

If a user tries to create a credit batch and the user provides a valid start and end date, then the transaction is successful and the credit batch is created.

- GIVEN - user provides a valid start and end date
- WHEN - user tries to create a credit batch
- THEN - transaction is successful, credit batch is created

If a user tries to create a credit batch and the user provides an invalid start and end date, then the transaction fails and the credit batch is NOT created.

- GIVEN - user provides an invalid start and end date
- WHEN - user tries to create a credit batch
- THEN - transaction fails, credit batch is NOT created

If a user tries to create a credit batch and the user includes retired credits with a retirement location, then the transaction is successful and the credit batch is created.

- GIVEN - user includes retired credits with a retirement location
- WHEN - user tries to create a credit batch
- THEN - transaction is successful, credit batch is created

If a user tries to create a credit batch and the user includes retired credits without a retirement location, then the transaction fails and the credit batch is NOT created.

- GIVEN - user includes retired credits without a retirement location
- WHEN - user tries to create a credit batch
- THEN - transaction fails, credit batch is NOT created

### Transfer Credits

If a user tries to transfer 20 credits and their tradable balance is more than or equal to 20 credits, then the transaction is successful and the credits are transferred.

- GIVEN - tradable balance is more than or equal to 20 credits
- WHEN - user tries to transfer 20 credits
- THEN - transaction is successful, credits are transferred

If a user tries to transfer 20 credits and their tradable balance is less than 20 credits, then the transaction fails and the credits are NOT transferred.

- GIVEN - tradable balance is less than 20 credits
- WHEN - user tries to transfer 20 credits
- THEN - transaction fails, credits are NOT transferred

If a user tries to transfer credits and the user provides a valid recipient address, then the transaction is successful and the credits are transferred.

- GIVEN - user provides a valid recipient address
- WHEN - user tries to transfer credits
- THEN - transaction is successful, credits are transferred

If a user tries to transfer credits and the user provides an invalid recipient address, then the transaction fails and the credits are NOT transferred.

- GIVEN - user provides an invalid recipient address
- WHEN - user tries to transfer credits
- THEN - transaction fails, credits are NOT transferred

If a user tries to transfer credits and the user provides a valid batch denomination, then the transaction is successful and the credits are transferred.

- GIVEN - user provides a valid batch denomination
- WHEN - user tries to transfer credits
- THEN - transaction is successful, credits are transferred

If a user tries to transfer credits and the user provides an invalid batch denomination, then the transaction fails and the credits are NOT transferred.

- GIVEN - user provides an invalid batch denomination
- WHEN - user tries to transfer credits
- THEN - transaction fails, credits are NOT transferred

If a user tries to retire 20 credits upon transfer and the user provides a retirement location, then the transaction is successful and the credits are transferred and retired.

- GIVEN - user provides a retirement location
- WHEN - user tries to retire 20 credits upon transfer
- THEN - transaction is successful, credits are transferred and retired

If a user tries to retire 20 credits upon transfer and the user does NOT provide a retirement location, then the transaction fails and the credits are NOT transferred or retired.

- GIVEN - user provides a retirement location
- WHEN - user tries to retire 20 credits upon transfer
- THEN - transaction fails, credits are NOT transferred or retired

### Retire Credits

If a user tries to retire 20 credits and their tradable balance is more than 20 credits, then the transaction is successful and the credits are retired.

- GIVEN - tradable balance is more than 20 credits
- WHEN - user tries to retire 20 credits
- THEN - transaction is successful, credits are retired

If a user tries to retire 20 credits and their tradable balance is less than 20 credits, then the transaction fails and the credits are NOT retired.

- GIVEN - tradable balance is less than 20 credits
- WHEN - user tries to retire 20 credits
- THEN - transaction fails, credits are NOT retired

If a user tries to retire credits and the user provides a valid batch denomination, then the transaction is successful and the credits are retired.

- GIVEN - user provides a valid batch denomination
- WHEN - user tries to retire credits
- THEN - transaction is successful, credits are retired

If a user tries to retire credits and the user provides an invalid batch denomination, then the transaction fails and the credits are NOT retired.

- GIVEN - user provides an invalid batch denomination
- WHEN - user tries to retire credits
- THEN - transaction fails, credits are NOT retired

### Cancel Credits

If a user tries to cancel 20 credits and their credit balance is more than or equal to 20 credits, then the transaction is successful and the credits are cancelled.

- GIVEN - credit balance is more than 20 credits
- WHEN - user tries to cancel 20 credits
- THEN - transaction is successful, credits are cancelled

If a user tries to cancel 20 credits and their credit balance is less than 20 credits, then the transaction fails and the credits are NOT cancelled.

- GIVEN - credit balance is less than 20 credits
- WHEN - user tries to cancel 20 credits
- THEN - transaction fails, credits are NOT cancelled

If a user tries to cancel credits and the user provides a valid batch denomination, then the transaction is successful and the credits are cancelled.

- GIVEN - user provides a valid batch denomination
- WHEN - user tries to cancel credits
- THEN - transaction is successful, credits are cancelled

If a user tries to cancel credits and the user provides an invalid batch denomination, then the transaction fails and the credits are NOT cancelled.

- GIVEN - user provides an invalid batch denomination
- WHEN - user tries to cancel credits
- THEN - transaction fails, credits are NOT cancelled

## Update Sell Order

If a user tries to update a sell order and the user provides a sell order id that does exist, then the transaction is successful and the sell order is updated.

- GIVEN - user provides a sell order id that does exist
- WHEN - user tries to update a sell order
- THEN - transaction is successful, sell order is updated

If a user tries to update a sell order and the user provides a sell order id that does NOT exist, then the transaction fails and the sell order is NOT updated.

- GIVEN - user provides a sell order id that does NOT exist
- WHEN - user tries to update a sell order
- THEN - transaction fails, sell order is NOT updated

If a user tries to update a sell order and the user provides a sell order id that they are the owner of, then the transaction is successful and the sell order is updated.

- GIVEN - user provides a sell order id that they are the owner of
- WHEN - user tries to update a sell order
- THEN - transaction is successful, sell order is updated

If a user tries to update a sell order and the user provides a sell order id that they are not the owner of, then the transaction fails and the sell order is NOT updated.

- GIVEN - user provides a sell order id that they are not the owner of
- WHEN - user tries to update a sell order
- THEN - transaction fails, sell order is NOT updated

If a user tries to update a sell order and the user provides a new quantity that is less than or equal to their credit balance, then the transaction is successful and the sell order is updated.

- GIVEN - user provides a new quantity that is less than or equal to their credit balance
- WHEN - user tries to update a sell order
- THEN - transaction is successful, sell order is updated

If a user tries to update a sell order and the user provides a new quantity that is more than their credit balance, then the transaction fails and the sell order is NOT updated.

- GIVEN - user provides a new quantity that is more than their credit balance
- WHEN - user tries to update a sell order
- THEN - transaction fails, sell order is NOT updated

If a user tries to update a sell order and the user provides a new ask price using an ask denom that is allowed, then the transaction is successful and the sell order is updated.

- GIVEN - user provides a new ask price using an ask denom that is allowed
- WHEN - user tries to update a sell order
- THEN - transaction is successful, sell order is updated

If a user tries to update a sell order and the user provides a new ask price using an ask denom that is NOT allowed, then the transaction fails and the sell order is NOT updated.

- GIVEN - user provides a new ask price using an ask denom that is NOT allowed
- WHEN - user tries to update a sell order
- THEN - transaction fails, sell order is NOT updated

## Allow Ask Denom

If a user tries to add an ask denom and the user submits a governance proposal to execute the AllowAskDenom message that is approved, then the transaction is successful and the ask denom is added.

- GIVEN - user submits a governance proposal to execute the AllowAskDenom message that is approved
- WHEN - user tries to add an ask denom
- THEN - transaction is successful, ask denom is added

If a user tries to add an ask denom and the user submits a governance proposal to execute the AllowAskDenom message that is NOT approved, then the transaction fails and the ask denom is NOT added.

- GIVEN - user submits a governance proposal to execute the AllowAskDenom message that is NOT approved
- WHEN - user tries to add an ask denom
- THEN - transaction fails, ask denom is NOT added

If a user tries to add an ask denom and the user submits a transaction with the AllowAskDenom message directly from their account, then the transaction fails and the ask denom is NOT added.

- GIVEN - user tries to add an ask denom and the user submits a transaction with the AllowAskDenom message directly from their account
- WHEN - user tries to add an ask denom
- THEN - transaction fails, ask denom is NOT added