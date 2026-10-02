# Test certificates and EUDI profiles

The generated certificates identify test services. Their organization names, addresses and registration numbers are fictional. Keys are held in software. The access certificate policy describes the test profile and carries no certification claim.

## Specification versions

Checked on 1 October 2026 against [ARF v3.0.0](https://github.com/eu-digital-identity-wallet/eudi-doc-architecture-and-reference-framework/releases/tag/v3.0.0), [CIR (EU) 2026/1731](https://eur-lex.europa.eu/legal-content/EN/TXT/PDF/?uri=CELEX:32026R1731) and the Commission's [standards tracker](https://github.com/eu-digital-identity-wallet/eudi-doc-standards-and-technical-specifications). The regulation's versions and adaptations take precedence over newer standalone specifications.

| Area | Version used | Source |
| --- | --- | --- |
| Issuance | OpenID4VCI 1.0 Final | [Specification](https://openid.net/specs/openid-4-verifiable-credential-issuance-1_0.html) |
| Presentation | OpenID4VP 1.0 Final | [Specification](https://openid.net/specs/openid-4-verifiable-presentations-1_0.html) |
| High assurance profile | HAIP 1.0 Final | [Specification](https://openid.net/specs/openid4vc-high-assurance-interoperability-profile-1_0.html) |
| PID and wallet provider certificates | ETSI TS 119 412-6 V1.1.1, September 2025 | [Specification](https://www.etsi.org/deliver/etsi_ts/119400_119499/11941206/01.01.01_60/ts_11941206v010101p.pdf) |
| Certificate issuer names | ETSI EN 319 412-2 V2.4.1, June 2025 | [Specification](https://www.etsi.org/deliver/etsi_en/319400_319499/31941202/02.04.01_60/en_31941202v020401p.pdf) |
| Legal person subjects | ETSI EN 319 412-3 V1.3.1, September 2023 | [Specification](https://www.etsi.org/deliver/etsi_en/319400_319499/31941203/01.03.01_60/en_31941203v010301p.pdf) |
| QCStatements | ETSI EN 319 412-5 V2.5.1, June 2025 | [Specification](https://www.etsi.org/deliver/etsi_en/319400_319499/31941205/02.05.01_60/en_31941205v020501p.pdf) |
| Access certificate policy | ETSI TS 119 411-8 V1.1.1, October 2025 | [Specification](https://www.etsi.org/deliver/etsi_ts/119400_119499/11941108/01.01.01_60/ts_11941108v010101p.pdf) |
| EUDI issuance profile | ETSI TS 119 472-3 V1.1.1, March 2026 | [Specification](https://www.etsi.org/deliver/etsi_ts/119400_119499/11947203/01.01.01_60/ts_11947203v010101p.pdf) |
| EUDI presentation profile | ETSI TS 119 472-2 V1.2.1, March 2026 | [Specification](https://www.etsi.org/deliver/etsi_ts/119400_119499/11947202/01.02.01_60/ts_11947202v010201p.pdf) |
| Attestation structure | ETSI TS 119 472-1 V1.2.1, February 2026 | [Specification](https://www.etsi.org/deliver/etsi_ts/119400_119499/11947201/01.02.01_60/ts_11947201v010201p.pdf) |
| Registration information | ETSI TS 119 475 V1.2.1, March 2026 | [Specification](https://www.etsi.org/deliver/etsi_ts/119400_119499/119475/01.02.01_60/ts_119475v010201p.pdf) |
| Trusted entity lists | ETSI TS 119 602 V1.1.1, November 2025 | [Specification](https://www.etsi.org/deliver/etsi_TS/119600_119699/119602/01.01.01_60/ts_119602v010101p.pdf) |
| Trust list signatures | ETSI TS 119 182-1 V1.2.1, July 2024 | [Specification](https://www.etsi.org/deliver/etsi_ts/119100_119199/11918201/01.02.01_60/ts_11918201v010201p.pdf) |
| EU PID attributes | PID Rulebook v1.7 | [Rulebook](https://github.com/eu-digital-identity-wallet/eudi-doc-attestation-rulebooks-catalog/blob/6d8f7f8422e5bf6c48186005b6835c078f762a67/rulebooks/pid/pid-rulebook.md) |
| German PID attributes | German PID Rulebook 1.0.0 consultation draft | [Rulebook](https://bmi.usercontent.opencode.de/eudi-wallet/eidas-2.0-architekturkonzept/content/features/PID/german-pid-rulebook/) |

ETSI has also published TS 119 412-6 V1.2.1, EN 319 412-3 V1.4.1 and EN 319 412-5 V2.6.1. The table retains the versions referenced by the regulation. TS 119 412-6 V1.2.1 replaces the explicit intermediate CA retrieval rule with a reference to EN 319 412-2. The generated intermediate path supports both certificate profiles.

OpenID4VCI 1.1 remains an optional draft feature level. It does not replace the 1.0 baseline. The German PID rulebook remains a consultation draft, so its national attributes may change.

## Signing roles

| Role | Material |
| --- | --- |
| PID issuance | Credential key, PID provider leaf with QcType `0.4.0.194126.1.1` |
| Wallet and key attestations | Separate wallet provider key and leaf with QcType `0.4.0.194126.1.2` |
| Signed issuer metadata and verifier requests | Separate access key and certificate with policy `0.4.0.194118.1.2` |
| Registrar responses and registration certificates | Separate registrar key and signing certificate |
| Credential status | Separate status key and signing certificate |
| Trust lists | Separate list operator key and signing certificate |

New wallets create a root CA and provider intermediates. Document signer leaves retain the ISO/IEC 18013-5 document signing purpose. The subject country matches the credential's `issuing_country`. AIA and CRL URLs identify the provider intermediate and its revocation list. Certificates are retained in the selected storage backend. A different subject, changed issuer URL or renewal receives a new serial number.

This hierarchy follows TS 119 412-6 V1.1.1 clause 4.4.3. The PID Rulebook's trust anchors are notified provider keys. The root permits one intermediate. It therefore differs from the direct IACA hierarchy in ISO/IEC 18013-5:2021 Annex B, whose root requires a path length of zero. The OpenID suite reports that difference as an ISO profile warning. Certificate signatures and trust paths are checked separately.

PID signatures include protected certificate references required by CIR (EU) 2026/1731 Annex I. SD-JWT uses `x5u` and `x5t#S256`. Mdoc uses `x5u` and SHA-256 `x5t`. Their URLs contain the certificate fingerprint and return PEM for JOSE or DER for COSE. The protected `iat` records signing time independently of the credential's issuance time. Published certificates remain available after renewal. Offline issuance has no certificate hosting endpoint.

Existing stored CAs keep their keys and chains. CAs with a path length of zero continue to sign leaves directly. Use a fresh wallet directory to exercise V1.1.1's intermediate CA retrieval rule. Memory storage retains certificates for the lifetime of that store. A seed retains keys across process restarts, while certificates receive fresh serial numbers.

## Discovery and trust lists

Both issuer discovery endpoints serve JSON by default and signed metadata for `Accept: application/jwt`. The signed form includes the access certificate in protected `x5c`. The `issuer_info` array contains registrar data and a registration certificate signed by the test registrar. Registration certificates use the identifier, legal name and country from the access certificate.

Trust lists publish issuance certificates, their provider CAs and status signing certificates. This keeps credentials verifiable across country overrides and certificate renewal. Protected `iat` and `x5t#S256` headers carry the signing time and certificate reference required by JAdES. They use English language code `en`, whole second UTC timestamps, postal addresses and a self pointer. An unchanged list keeps its signed instance until it expires. Changed content or expiry advances the sequence number. Append `/history` to a trust list URL to list its retained instances, then `/history/<sequence>` to retrieve one.

The schema is ETSI's [published JSON binding](https://forge.etsi.org/rep/esi/x19_60201_lists_of_trusted_entities), revision `e84f427f0cde99513b574ef4b5a155ac4a38eab6` from 13 November 2025. The PID and wallet provider lists follow Annexes D and E. Their fictional provider entries are for local interoperability tests.

## Public PID provider comparison

The Bundesdruckerei [demo](https://demo.pid-provider.bundesdruckerei.de/) and [preproduction](https://preprod.pid-provider.bundesdruckerei.de/) deployments publish separate credential, status and access certificate material. Their PID paths use P-521 CAs and P-256 signing leaves. Signed issuer metadata uses an access certificate.

EUDI Dev exercises those roles with P-256 keys. Algorithm choices follow the applicable profile. Public documentation samples may contain expired certificates or older attribute names. The versioned rulebooks determine generated data types and names.

## Test scope

The toolkit tests protocol exchanges, signatures, certificate structure and generated data. Registration certificates simulate provider registration. Their status and revocation lifecycle is not implemented. Official trust, certified hardware protection and physical presence checks require the corresponding ecosystem services. Configured key attestation assurance values simulate a test scenario. See [spec compliance](spec-compliance.md) and [conformance results](conformance-results.md) for implemented checks and remaining protocol limits.
