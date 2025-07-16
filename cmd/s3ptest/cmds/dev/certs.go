package dev

import (
    "log"

    "github.com/spf13/cobra"
    "s3p/internal/filecrypto"
    u "s3p/internal/utils"
)

// Root Certs Command
var certsCmd = &cobra.Command{
    Use:   "certs",
    Short: "Certificate commands",
    Long:  "Certificate commands",
}

func getCertsCmd() *cobra.Command {
    certsCmd.AddCommand(genPairCmd)
    certsCmd.AddCommand(cmpPairCmd)
    certsCmd.AddCommand(genPairEnc)
    return certsCmd
}

// certs.gen-pair test output
var genPairCmd = &cobra.Command{
    Use:   "gen-pair",
    Short: "generate a certificate pair",
    Long:  "generate a certificate pair",
    Run:   useGenPair,
}

func useGenPair(cmd *cobra.Command, args []string) {
    u.PrintFatal("DEV - %s\n\n", u.GreenString("Generate a certificate:"))

    certs, err := filecrypto.CreateCertificate(4096)
    if err != nil {
        log.Fatal(err)
    }

    prv, err := certs.PrivateKeyPEM()
    if err != nil {
        log.Fatal(err)
    }
    u.PrintFatal("Private Key PEM:\n%s\n\n", prv)

    pub, err := certs.PublicKeyPEM()
    if err != nil {
        log.Fatal(err)
    }
    u.PrintFatal("Public Key PEM:\n%s\n\n", pub)
}

// certs.gen-pair test output
var cmpPairCmd = &cobra.Command{
    Use:   "gen-cmp",
    Short: "generate certificates and test comparison",
    Long:  "generate an RSA certificate key pair and print them to screen",
    Run:   useGenPairCompare,
}

func useGenPairCompare(cmd *cobra.Command, args []string) {
    u.PrintFatal("DEV - %s\n\n", u.GreenString("Generate two certificates and test comparison features:"))

    certs1, err := filecrypto.CreateCertificate(4096)
    if err != nil {
        log.Fatal(err)
    }

    certs2, err := filecrypto.CreateCertificate(4096)
    if err != nil {
        log.Fatal(err)
    }

    genPairCompareToScreen(certs1, certs2, "A != B", false)
    genPairCompareToScreen(certs1, certs1, "A == A", true)
    genPairCompareToScreen(certs2, certs1, "B != A", false)
    genPairCompareToScreen(certs2, certs2, "B == B", true)
}

func genPairCompareToScreen(cert1 *filecrypto.Certificate, cert2 *filecrypto.Certificate, name string, expected bool) {
    isMatch := filecrypto.CompareCertificates(cert1, cert2)
    if isMatch == expected {
        u.PrintFatal("%s: %s\n", name, u.GreenString("Pass"))
    } else {
        u.PrintFatal("%s: %s\n", name, u.RedString("Fail"))
    }
}

// certs.gen-pair test output
var genPairEnc = &cobra.Command{
    Use:   "gen-enc",
    Short: "generate certificates and test encryption/decryption",
    Long:  "generate an RSA certificate pair and test encryption/decryption functions",
    Run:   useGenPairEnc,
}

func useGenPairEnc(cmd *cobra.Command, args []string) {
    u.PrintFatal("DEV - %s\n\n", u.GreenString("Generate a certificate and test encryption/decryption functions:"))

    certs, err := filecrypto.CreateCertificate(4096)
    if err != nil {
        log.Fatal(err)
    }

    s := "THIS IS A TEST STRING"
    u.PrintFatal("Test data:\n%s\n\n", s)

    enc, err := filecrypto.Encrypt(certs, []byte(s))
    if err != nil {
        log.Fatal(err)
    }
    u.PrintFatal("Encrypting Text: \n%x\n\n", enc)

    dec, err := filecrypto.Decrypt(certs, enc)
    if err != nil {
        log.Fatal(err)
    }

    u.PrintFatal("Decrypting Text: \n%s\n\n", string(dec))
    if string(dec) == s {
        u.PrintFatal("%s\n", u.GreenString("PASS"))
        return
    }
    u.PrintFatal("%s\n", u.RedString("FAIL"))
}
