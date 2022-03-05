package composites

import (
	"ca-library-app/internal/adapters/api"
	"ca-library-app/internal/adapters/api/book"
	book3 "ca-library-app/internal/adapters/db/book"
	book2 "ca-library-app/internal/domain/book"
)

type BookComposite struct {
	Storage book2.Storage
	Service book.Service
	Handler api.Handler
}

func NewBookComposite(mongoComposite *MongoDBComposite) (*BookComposite, error) {
	storage := book3.NewStorage(mongoComposite.db)
	service := book2.NewService(storage)
	handler := book.NewHandler(service)
	return &BookComposite{
		Storage: storage,
		Service: service,
		Handler: handler,
	}, nil
}
