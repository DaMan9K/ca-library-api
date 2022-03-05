package book

import "ca-library-app/internal/domain/author"

type Book struct {
	UUID   string        `json:"uuid,omitempty"`
	Name   string        `json:"name,omitempty"`
	Year   int           `json:"year,omitempty"`
	Author author.Author `json:"author"`
	Busy   bool          `json:"busy,omitempty"`
	Owner  string        `json:"owner,omitempty"`
}
