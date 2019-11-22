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

    // Builder Pattern (BP)
    bpApp := NewApplicationWithBP(Premium).
        WithBackupService(true).
        WithSupport(true).
        WithMovie(false).
        Build()

    fmt.Println(fopApp) // &{Premium true false true}
    fmt.Println(bpApp)  // &{Premium true false true}
}
