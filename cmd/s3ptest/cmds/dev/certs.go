package dev

import (
    "fmt"
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
    Run:   useCerts,
}

func getCertsCmd() *cobra.Command {
    certsCmd.AddCommand(genPairCmd)

    return certsCmd
}

func useCerts(cmd *cobra.Command, args []string) {
    fmt.Println("dev.certs called")
}

// certs.gen-pair test output
var genPairCmd = &cobra.Command{
    Use:   "gen-pair",
    Short: "generate an RSA key pair and print them to screen",
    Long:  "generate an RSA key pair and print them to screen",
    Run:   useGenPair,
}

func useGenPair(cmd *cobra.Command, args []string) {
    u.PrintFatal("DEV - %s\n\n", u.GreenString("Generate an RSA key pair and print them to the screen:"))

    certs, err := filecrypto.GenerateNewPair()
    if err != nil {
        log.Fatal(err)
    }

    u.PrintFatal("Private Key:\n%s\n\nPublic Key:\n%s\n", string(certs.Prv), string(certs.Pub))
}
