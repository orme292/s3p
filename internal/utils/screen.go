package utils

import (
    "fmt"
    "log"
    "os"
)

func PrintFatal(f string, a ...interface{}) {
    _, err := fmt.Fprintf(os.Stderr, f, a...)
    if err != nil {
        log.Fatal(err)
    }
}

func GreenString(s string) string {
    return fmt.Sprintf("\033[1;32m%s\033[0m", s)
}

func RedString(s string) string {
    return fmt.Sprintf("\033[1;31m%s\033[0m", s)
}

func CyanString(s string) string {
    return fmt.Sprintf("\033[1;36m%s\033[0m", s)
}
