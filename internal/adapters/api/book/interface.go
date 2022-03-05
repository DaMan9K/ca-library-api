package book

import (
	"ca-library-app/internal/domain/book"
	"context"
)

type Service interface {
	GetBookByUUID(ctx context.Context, uuid string) *book.Book
	GetAllBooks(ctx context.Context, limit, offset int) []*book.Book
	CreateBook(ctx context.Context, dto *book.CreateBookDTO) *book.Book
}
