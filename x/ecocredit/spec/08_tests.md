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

If a user tries to create a credit class and the user provides a credit type that exists, then the transaction is successful and the credit class is created.

- GIVEN - user provides a valid credit type
- WHEN - user tries to create a credit class
- THEN - transaction is successful, credit class is created

If a user tries to create a credit class and the user provide a credit type that does not exist, then the transaction fails and the credit class is NOT created.

- GIVEN - user provides an invalid credit type
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

If a user tries to create a project and the credit class exists, then the transaction is successful and the project is created.

- GIVEN - credit class exists
- WHEN - user tries to create a project
- THEN - transaction is successful, project is created

If a user tries to create a project and the credit class does NOT exist, then the transaction fails and the project is NOT created.

- GIVEN - credit class does NOT exist
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

If a user tries to create a credit batch and the user provides a project that exists, then the transaction is successful and the credit batch is created.

- GIVEN - user provides a valid project
- WHEN - user tries to create a credit batch
- THEN - transaction successful, credit batch is created

If a user tries to create a credit batch and the user provides a project that does not exist, then the transaction fails and the credit batch is NOT created.

- GIVEN - user provides an invalid project
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

If a user tries to transfer credits and the user provides a batch denom that exists, then the transaction is successful and the credits are transferred.

- GIVEN - user provides a batch denom that exists
- WHEN - user tries to transfer credits
- THEN - transaction is successful, credits are transferred

If a user tries to transfer credits and the user provides a batch denom that does not exist, then the transaction fails and the credits are NOT transferred.

- GIVEN - user provides a batch denom that does not exist
- WHEN - user tries to transfer credits
- THEN - transaction fails, credits are NOT transferred

### Retire Credits

If a user tries to retire 20 credits and their tradable balance is more than 20 credits, then the transaction is successful and the credits are retired.

- GIVEN - tradable balance is more than 20 credits
- WHEN - user tries to retire 20 credits
- THEN - transaction is successful, credits are retired

If a user tries to retire 20 credits and their tradable balance is less than 20 credits, then the transaction fails and the credits are NOT retired.

- GIVEN - tradable balance is less than 20 credits
- WHEN - user tries to retire 20 credits
- THEN - transaction fails, credits are NOT retired

If a user tries to retire credits and the user provides a batch denom that exists, then the transaction is successful and the credits are retired.

- GIVEN - user provides a batch denom that exists
- WHEN - user tries to retire credits
- THEN - transaction is successful, credits are retired

If a user tries to retire credits and the user provides a batch denom that does not exist, then the transaction fails and the credits are NOT retired.

- GIVEN - user provides a batch denom that does not exist
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

If a user tries to cancel credits and the user provides a batch denom that exists, then the transaction is successful and the credits are cancelled.

- GIVEN - user provides a batch denom that exists
- WHEN - user tries to cancel credits
- THEN - transaction is successful, credits are cancelled

If a user tries to cancel credits and the user provides a batch denom that does not exist, then the transaction fails and the credits are NOT cancelled.

- GIVEN - user provides a batch denom that does not exist
- WHEN - user tries to cancel credits
- THEN - transaction fails, credits are NOT cancelled
