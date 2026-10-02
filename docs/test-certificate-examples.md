# Test certificate examples

This public reference set contains the root CA, provider intermediates and signing certificates described in [test certificates](test-certificates.md#certificate-contents). Each entry includes the complete PEM certificate and decoded X.509 contents, including serial number, validity, public key, extensions and signature.

The certificates are generated through the wallet signing APIs at source revision [`44ef415d1ae3`](https://github.com/dominikschlosser/eudi-dev/tree/44ef415d1ae36e69551bf2a2c9de7f850d2d48ba). The reference issuer is `https://eudi-test.dev`, the country is `NL`, and the trust list operator is `EUDI Dev Wallet`. The names and registration identifiers are fictional. The public certificates carry no official trust.

The decoded values below belong to these reference PEM files. A running wallet, including the public demo, has its own keys, serial numbers, timestamps and signatures. Certificate profiles and configurable values are described in the [profile tables](test-certificates.md#certificate-contents). The [localhost special case](test-certificates.md#localhost-special-case) gives the local issuer URL and corresponding certificate URLs.

From the repository root, inspect a certificate with:

```bash
openssl x509 -in docs/assets/test-certificates/pid-signer.pem -noout -text
openssl asn1parse -in docs/assets/test-certificates/pid-signer.pem -i
```

The QCStatements payloads are decoded separately below because OpenSSL displays their extension values as binary text in its X.509 output.

## Root CA

[Complete PEM certificate](assets/test-certificates/root-ca.pem)

<details>
<summary>Full decoded certificate contents</summary>

```text
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            e4:31:53:0c:6b:08:a2:06:3b:f2:9b:70:78:9b:5e:e5
        Signature Algorithm: ecdsa-with-SHA256
        Issuer: C=NL, O=EUDI Dev Test CA, CN=OID4VC Dev Wallet CA
        Validity
            Not Before: Oct  2 16:21:20 2026 GMT
            Not After : Sep 29 17:21:20 2036 GMT
        Subject: C=NL, O=EUDI Dev Test CA, CN=OID4VC Dev Wallet CA
        Subject Public Key Info:
            Public Key Algorithm: id-ecPublicKey
                Public-Key: (256 bit)
                pub:
                    04:f9:52:4e:0e:55:9f:07:53:df:94:7d:0a:38:8f:
                    8e:0b:cb:4f:38:d0:34:c0:fc:35:b6:05:88:d5:f7:
                    eb:5d:64:2f:20:69:26:18:09:b6:4a:0d:0b:61:14:
                    79:ad:c3:e9:ca:e4:d3:f5:7b:94:bf:59:20:70:be:
                    06:ff:10:26:4f
                ASN1 OID: prime256v1
                NIST CURVE: P-256
        X509v3 extensions:
            X509v3 Key Usage: critical
                Certificate Sign, CRL Sign
            X509v3 Basic Constraints: critical
                CA:TRUE, pathlen:1
            X509v3 Subject Key Identifier:
                B2:10:CA:15:53:9E:5E:4E:73:EF:5D:67:67:38:26:59:EB:DA:90:52
            X509v3 Issuer Alternative Name:
                URI:https://github.com/dominikschlosser/eudi-dev
    Signature Algorithm: ecdsa-with-SHA256
    Signature Value:
        30:44:02:20:74:76:d2:9c:57:d3:38:7c:9b:9d:8c:39:6a:20:
        cf:01:5c:01:e0:c5:88:7e:f5:26:40:f6:ec:a4:07:f3:3e:16:
        02:20:14:1f:be:5f:4f:3b:b8:78:52:0b:24:9c:45:c0:a8:10:
        f8:8c:16:c4:35:e9:8b:9a:99:e4:a8:a3:82:61:d7:c9
```

</details>

## PID provider CA

[Complete PEM certificate](assets/test-certificates/pid-ca.pem)

<details>
<summary>Full decoded certificate contents</summary>

```text
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            b6:3c:6d:8d:f5:41:7a:d8:b5:88:fd:34:48:57:11:01
        Signature Algorithm: ecdsa-with-SHA256
        Issuer: C=NL, O=EUDI Dev Test CA, CN=OID4VC Dev Wallet CA
        Validity
            Not Before: Oct  2 16:21:20 2026 GMT
            Not After : Oct  1 17:21:20 2031 GMT
        Subject: C=NL, O=EUDI Dev Test Provider, CN=EUDI Dev Test pid CA NL, organizationIdentifier=NTRNL-00000000
        Subject Public Key Info:
            Public Key Algorithm: id-ecPublicKey
                Public-Key: (256 bit)
                pub:
                    04:f8:b0:75:1e:61:16:33:cd:1e:1b:8a:7c:eb:d6:
                    6a:99:87:04:76:d0:bd:cc:32:0b:ae:83:46:5d:c3:
                    83:71:1f:fc:19:59:80:53:b4:ac:f2:b9:00:06:4d:
                    ed:50:de:d4:31:6e:12:e8:28:53:7f:cb:bc:38:55:
                    7b:f0:0c:48:d8
                ASN1 OID: prime256v1
                NIST CURVE: P-256
        X509v3 extensions:
            X509v3 Key Usage: critical
                Certificate Sign, CRL Sign
            X509v3 Basic Constraints: critical
                CA:TRUE, pathlen:0
            X509v3 Subject Key Identifier:
                B7:7C:03:F6:04:9B:16:F6:7D:33:E7:5F:9F:12:89:6C:AA:99:7A:21
            X509v3 Authority Key Identifier:
                B2:10:CA:15:53:9E:5E:4E:73:EF:5D:67:67:38:26:59:EB:DA:90:52
            Authority Information Access:
                CA Issuers - URI:https://eudi-test.dev/api/certificates/ca.der
            X509v3 CRL Distribution Points:
                Full Name:
                  URI:https://eudi-test.dev/api/crl

            X509v3 Issuer Alternative Name:
                URI:https://github.com/dominikschlosser/eudi-dev
    Signature Algorithm: ecdsa-with-SHA256
    Signature Value:
        30:45:02:21:00:8a:fc:23:a4:84:e1:63:2b:92:00:59:1c:74:
        d7:c5:d8:25:88:e9:79:04:49:82:98:66:46:ea:ee:75:4c:b9:
        62:02:20:2f:4c:16:43:29:f9:78:6e:e0:a4:38:00:ba:3c:97:
        65:ab:38:c2:1d:00:c3:8d:ca:64:f9:2c:c9:08:43:8f:a1
```

</details>

## Wallet provider CA

[Complete PEM certificate](assets/test-certificates/wallet-provider-ca.pem)

<details>
<summary>Full decoded certificate contents</summary>

```text
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            e1:ef:a2:48:0a:91:5e:ef:14:ec:50:79:ca:57:ba:64
        Signature Algorithm: ecdsa-with-SHA256
        Issuer: C=NL, O=EUDI Dev Test CA, CN=OID4VC Dev Wallet CA
        Validity
            Not Before: Oct  2 16:21:20 2026 GMT
            Not After : Oct  1 17:21:20 2031 GMT
        Subject: C=NL, O=EUDI Dev Test Provider, CN=EUDI Dev Test wallet CA NL, organizationIdentifier=NTRNL-00000000
        Subject Public Key Info:
            Public Key Algorithm: id-ecPublicKey
                Public-Key: (256 bit)
                pub:
                    04:6a:0b:e9:66:76:52:f6:c2:60:85:b7:c5:4b:d2:
                    59:96:97:22:74:1b:81:df:fe:23:4f:a1:d0:c7:6f:
                    cb:0a:aa:b2:26:80:4b:f9:6c:98:c9:4e:76:3b:39:
                    30:f7:a5:60:ce:97:03:9b:9d:7c:9c:04:e1:7d:74:
                    aa:7d:49:a7:07
                ASN1 OID: prime256v1
                NIST CURVE: P-256
        X509v3 extensions:
            X509v3 Key Usage: critical
                Certificate Sign, CRL Sign
            X509v3 Basic Constraints: critical
                CA:TRUE, pathlen:0
            X509v3 Subject Key Identifier:
                6A:B1:1C:F3:DF:76:DE:84:AC:42:6E:2E:74:2A:ED:66:8C:E9:4B:26
            X509v3 Authority Key Identifier:
                B2:10:CA:15:53:9E:5E:4E:73:EF:5D:67:67:38:26:59:EB:DA:90:52
            Authority Information Access:
                CA Issuers - URI:https://eudi-test.dev/api/certificates/ca.der
            X509v3 CRL Distribution Points:
                Full Name:
                  URI:https://eudi-test.dev/api/crl

            X509v3 Issuer Alternative Name:
                URI:https://github.com/dominikschlosser/eudi-dev
    Signature Algorithm: ecdsa-with-SHA256
    Signature Value:
        30:45:02:21:00:a6:c5:65:0f:75:5a:6e:07:96:70:05:63:26:
        d0:b1:6a:80:fc:e8:86:5c:df:53:c1:97:78:10:3c:f1:91:58:
        d9:02:20:2c:b0:b3:cf:f6:ab:58:aa:05:14:6b:36:7c:4b:02:
        a7:5f:e8:70:8c:d5:e2:f5:ac:de:d5:1a:a1:8a:41:c8:fe
```

</details>

## Local credential provider CA

[Complete PEM certificate](assets/test-certificates/local-ca.pem)

<details>
<summary>Full decoded certificate contents</summary>

```text
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            b9:0d:ce:b5:a0:88:f0:c9:d7:c5:85:50:48:1e:a5:3d
        Signature Algorithm: ecdsa-with-SHA256
        Issuer: C=NL, O=EUDI Dev Test CA, CN=OID4VC Dev Wallet CA
        Validity
            Not Before: Oct  2 16:21:20 2026 GMT
            Not After : Oct  1 17:21:20 2031 GMT
        Subject: C=NL, O=EUDI Dev Test Provider, CN=EUDI Dev Test local CA NL, organizationIdentifier=NTRNL-00000000
        Subject Public Key Info:
            Public Key Algorithm: id-ecPublicKey
                Public-Key: (256 bit)
                pub:
                    04:fc:26:a7:18:b5:bb:3c:2a:24:6f:7b:c7:07:1f:
                    3a:b3:c6:8b:e9:19:a7:fd:e1:67:c4:90:2f:e0:07:
                    ce:93:e2:6a:c8:43:9e:ba:87:73:e6:26:a5:1d:f4:
                    c4:b3:38:b2:99:07:c1:e7:32:b4:9c:15:67:06:fb:
                    cd:f3:c8:05:a6
                ASN1 OID: prime256v1
                NIST CURVE: P-256
        X509v3 extensions:
            X509v3 Key Usage: critical
                Certificate Sign, CRL Sign
            X509v3 Basic Constraints: critical
                CA:TRUE, pathlen:0
            X509v3 Subject Key Identifier:
                FF:52:94:1C:83:CB:4D:B5:D1:07:65:7A:F7:86:02:52:3F:EE:51:18
            X509v3 Authority Key Identifier:
                B2:10:CA:15:53:9E:5E:4E:73:EF:5D:67:67:38:26:59:EB:DA:90:52
            Authority Information Access:
                CA Issuers - URI:https://eudi-test.dev/api/certificates/ca.der
            X509v3 CRL Distribution Points:
                Full Name:
                  URI:https://eudi-test.dev/api/crl

            X509v3 Issuer Alternative Name:
                URI:https://github.com/dominikschlosser/eudi-dev
    Signature Algorithm: ecdsa-with-SHA256
    Signature Value:
        30:45:02:20:06:08:22:73:4a:14:ce:9d:aa:82:74:f2:11:26:
        19:6e:6b:36:b9:0e:ae:79:f5:3d:45:79:b3:13:ab:94:62:be:
        02:21:00:a7:91:a1:2b:21:f1:fd:78:27:7b:bd:5b:d5:f5:79:
        40:52:2d:93:5a:89:a6:df:31:d5:00:c7:07:2e:6d:17:58
```

</details>

## PID signer

[Complete PEM certificate](assets/test-certificates/pid-signer.pem)

<details>
<summary>Full decoded certificate contents</summary>

```text
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            e9:74:7b:65:33:73:b1:58:77:30:6d:9c:01:24:8f:48
        Signature Algorithm: ecdsa-with-SHA256
        Issuer: C=NL, O=EUDI Dev Test Provider, CN=EUDI Dev Test pid CA NL, organizationIdentifier=NTRNL-00000000
        Validity
            Not Before: Oct  2 16:21:20 2026 GMT
            Not After : Oct  2 17:21:20 2027 GMT
        Subject: C=NL, O=EUDI Dev Test Provider, CN=EUDI Dev Wallet PID Provider (pid), organizationIdentifier=NTRNL-00000000
        Subject Public Key Info:
            Public Key Algorithm: id-ecPublicKey
                Public-Key: (256 bit)
                pub:
                    04:76:a0:5b:ee:29:b4:a4:b2:f0:f3:83:5c:33:c8:
                    b2:fe:2a:56:08:f0:dd:95:63:38:b9:09:f5:28:23:
                    60:16:03:2c:71:f8:ce:a3:e9:51:c8:2b:5d:77:88:
                    4d:e4:23:01:bf:cf:93:fa:ba:e2:94:b1:07:97:ba:
                    a4:6d:d5:d8:5a
                ASN1 OID: prime256v1
                NIST CURVE: P-256
        X509v3 extensions:
            X509v3 Key Usage: critical
                Digital Signature
            X509v3 Subject Key Identifier:
                DE:DE:D0:97:4C:24:23:E2:F7:B1:6D:D5:C4:9D:E2:99:F1:4E:06:50
            X509v3 Authority Key Identifier:
                B7:7C:03:F6:04:9B:16:F6:7D:33:E7:5F:9F:12:89:6C:AA:99:7A:21
            Authority Information Access:
                CA Issuers - URI:https://eudi-test.dev/api/certificates/providers/pid/NL.der
            X509v3 Subject Alternative Name:
                DNS:eudi-test.dev, URI:https://eudi-test.dev
            X509v3 CRL Distribution Points:
                Full Name:
                  URI:https://eudi-test.dev/api/crl/providers/pid/NL

            X509v3 Issuer Alternative Name:
                URI:https://github.com/dominikschlosser/eudi-dev
            X509v3 Extended Key Usage: critical
                1.0.18013.5.1.2
            qcStatements:
                0.0......F..0.......N..
    Signature Algorithm: ecdsa-with-SHA256
    Signature Value:
        30:44:02:20:70:94:7f:6c:f9:d5:56:1d:fe:07:49:fc:94:ec:
        5c:6b:aa:c2:54:e8:c3:9b:79:1b:48:75:f6:cf:ee:66:3f:1d:
        02:20:7e:c1:76:06:26:4b:79:8e:ff:26:45:f6:02:84:d6:c4:
        18:97:90:ee:6e:37:b9:86:29:3e:30:06:0a:25:95:c1
```

QCStatements extension value:

```text
DER: 30153013060604008e4601063009060704008bec4e0101
    0:d=0  hl=2 l=  21 cons: SEQUENCE
    2:d=1  hl=2 l=  19 cons:  SEQUENCE
    4:d=2  hl=2 l=   6 prim:   OBJECT            :0.4.0.1862.1.6
   12:d=2  hl=2 l=   9 cons:   SEQUENCE
   14:d=3  hl=2 l=   7 prim:    OBJECT            :0.4.0.194126.1.1
```

</details>

## Wallet provider signer

[Complete PEM certificate](assets/test-certificates/wallet-provider-signer.pem)

<details>
<summary>Full decoded certificate contents</summary>

```text
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            a5:85:cc:a2:3d:b0:0d:a4:ff:ab:0e:e8:6b:6d:cc:d7
        Signature Algorithm: ecdsa-with-SHA256
        Issuer: C=NL, O=EUDI Dev Test Provider, CN=EUDI Dev Test wallet CA NL, organizationIdentifier=NTRNL-00000000
        Validity
            Not Before: Oct  2 16:21:20 2026 GMT
            Not After : Oct  2 17:21:20 2027 GMT
        Subject: C=NL, O=EUDI Dev Test Provider, CN=EUDI Dev Wallet Provider (wallet-provider), organizationIdentifier=NTRNL-00000000
        Subject Public Key Info:
            Public Key Algorithm: id-ecPublicKey
                Public-Key: (256 bit)
                pub:
                    04:ef:7a:9c:b9:4d:61:0b:4c:fd:db:e8:18:db:d1:
                    12:88:6a:cd:3b:cc:73:3d:e5:16:e9:ef:b2:a7:99:
                    89:87:4d:c9:e5:36:46:19:a4:6b:e9:c6:56:c0:d4:
                    64:77:33:55:54:34:46:cb:3d:af:82:e5:41:05:2e:
                    7a:40:7b:29:26
                ASN1 OID: prime256v1
                NIST CURVE: P-256
        X509v3 extensions:
            X509v3 Key Usage: critical
                Digital Signature
            X509v3 Subject Key Identifier:
                92:AC:A0:0B:08:F5:12:BE:C1:DA:A7:EF:01:A6:CC:11:34:87:8B:4A
            X509v3 Authority Key Identifier:
                6A:B1:1C:F3:DF:76:DE:84:AC:42:6E:2E:74:2A:ED:66:8C:E9:4B:26
            Authority Information Access:
                CA Issuers - URI:https://eudi-test.dev/api/certificates/providers/wallet/NL.der
            X509v3 Subject Alternative Name:
                DNS:eudi-test.dev, URI:https://eudi-test.dev
            X509v3 CRL Distribution Points:
                Full Name:
                  URI:https://eudi-test.dev/api/crl/providers/wallet/NL

            X509v3 Issuer Alternative Name:
                URI:https://github.com/dominikschlosser/eudi-dev
            qcStatements:
                0.0......F..0.......N..
    Signature Algorithm: ecdsa-with-SHA256
    Signature Value:
        30:45:02:21:00:8f:9f:e2:d4:31:78:48:38:a6:1f:4c:d3:1f:
        80:57:81:2d:78:81:9b:9a:4b:7c:30:20:04:be:64:f7:27:40:
        33:02:20:50:3b:03:fa:3b:80:c6:7c:41:4a:39:97:5e:01:fc:
        b6:3b:cf:d3:ab:24:07:ac:8a:41:6e:33:bf:98:4a:8c:23
```

QCStatements extension value:

```text
DER: 30153013060604008e4601063009060704008bec4e0102
    0:d=0  hl=2 l=  21 cons: SEQUENCE
    2:d=1  hl=2 l=  19 cons:  SEQUENCE
    4:d=2  hl=2 l=   6 prim:   OBJECT            :0.4.0.1862.1.6
   12:d=2  hl=2 l=   9 cons:   SEQUENCE
   14:d=3  hl=2 l=   7 prim:    OBJECT            :0.4.0.194126.1.2
```

</details>

## Local credential signer

[Complete PEM certificate](assets/test-certificates/local-signer.pem)

<details>
<summary>Full decoded certificate contents</summary>

```text
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            9f:3e:67:41:5b:f6:f3:e1:3d:cc:85:37:34:70:b5:51
        Signature Algorithm: ecdsa-with-SHA256
        Issuer: C=NL, O=EUDI Dev Test Provider, CN=EUDI Dev Test local CA NL, organizationIdentifier=NTRNL-00000000
        Validity
            Not Before: Oct  2 16:21:20 2026 GMT
            Not After : Oct  2 17:21:20 2027 GMT
        Subject: C=NL, O=EUDI Dev Test Provider, CN=EUDI Dev Wallet Issuer (local), organizationIdentifier=NTRNL-00000000
        Subject Public Key Info:
            Public Key Algorithm: id-ecPublicKey
                Public-Key: (256 bit)
                pub:
                    04:76:a0:5b:ee:29:b4:a4:b2:f0:f3:83:5c:33:c8:
                    b2:fe:2a:56:08:f0:dd:95:63:38:b9:09:f5:28:23:
                    60:16:03:2c:71:f8:ce:a3:e9:51:c8:2b:5d:77:88:
                    4d:e4:23:01:bf:cf:93:fa:ba:e2:94:b1:07:97:ba:
                    a4:6d:d5:d8:5a
                ASN1 OID: prime256v1
                NIST CURVE: P-256
        X509v3 extensions:
            X509v3 Key Usage: critical
                Digital Signature
            X509v3 Subject Key Identifier:
                DE:DE:D0:97:4C:24:23:E2:F7:B1:6D:D5:C4:9D:E2:99:F1:4E:06:50
            X509v3 Authority Key Identifier:
                FF:52:94:1C:83:CB:4D:B5:D1:07:65:7A:F7:86:02:52:3F:EE:51:18
            Authority Information Access:
                CA Issuers - URI:https://eudi-test.dev/api/certificates/providers/local/NL.der
            X509v3 Subject Alternative Name:
                DNS:eudi-test.dev, URI:https://eudi-test.dev
            X509v3 CRL Distribution Points:
                Full Name:
                  URI:https://eudi-test.dev/api/crl/providers/local/NL

            X509v3 Issuer Alternative Name:
                URI:https://github.com/dominikschlosser/eudi-dev
            X509v3 Extended Key Usage: critical
                1.0.18013.5.1.2
    Signature Algorithm: ecdsa-with-SHA256
    Signature Value:
        30:45:02:21:00:ef:d8:8c:e4:37:2e:8a:7b:42:78:a4:2a:92:
        81:fc:15:bc:44:66:89:d7:7f:49:c3:dc:bc:79:c7:25:d8:cd:
        7d:02:20:38:5a:3c:cf:1a:96:f3:db:5b:ac:db:d8:fd:2c:96:
        e8:c4:c8:3d:06:30:5d:4a:51:1e:6f:f2:76:82:11:b2:e4
```

</details>

## Access signer

[Complete PEM certificate](assets/test-certificates/access-signer.pem)

<details>
<summary>Full decoded certificate contents</summary>

```text
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            3b:47:59:b6:82:23:82:c9:f2:63:c9:66:57:51:db:40
        Signature Algorithm: ecdsa-with-SHA256
        Issuer: C=NL, O=EUDI Dev Test CA, CN=OID4VC Dev Wallet CA
        Validity
            Not Before: Oct  2 16:21:20 2026 GMT
            Not After : Oct  2 17:21:20 2027 GMT
        Subject: C=NL, O=EUDI Dev Test Provider, CN=EUDI Dev Test Access, organizationIdentifier=NTRNL-00000000
        Subject Public Key Info:
            Public Key Algorithm: id-ecPublicKey
                Public-Key: (256 bit)
                pub:
                    04:1e:5b:89:28:23:14:7d:47:16:b1:58:a6:0a:38:
                    d9:11:1b:8b:42:34:e0:27:92:19:f5:43:50:bd:39:
                    7e:6a:ae:14:ee:28:8d:da:c3:63:b7:e7:09:cc:5e:
                    2e:20:f4:7e:fe:1a:0d:02:3a:7a:7c:a0:69:8d:ac:
                    7f:7b:f0:d4:63
                ASN1 OID: prime256v1
                NIST CURVE: P-256
        X509v3 extensions:
            X509v3 Key Usage: critical
                Digital Signature
            X509v3 Subject Key Identifier:
                B1:59:68:B9:00:60:AE:A0:AE:BA:B4:1B:51:57:74:35:B2:39:56:57
            X509v3 Authority Key Identifier:
                B2:10:CA:15:53:9E:5E:4E:73:EF:5D:67:67:38:26:59:EB:DA:90:52
            Authority Information Access:
                CA Issuers - URI:https://eudi-test.dev/api/certificates/ca.der
            X509v3 Subject Alternative Name:
                DNS:eudi-test.dev, URI:https://eudi-test.dev
            X509v3 CRL Distribution Points:
                Full Name:
                  URI:https://eudi-test.dev/api/crl

            X509v3 Issuer Alternative Name:
                URI:https://github.com/dominikschlosser/eudi-dev
            X509v3 Certificate Policies:
                Policy: 0.4.0.194118.1.2
                  CPS: https://github.com/dominikschlosser/eudi-dev/blob/main/docs/test-certificates.md
    Signature Algorithm: ecdsa-with-SHA256
    Signature Value:
        30:46:02:21:00:ef:79:01:76:d2:e6:d0:81:61:df:8d:4f:05:
        ca:00:03:25:6f:98:89:76:d7:eb:f8:72:8d:19:9b:9b:83:3a:
        da:02:21:00:9a:4c:44:8d:11:61:28:cf:bd:eb:3a:b3:56:fd:
        27:56:59:00:1d:a0:46:2f:5c:f4:4a:3a:7f:a6:bf:f4:d3:d8
```

</details>

## Registrar signer

[Complete PEM certificate](assets/test-certificates/registrar-signer.pem)

<details>
<summary>Full decoded certificate contents</summary>

```text
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            0e:85:fb:54:45:a2:89:3d:f9:e8:62:45:fe:76:32:75
        Signature Algorithm: ecdsa-with-SHA256
        Issuer: C=NL, O=EUDI Dev Test CA, CN=OID4VC Dev Wallet CA
        Validity
            Not Before: Oct  2 16:21:20 2026 GMT
            Not After : Oct  2 17:21:20 2027 GMT
        Subject: C=NL, O=EUDI Dev Test Provider, CN=EUDI Dev Test Registrar, organizationIdentifier=NTRNL-00000000
        Subject Public Key Info:
            Public Key Algorithm: id-ecPublicKey
                Public-Key: (256 bit)
                pub:
                    04:0a:59:20:2b:13:3d:31:5f:65:40:e1:32:42:d5:
                    62:7f:03:6b:25:75:c8:70:00:f7:9f:92:55:56:ef:
                    e2:95:6c:81:3f:57:92:cb:85:07:2c:9c:87:2b:43:
                    42:01:a5:7a:06:1d:82:26:83:d3:92:be:df:e3:12:
                    d8:1c:ab:14:77
                ASN1 OID: prime256v1
                NIST CURVE: P-256
        X509v3 extensions:
            X509v3 Key Usage: critical
                Digital Signature
            X509v3 Subject Key Identifier:
                8C:3B:ED:92:33:D2:5A:29:38:85:D3:B3:1C:84:9F:F5:D1:5B:00:15
            X509v3 Authority Key Identifier:
                B2:10:CA:15:53:9E:5E:4E:73:EF:5D:67:67:38:26:59:EB:DA:90:52
            Authority Information Access:
                CA Issuers - URI:https://eudi-test.dev/api/certificates/ca.der
            X509v3 Subject Alternative Name:
                DNS:eudi-test.dev, URI:https://eudi-test.dev
            X509v3 CRL Distribution Points:
                Full Name:
                  URI:https://eudi-test.dev/api/crl

            X509v3 Issuer Alternative Name:
                URI:https://github.com/dominikschlosser/eudi-dev
    Signature Algorithm: ecdsa-with-SHA256
    Signature Value:
        30:44:02:20:1d:17:d4:5d:32:f8:da:0e:19:ba:51:ee:1f:62:
        1b:3c:fe:bd:73:d4:09:81:69:e1:88:dc:07:d6:36:f1:65:a5:
        02:20:22:e8:e6:8e:f2:70:b4:87:5c:fa:28:d5:cf:9f:c8:1f:
        80:a8:a6:b1:1b:e4:f3:36:eb:73:4d:3a:c2:65:f4:47
```

</details>

## Status signer

[Complete PEM certificate](assets/test-certificates/status-signer.pem)

<details>
<summary>Full decoded certificate contents</summary>

```text
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            09:40:10:82:6a:56:da:50:2a:e9:20:dd:0a:c6:3e:df
        Signature Algorithm: ecdsa-with-SHA256
        Issuer: C=NL, O=EUDI Dev Test CA, CN=OID4VC Dev Wallet CA
        Validity
            Not Before: Oct  2 16:21:20 2026 GMT
            Not After : Oct  2 17:21:20 2027 GMT
        Subject: C=NL, O=EUDI Dev Test Provider, CN=EUDI Dev Status List Signer, organizationIdentifier=NTRNL-00000000
        Subject Public Key Info:
            Public Key Algorithm: id-ecPublicKey
                Public-Key: (256 bit)
                pub:
                    04:2f:57:54:09:54:e8:d5:8b:ab:61:86:7f:ea:fb:
                    b7:58:e5:2b:00:60:3f:66:69:ea:75:53:9a:01:79:
                    24:40:ef:d7:4c:ef:06:cf:12:3b:7e:94:52:db:4c:
                    d3:f9:09:08:ea:02:c5:f8:93:dc:46:56:bd:80:11:
                    11:47:e1:14:35
                ASN1 OID: prime256v1
                NIST CURVE: P-256
        X509v3 extensions:
            X509v3 Key Usage: critical
                Digital Signature
            X509v3 Subject Key Identifier:
                38:39:FD:9E:76:C7:D0:A9:53:97:BF:9C:B8:6A:EB:20:2A:C2:4A:AC
            X509v3 Authority Key Identifier:
                B2:10:CA:15:53:9E:5E:4E:73:EF:5D:67:67:38:26:59:EB:DA:90:52
            X509v3 CRL Distribution Points:
                Full Name:
                  URI:https://eudi-test.dev/api/crl

            X509v3 Issuer Alternative Name:
                URI:https://github.com/dominikschlosser/eudi-dev
    Signature Algorithm: ecdsa-with-SHA256
    Signature Value:
        30:46:02:21:00:a8:e0:fb:b5:f7:c9:67:1e:83:f6:ff:f8:d4:
        f5:05:cc:95:4e:59:bb:c9:28:60:e8:8f:f3:b6:ac:1f:63:ad:
        08:02:21:00:94:1a:a0:90:c8:0e:4d:c6:3c:84:d6:d4:f2:30:
        7c:35:0f:62:20:3b:34:ba:82:d0:92:b2:8e:56:3d:e6:12:f6
```

</details>

## Trust list signer

[Complete PEM certificate](assets/test-certificates/trust-list-signer.pem)

<details>
<summary>Full decoded certificate contents</summary>

```text
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            76:79:e0:5b:9d:09:14:43:3d:a3:ba:e2:b9:d0:5a:c5
        Signature Algorithm: ecdsa-with-SHA256
        Issuer: C=NL, O=EUDI Dev Test CA, CN=OID4VC Dev Wallet CA
        Validity
            Not Before: Oct  2 16:21:20 2026 GMT
            Not After : Oct  2 17:21:20 2027 GMT
        Subject: C=NL, O=EUDI Dev Wallet, CN=EUDI Dev Test List Operator, organizationIdentifier=NTRNL-00000000
        Subject Public Key Info:
            Public Key Algorithm: id-ecPublicKey
                Public-Key: (256 bit)
                pub:
                    04:65:6c:56:cc:2a:67:44:55:38:dc:dd:d8:a0:fd:
                    ee:f9:14:39:65:95:48:52:5b:d9:42:ba:d8:f3:7f:
                    66:05:cb:25:e2:0c:e7:1d:37:0a:b6:d9:d3:27:e2:
                    9a:72:4b:b1:e0:ec:2a:de:c0:e4:62:c7:82:93:39:
                    ed:69:d4:ee:49
                ASN1 OID: prime256v1
                NIST CURVE: P-256
        X509v3 extensions:
            X509v3 Key Usage: critical
                Digital Signature
            X509v3 Subject Key Identifier:
                F8:CA:D1:10:FE:8C:DC:00:29:4C:0D:6E:91:2A:55:4D:EB:92:DF:7D
            X509v3 Authority Key Identifier:
                B2:10:CA:15:53:9E:5E:4E:73:EF:5D:67:67:38:26:59:EB:DA:90:52
            Authority Information Access:
                CA Issuers - URI:https://eudi-test.dev/api/certificates/ca.der
            X509v3 Subject Alternative Name:
                DNS:eudi-test.dev, URI:https://eudi-test.dev
            X509v3 CRL Distribution Points:
                Full Name:
                  URI:https://eudi-test.dev/api/crl

            X509v3 Issuer Alternative Name:
                URI:https://github.com/dominikschlosser/eudi-dev
    Signature Algorithm: ecdsa-with-SHA256
    Signature Value:
        30:45:02:20:66:95:ae:16:5e:03:26:d8:d5:9f:79:a1:38:1d:
        46:85:c2:22:d0:a4:d9:01:c7:6e:46:f5:ef:c0:a6:33:c9:4c:
        02:21:00:ad:ac:87:83:dd:13:7b:3e:ee:35:1a:b3:07:0e:09:
        50:95:76:ba:fc:20:f7:15:76:f5:5a:40:00:17:ff:42:4e
```

</details>
