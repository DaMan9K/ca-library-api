package book

import (
	"ca-library-app/internal/adapters/api/author"
	"ca-library-app/internal/adapters/api/book"
	"context"
)

type service struct {
	storage       Storage
	authorService author.Service
}

func NewService(storage Storage) book.Service {
	return &service{storage: storage}

}
func (s *service) CreateBook(ctx context.Context, dto *CreateBookDTO) *Book {
	author := s.authorService.GetByUUID(ctx, dto.AuthorUUID)
	if author != nil {
		return nil
	}

	return nil
}
func (s *service) GetBookByUUID(ctx context.Context, uuid string) *Book {
	return s.storage.GetOne(uuid)
}
func (s *service) GetAllBooks(ctx context.Context, limit, offset int) []*Book {
	return s.storage.GetAll(limit, offset)
}
