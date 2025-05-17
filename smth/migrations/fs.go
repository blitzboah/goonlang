package migrations

import (
    "embed"
    "log"
)

//go:embed *.sql
var FS embed.FS

func init() {
    files, err := FS.ReadDir(".")
    if err != nil {
        log.Fatal(err)
    }
    for _, file := range files {
        log.Println("Embedded file:", file.Name())
    }
}
