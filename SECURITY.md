# Security

**Do not open up a public GitHub issue when reporting a security vulnerability. Please report security vulnerabilities by email from a secure email address to security@regen.network.**

Regen Ledger is built on top of Cosmos SDK and Tendermint Core. Please refer to the security policy for each of these projects if the security vulnerability is not specific to Regen Ledger:

- [Cosmos SDK Security Policy](https://github.com/cosmos/cosmos-sdk/blob/main/SECURITY.md)
- [Tendermint Core Security Policy](https://github.com/tendermint/tendermint/blob/master/SECURITY.md)

Significant security vulnerabilities within Regen Ledger are most likely to occur in the following packages:

- [`/app`](https://github.com/regen-network/regen-ledger/tree/master/app)
- [`/types`](https://github.com/regen-network/regen-ledger/tree/master/types)
- [`x/data`](https://github.com/regen-network/regen-ledger/tree/master/x/data)
- [`x/ecocredit`](https://github.com/regen-network/regen-ledger/tree/master/x/ecocredit)

## Disclosure Process

The Regen Ledger team uses the following disclosure process:

1. After a security report is received, the Regen Ledger team works to verify the issue and confirm its severity level using Common Vulnerability Scoring System (CVSS).
1. The Regen Ledger team collaborates with the Cosmos SDK and Tendermint teams to determine the vulnerability’s potential impact on the other applications and services in the Cosmos ecosystem.
1. Patches are prepared in private repositories for eligible releases. See [Release Process](https://github.com/regen-network/regen-ledger/blob/master/RELEASE_PROCESS.md) for more information about eligible releases.
1. If it is determined that a CVE-ID is required, we request a CVE through a CVE Numbering Authority.
1. We notify the community that a security release is coming to give users time to prepare their systems for the update. Notifications can include forum posts, tweets, and emails to partner projects and validators.
1. 24 hours after the notification, fixes are applied publicly and the new release is issued.
1. After a release is available, we notify the community again through the same channels. We also publish a Security Advisory on Github and publish the CVE, as long as the Security Advisory and the CVE do not include information on how to exploit these vulnerabilities beyond the information that is available in the patch.
1. One week after the release goes out, we publish a postmortem with details about and our response to the vulnerability.

This process can take some time. Every possible effort is made to handle the security vulnerability in a timely a manner. However, it's important that we follow this security process to ensure that disclosures are handled consistently and to keep Regen Ledger and its partner projects as secure as possible.

### Disclosure Communications

Communications to partners usually include the following details:

1. Affected version or versions
1. New release version
1. Impact on user funds
1. For timed releases, a date and time that the new release will be made available
1. Impact on the partners if upgrades are not completed in a timely manner
1. Potential required actions if an adverse condition arises during the security release process

An example notice looks like the following:

```text
Dear Regen Ledger partners,

A critical security vulnerability has been identified in Regen Ledger vX.X.X.

User funds are NOT at risk; however, the vulnerability can result in a chain halt.

This notice is to inform you that on [[**March 1 at 1pm EST/6pm UTC**]], we will be releasing Regen Ledger vX.X.Y to fix the security vulnerability.

We ask all validators to upgrade their nodes ASAP.

If the chain halts, validators with sufficient voting power must upgrade and come online for the chain to resume.
```

### Example Timeline

The following timeline is an example of triage and response. Each task identifies the required roles and team members; however, multiple people may play each role and each person may play multiple roles.

#### 24+ Hours Before Release Time

1. Request CVE number (ADMIN)
1. Gather emails and other contact info for validators (COMMS LEAD)
1. Test fixes on a testnet  (REGEN LEDGER ENG)
1. Write “Security Advisory” for forum (REGEN LEDGER LEAD)

#### 24 Hours Before Release Time

1. Post “Security Advisory” pre-notification on forum (REGEN LEDGER LEAD)
1. Post Tweet linking to forum post (COMMS LEAD)
1. Announce security advisory/link to post in various other social channels (Telegram, Discord) (COMMS LEAD)
1. Send emails to partner projects or other users (PARTNERSHIPS LEAD)

#### Release Time

1. Cut Regen Ledger release for eligible version (REGEN LEDGER ENG)
1. Post “Security release” on forum (REGEN LEDGER LEAD)
1. Post new Tweet linking to forum post (COMMS LEAD)
1. Remind everyone on social channels (Telegram, Discord) that the release is out (COMMS LEAD)
1. Send emails to validators and other users (COMMS LEAD)
1. Publish Security Advisory and CVE if the CVE has no sensitive information (ADMIN)

#### After Release Time

1. Write forum post with exploit details (REGEN LEDGER LEAD)

#### 7 Days After Release Time

1. Publish CVE if it has not yet been published (ADMIN)
1. Publish forum post with exploit details (REGEN LEDGER ENG, REGEN LEDGER LEAD)
