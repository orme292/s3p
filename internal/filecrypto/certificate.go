package filecrypto

import (
    "bytes"
    "crypto/rand"
    "crypto/rsa"
    "crypto/x509"
    "encoding/pem"
    "errors"
)

type Certificate struct {
    PrivateKey *rsa.PrivateKey
    PublicKey  *rsa.PublicKey
    PEM        []byte
}

func (c *Certificate) PrivateKeyPEM() (string, error) {
    if c.PrivateKey == nil {
        return "", errors.New("private key is nil")
    }

    privDER := x509.MarshalPKCS1PrivateKey(c.PrivateKey)
    var buf bytes.Buffer
    err := pem.Encode(&buf, &pem.Block{
        Type:  "RSA PRIVATE KEY",
        Bytes: privDER,
    })
    if err != nil {
        return "", err
    }

    return buf.String(), nil
}

// PublicKeyPEM returns the public key as a PEM-encoded string.
func (c *Certificate) PublicKeyPEM() (string, error) {
    if c.PublicKey == nil {
        return "", errors.New("public key is nil")
    }

    pubDER, err := x509.MarshalPKIXPublicKey(c.PublicKey)
    if err != nil {
        return "", err
    }

    var buf bytes.Buffer
    err = pem.Encode(&buf, &pem.Block{
        Type:  "PUBLIC KEY",
        Bytes: pubDER,
    })
    if err != nil {
        return "", err
    }

    return buf.String(), nil
}

func CreateCertificate(bits int) (*Certificate, error) {
    prv, err := rsa.GenerateKey(rand.Reader, bits)
    if err != nil {
        return nil, err
    }

    ASN1DER := x509.MarshalPKCS1PrivateKey(prv)
    block := &pem.Block{
        Type:  "RSA PRIVATE KEY",
        Bytes: ASN1DER,
    }

    var buf bytes.Buffer
    if err := pem.Encode(&buf, block); err != nil {
        return nil, err
    }

    return &Certificate{
        PrivateKey: prv,
        PublicKey:  &prv.PublicKey,
        PEM:        buf.Bytes(),
    }, nil
}

func ImportCertificate(pemData []byte) (*Certificate, error) {
    block, _ := pem.Decode(pemData)
    if block == nil || block.Type != "RSA PRIVATE KEY" {
        return nil, errors.New("failed to decode PEM block")
    }

    prv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
    if err != nil {
        return nil, err
    }

    return &Certificate{
        PrivateKey: prv,
        PublicKey:  &prv.PublicKey,
        PEM:        pemData,
    }, nil
}

func CompareCertificates(cert1, cert2 *Certificate) bool {
    if cert1 == nil || cert2 == nil {
        return false
    }

    return cert1.PublicKey.N.Cmp(cert2.PublicKey.N) == 0 &&
        cert1.PublicKey.E == cert2.PublicKey.E
}
