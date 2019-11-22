package main

import (
    "fmt"
)

func main() {
    // Functional Option Pattern (FOP)
    fopApp := NewApplicationWithFOP(Premium,
        WithBackupService(true),
        WithSupport(true),
        WithMovie(false),
    )

    fmt.Println(fopApp)

    // Builder Pattern (BP)
    bpApp := NewApplicationWithBP(Premium).
        WithBackupService(true).
        WithSupport(true).
        WithMovie(false).
        Build()

    fmt.Println(bpApp)
}
