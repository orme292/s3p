package filecrypto

import (
    "crypto/rand"
    "crypto/rsa"
    "crypto/x509"
    "encoding/pem"
    "errors"
)

// implementation: https://gist.github.com/miguelmota/3ea9286bd1d3c2a985b67cac4ba2130a

type CertPair struct {
    Prv []byte
    Pub []byte
}

func GenerateNewPair() (*CertPair, error) {
    prv, pub, err := newKeyPair()
    if err != nil {
        return nil, err
    }

    pair := &CertPair{
        Prv: privateToBytes(prv),
        Pub: publicToBytes(pub),
    }

    if pair.Prv == nil || pair.Pub == nil {
        return nil, errors.New("error occurred when converting to pem")
    }

    return pair, nil
}

func newKeyPair() (*rsa.PrivateKey, *rsa.PublicKey, error) {
    // generate a private key
    privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
    if err != nil {
        return nil, nil, err
    }

    // validate the key
    err = privateKey.Validate()
    if err != nil {
        return nil, nil, err
    }

    return privateKey, &privateKey.PublicKey, nil
}

func privateToBytes(priv *rsa.PrivateKey) []byte {
    prvBytes := pem.EncodeToMemory(&pem.Block{
        Type:  "RSA PRIVATE KEY",
        Bytes: x509.MarshalPKCS1PrivateKey(priv),
    })

    return prvBytes
}

func publicToBytes(pub *rsa.PublicKey) []byte {
    pubASN1, err := x509.MarshalPKIXPublicKey(pub)
    if err != nil {
        return nil
    }

    pubBytes := pem.EncodeToMemory(&pem.Block{
        Type:  "RSA PUBLIC KEY",
        Bytes: pubASN1,
    })

    return pubBytes
}
