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

	fmt.Printf("%+v\n", fopApp) // &{Course:Premium SubscribeSupportService:true SubscribeMovieService:false SubscribeBackupService:true}
	fmt.Printf("%+v\n", bpApp)  // &{Course:Premium SubscribeSupportService:true SubscribeMovieService:false SubscribeBackupService:true}
}
