package filecrypto

import (
    "bytes"
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "crypto/rsa"
    "crypto/sha256"
    "encoding/binary"
    "io"
)

func Encrypt(cert *Certificate, plaintext []byte) ([]byte, error) {
    aesKey := make([]byte, 32)
    if _, err := rand.Read(aesKey); err != nil {
        return nil, err
    }

    hash := sha256.New()
    encAESkey, err := rsa.EncryptOAEP(hash, rand.Reader, cert.PublicKey, aesKey, nil)
    if err != nil {
        return nil, err
    }

    block, err := aes.NewCipher(aesKey)
    if err != nil {
        return nil, err
    }

    nonce := make([]byte, 12)
    if _, err := rand.Read(nonce); err != nil {
        return nil, err
    }

    aesgcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }

    cipherText := aesgcm.Seal(nil, nonce, plaintext, nil)

    var buf bytes.Buffer
    if err := binary.Write(&buf, binary.BigEndian, uint16(len(encAESkey))); err != nil {
        return nil, err
    }

    buf.Write(encAESkey)
    buf.Write(nonce)
    buf.Write(cipherText)

    return buf.Bytes(), nil
}

func Decrypt(cert *Certificate, data []byte) ([]byte, error) {
    buf := bytes.NewReader(data)

    var rsaKeyLen uint16
    if err := binary.Read(buf, binary.BigEndian, &rsaKeyLen); err != nil {
        return nil, err
    }

    encAESkey := make([]byte, rsaKeyLen)
    if _, err := io.ReadFull(buf, encAESkey); err != nil {
        return nil, err
    }

    hash := sha256.New()
    aesKey, err := rsa.DecryptOAEP(hash, rand.Reader, cert.PrivateKey, encAESkey, nil)
    if err != nil {
        return nil, err
    }

    nonce := make([]byte, 12)
    if _, err := io.ReadFull(buf, nonce); err != nil {
        return nil, err
    }

    ciphertext, err := io.ReadAll(buf)
    if err != nil {
        return nil, err
    }

    block, err := aes.NewCipher(aesKey)
    if err != nil {
        return nil, err
    }

    aesgcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }

    plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
    if err != nil {
        return nil, err
    }

    return plaintext, nil
}
