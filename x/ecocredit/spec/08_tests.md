# Tests

This document includes acceptance tests for the ecocredit module.

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

