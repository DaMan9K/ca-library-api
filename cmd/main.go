package main

import (
	author3 "ca-library-app/internal/adapters/api/author"
	book3 "ca-library-app/internal/adapters/api/book"
	author2 "ca-library-app/internal/adapters/db/author"
	book2 "ca-library-app/internal/adapters/db/book"
	"ca-library-app/internal/composites"
	"ca-library-app/internal/domain/author"
	"ca-library-app/internal/domain/book"
)

func main() {
	//entry point
	bookComposite, err := composites.NewBookComposite()
	if err != nil {
		return nil
	}

	authorComposite, err := composites.NewAuthorComposite()
	if err != nil {
		return nil
	}

}
